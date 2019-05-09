const CHECK_EMAIL = 'CHECK_EMAIL'
const SIGNUP = 'SIGNUP'

const initialState = {
  name: '',
  email: '',
  pic: '',
  isSignedIn: false,
  tok: ''
}

export function signup (email, name, pic) {
  return {
    type: SIGNUP,
    payload: {
      email,
      name,
      pic
    }
  }
}

export function checkEmail (email) {
  return {
    type: CHECK_EMAIL,
    payload: {
      email
    }
  }
}

export default function UserReducer (state = initialState, action) {
  switch (action.type) {
    case SIGNUP:
      return {
        ...state,
        email: action.payload.email,
        name: action.payload.name,
        pic: action.payload.pic
      }
    case CHECK_EMAIL:
      return { ...state, email: action.payload.email }
    default:
      return state
  }
}
