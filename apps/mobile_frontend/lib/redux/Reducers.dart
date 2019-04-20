import '../models/AppState.dart';
import './Actions.dart';

AppState appReducer(AppState state, action){
  return new AppState(
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