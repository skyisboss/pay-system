import { AnyObject } from 'antd/es/_util/type'

// 商户列表
export const getUserList = (data: { page: number; pageSize: number; status: number }) => {
  return makeRequest<ApiType.UserListItem, 'list'>({ method: 'post', url: '/user/list', data })
}

// 搜索商户
export const searchUserList = (data: { search: string; status: number }) => {
  return makeRequest<ApiType.UserListItem, 'list'>({ method: 'post', url: '/user/search', data })
}

export const addUser = (data: AnyObject) => {
  return makeRequest({ method: 'post', url: '/user/add', data })
}

export const editUser = (data: AnyObject) => {
  return makeRequest({ method: 'post', url: '/user/edit', data })
}
