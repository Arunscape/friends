import 'package:flutter/material.dart';

import '../actions/LoginAction.dart';
import '../AppState.dart';

class SignInWithGoogle extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return new RaisedButton(
      child: Text('Login with Google'),
      onPressed: () {
        login();
        var s = new AppState();
        if (s.isLoggedin) {
          Navigator.pushNamed(context, '/');
        }
      },
    );
  }
}