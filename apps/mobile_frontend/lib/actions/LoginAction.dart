import '../AppState.dart';

/*
import 'package:flutter_dotenv/flutter_dotenv.dart';
import 'package:google_sign_in/google_sign_in.dart';
GoogleSignIn _googleSignIn = GoogleSignIn(
  serverClientId: DotEnv().env['GOOGLE_SERVER_CLIENT_ID'],
  scopes: <String>[
    'profile',
    'email',
    'openid',
    // 'https://www.googleapis.com/auth/contacts.readonly',
  ],
);
*/

Future<void> login() async {
  // try {
  //   print("👀👀👀👀👀👀👀👀👀👀👀👀👀👀👀👀👀👀");
  //   print(DotEnv().env['GOOGLE_SERVER_CLIENT_ID']);
  //   print("👀👀👀👀👀👀👀👀👀👀👀👀👀👀👀👀👀👀");
  //   await _googleSignIn.signOut();
  //   await _googleSignIn.signIn();
  //   var s = new AppState();
  //   s.isLoggedin = true;
  // } catch (error) {
  //   print(error);
  // }

  // TEMP: BYPASS LOGIN
  var s = new AppState();
  s.isLoggedin = true;
}
