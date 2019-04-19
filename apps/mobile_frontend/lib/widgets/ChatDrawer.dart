import 'package:flutter/material.dart';

class ChatDrawer extends StatelessWidget {
  void Function(String title) setTitle;
  ChatDrawer(void Function(String title) setTitle) {
    this.setTitle = setTitle;
  }

  @override
  Widget build(BuildContext context) {
    return ListView(
      padding: EdgeInsets.zero,
      children: <Widget>[
        DrawerHeader(
          child: Text('Group Select'),
          decoration: BoxDecoration(
            color: Colors.blue,
          ),
        ),
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
