// https://flutterbyexample.com/flutter-redux-setup
// import 'package:redux/redux.dart';
// import 'package:flutter_redux/flutter_redux.dart';
// import 'package:redux_logging/redux_logging.dart';
// import 'reducers/ducks.dart';
import 'package:flutter/material.dart';
import 'package:flutter_dotenv/flutter_dotenv.dart';
import 'pages/LoginPage/LoginPage.dart';
import 'pages/HomePage/HomePage.dart';
import 'AppState.dart';

Future main() async {
  await DotEnv().load('.env');
  runApp(new Friends());
}

class Friends extends StatelessWidget {
  // final store = new Store<AppState>(
  //   appReducer,
  //   initialState: new AppState(),
  //   middleware: [new LoggingMiddleware.printer()],
  // );

  @override
  Widget build(BuildContext context) {
    // Wrap your MaterialApp in a StoreProvider
    // return new StoreProvider(
    //   // new
    //   store: store, // new
    //   child: new MaterialApp(
    //     title: title,
    //     home: new LoginPage(),
    //   ),
    // );

    var s = new AppState();
    return new MaterialApp(
      
      title: 'friends',

      initialRoute: s.isLoggedin ? '/' : '/login',
      routes:{
        '/': (context) => new HomePage(),
        '/login': (context) => new LoginPage(),
      }
    );
  }
}