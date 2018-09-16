import Comment from './App.vue'

let API = location.origin
let auth: string

export function setAPI (url: string) {
  API = url
}
export function getAPI () {
  return API
}

export function setAuth (authCode: string) {
  auth = authCode
}
export function getAuth () {
  return auth
}

export default Comment
