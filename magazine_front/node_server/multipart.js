"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.createMultipartUploader = exports.parseRequestedFilename = exports.isMultipartUploadFilename = exports.isMultipartRequest = void 0;
const assert_1 = __importDefault(require("assert"));
const fs_1 = __importDefault(require("fs"));
const path_1 = __importDefault(require("path"));
const busboy_1 = __importDefault(require("busboy"));
function genRandInt(low, high) {
    return Math.round(low + (Math.random() * (high - low)));
}
const multipartUploadFilenameRegex = /^mpupload_\d{6}_(.+)/;
const REJECTFILE_MIMEEXT = 'bad_mimetype_extension_pair';
const REJECTFILE_UNKNOWN = 'unknown';
const REJECTFILE_LIMIT = 'file_size_too_large';
function toMultipartUploadFilename(requestedFilename) {
    const r = genRandInt(100000, 999999);
    return `mpupload_${r}_${requestedFilename}`;
}
function isMultipartRequest(req) {
    return req && req.headers && req.headers['content-type'].startsWith('multipart/form-data; boundary=');
}
exports.isMultipartRequest = isMultipartRequest;
function isMultipartUploadFilename(filename) {
    return multipartUploadFilenameRegex.test(filename);
}
exports.isMultipartUploadFilename = isMultipartUploadFilename;
function parseRequestedFilename(f) {
    const filename = path_1.default.basename(f);
    const result = multipartUploadFilenameRegex.exec(filename);
    if (!result)
        return null;
    return result[1];
}
exports.parseRequestedFilename = parseRequestedFilename;
function createMultipartUploader(uploaderOpts) {
    assert_1.default.ok(uploaderOpts.fileSizeLimit > 0);
    assert_1.default.ok(uploaderOpts.filesLimit > 0);
    assert_1.default.ok(fs_1.default.statSync(uploaderOpts.tempdir).isDirectory(), 'argument should resolve to directory');
    function isGoodMimetypeExtensionPair(filename, mimetype) {
        var _a;
        const extensions = (_a = uploaderOpts.mimeTypeExtensionPairs.find(pair => pair.mimeType === mimetype)) === null || _a === void 0 ? void 0 : _a.extensions;
        // if (!extensions) return false;
        // return !!extensions.find(ex => filename.endsWith(ex));
        //NOTE: this allows any file to upload
        return true;
    }
    return (req) => {
        return new Promise((resolve, reject) => {
            if (!isMultipartRequest(req)) {
                return reject(new Error('Only multipart requests'));
            }
            const busboy = new busboy_1.default({
                headers: req.headers,
                limits: {
                    fileSize: uploaderOpts.fileSizeLimit,
                    files: uploaderOpts.filesLimit
                }
            });
            const result = [];
            function acceptFile(index, requestedFilename, uploadedFilename) {
                if (result[index])
                    return;
                result[index] = {
                    didUpload: true,
                    uploaded: {
                        requestedFilename,
                        uploadedFilename
                    }
                };
            }
            function rejectFile(index, requestedFilename, reason) {
                if (result[index])
                    return;
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
                const uploadFilepath = path_1.default.join(uploaderOpts.tempdir, uploadFilename);
                file.on('error', error => {
                    rejectFile(index, filename, REJECTFILE_UNKNOWN);
                });
                file.on('limit', error => {
                    rejectFile(index, filename, REJECTFILE_LIMIT);
                });
                file.on('end', error => {
                    acceptFile(index, filename, uploadFilename);
                });
                file.pipe(fs_1.default.createWriteStream(uploadFilepath));
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
exports.createMultipartUploader = createMultipartUploader;
//# sourceMappingURL=data:application/json;base64,eyJ2ZXJzaW9uIjozLCJmaWxlIjoibXVsdGlwYXJ0LmpzIiwic291cmNlUm9vdCI6IiIsInNvdXJjZXMiOlsibXVsdGlwYXJ0LnRzIl0sIm5hbWVzIjpbXSwibWFwcGluZ3MiOiI7Ozs7OztBQUFBLG9EQUE0QjtBQUM1Qiw0Q0FBb0I7QUFDcEIsZ0RBQXdCO0FBQ3hCLG9EQUE0QjtBQXFCNUIsU0FBUyxVQUFVLENBQUMsR0FBVyxFQUFFLElBQVk7SUFDM0MsT0FBTyxJQUFJLENBQUMsS0FBSyxDQUFFLEdBQUcsR0FBRyxDQUFFLElBQUksQ0FBQyxNQUFNLEVBQUUsR0FBRyxDQUFDLElBQUksR0FBRyxHQUFHLENBQUMsQ0FBRSxDQUFFLENBQUM7QUFDOUQsQ0FBQztBQUVELE1BQU0sNEJBQTRCLEdBQUcsc0JBQXNCLENBQUM7QUFFNUQsTUFBTSxrQkFBa0IsR0FBRyw2QkFBNkIsQ0FBQztBQUN6RCxNQUFNLGtCQUFrQixHQUFHLFNBQVMsQ0FBQztBQUNyQyxNQUFNLGdCQUFnQixHQUFHLHFCQUFxQixDQUFDO0FBRS9DLFNBQVMseUJBQXlCLENBQUMsaUJBQXlCO0lBQzFELE1BQU0sQ0FBQyxHQUFHLFVBQVUsQ0FBQyxNQUFNLEVBQUUsTUFBTSxDQUFDLENBQUM7SUFDckMsT0FBTyxZQUFZLENBQUMsSUFBSSxpQkFBaUIsRUFBRSxDQUFDO0FBQzlDLENBQUM7QUFFRCxTQUFnQixrQkFBa0IsQ0FBQyxHQUFRO0lBQ3pDLE9BQU8sR0FBRyxJQUFJLEdBQUcsQ0FBQyxPQUFPLElBQUksR0FBRyxDQUFDLE9BQU8sQ0FBQyxjQUFjLENBQUMsQ0FBQyxVQUFVLENBQUMsZ0NBQWdDLENBQUMsQ0FBQztBQUN4RyxDQUFDO0FBRkQsZ0RBRUM7QUFFRCxTQUFnQix5QkFBeUIsQ0FBQyxRQUFnQjtJQUN4RCxPQUFPLDRCQUE0QixDQUFDLElBQUksQ0FBQyxRQUFRLENBQUMsQ0FBQztBQUNyRCxDQUFDO0FBRkQsOERBRUM7QUFFRCxTQUFnQixzQkFBc0IsQ0FBQyxDQUFTO0lBQzlDLE1BQU0sUUFBUSxHQUFHLGNBQUks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