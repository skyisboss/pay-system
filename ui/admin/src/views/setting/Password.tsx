import { Button, Form, Input, notification } from 'antd'

interface Props extends WithClassName {}

const AccountView = memo((props: Props) => {
  const { className } = props
  const [form] = Form.useForm()
  const store = useUserStore()
  const navigate = useNavigate()
  const formItems = [
    {
      label: '输入旧密码',
      name: 'old_password',
      rule: [{ required: true, message: '请输入旧密码' }],
    },
    {
      label: '输入新密码',
      name: 'new_password',
      rule: [{ required: true, message: '请输入新密码' }],
    },
    {
      label: '输入安全码',
      name: 'authcode',
      rule: [{ required: true, message: '请输入安全码' }],
    },
  ]

  const { run: onFinish, loading } = useRequest((params: ApiType.ChangePwd) => changePwd(params), {
    manual: true,
    onSuccess(res) {
      notification.open({
        placement: 'top',
        message: res.msg,
        type: res?.success ? 'success' : 'error',
      })

      if (res?.success) {
        store.setLogout()
        setTimeout(() => navigate('/login'), 600)
      }
    },
  })

  return (
    <div className={className} style={{ width: '800px' }}>
      <Form form={form} onFinish={onFinish} layout="vertical" labelAlign="left" size="large" wrapperCol={{ span: 24 }}>
        {formItems.map((x, i) => (
          <Form.Item key={i} label={x.label} name={x.name} rules={x.rule}>
            <Input placeholder={x.label} />
          </Form.Item>
        ))}

        <Form.Item className="mt-8" wrapperCol={{ span: 24, offset: 10 }}>
          <Button type="primary" htmlType="submit" loading={loading}>
            保存配置
          </Button>
        </Form.Item>
      </Form>
    </div>
  )
})

const Account = styled(AccountView)``

Account.displayName = 'Account'
AccountView.displayName = 'AccountView'

export default Account
