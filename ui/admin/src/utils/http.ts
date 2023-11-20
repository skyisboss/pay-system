import axios, { AxiosInstance, AxiosRequestConfig } from 'axios'
import { message } from 'antd'

const devURL = '/admin/api'
const testURL = 'http://127.0.0.1:52033/api/admin/system'
const instance: AxiosInstance = axios.create({
  baseURL: testURL,
  timeout: 20000,
  headers: {
    'content-type': 'application/json;charset=UTF-8',
  },
})

// 请求拦截器
instance.interceptors.request.use(
  config => {
    // 在请求发送之前做一些处理 // 让每个请求携带token-- ['X-Token']为自定义key 请根据实际情况自行修改
    const token = localStorage.getItem('x-token')
    if (token) {
      config.headers['x-token'] = `Bearer ${token}`
    }
    return config
  },
  err => {
    console.log(err)
  },
)

// 响应拦截器
instance.interceptors.response.use(
  response => {
    const {
      data,
      config: { url },
      status,
    } = response
    if (status !== 200) {
      message.error('状态错误')
      console.warn('状态错误', url)
    }
    if (data?.err === 500) {
      message.error('请求eg错误')
      console.warn('请求错误', url)
    }

    return data
  },
  error => {
    // token错误，退出登录
    if (error.response.status === 502 && location.pathname != '/login') {
      message.error('token失效或不存在')
      localStorage.removeItem('x-token')
      window.location.replace('/login')
      return
    }
    return Promise.reject(new Error(error.message))
  },
)

export default instance

/**
 * 通用的列表返回值
 */
interface IResponse<T = any> {
  err?: number
  msg?: string
  success?: boolean
  // rows?: T
  data?: T
  total?: number
  size?: number
  page?: number
}
interface IResponseList<T = any> extends Omit<IResponse, 't'> {
  rows?: T[]
  total?: number
}
interface IResponseRows<T = any> extends Omit<IResponse, 'data'> {
  data?: {
    rows: T[]
    total?: number
    size?: number
    page?: number
  }
}
/**
 * 通用http请求
 * @param config
 * @returns
 */
export function makeRequest<T = any, R = undefined>(config: AxiosRequestConfig & { noToast?: boolean }) {
  return instance.request<T, R extends 'list' ? IResponseRows<T> : IResponse<T>>(config)
}
