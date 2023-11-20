// 列表
export const getOrderList = (param: { page: number; pageSize?: number; status: number }) => {
  const data = { ...param, ...{ pageSize: param?.pageSize ?? 20 } }
  return makeRequest<ApiType.OrderItem, 'list'>({ method: 'post', url: '/order/list', data })
}

export const searchOrderList = (data: { search: string; page: number; status: number }) => {
  return makeRequest<ApiType.OrderItem, 'list'>({ method: 'post', url: '/order/search', data })
}
