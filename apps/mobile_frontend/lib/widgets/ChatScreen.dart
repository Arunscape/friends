import 'package:flutter/material.dart';

import './Message.dart';

class ChatScreen extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return ListView(
      padding: EdgeInsets.zero,
      physics: ScrollPhysics(),
      children: <Widget>[
        Center(
          child: ListBody(children: [
            new Message("Hey look, a text"),
            new Message("This is a replay"),
            new Message("And a third message, cause why not."),
          ]),
        )
      ],
    );
  }
}
