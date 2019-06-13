import 'dart:convert';

import '../../state/loginState.dart';
import 'package:http/http.dart' as http;

final String myemail = "arun@woosaree.xyz";
final String AUTH_URL = "http://auth.dev.friends.reckhard.ca";

var client = new http.Client();

Future<bool> isUser(email) async {


  final res = await client.post(AUTH_URL + '/isuser', body: json.encode({
    'Email': email
  }));


  if (res.statusCode == 200){
    print(res.body);
    return true;
  }

return false;
}

String signIn(){

  const token = "";
  return token;
}

bool signUp(email){
  return true;
}

bool validate(){
  return false;
}

User upgrade(){

}

Future<User> login(email) async {

 if ( await isUser(email)){
   signIn();
   
 } else{
   signUp(email);
 }
 //wait for email
 bool err = validate();
 
 
 if (!err){
   return upgrade();
 }
  return User(name: "REEEEE");
}

main() async{
  final User u = await login(myemail);
  print(u);
}