import { createStore, combineReducers } from 'redux'
import user from 'modules/Auth/reducer.js'
import chat from 'modules/Chat/reducer.js'

export default createStore(combineReducers({
  chat,
  user
}))
