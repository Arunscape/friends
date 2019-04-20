import 'package:flutter/foundation.dart';

@immutable
class LoginAction {

  bool authenticated;

  LoginAction(this.authenticated);
}