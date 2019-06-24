/*
 * No need to validate the token at all on the front end. Users could fake the whole front end anyways
 */

export function setToken (token) {
  localStorage.setItem('tok', token)
}

export function clearToken () {
  localStorage.removeItem('tok')
}

export function getTokenForHttp () {
  return localStorage.getItem('tok') || undefined
}

export function isTempTokenValid () {
  let data = getTokenData()
  return data.exp && data.exp > Date.now()
}

export function isTokenValid () {
  let data = getTokenData()
  return data.exp && data.exp > Date.now() + 1000 * 60 * 60
}

export function userHasPermission (permission) {
  let data = getTokenData()
  return data.permissions && data.permissions[permission]
}

function parseJWT (token) {
  let userData64Dirty = token.split('.')[1]
  let userData64 = userData64Dirty.replace(/-/g, '+').replace(/_/g, '/')
  let userDataJson = decodeURIComponent(atob(userData64).split('').map(c => '%' + ('00' + c.charCodeAt(0).toString(16)).slice(-2)).join(''))
  return JSON.parse(userDataJson)
}

export function getTokenData () {
  let tok = localStorage.getItem('tok')
  if (tok) {
    return parseJWT(tok)
  }
  return {}
}
