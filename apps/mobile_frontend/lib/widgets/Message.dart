import 'package:flutter/material.dart';

class Message extends StatelessWidget {
  String body;
  Message(String body) {
    this.body = body;
  }

  @override
  Widget build(BuildContext context) {
    return new Center(child: new Text(this.body));
  }
}
