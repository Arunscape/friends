import 'package:flutter/material.dart';

import 'Message.dart';

class ChatScreen extends StatelessWidget {
  final List<String> msgs = ["One", "Two", "Three"];

  @override
  Widget build(BuildContext context) {
    return Container(
      height: double.maxFinite,
      color: Colors.deepOrange,
      child: ListView.builder(
        reverse: true,
        itemBuilder: (context, position) {
          return Card(
            child: new Message(msgs[position]),
          );
        },
        itemCount: msgs.length,
      ),
    );
  }
}
