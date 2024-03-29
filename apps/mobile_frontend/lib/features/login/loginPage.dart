import 'package:flutter/material.dart';
import 'package:flutter_hooks/flutter_hooks.dart';
import 'package:functional_widget_annotation/functional_widget_annotation.dart';
import 'package:email_validator/email_validator.dart';
import 'package:provider/provider.dart';

import 'package:mobile_frontend/state/currentUser.dart';

part 'loginPage.g.dart';


@widget
Widget loginPage(){
  const appTitle = 'Login';

    return new MaterialApp(
      title: appTitle,
      home: Scaffold(
        appBar: AppBar(
          title: Text(appTitle),
        ),
        body: LoginForm(),
      ),
    );
}

// @hwidget
// Widget loginForm(){
  class LoginForm extends StatelessWidget{

  static final _formKey = GlobalKey<FormState>();

  // final contxt = useContext();

  Widget build(context){

  final user = Provider.of<CurrentUser>(context);
  return new Form(
    key: _formKey,
    child: Column(
      children: <Widget>[
        TextFormField(
          validator: (value){
            if (!EmailValidator.validate(value)){
              return 'Please enter a valid email address';
            }
            user.email = value;
          },
        ),
        Padding(
            padding: const EdgeInsets.symmetric(vertical: 16.0),
            child: RaisedButton(
              onPressed: () {
                // Validate will return true if the form is valid, or false if
                // the form is invalid.
                if (_formKey.currentState.validate()) {
                  // If the form is valid, we want to show a Snackbar
                  Scaffold.of(context)
                      .showSnackBar(SnackBar(content: Text('Logging in ${user.email}...' )));
                }
              },
              child: Text('Submit'),
            ),
          )
      ],
    )

  );
  }
}