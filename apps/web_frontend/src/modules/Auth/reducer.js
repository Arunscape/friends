const CHECK_EMAIL = 'CHECK_EMAIL'

const initialState = {
  name: '',
  email: '',
  pic: '',
  isSignedIn: false,
  tok: ''
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
    case CHECK_EMAIL:
      return { ...state, email: action.payload.email }
    default:
      return state
  }
}
