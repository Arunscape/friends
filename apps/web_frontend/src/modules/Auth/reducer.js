import { REPLACE_TOKEN } from 'app/store'

const CHECK_EMAIL = 'CHECK_EMAIL'
const SIGNUP = 'SIGNUP'
const SIGNIN = 'SIGNIN'

const initialState = {
  name: '',
  email: '',
  pic: ''
}

export const signin = (email) => action(SIGNIN, { email })
export const signup = (email, name, pic) => action(SIGNUP, { email, name, pic })
export const checkEmail = (email) => action(CHECK_EMAIL, { email })

function action (type, payload) {
  return { type, payload }
}

export default function UserReducer (state = initialState, action) {
  switch (action.type) {
    case REPLACE_TOKEN:
      return {
        name: action.payload.name,
        email: action.payload.email,
        pic: action.payload.picture,
        permissions: action.payload.permissions
      }
    case SIGNUP:
      return {
        ...state,
        email: action.payload.email,
        name: action.payload.name,
        pic: action.payload.pic
      }
    case CHECK_EMAIL:
    case SIGNIN:
      return { ...state, email: action.payload.email }
    default:
      return state
  }
}
