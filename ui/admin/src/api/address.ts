// 列表
export const getAddressList = (param: { page: number; pageSize?: number; status: number }) => {
  const data = { ...param, ...{ pageSize: param?.pageSize ?? 20 } }
  return makeRequest<ApiType.AddressItem, 'list'>({ method: 'post', url: '/address/list', data })
}

export const createAddress = (data: { symbol: string; number: number }) => {
  return makeRequest({ method: 'post', url: '/address/create', data })
}

export const searchAddressList = (data: { search: string; page: number; status: number }) => {
  return makeRequest<ApiType.AddressItem, 'list'>({ method: 'post', url: '/address/search', data })
}
