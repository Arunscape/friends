import 'package:flutter/material.dart';
import '../../AppState.dart';
import '../../pages/HomePage.dart';

class SignInWithGoogle extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return new RaisedButton(
      child: Text('Login with Google'),
      onPressed: () {
        var s = new AppState();
        s.isLoggedin = true;
        Navigator.push(
          context,
          MaterialPageRoute(builder: (context) => new HomePage()),
        );
      },
    );
  }
}
