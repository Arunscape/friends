import { createStore, combineReducers } from 'redux'
import { composeWithDevTools } from 'redux-devtools-extension'

import user from 'modules/Auth/reducer.js'
import chat from 'modules/Chat/reducer.js'
import settings from 'modules/Settings/reducer.js'

export default (preloadedState) => createStore(combineReducers({
  chat,
  user,
  settings
}), preloadedState, composeWithDevTools())
