import PageTitle from '@/components/PageTitle'
import { PlusOutlined } from '@ant-design/icons'
import { Button, Divider, Popconfirm, Radio, Table, Tag, Modal, Form, message, InputNumber, Select } from 'antd'
import Search from 'antd/es/input/Search'
import { ColumnsType } from 'antd/es/table'

interface Props extends WithClassName {}

const AddressView = memo((props: Props) => {
  const { className } = props
  const [total, setTotal] = useState(0)
  const [page, setPage] = useState(1)
  const [pageSize, setPageSize] = useState(10)
  const [status, setStatus] = useState(0)
  const [dataList, setDataList] = useState<ApiType.AddressItem[]>([])

  const columns: ColumnsType<ApiType.AddressItem> = [
    {
      title: 'ID',
      dataIndex: 'id',
    },
    {
      title: '地址',
      dataIndex: 'address',
      render: (_, record) => <span className="sc-text" dangerouslySetInnerHTML={{ __html: record.address }} />,
    },
    {
      title: '币种',
      dataIndex: 'symbol',
    },
    {
      title: '状态',
      dataIndex: 'status',
      render: (_, record) => {
        const text = options.find(x => x.value === record.status)
        const color = ['green', 'orange']
        return <Tag color={color[record.status]}>{text?.label}</Tag>
      },
    },
    {
      title: '时间',
      dataIndex: 'created_at',
      render: (_, record) => format(Number(record.created_at), 'yyyy-MM-dd HH:ii:ss'),
    },
    {
      title: '操作',
      dataIndex: 'actions',
      render: (_, record) => (
        <Popconfirm title="确定删除操作吗" onConfirm={() => {}} onCancel={() => {}} okText="确定" cancelText="取消">
          <a>删除</a>
        </Popconfirm>
      ),
    },
  ]
  const options = [
    { label: '全部', value: 0 },
    { label: '已使用', value: 1 },
    { label: '未使用', value: 2 },
  ]

  const { run: runList, loading: loading1 } = useRequest(param => getAddressList(param), {
    refreshDeps: [page, status],
    manual: true,
    onSuccess(res) {
      const data = res?.data?.rows ?? []
      const total = res?.data?.total ?? 0
      setDataList(orderBy(data, 'id', 'desc'))
      setTotal(total)
    },
  })

  const { run: runSearch, loading: loading2 } = useRequest(param => searchAddressList(param), {
    manual: true,
    onSuccess(res) {
      const data = res?.data?.rows ?? []
      const total = res?.data?.total ?? 0
      setDataList(orderBy(data, 'id', 'desc'))
      setTotal(total)
      setPage(1)
    },
  })

  const loading = loading1

  const { run: onFinish } = useRequest(data => createAddress(data), {
    manual: true,
    onSuccess(res) {
      message.success(res?.msg ?? 'message')
      if (res?.success) {
        Modal.destroyAll()
      }
    },
  })

  const [form] = Form.useForm()
  const handleOpenModel = () => {
    Modal.confirm({
      title: '创建地址',
      onOk: () => {},
      okText: '确定',
      cancelText: '取消',
      footer: <></>,
      content: (
        <Form form={form} layout="vertical" onFinish={onFinish}>
          <Form.Item label="选择币种" name="symbol" rules={[{ required: true, message: '请选择币种' }]}>
            <Select
              onChange={() => {}}
              placeholder="请选择币种"
              options={[
                { value: 'erc20', label: 'USDT-ERC20 以太坊网络' },
                { value: 'trc20', label: 'USDT-TRC20 TRX网络' },
                { value: 'bep20', label: 'USDT-BEP20 币安网络' },
              ]}
            />
          </Form.Item>
          <Form.Item label="创建数量" name="number" rules={[{ required: true, message: '请输入创建数量' }]}>
            <InputNumber className="w-full" min={1} max={10} placeholder="请输入创建数量" />
          </Form.Item>

          <div className="flex justify-center space-x-8 mt-8">
            <Button onClick={() => Modal.destroyAll()}>取消</Button>
            <Button type="primary" htmlType="submit" children="确定" />
          </div>
        </Form>
      ),
    })
  }

  useEffect(() => {
    runList({ page, status })
  }, [page, status])

  return (
    <div className={className}>
      <div>
        <PageTitle total={total} />
        <div className="flex justify-between">
          <Button type="primary" icon={<PlusOutlined />} onClick={handleOpenModel}>
            创建地址
          </Button>
          <div className="flex space-x-8">
            <Radio.Group
              options={options}
              onChange={e => {
                setStatus(Number(e.target.value))
              }}
              defaultValue={0}
              optionType="button"
              className="flex-none"
            />

            <Search
              className={className}
              placeholder="搜索内容"
              onSearch={value => {
                const search = value.trim()
                if (search) {
                  runSearch({ search, status })
                } else {
                  runList({ page, pageSize, status })
                  // runUserList({ page: 1, status })
                }
              }}
              allowClear
              enterButton
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

const Address = styled(AddressView)`
  .sc-text {
    em {
      color: #1677ff;
      border: 1px dashed #1677ff;
    }
  }
`

Address.displayName = 'Address'
AddressView.displayName = 'AddressView'

export default Address
