import 'package:flutter/material.dart';
import 'package:flutter_hooks/flutter_hooks.dart';
import 'package:functional_widget_annotation/functional_widget_annotation.dart';

import 'features/home/homePage.dart';
import 'features/login/loginPage.dart';

// $ flutter packages pub run build_runner watch
part 'main.g.dart';

void main() => runApp(friends());

@FunctionalWidget(widgetType: FunctionalWidgetType.stateless)
Widget friends() => new MaterialApp(title: 'friends', initialRoute: '/login', routes: {
          '/': (context) => new HomePage(),
          '/login': (context) => new LoginPage(),
        });