import 'package:flutter/material.dart';

import '../../pages/HomePage/HomePage.dart';

class User {
  String name;
  String email;
  String profilePic;

  User(this.name, this.email, this.profilePic);
}

class UserInfoWidget extends StatelessWidget {
  final User usr;

  UserInfoWidget(this.usr);

  @override
  Widget build(BuildContext context) {
    return UserAccountsDrawerHeader(
        accountName: Text(usr.name),
        accountEmail: Text(usr.email),
        otherAccountsPictures: <Widget>[
          IconButton(
            icon: Icon(Icons.settings),
            onPressed: () {
              Navigator.push(
                context,
                // TODO: when the settings page is created, make this button lead there
                MaterialPageRoute(builder: (context) => new HomePage()),
              );
            },
          )
        ],
        currentAccountPicture:
            CircleAvatar(backgroundImage: NetworkImage(this.usr.profilePic)));
  }
}
