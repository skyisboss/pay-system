import { BellOutlined, ExclamationCircleFilled, TranslationOutlined, UserOutlined } from '@ant-design/icons'
import { Avatar, Badge, Dropdown, Layout, Modal } from 'antd'
interface Props extends WithClassName {}

const { confirm, info } = Modal
const HeaderView = memo((props: Props) => {
  const { className } = props
  const store = useUserStore()
  const navigate = useNavigate()

  const items = [
    {
      key: 'authcode',
      label: '验证码',
      onClick: () => {
        info({
          title: '验证码管理',
          icon: <ExclamationCircleFilled />,
          width: 500,
          content: (
            <div>
              <img src="https://chart.googleapis.com/chart?cht=qr&chs=380&chl=otpauth://totp/贝壳支付系统:admin?secret=GVKWYEQPGWHQPI265U62R5WJTJXFFS3W&issuer=贝壳支付系统" />
            </div>
          ),
          okText: '确定',
          onOk() {
            console.log('OK')
          },
        })
      },
    },
    {
      key: 'account',
      label: '账号设置',
      onClick: () => {
        navigate('/setting/password')
      },
    },
    {
      type: 'divider',
    },
    {
      key: 'logout',
      label: '注销登录',
      onClick: () =>
        confirm({
          title: '确定退出登录吗?',
          icon: <ExclamationCircleFilled />,
          cancelText: '取消',
          okText: '确定',
          onOk: () => runLogout(),
        }),
    },
  ] as any

  const { run: runLogout } = useRequest(() => logoutAdmin(), {
    manual: true,
    onSuccess(res) {
      if (res?.success) {
        store.setLogout()
        navigate('/login')
      }
    },
  })

  return (
    <Layout.Header className={className}>
      <div className="logo">贝壳支付系统</div>
      <div className="topbar flex justify-end space-x-4">
        <div className="item">
          <TranslationOutlined />
        </div>
        <div className="item">
          <Badge size="small" count={5}>
            <BellOutlined style={{ color: '#fff' }} />
          </Badge>
        </div>

        <div className="item">
          <Dropdown menu={{ items }} placement="bottom">
            <div className="flex items-center space-x-1">
              <Avatar icon={<UserOutlined />} size="small" style={{ backgroundColor: '#87d068' }} />
              <span>admin</span>
            </div>
          </Dropdown>
        </div>
      </div>
    </Layout.Header>
  )
})

const Header = styled(HeaderView)`
  .item {
    cursor: pointer;
    padding: 0 12px;
    &:hover {
      background: #252a3d;
    }
  }
`

Header.displayName = 'Header'
HeaderView.displayName = 'HeaderView'

export default Header
