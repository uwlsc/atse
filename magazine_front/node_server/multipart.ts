import assert from 'assert';
import fs from 'fs';
import path from 'path';
import Busboy from 'busboy';

export interface CreateMultipartUploaderOptions {
  tempdir: string;
  mimeTypeExtensionPairs: { mimeType: string, extensions: string[] }[];
  fileSizeLimit: number;
  filesLimit: number;
}

export type MultipartUploadResult = {
    didUpload: boolean;
    uploaded?: {
      requestedFilename: string;
      uploadedFilename: string;
    };
    rejected?: {
      requestedFilename: string;
      reason: string;
    };
}[];

function genRandInt(low: number, high: number) {
  return Math.round( low + ( Math.random() * (high - low) ) );
}

const multipartUploadFilenameRegex = /^mpupload_\d{6}_(.+)/;

const REJECTFILE_MIMEEXT = 'bad_mimetype_extension_pair';
const REJECTFILE_UNKNOWN = 'unknown';
const REJECTFILE_LIMIT = 'file_size_too_large';

function toMultipartUploadFilename(requestedFilename: string) {
  const r = genRandInt(100000, 999999);
  return `mpupload_${r}_${requestedFilename}`;
}

export function isMultipartRequest(req: any) {
  return req && req.headers && req.headers['content-type'].startsWith('multipart/form-data; boundary=');
}

export function isMultipartUploadFilename(filename: string) {
  return multipartUploadFilenameRegex.test(filename);
}

export function parseRequestedFilename(f: string): string | null {
  const filename = path.basename(f);
  const result = multipartUploadFilenameRegex.exec(filename);
  if (!result) return null;

  return result[1];
}

export function createMultipartUploader(uploaderOpts: CreateMultipartUploaderOptions) {
  assert.ok(uploaderOpts.fileSizeLimit > 0); assert.ok(uploaderOpts.filesLimit > 0);
  assert.ok(fs.statSync(uploaderOpts.tempdir).isDirectory(), 'argument should resolve to directory');

  function isGoodMimetypeExtensionPair(filename: string, mimetype: string) {
    const extensions = uploaderOpts.mimeTypeExtensionPairs.find(pair => pair.mimeType === mimetype)?.extensions;
    // if (!extensions) return false;

    // return !!extensions.find(ex => filename.endsWith(ex));

    //NOTE: this allows any file to upload
    return true;
  }

  return (req: any) => {
    return new Promise<MultipartUploadResult>((resolve, reject) => {
      if (!isMultipartRequest(req)) {
        return reject(new Error('Only multipart requests')); }

      const busboy = new Busboy({ 
        headers: req.headers, 
        limits: { 
          fileSize: uploaderOpts.fileSizeLimit, 
          files: uploaderOpts.filesLimit 
        } 
      });

      const result: MultipartUploadResult = [];

      function acceptFile(index: number, requestedFilename: string, uploadedFilename: string) {
        if (result[index]) return;
        result[index] = {
          didUpload: true,
          uploaded: {
            requestedFilename,
            uploadedFilename
          }
        };
      }

      function rejectFile(index: number, requestedFilename: string, reason: string) {
        if (result[index]) return;
        result[index] = {
          didUpload: false,
          rejected: {
            requestedFilename,
            reason
          }
        };
      }

      let N = 0;
      busboy.on('file', (fieldname, file, filename, encoding, mimetype) => {
        const index = N++;

        if (!isGoodMimetypeExtensionPair(filename, mimetype)) {
          rejectFile(index, filename, REJECTFILE_MIMEEXT);
          file.resume();
          return;
        }

        const uploadFilename = toMultipartUploadFilename(filename); 
        const uploadFilepath = path.join(uploaderOpts.tempdir, uploadFilename);

        file.on('error', error => {
          rejectFile(index, filename, REJECTFILE_UNKNOWN);
        });

        file.on('limit', error => {
          rejectFile(index, filename, REJECTFILE_LIMIT);
        });

        file.on('end', error => {
          acceptFile(index, filename, uploadFilename);
        });

        file.pipe(fs.createWriteStream(uploadFilepath));
      });

      busboy.on('finish', () => {
        return resolve(result);
      });

      busboy.on('error', error => {
        return reject(error);
      });

      req.pipe(busboy);
    });
  };
}
