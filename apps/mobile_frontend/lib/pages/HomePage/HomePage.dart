import 'package:flutter/cupertino.dart';
import 'package:flutter/material.dart';

import '../../widgets/ChatDrawer.dart';
import '../../widgets/ChatScreen.dart';

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
  String title = 'Title';

  void setTitle(String title) {
    setState(() {
      this.title = title;
    });
  }

  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      home: DefaultTabController(
        length: 3,
        child: Scaffold(
          appBar: AppBar(
            bottom: TabBar(
              tabs: [
                Tab(icon: Icon(Icons.directions_car)),
                Tab(icon: Icon(Icons.directions_transit)),
              ],
            ),
            title: Text(this.title),
          ),
          body: TabBarView(
            children: [
              Center(child: new ChatScreen()),
              Center(child: new Text('Agenda tab')),
            ],
          ),
          drawer: Drawer(child: new ChatDrawer(this.setTitle)),
        ),
      ),
    );
  }
}
