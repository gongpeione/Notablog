import axios from 'axios'
import { getAPI } from './index'

const API = getAPI()

const fetch = axios.create({
  baseURL: process.env.NODE_ENV === 'production' ? API : '/api',
  timeout: 10000,
  headers: {},
  withCredentials: true
})

fetch.interceptors.response.use((res: any) => {
  return res.data
}, err => {
  throw err
})

export default fetch
