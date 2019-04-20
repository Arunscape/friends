import 'package:flutter/foundation.dart';

@immutable
class LoginAction {

  final bool authenticated = true;

  LoginAction();
}

@immutable
class LogoutAction {

  final bool authenticated = false;

  LogoutAction();
}