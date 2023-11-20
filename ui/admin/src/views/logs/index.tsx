import PageTitle from '@/components/PageTitle'
import { Button, Divider, Modal, Radio } from 'antd'
import Table, { ColumnsType } from 'antd/es/table'

interface Props extends WithClassName {}
const { info } = Modal
const options = [
  { label: '全部', value: 0 },
  { label: '登录', value: 1 },
  { label: '操作', value: 2 },
]
const columns: ColumnsType<ApiType.LogsItem> = [
  {
    title: 'ID',
    dataIndex: 'id',
  },
  {
    title: '用户',
    dataIndex: 'username',
  },

  {
    title: '类型',
    dataIndex: 'action_type',
    render: (_, record) => {
      return <div>{options.find(x => x.value === record.action_type)?.label}</div>
    },
  },
  {
    title: '时间',
    dataIndex: 'created_at',
    render: (_, record) => {
      return <div>{format(new Date(record.created_at), 'yyyy-MM-dd HH:ii:ss')}</div>
    },
  },
  {
    title: '对象',
    dataIndex: 'action_func',
  },
  {
    title: '详情',
    dataIndex: 'params',
    render: (_, record) => {
      const obj = JSON.parse(record.params)
      const formattedStr = JSON.stringify(
        {
          IP: record.ip,
          对象: record.action_func,
          参数: obj,
        },
        null,
        2,
      )
      return (
        <Button
          size="small"
          onClick={() =>
            info({
              width: 800,
              title: '查看详情',
              content: <pre>{formattedStr}</pre>,
              okText: '确定',
            })
          }
        >
          查看详情
        </Button>
      )
    },
  },
]
const LogsView = memo((props: Props) => {
  const { className } = props
  const [total, setTotal] = useState(0)
  const [page, setPage] = useState(1)
  const [status, setStatus] = useState(0)
  const [dataList, setDataList] = useState<ApiType.LogsItem[]>([])

  const { loading } = useRequest(() => getLogsList({ action_type: status, page: page }), {
    refreshDeps: [status, page],
    onSuccess(res) {
      const data = res?.data?.rows ?? []
      const total = res?.data?.total ?? 0
      setDataList(orderBy(data, 'id', 'desc'))
      setTotal(total)
    },
  })

  return (
    <div className={className}>
      <div>
        <PageTitle total={total} />
        <Radio.Group
          options={options}
          onChange={e => setStatus(Number(e.target.value))}
          defaultValue={0}
          optionType="button"
          className="flex-none"
        />
      </div>
      <Divider />

      <Table
        loading={loading}
        columns={columns}
        dataSource={dataList}
        pagination={{
          total: total ?? 0,
          defaultCurrent: page,
          onChange(page) {
            setPage(page)
          },
        }}
      />
    </div>
  )
})

const Logs = styled(LogsView)``

Logs.displayName = 'Logs'
LogsView.displayName = 'LogsView'
export default Logs
