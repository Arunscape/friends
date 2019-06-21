import { JUrl, JPost } from 'services/http/api'
import { checkEmail, signin as signinAction, signup as signupAction } from './reducer'
import { setToken, clearToken } from 'services/security/token'

export async function checkUser (dispatch, Email) {
  dispatch(checkEmail(Email))
  try {
    await JPost(JUrl('auth', 'isuser'), { Email })
    return true
  } catch (err) { } // 404 means not a user
  return false
}

export async function signup (dispatch, Email, Name, Pic) {
  dispatch(signupAction(Email, Name, Pic))
  const { Tok } = await JPost(JUrl('auth', 'signup'), { Email, Name, Pic })
  setToken(Tok)
}

export async function signin (dispatch, Email) {
  dispatch(signinAction(Email))
  const { Tok } = await JPost(JUrl('auth', 'signin'), { Email })
  setToken(Tok)
}

export async function upgrade (dispatch) {
  try {
    let res = await JPost(JUrl('auth', 'upgrade'))
    setToken(res.Tok)
    return true
  } catch (err) { } // error means we can't upgrade yet
  return false
}

export async function signout (dispatch) {
  await JPost(JUrl('auth', 'signout'))
  clearToken()
}
