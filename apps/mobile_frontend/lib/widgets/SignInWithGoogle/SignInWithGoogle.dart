import 'package:flutter/material.dart';

class SignInWithGoogle extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return new RaisedButton(
      child: Text('Login with Google'),
      onPressed: () {
        // Navigator.push(
        //   context,
        //   MaterialPageRoute(builder: (context) => SecondRoute()),
        // );
      },
    );
  }
}
