import 'package:flutter/material.dart';
import 'package:flutter_hooks/flutter_hooks.dart';
import 'package:functional_widget_annotation/functional_widget_annotation.dart';
import 'package:provider/provider.dart';

import 'features/home/homePage.dart';
import 'features/login/loginPage.dart';
import 'state/currentUser.dart';


// $ flutter packages pub run build_runner watch
part 'main.g.dart';


// void main()=> runApp(friends());
void main()=> runApp(providerWrapper());

@widget
Widget providerWrapper() => new MultiProvider(
  providers: [
    ChangeNotifierProvider<CurrentUser>.value(value: new CurrentUser()),
  ],
  child: friends(),
);

@hwidget
Widget friends() {

  return  new MaterialApp(title: 'friends', initialRoute: '/login', routes: {
          '/': (context) => new HomePage(),
          '/login': (context) => new LoginPage(),
        });
}


