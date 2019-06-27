import 'dart:io';
import 'dart:convert';

import '../../state/loginState.dart';
import 'package:http/http.dart' as http;

const String myemail = "arunscape@gmail.com";
const String myname = "Arun Woosaree";
const String mypic = "https://avatars1.githubusercontent.com/u/8227297";
final String mysignupinfo = json.encode(
{
   "Email": myemail,
   "Name":  myname,
   "Picture": mypic,
}
);
const String AUTH_URL = "http://auth.dev.friends.reckhard.ca";

var client = new http.Client();

var token;

Future<bool> isUser(email) async {


   final res = await client.post(AUTH_URL + '/isuser', body: json.encode({
    'Email': email
  }));


  if (res.statusCode == 200){
    // print(res.body);
    return true;
  }

return false;
}

String signIn(){

  const token = "";
  return token;
}

Future<String> signUp(email) async {

  final String name = "getName";
  final String pic = "getPic";
  final String body = json.encode(
    {
      "Email": email,
      "Name":  name,
      "Picture": pic,
    }
    );

  // final res = await client.post(AUTH_URL + '/signup', body: body);
  final res = await client.post(AUTH_URL + '/signup', body: mysignupinfo);

  if (res.statusCode == 200){
    // print(res.body);
    token = res.body;
    return res.body;
  } else{
    return "REEEEEEEE";
  }
}

Future<User> validate() async{

  final body = json.encode({
    "Tok": token
  });
  final res = await client.post(AUTH_URL + '/validate', body: body);
  print("validating...");
  print(res.body);
  print(res.statusCode);
  return new User(name: "validation succeeded");
}

User upgrade(){

}

Future<User> login(email) async {

  String _token;

 if ( await isUser(email)){
   _token = await signIn();
   
 } else{
   _token = await signUp(email);
 }
  return new User(name: "email should have sent");
}

main() async{
  final User u = await login(myemail);
  print(u.name);
  await validate();
  exit(0);
}