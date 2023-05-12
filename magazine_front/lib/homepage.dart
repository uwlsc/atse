import 'dart:io';
import 'package:flutter/material.dart';
import 'package:google_nav_bar/google_nav_bar.dart';
import 'package:file_picker/file_picker.dart';
import 'package:flutter_pdfview/flutter_pdfview.dart';
import 'package:path_provider/path_provider.dart';
import 'package:flutter_easyloading/flutter_easyloading.dart';
import './api.dart' as api;

class Homepage extends StatefulWidget{
  const Homepage({Key? key}) : super(key:key);

  @override
  State<Homepage> createState() => _HomePageState();
}

class _HomePageState extends State<Homepage>{
  void pickFile() async
  {
    FilePickerResult? result = await FilePicker.platform.pickFiles();

    if (result != null) 
    {
      try
      {
        PlatformFile file = result.files.first;

        EasyLoading.show(status: 'Converting...');

        //1. upload to server to convert
        String convertedFilename = await api.convertFile(file.path!);

        //2. Download the converted file from server
        List<int> bytes = await api.downloadFile(convertedFilename);

        //3. Store the downloaded file in application temp folder
        Directory tempDir = await getTemporaryDirectory();

        String cacheFilename = tempDir.path + '/' + convertedFilename;
        
        final output = File(cacheFilename);
        await output.writeAsBytes(bytes);

        EasyLoading.dismiss();

        //4. Show pdf
        Navigator.push(
          context,
          MaterialPageRoute(builder: (context) => ViewPage(pdfPath: cacheFilename)),
        );
      }
      catch(error)
      {
        EasyLoading.dismiss();
        print('Error occured ${error}');
        EasyLoading.showError('${error}');
      }    
    }
    else 
    {
      // User canceled the file picker.
    }
  }

  @override
  Widget build(BuildContext context)
  {
    Color primaryColor = Theme.of(context).primaryColor;

    return Scaffold(
      appBar: AppBar(title: Text('PDF Converter')),
      body: Center(
        child: GestureDetector(
              child: Container(
              padding: EdgeInsets.all(40.0),
              decoration: BoxDecoration(
                border: Border.all(
                  color: primaryColor,
                  width: 2.0,
                ),
                borderRadius: BorderRadius.circular(5.0),
              ),
            child: Text(
              '+',
              style: TextStyle(fontSize: 32, color: primaryColor),
            ),
          ),
          onTap: pickFile,
        )
      //   child:  
      )
    );
  }
}


class ViewPage extends StatelessWidget 
{
  final String pdfPath;

  const ViewPage({Key? key, required this.pdfPath}) : super(key: key);

  @override
  Widget build(BuildContext context)
  {
    return Scaffold(
      appBar: AppBar(title: Text('PDF Converter')),
      body: PDFView(filePath: pdfPath)    
    );
  }
}