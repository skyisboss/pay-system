export const getNotifyList = (data: { status: number; notifyType: number; page: number; pageSize?: number }) => {
  data.pageSize = data?.pageSize ?? 20
  return makeRequest<ApiType.NotifyItem, 'list'>({ method: 'post', url: '/notify/list', data })
}
