import 'package:flutter/material.dart';

class Message {
  String body;
  String author;

  Message(this.body, this.author);

  bool isMine() {
    return false;
  }
}

class MessageWidget extends StatelessWidget {
  final Message msg;

  MessageWidget(this.msg);

  @override
  Widget build(BuildContext context) {
    return new FractionallySizedBox(
        widthFactor: 0.7,
        alignment: Alignment(this.msg.isMine() ? 1 : -1, 0),
        child: Card(
            child: Container(
                padding: EdgeInsets.all(20),
                color: this.msg.isMine() ? Colors.lightBlue : Colors.amber,
                child: new Center(child: new Text(this.msg.body)))));
  }
}
