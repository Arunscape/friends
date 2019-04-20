import 'package:flutter/cupertino.dart';
import 'package:flutter/material.dart';
import 'package:mobile_frontend/widgets/chat/ChatScreen.dart';
import 'package:mobile_frontend/widgets/chat/Message.dart';
import 'package:mobile_frontend/widgets/chat/UserInfoWidget.dart';

import '../../widgets/chat/ChatDrawer.dart';

class HomePage extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return new HomePageInternal();
  }
}

class HomePageInternal extends StatefulWidget {
  @override
  HomeState createState() => HomeState();
}

class HomeState extends State {
  String title = 'Friends';
  List<Message> msgs = [new Message("Body one", "Jacob")];
  List<Group> groups = [
    new Group("Engineering Friends", "sjfakjsdnfjksdn"),
    new Group("Programming Friends", "dsjlfakdlkfjasl"),
    new Group("Highschool Friends", "slafjdklnvjkasd"),
  ];
  User usr = new User("Jacob Reckhard", "jacob@reckhard.ca",
      "https://lh3.googleusercontent.com/a-/AAuE7mDJpoJdWan5dsUF0hKdoSlJoNh88Z3Nmt_DG6ju=s192");

  void setTitle(String title) {
    setState(() {
      this.title = title;
    });
  }

  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      theme: ThemeData.dark(),
      home: DefaultTabController(
        length: 3,
        child: Scaffold(
          appBar: AppBar(
            bottom: TabBar(
              tabs: [
                Tab(icon: Icon(Icons.message)),
                Tab(icon: Icon(Icons.calendar_view_day)),
              ],
            ),
            title: Text(this.title),
          ),
          body: TabBarView(
            children: [
              Center(child: new ChatScreen(msgs: this.msgs)),
              Center(child: new Text('Agenda tab')),
            ],
          ),
          drawer: Drawer(
              child: new ChatDrawer(this.setTitle, this.usr, this.groups)),
        ),
      ),
    );
  }
}
