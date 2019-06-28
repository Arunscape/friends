import 'package:flutter/material.dart';

class User{
 String email;
 String name;
 String avatarURL;
 User({this.email="", this.name="", this.avatarURL=""});
}


class CurrentUser with ChangeNotifier{

  bool _authenticated;
  String _email;
  String _token;
  String _firstName;
  String _lastName;
  String _fullName;
  String _phoneNumber;
  String _avatarURL;

  bool get authenticated => _authenticated;
  String get email => _email;
  set email(e){
    this._email = e;
    notifyListeners();

  }
  String get token => _token;
  set token(t){
    this._token = t;
    notifyListeners();
  }
  String get firstName => _firstName;
   set firstName(f){
    this._firstName = f;
    updateFullName();
    notifyListeners();
  }
  String get lastName => _lastName;
  set lastName(l){
    this._firstName = l;
    updateFullName();
    notifyListeners();
  }
  String get fullname => _fullName;
  String get phoneNumber => _phoneNumber;
  set phoneNumber(p){
    this._phoneNumber = p;
    notifyListeners();
  }
  String get avatarURL => _avatarURL;
  set avatarURL(a){
    this._avatarURL = a;
    notifyListeners();
  }

  updateFullName(){
    this._fullName = "${this.firstName} ${this.lastName}";
  }

  CurrentUser({authenticated=false, email, token, firstName, lastName, phoneNumber=null, avatarURL });
}