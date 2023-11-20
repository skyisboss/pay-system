export const getLogsList = (data: { action_type: number; page: number; pageSize?: number }) => {
  if (!data?.pageSize) {
    data.pageSize = 10
  }
  return makeRequest<ApiType.LogsItem, 'list'>({ method: 'post', url: '/logs/list', data })
}
