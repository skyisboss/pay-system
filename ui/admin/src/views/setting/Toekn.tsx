import { Button, Input, InputNumber, Radio, Spin, Tabs, notification } from 'antd'
import Form, { Rule } from 'antd/es/form'

interface Props extends WithClassName {}

interface FormItemInput {
  label: string
  name: string
  rule: Rule[]
  type?: string
  options?: any[]
}

const formItem: FormItemInput[] = [
  {
    label: '是否启用',
    name: 'status',
    type: 'radio',
    options: [
      { label: '开启', value: 1 },
      { label: '暂停', value: 0 },
    ],
    rule: [{ required: true, message: '' }],
  },
  {
    label: '最小确认区块',
    name: 'min_confirm',
    rule: [{ required: true, message: '输入最小确认区块' }],
  },
  {
    label: '最小可用地址',
    name: 'min_address',
    rule: [{ required: true, message: '输入最小可用地址' }],
  },
  {
    label: '循环检测时间',
    name: 'interval_time',
    rule: [{ required: true, message: '输入循环检测时间' }],
  },
  {
    label: '手续费方式',
    name: 'withdraw_fee_type',
    type: 'radio',
    options: [
      { label: '百分比', value: '1' },
      { label: '每笔', value: '2' },
    ],
    rule: [{ required: true, message: '' }],
  },
  {
    label: '结算手续费',
    name: 'withdraw_fee',
    rule: [{ required: true, message: '输入结算手续费' }],
  },
  {
    label: '最小结算金额',
    name: 'min_withdraw',
    rule: [{ required: true, message: '输入最小结算金额' }],
  },
]

const SystemView = memo((props: Props) => {
  const { className } = props
  const [tokenList, setTokenList] = useState<ApiType.TokenItem[]>([])

  const { run: runGetConfig, loading } = useRequest(() => getConfig(), {
    manual: true,
    onSuccess(res) {
      console.log(res?.data?.rows)
      if (res?.success) {
        setTokenList(res?.data?.rows ?? [])
      }
    },
  })

  const { run: handleSubmit, loading: onFinishLoading } = useRequest(params => saveConfig(params), {
    manual: true,
    onSuccess(res) {
      notification.open({
        placement: 'top',
        message: res.msg,
        type: res?.success ? 'success' : 'error',
      })
    },
  })

  useEffect(() => {
    runGetConfig()
  }, [])

  return (
    <div className={className + ' w-full'}>
      {loading ? (
        <div className="w-full flex justify-center">
          <Spin />
        </div>
      ) : (
        <Tabs
          defaultActiveKey="1"
          className="w-full"
          tabPosition="right"
          items={tokenList.map((item, index) => {
            return {
              label: item.title,
              key: item.symbol,
              children: (
                <div className="px-4" key={index}>
                  <h2>{item.title}</h2>
                  <FormItemView
                    formItem={formItem}
                    initialValues={item}
                    onFinish={handleSubmit}
                    loading={onFinishLoading}
                  />
                </div>
              ),
            }
          })}
        />
      )}
    </div>
  )
})

const System = styled(SystemView)`
  display: flex;
`

System.displayName = 'System'
SystemView.displayName = 'SystemView'

export default System

interface FormViewProps extends Props {
  onFinish: (values: ApiType.TokenItem) => void
  formItem: FormItemInput[]
  initialValues: any
  loading: boolean
}
const FormItemView = (props: FormViewProps) => {
  const { onFinish, formItem, initialValues, loading } = props
  const [form] = Form.useForm()

  return (
    <Form
      form={form}
      onFinish={onFinish}
      layout="vertical"
      labelAlign="left"
      size="large"
      initialValues={initialValues}
    >
      <Form.Item name="id" style={{ display: 'none' }}>
        <Input />
      </Form.Item>
      <Form.Item name="symbol" style={{ display: 'none' }}>
        <Input />
      </Form.Item>
      {formItem.map((item, index) => (
        <Form.Item key={index} label={item.label} name={item.name} rules={item.rule}>
          {item?.type === 'radio' ? (
            <Radio.Group>
              {item?.options?.map((x, index) => {
                const value = typeof x.value === 'boolean' ? x.value : Number(x.value)
                return (
                  <Radio key={index} value={value}>
                    {x.label}
                  </Radio>
                )
              })}
            </Radio.Group>
          ) : (
            <InputNumber className="w-full" placeholder={`输入${item.label}`} />
          )}
        </Form.Item>
      ))}
      <Form.Item className="mt-8 flex justify-center">
        <Button type="primary" htmlType="submit" loading={loading}>
          保存配置
        </Button>
      </Form.Item>
    </Form>
  )
}
