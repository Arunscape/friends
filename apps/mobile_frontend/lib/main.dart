import 'package:flutter/material.dart';
import 'package:flutter_hooks/flutter_hooks.dart';
import 'package:functional_widget_annotation/functional_widget_annotation.dart';

import 'features/home/homePage.dart';
import 'features/login/loginPage.dart';
import 'state/loginState.dart';


// $ flutter packages pub run build_runner watch
part 'main.g.dart';



// void main() => runApp(app());

// @widget
// Widget app(){

//   return new loginNotifierProvider<User>.value(
//     value: loginNotifier,
//     child: friends()
//   );
// }

void main()=> runApp(friends());

@hwidget
Widget friends() {

  return  new MaterialApp(title: 'friends', initialRoute: '/login', routes: {
          '/': (context) => new HomePage(),
          '/login': (context) => new LoginPage(),
        });
}


