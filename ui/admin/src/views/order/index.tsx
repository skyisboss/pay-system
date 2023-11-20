import PageTitle from '@/components/PageTitle'
import { Divider, Radio, Table, Tag } from 'antd'
import { AnyObject } from 'antd/es/_util/type'
import Search from 'antd/es/input/Search'
import { ColumnsType } from 'antd/es/table'

interface Props extends WithClassName {}

const options = [
  { label: '全部', value: 0 },
  { label: '待处理', value: 1 },
  { label: '处理中', value: 2 },
  { label: '已完成', value: 3 },
]

const OrderView = memo((props: Props) => {
  const { className } = props
  const [total, setTotal] = useState(0)
  const [page, setPage] = useState(1)
  const [searchPage, setSearchPage] = useState(1)
  const [status, setStatus] = useState(0)
  const [dataList, setDataList] = useState<ApiType.OrderItem[]>([])

  const columns: ColumnsType<ApiType.OrderItem> = [
    {
      title: '订单',
      dataIndex: 'tx_id',
      render: (_, record) => {
        return (
          <div className="flex flex-col flex-auto">
            <div>{format(Number(record.created_at), 'yyyy-MM-dd HH:ii:ss')}</div>
            <div>
              <span className="sc-text" dangerouslySetInnerHTML={{ __html: record.tx_id }} />
            </div>
          </div>
        )
      },
    },
    {
      title: '交易',
      dataIndex: 'tx_info',
      render: (_, record) => {
        return (
          <div className="flex flex-col flex-auto">
            <div>
              从: <span className="sc-text" dangerouslySetInnerHTML={{ __html: record.from_address }} />
            </div>
            <div>
              到: <span className="sc-text" dangerouslySetInnerHTML={{ __html: record.to_address }} />
            </div>
          </div>
        )
      },
    },
    {
      title: '金额',
      dataIndex: 'amount',
      render: (_, record) => {
        return (
          <div className="flex flex-col flex-auto">
            <div>{record.symbol}</div>
            <div>{record.amount}</div>
          </div>
        )
      },
    },
    {
      title: '状态',
      dataIndex: 'status',
      render: (_, record) => {
        const statusText = options.find(x => x.value === record.handle_status) as AnyObject
        const color = ['', '', 'orange', 'green']
        return (
          <div className="flex flex-col flex-auto">
            <div>{format(Number(record.handle_time), 'yyyy-MM-dd HH:ii:ss')}</div>
            <div>
              <Tag color={color[record.handle_status ?? 0]}>{statusText?.label}</Tag>
            </div>
          </div>
        )
      },
    },
    {
      title: '操作',
      dataIndex: 'actions',
      render: (_, record) => <a onClick={() => {}}>编辑</a>,
    },
  ]

  const { run: runList, loading: loading1 } = useRequest(param => getOrderList(param), {
    refreshDeps: [page, status],
    manual: true,
    onSuccess(res) {
      const data = res?.data?.rows ?? []
      const total = res?.data?.total ?? 0
      setDataList(orderBy(data, 'id', 'desc'))
      setTotal(total)
    },
  })

  const { run: runSearch, loading: loading2 } = useRequest(param => searchOrderList(param), {
    manual: true,
    onSuccess(res) {
      const data = res?.data?.rows ?? []
      const total = res?.data?.total ?? 0
      setDataList(orderBy(data, 'id', 'desc'))
      setTotal(total)
      setPage(1)
    },
  })

  const loading = loading1 || loading2

  useEffect(() => {
    runList({ page, status })
  }, [page, status])

  return (
    <div className={className}>
      <div>
        <PageTitle total={total} />
        <div className="flex justify-between">
          <Radio.Group
            disabled={loading}
            options={options}
            onChange={e => {
              setStatus(Number(e.target.value))
            }}
            defaultValue={0}
            optionType="button"
            className="flex-none"
          />

          <div>
            <Search
              className={className}
              placeholder="搜索内容"
              onSearch={value => {
                const search = value.trim()
                if (search) {
                  runSearch({ search, status, page: searchPage })
                } else {
                  setSearchPage(1)
                  runList({ page: 1, status })
                }
              }}
              allowClear
              enterButton
              disabled={loading}
            />
          </div>
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

const Order = styled(OrderView)`
  .sc-text {
    em {
      color: #1677ff;
      border: 1px dashed #1677ff;
    }
  }
`

Order.displayName = 'Order'
OrderView.displayName = 'OrderView'

export default Order
