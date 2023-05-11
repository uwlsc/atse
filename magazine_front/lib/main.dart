import 'dart:async';
import 'dart:io';
import 'dart:typed_data';

import 'package:flutter/material.dart';
import 'package:image_picker/image_picker.dart';
import 'package:pdf/pdf.dart';
import 'package:pdf/widgets.dart' as pw;
import 'package:path_provider/path_provider.dart';
import 'package:open_file/open_file.dart';

void main() {
  runApp(const PdfConverterApp());
}

class PdfConverterApp extends StatelessWidget {
  const PdfConverterApp({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      title: 'PDF Converter',
      theme: ThemeData(
        primarySwatch: Colors.blue,
      ),
      home: const PdfConvertScreen(),
    );
  }
}

class PdfConvertScreen extends StatefulWidget {
  const PdfConvertScreen({Key? key}) : super(key: key);

  @override
  _PdfConvertScreenState createState() => _PdfConvertScreenState();
}

class _PdfConvertScreenState extends State<PdfConvertScreen> {
  File? _imageFile;
  Uint8List? _pdfBytes;

  Future<void> _pickImage() async {
    // ignore: deprecated_member_use
    final imageFile = await ImagePicker().getImage(source: ImageSource.gallery);
    if (imageFile != null) {
      setState(() {
        _imageFile = File(imageFile.path);
        _pdfBytes = null;
      });
    }
  }

  Future<void> _convertToPdf() async {
    if (_imageFile == null) {
      return;
    }

    final imageBytes = await _imageFile!.readAsBytes();

    final pdf = pw.Document();
    pdf.addPage(pw.Page(build: (pw.Context context) {
      return pw.Center(
        child: pw.Image(pw.MemoryImage(imageBytes)),
      );
    }));

    final pdfBytes = await pdf.save();

    final directory = await getApplicationDocumentsDirectory();
    final pdfFile = File('${directory.path}/pdf.pdf');
    await pdfFile.writeAsBytes(pdfBytes);

    setState(() {
      _pdfBytes = pdfBytes;
    });
    print('PDF conversion complete.');
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: const Text('PDF Converter'),
      ),
      body: SingleChildScrollView(
        child: ConstrainedBox(
          constraints: BoxConstraints(
            minHeight: MediaQuery.of(context).size.height,
          ),
          child: Center(
            child: Column(
              mainAxisAlignment: MainAxisAlignment.center,
              children: <Widget>[
                if (_imageFile != null)
                  AspectRatio(
                    aspectRatio: 1,
                    child: Image.file(_imageFile!),
                  ),
                const SizedBox(height: 16),
                ElevatedButton(
                  onPressed: _pickImage,
                  child: const Text('Pick Image'),
                ),
                const SizedBox(height: 16),
                ElevatedButton(
                  onPressed: _convertToPdf,
                  child: const Text('Convert to PDF'),
                ),
                if (_pdfBytes != null) const SizedBox(height: 16),
                ElevatedButton(
                  onPressed: () async {
                    final directory = await getApplicationDocumentsDirectory();
                    final pdfFile = File('${directory.path}/pdf.pdf');
                    await OpenFile.open(pdfFile.path);
                  },
                  child: const Text('Open PDF'),
                ),
              ],
            ),
          ),
        ),
      ),
    );
  }
}
