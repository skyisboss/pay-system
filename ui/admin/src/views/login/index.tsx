import { Alert, Button, Card, Checkbox, Form, Input } from 'antd'

interface Props extends WithClassName {}

const LoginView = memo((props: Props) => {
  const { className } = props
  const title = '登录后台'
  const [form] = Form.useForm()
  const navigate = useNavigate()
  const store = useUserStore()
  const isLogin = useIsLogin()
  const [error, setError] = useState('')

  const { run, loading } = useRequest(data => loginAdmin(data), {
    manual: true,
    onSuccess(res) {
      if (res?.success) {
        store.setLogin(res?.data?.token ?? '')
        navigate('/')
      } else {
        setError(res?.msg ?? '发生错误')
      }
    },
  })
  const handleSubmit = (values: any) => {
    if (!isLogin) {
      setError('')
      run(values)
    }
  }

  useEffect(() => {
    document.title = title
  }, [])

  return (
    <div className={className}>
      <div className="flex justify-center items-center h-full">
        <Card title={<h2 className="text-center my-4">{title}</h2>} style={{ width: 440 }}>
          <Form
            onFinish={handleSubmit}
            labelAlign="left"
            size="large"
            wrapperCol={{ span: 24 }}
            style={{ maxWidth: 600 }}
          >
            <Form.Item<ApiUser.LoginParam> name="username" rules={[{ required: true, message: '请输入账号' }]}>
              <Input placeholder="账号" />
            </Form.Item>

            <Form.Item<ApiUser.LoginParam> name="password" rules={[{ required: true, message: '请输入密码' }]}>
              <Input.Password placeholder="密码" />
            </Form.Item>

            <Form.Item<ApiUser.LoginParam> name="authcode" rules={[{ required: true, message: '请输入验证码' }]}>
              <Input placeholder="验证码" />
            </Form.Item>

            <Form.Item wrapperCol={{ span: 24 }} className="mt-12">
              <Button block type="primary" htmlType="submit" loading={loading}>
                登录
              </Button>
            </Form.Item>
          </Form>

          {error && <Alert message={error} type="error" showIcon closable />}
        </Card>
      </div>
    </div>
  )
})

const Login = styled(LoginView)`
  background-color: #fafafa;
  height: 100%;
`

Login.displayName = 'Login'
LoginView.displayName = 'LoginView'

export default Login
