/*
 * A collection of convenince functions for doing http requests with json
 * Each expects the incomming data to be json, then returns a promise that resolves to that json
 */
const STANDARD_HEADERS = { 'Content-type': 'application/json; charset=UTF-8' }
export function JUrl (service, endpoint) {
  const server = 'dev.friends.reckhard.ca'
  const services = ['auth', 'image', 'msg', 'email']
  if (!services.includes(service)) {
    console.warn(`Unrecognized service: ${service}`)
  }
  return `http://${service}.${server}/${endpoint}`
}

export function JGet (url, headers) {
  return JHttp(url, {}, 'get', headers || STANDARD_HEADERS)
}

export function JPost (url, body, headers) {
  return JHttp(url, body, 'post', headers || STANDARD_HEADERS)
}

export function JPut (url, body, headers) {
  return JHttp(url, body, 'put', headers || STANDARD_HEADERS)
}

export function JDelete (url, body, headers) {
  return JHttp(url, body, 'delete', headers || STANDARD_HEADERS)
}

function JHttp (url, body, method, headers) {
  return fetch(url, {
    method,
    headers,
    body: JSON.stringify(body)
  }
  ).then(errorOnStatus).then(convertToJson)
}

function convertToJson (response) {
  if ([204, 205].includes(response.status)) {
    return { status: response.status }
  }
  return { ...response.json(), status: response.status }
}
function errorOnStatus (response) {
  const status = response.status
  if (!(status >= 200 && status < 300)) {
    return Promise.reject(new Error(response))
  }
  return response
}
