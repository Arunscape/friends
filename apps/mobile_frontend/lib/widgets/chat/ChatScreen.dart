import 'package:flutter/material.dart';
import 'package:mobile_frontend/widgets/chat/Message.dart';

import 'ChatSender.dart';

class ChatScreen extends StatelessWidget {
  final List<Message> msgs;

  ChatScreen({this.msgs});

  @override
  Widget build(BuildContext context) {
    return Column(children: [
      new Expanded(
          child: ListView.builder(
        reverse: true,
        itemBuilder: (context, position) {
          return new MessageWidget(msgs[position]);
        },
        itemCount: msgs.length,
      )),
      ChatSender(),
    ]);
  }
}
