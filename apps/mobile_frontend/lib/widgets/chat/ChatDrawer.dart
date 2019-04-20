import 'package:flutter/material.dart';
import 'package:mobile_frontend/widgets/chat/UserInfoWidget.dart';

class Group {
  final String name;
  final String id;
  Group(this.name, this.id);
}

class ChatDrawer extends StatelessWidget {
  final void Function(String title) setTitle;
  final User usr;
  final List<Group> groups;
  ChatDrawer(this.setTitle, this.usr, this.groups);

  @override
  Widget build(BuildContext context) {
    return new ListView.builder(
        padding: EdgeInsets.zero,
        itemCount: groups.length + 1,
        itemBuilder: (BuildContext context, int index) {
          if (index == 0) {
            return UserInfoWidget(this.usr);
          }
          return ListTile(
            title: new Text(groups[index - 1].name),
            onTap: () {
              // TODO: actually set the group instead of the title
              this.setTitle(groups[index - 1].name);
              Navigator.pop(context);
            },
          );
        });
  }
}
