import 'package:flutter/material.dart';
import 'package:mobile_frontend/widgets/chat/Message.dart';

class ChatScreen extends StatelessWidget {
  final List<Message> msgs;
  ChatScreen({this.msgs});

  @override
  Widget build(BuildContext context) {
    return Container(
      height: double.maxFinite,
      child: ListView.builder(
        reverse: true,
        itemBuilder: (context, position) {
          return new MessageWidget(msgs[position]);
        },
        itemCount: msgs.length,
      ),
    );
  }
}
