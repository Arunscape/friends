class AppState {
  static final AppState _singleton = new AppState._internal();
  bool isLoggedin;

  factory AppState() {
    return _singleton;
  }

  AppState._internal();
}
