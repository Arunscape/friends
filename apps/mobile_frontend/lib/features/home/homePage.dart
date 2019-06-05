import 'package:flutter/material.dart';
import 'package:flutter_hooks/flutter_hooks.dart';
import 'package:functional_widget_annotation/functional_widget_annotation.dart';

part 'homePage.g.dart';

@hwidget
Widget homePage() => MaterialApp(
      title: 'This is the Homepage',
      home: Scaffold(
        appBar: AppBar(
          title: Text('Homepage'),
        ),
        body: Center(
          child: Text('Hello World'),
        ),
      ),
    );
