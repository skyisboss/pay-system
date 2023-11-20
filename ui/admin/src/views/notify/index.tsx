import PageTitle from '@/components/PageTitle'
import { Button, Divider, Modal, Radio, Select } from 'antd'
import Table, { ColumnsType } from 'antd/es/table'

interface Props extends WithClassName {}
const { info } = Modal
const options = [
  { label: '全部', value: -1 },
  { label: '待处理', value: 0 },
  { label: '失败', value: 1 },
  { label: '完成', value: 2 },
]
const notifyTypeOption = [
  { value: 1, label: '充值通知' },
  { value: 2, label: '提款通知' },
]

const columns: ColumnsType<ApiType.NotifyItem> = [
  {
    title: 'ID',
    dataIndex: 'id',
  },
  {
    title: '商户',
    dataIndex: 'product_name',
  },
  {
    title: '类型',
    dataIndex: 'notify_type',
    render: (_, record) => {
      const text = notifyTypeOption.find(x => x.value === record.notify_type)
      return text?.label
    },
  },
  {
    title: '回调通知',
    dataIndex: 'send_url',
    render: (_, record) => {
      const obj = JSON.parse(record.send_body)
      const formattedStr = JSON.stringify(obj, null, 2)
      return (
        <div className="flex flex-col">
          <div>{record.send_url}</div>
          <div>
            <a
              onClick={() =>
                info({
                  width: 500,
                  title: '通知内容',
                  content: <pre>{formattedStr}</pre>,
                  okText: '确定',
                })
              }
            >
              通知内容
            </a>
          </div>
        </div>
      )
    },
  },
  {
    title: '处理状态',
    dataIndex: 'handle_status',
    render: (_, record) => {
      const handleStatus = options.find(x => x.value === record.handle_status)

      return (
        <div className="flex flex-col">
          <div className="flex">
            <span>时间</span>
            <span>{record.handle_time}</span>
          </div>
          <div className="flex space-x-2">
            <span>状态: {handleStatus?.label}</span>
            <a
              onClick={() =>
                info({
                  width: 800,
                  title: '响应',
                  content: <pre>{record.handle_msg}</pre>,
                  okText: '确定',
                })
              }
            >
              响应内容
            </a>
          </div>
        </div>
      )
    },
  },
  {
    title: '创建时间',
    dataIndex: 'created_at',
    render: (_, record) => {
      return <div>{format(record.created_at, 'yyyy-MM-dd HH:ii:ss')}</div>
    },
  },
]

const NotifyView = memo((props: Props) => {
  const { className } = props
  const [total, setTotal] = useState(0)
  const [page, setPage] = useState(1)
  const [status, setStatus] = useState(0)
  const [notifyType, setNotifyType] = useState(1)
  const [dataList, setDataList] = useState<ApiType.NotifyItem[]>([])

  const { run: runList, loading } = useRequest(param => getNotifyList(param), {
    refreshDeps: [page, status, notifyType],
    manual: true,
    onSuccess(res) {
      const data = res?.data?.rows ?? []
      const total = res?.data?.total ?? 0
      setDataList(orderBy(data, 'id', 'desc'))
      setTotal(total)
    },
  })

  useEffect(() => {
    runList({ page, status, notifyType })
  }, [page, status, notifyType])

  return (
    <div className={className}>
      <div>
        <PageTitle total={total} />
        <div className="flex space-x-4">
          <Select
            defaultValue={notifyType}
            style={{ width: 120 }}
            onChange={e => setNotifyType(e)}
            options={notifyTypeOption}
          />
          <Radio.Group
            options={options}
            onChange={e => setStatus(Number(e.target.value))}
            defaultValue={0}
            optionType="button"
            className="flex-none"
          />
        </div>
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

const Notify = styled(NotifyView)``

Notify.displayName = 'Notify'
NotifyView.displayName = 'NotifyView'
export default Notify
