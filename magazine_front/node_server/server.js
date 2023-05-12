const path = require('path');
const fs = require('fs');
const http = require('http');
const url = require('url');
const multipart = require('./multipart');
const mime = require('mime-types')

const multipartUploader = multipart.createMultipartUploader({
  tempdir: path.join(__dirname, 'out'),
  mimeTypeExtensionPairs: [],
  fileSizeLimit: 50 * 1024 * 1024,/*50 mb*/
  filesLimit: 1
});

const libre = require('libreoffice-convert');
libre.convertAsync = require('util').promisify(libre.convert);

function createUploadPath(filename)
{
  return path.join(__dirname, 'out', filename);
}

async function toPdf(infile)
{
  const outfile = infile + '.pdf';
  const ext = '.pdf'
  const inputPath = createUploadPath(infile);
  const outputPath = createUploadPath(outfile);

  // Read file
  const docxBuf = fs.readFileSync(inputPath);

  // Convert it to pdf format with undefined filter (see Libreoffice docs about filter)
  let pdfBuf = await libre.convertAsync(docxBuf, ext, undefined);

  // Here in done you have pdf file which you can save or transfer in another stream
  fs.writeFileSync(outputPath, pdfBuf);

  return outfile;
}

function isUploadRequest(req)
{
  const filename = getFilenameFromUrlPath(req.path);
  return filename.length == 0 && req.method.toLowerCase() == 'post';
}

function isDownloadRequest(req)
{
  const filename = getFilenameFromUrlPath(req.path);
  return req.method.toLowerCase()  == 'get' && filename.length > 0;
}

function getFilenameFromUrlPath(path)
{
  return path.split('/').join('');
}

async function uploadToServer(req, res)
{
  if (!multipart.isMultipartRequest(req))
  {
    throw new Error('Only multipart request');
  }

  const uploadResult = await multipartUploader(req);

  if (uploadResult[0].didUpload)
  {
    const toConvertFilename = uploadResult[0].uploaded.uploadedFilename;
    const convertedFilename = await toPdf(toConvertFilename);

    res.writeHead(200, {'Content-Type': 'text/plain'});
    res.write(convertedFilename);
    res.end();
  }
  else
  {
    throw new Error('Error uploading something went wrong');
  }
}

async function uploadToClient(req, res)
{
  const filename = getFilenameFromUrlPath(req.path);

  const fd = fs.openSync(createUploadPath(filename), 'r');
  const readStream = fs.createReadStream('', { fd });
  res.writeHead(200, { 'Content-Type': mime.lookup(filename) });
  readStream.pipe(res);  
}

http.createServer(async function (req, res) 
{
  console.log('request recieved');
  const parsedUrl = url.parse(req.url, true);
  req.path = (parsedUrl.pathname || '');
  req.path = req.path.replace(/%20/g, ' ');

  console.log(req.path);
  try
  {
    if (isUploadRequest(req))
    {
      await uploadToServer(req, res);
    }
    else if (isDownloadRequest(req))
    {
      await uploadToClient(req, res);
    }
    else
    {
      res.writeHead(404, {'Content-Type': 'text/plain'});
      res.write('Unknown route');
      res.end();
    }
  }
  catch(error)
  {
    console.error(error);
    res.writeHead(400, {'Content-Type': 'text/plain'});
    res.write('Error: ' + error.message || error);
    res.end();
  }

}).listen(8080, () =>
{
  console.log('listening 8080');
});