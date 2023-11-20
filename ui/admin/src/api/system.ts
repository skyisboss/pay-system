export const loginAdmin = (data: ApiUser.LoginParam) => {
  return makeRequest({ method: 'post', url: '/login', data })
}

export const logoutAdmin = () => {
  return makeRequest({ method: 'get', url: '/logout' })
}

// 修改密码
export const changePwd = (data: ApiType.ChangePwd) => {
  return makeRequest({ method: 'post', url: '/setting/password', data })
}

// 获取系统参数
export const getConfig = () => {
  return makeRequest<ApiType.TokenItem, 'list'>({ method: 'get', url: '/setting/token' })
}

export const saveConfig = (data: ApiType.TokenItem) => {
  return makeRequest({ method: 'post', url: '/setting/token', data })
}
