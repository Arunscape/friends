import { REPLACE_TOKEN } from 'app/store'

import Purple from '@material-ui/core/colors/purple'
import Orange from '@material-ui/core/colors/orange'

const SAVE_SETTINGS = 'SAVE_SETTINGS'

const initialState = {
  palette: {
    type: 'dark',
    primary: Purple,
    secondary: Orange
  }
}

export const saveSettings = (type, primary, secondary) => action(SAVE_SETTINGS, { type, primary, secondary })

function action (type, payload) {
  return { type, payload }
}

function getSettingsOrDefault (settings) {
  if (!settings.palette) {
    settings.palette = initialState.palette
  }
  return settings
}

export default function SettingsReducer (state = initialState, action) {
  switch (action.type) {
    case REPLACE_TOKEN:
      return getSettingsOrDefault(action.payload.settings)
    case SAVE_SETTINGS:
      return {
        ...state,
        palette: {
          type: action.payload.type,
          primary: action.payload.primary,
          secondary: action.payload.secondary
        }
      }
    default:
      return state
  }
}
