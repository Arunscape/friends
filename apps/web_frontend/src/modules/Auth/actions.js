import { JUrl, JPost } from 'api'
import { checkEmail, signup as signupAction } from './reducer'

export async function checkUser (dispatch, Email) {
  dispatch(checkEmail(Email))
  try {
    await JPost(JUrl('auth', 'isuser'), { Email })
    return true
  } catch (err) { } // 404 means not a user
  return false
}

function signinorupend (data) {
  localStorage.setItem('tok', data.Tok)
}
export async function signup (dispatch, Email, Name, Pic) {
  dispatch(signupAction(Email, Name, Pic))
  signinorupend(await JPost(JUrl('auth', 'signup'), { Email, Name, Pic }))
}

export async function signin (dispatch, Email) {
  dispatch(signupAction(Email))
  signinorupend(await JPost(JUrl('auth', 'signin'), { Email }))
}

export async function upgrade (dispatch) {
  try {
    const tok = localStorage.getItem('tok')
    let res = await JPost(JUrl('auth', 'upgrade'), { Tok: tok })
    localStorage.setItem('tok', res.Tok)
    return true
  } catch (err) { } // error means we can't upgrade yet
  return false
}

export async function signout (dispatch) {
  await JPost(JUrl('auth', 'signout'), { Tok: localStorage.getItem('tok') })
  localStorage.setItem('tok', undefined)
}
