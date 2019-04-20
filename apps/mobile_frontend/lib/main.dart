import 'package:redux/redux.dart';
import 'package:flutter_redux/flutter_redux.dart';
import 'package:redux_thunk/redux_thunk.dart';

import 'package:flutter/material.dart';
import 'package:flutter_dotenv/flutter_dotenv.dart';
import 'pages/LoginPage/LoginPage.dart';
import 'pages/HomePage/HomePage.dart';

import 'models/AppState.dart';
import 'redux/Actions.dart';
import 'redux/Reducers.dart';

Future main() async {
  await DotEnv().load('.env');
  runApp(new Friends());
}

class Friends extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    final Store<AppState> store = new Store<AppState>(
      appReducer,
      initialState: new AppState.initialState(),
      // middleware:
    );

    // Wrap your MaterialApp in a StoreProvider
    return new StoreProvider<AppState>(
        store: store,
        child:
            new MaterialApp(title: 'friends', initialRoute: '/login', routes: {
          '/': (context) => new HomePage(),
          '/login': (context) => new LoginPage(),
        }));
  }
}
