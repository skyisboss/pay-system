export const getCounts = (data: { start: number; end: number }) => {
  return makeRequest<ApiType.HomeCounts>({ method: 'post', url: '/home/count', data })
}

export const getRecords = (data: { type: number }) => {
  return makeRequest<ApiType.HomeRecordItem, 'list'>({ method: 'post', url: '/home/record', data })
}
