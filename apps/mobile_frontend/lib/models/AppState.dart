import 'package:flutter/foundation.dart';

@immutable
class AppState {

  bool authenticated;

  AppState({
    this.authenticated
  });

  AppState.initialState(): 
  authenticated=false
  ;
}