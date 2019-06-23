import { saveSettings as saveSettingsAction } from './reducer'

import { JPost, JUrl } from 'services/http/api'
import { setToken } from 'services/security/token'

export async function saveSettings (dispatch, type, primary, secondary) {
  dispatch(saveSettingsAction(type, primary, secondary))
  let { Tok } = await JPost(JUrl('auth', 'set-user-preferences'), { Settings: JSON.stringify({ palette: { type, primary, secondary } }) })
  setToken(Tok)
}
