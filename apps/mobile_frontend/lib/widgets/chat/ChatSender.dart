import 'package:flutter/material.dart';

class ChatSender extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return new Card(
        child: Container(
            padding: EdgeInsets.fromLTRB(10, 0, 0, 0),
            child: Row(children: [
              new Expanded(child: new TextField()),
              IconButton(icon: Icon(Icons.send), onPressed: () {})
            ])));
  }
}
