import { JUrl, JPost } from 'api'
import { checkEmail } from './reducer'

export async function checkUser (dispatch, Email) {
  dispatch(checkEmail(Email))
  try {
    await JPost(JUrl('auth', 'isuser'), { Email })
    return true
  } catch (err) {
  }
  return false
}
