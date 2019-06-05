import 'package:flutter/material.dart';
import 'package:flutter_hooks/flutter_hooks.dart';
import 'package:functional_widget_annotation/functional_widget_annotation.dart';

part 'loginPage.g.dart';

@hwidget
Widget loginPage() {
  
  final email = useState("");
  
return new MaterialApp(
      title: 'This is the LoginPage',
      home: Scaffold(
        appBar: AppBar(
          title: Text(email.value),
        ),
        body: Center(
          child: TextField(
          obscureText: false,
          decoration: InputDecoration(
              contentPadding: EdgeInsets.fromLTRB(20.0, 15.0, 20.0, 15.0),
              hintText: "Email",
              border:
                  OutlineInputBorder(borderRadius: BorderRadius.circular(32.0)),
                      ),
          onChanged: (e){
            email.value = e;
          },

        ),
        ),
      ),
    );
}

