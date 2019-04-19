import 'package:flutter/material.dart';
import 'package:mobile_frontend/widgets/chat/UserInfoWidget.dart';

class ChatDrawer extends StatelessWidget {
  final void Function(String title) setTitle;
  final User usr;
  ChatDrawer(this.setTitle, this.usr);

  @override
  Widget build(BuildContext context) {
    return ListView(
      padding: EdgeInsets.zero,
      children: <Widget>[
        UserInfoWidget(this.usr),
        ListTile(
          title: Text('Engineering Friends!'),
          onTap: () {
            this.setTitle('Engineering Friends!');
            Navigator.pop(context);
          },
        ),
        ListTile(
          title: Text('Programming Friends!'),
          onTap: () {
            this.setTitle('Programming Friends!');
            Navigator.pop(context);
          },
        ),
        ListTile(
          title: Text('High School Friends!'),
          onTap: () {
            this.setTitle('High School Friends!');
            Navigator.pop(context);
          },
        ),
      ],
    );
  }
}
