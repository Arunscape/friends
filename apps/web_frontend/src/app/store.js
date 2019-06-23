import { createStore, combineReducers } from 'redux'
import { composeWithDevTools } from 'redux-devtools-extension'

import user from 'modules/Auth/reducer.js'
import chat from 'modules/Chat/reducer.js'
import settings from 'modules/Settings/reducer.js'

export function getInitalStateFromToken (tok) {
  console.log(tok)
  return {
    user: {
      name: tok.name,
      email: tok.email,
      pic: tok.picture,
      permissions: tok.permissions
    },
    chat: {
      groups: tok.groups
    },
    settings: JSON.parse(tok.settings)
  }
}


export default (preloadedState) => createStore(combineReducers({
  chat,
  user,
  settings
}), preloadedState, composeWithDevTools())
