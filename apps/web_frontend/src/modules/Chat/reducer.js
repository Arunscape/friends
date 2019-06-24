import { REPLACE_TOKEN } from 'app/store'

const initialState = {}

export default function ChatReducer (state = initialState, action) {
  switch (action) {
    case REPLACE_TOKEN:
      return {
        groups: action.payload.groups
      }
    default:
      return state
  }
}
