import { createStore, combineReducers } from 'redux'
import user from 'reducers/user'
import chat from 'reducers/chat'

export default createStore(combineReducers({
  chat,
  user
}))
