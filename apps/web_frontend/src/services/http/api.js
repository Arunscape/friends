import { getTokenForHttp } from 'services/security/token'
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

export async function JGet (url, headers) {
  return JHttp(url, {}, 'get', headers || STANDARD_HEADERS)
}

export async function JPost (url, body, headers) {
  return JHttp(url, body, 'post', headers || STANDARD_HEADERS)
}

export async function JPut (url, body, headers) {
  return JHttp(url, body, 'put', headers || STANDARD_HEADERS)
}

export async function JDelete (url, body, headers) {
  return JHttp(url, body, 'delete', headers || STANDARD_HEADERS)
}

async function JHttp (url, body = {}, method, headers) {
  body.Tok = getTokenForHttp()
  let data = await fetch(url, {
    method,
    headers,
    body: JSON.stringify(body)
  })
  return convertToJson(await errorOnStatus(data))
}

async function convertToJson (response) {
  if ([204, 205].includes(response.status)) {
    return { status: response.status }
  }
  let data = await response.json()
  console.log(data)
  return { ...data, status: response.status }
}
async function errorOnStatus (response) {
  const status = response.status
  if (!(status >= 200 && status < 300)) {
    return Promise.reject(new Error(response))
  }
  return response
}
