import 'package:flutter/material.dart';
import 'package:flutter_hooks/flutter_hooks.dart';
import 'package:functional_widget_annotation/functional_widget_annotation.dart';
import 'package:email_validator/email_validator.dart';

part 'loginPage.g.dart';


@widget
Widget loginPage(){
  final appTitle = 'Login';

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

@hwidget
Widget loginForm(){
  final _formKey = GlobalKey<FormState>();

  final context = useContext();

  return new Form(
    key: _formKey,
    child: Column(
      children: <Widget>[
        TextFormField(
          validator: (value){
            if (!EmailValidator.validate(value)){
              return 'Please enter a valid email address';
            }
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
                      .showSnackBar(SnackBar(content: Text('Logging in...')));
                }
              },
              child: Text('Submit'),
            ),
          )
      ],
    )

  );
}



// class LoginForm extends StatefulWidget {
//   @override
//   LoginFormState createState() {
//     return LoginFormState();
//   }
// }



// // Create a corresponding State class. This class will hold the data related to
// // the form.
// class LoginFormState extends State<LoginForm> {
//   // Create a global key that will uniquely identify the Form widget and allow
//   // us to validate the form
//   //
//   // Note: This is a GlobalKey<FormState>, not a GlobalKey<LoginFormState>!
//   final _formKey = GlobalKey<FormState>();

//   @override
//   Widget build(BuildContext context) {
//     // Build a Form widget using the _formKey we created above
//     return Form(
//       key: _formKey,
//       child: Column(
//         crossAxisAlignment: CrossAxisAlignment.start,
//         children: <Widget>[
//           TextFormField(
//             validator: (value) {
//               if (value.isEmpty) {
//                 return 'Please enter some text';
//               }
//             },
//           ),
//           Padding(
//             padding: const EdgeInsets.symmetric(vertical: 16.0),
//             child: RaisedButton(
//               onPressed: () {
//                 // Validate will return true if the form is valid, or false if
//                 // the form is invalid.
//                 if (_formKey.currentState.validate()) {
//                   // If the form is valid, we want to show a Snackbar
//                   Scaffold.of(context)
//                       .showSnackBar(SnackBar(content: Text('Processing Data')));
//                 }
//               },
//               child: Text('Submit'),
//             ),
//           ),
//         ],
//       ),
//     );
//   }
// }
