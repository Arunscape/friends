import '../models/AppState.dart';
import './Actions.dart';

AppState appReducer(AppState state, action){
  return new AppState(
    authenticated: AuthenticationReducer(state.authenticated, action),
    // example: Exampleeducer(state.example, action),
  );
}

// List<ExampleItem> ExampleReducer(state, action){
  
//   if (action is ExampleAction){
//     return []
//             ..addAll(state)
//             ..add(new ExampleItem(
//                 var1: action.var1, 
//                 var2: action.var2));
//   }
//   return state;
// }

bool AuthenticationReducer(state, action){

  if (action is LoginAction){
    return true;
  }

  if(action is LogoutAction){
    return false;
  }
}