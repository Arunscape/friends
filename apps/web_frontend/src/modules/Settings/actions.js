import { saveSettings as saveSettingsAction } from './reducer'

export function saveSettings (dispatch, type, primary, secondary) {
  dispatch(saveSettingsAction(type, primary, secondary))
  // TODO: make http to actually save user colors to server
}
