import 'dart:convert';

import 'package:http/http.dart' as http;

Future<String> convertFile(String filepath) async
{
  var uri = Uri.parse('http://10.0.0.197:8080/');

  var request = http.MultipartRequest('POST', uri);
  var reqOption = await http.MultipartFile.fromPath(
      'file',
      filepath
  );
  request.files.add(reqOption);

  var res = await request.send();


  final bytes = await res.stream.toBytes();
  final result = utf8.decode(bytes);

  if (res.statusCode == 200)
  {
    return result;
  }

  throw Exception(result);
}

Future<List<int>> downloadFile(String filename) async
{
  var uri = Uri(
    scheme: 'http',
    host: '10.0.0.197',
    port: 8080,
    path: filename
  );

  print(uri);

  final response = await http.get(uri);

  if (response.statusCode == 200)
  {
    return response.bodyBytes.toList();
  }
  else
  {
    final body = utf8.decode(response.bodyBytes);
    throw Exception(body);
  }
}


//   ..files.add(await http.MultipartFile.fromPath(
//       'package', 'build/package.tar.gz',
//       contentType: MediaType('application', 'x-tar')));
// var response = await request.send();
// if (response.statusCode == 200) print('Uploaded!');