// https://flutterbyexample.com/flutter-redux-setup
import 'package:flutter/material.dart';
import 'package:redux/redux.dart';
import 'package:flutter_redux/flutter_redux.dart';
import 'package:redux_logging/redux_logging.dart';
import 'reducers/ducks.dart';

import 'AppState.dart';
import 'pages/Login.dart';

void main() => runApp(new Friends());

class Friends extends StatelessWidget {
  String title = 'Friends';
  final store = new Store<AppState>(
    appReducer,
    initialState: new AppState(),
    middleware: [new LoggingMiddleware.printer()],
  );

  @override
  Widget build(BuildContext context) {
    // Wrap your MaterialApp in a StoreProvider
    return new StoreProvider(
      // new
      store: store, // new
      child: new MaterialApp(
        title: title,
        home: new LoginPage(),
      ),
    );
  }
}
