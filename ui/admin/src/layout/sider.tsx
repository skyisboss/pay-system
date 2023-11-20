import { DashboardOutlined, FileProtectOutlined, SettingOutlined, TeamOutlined } from '@ant-design/icons'
import { Layout, Menu } from 'antd'
interface Props extends WithClassName {}

const menuData = [
  { label: '控制台', key: 'home', icon: <DashboardOutlined />, path: '/' },
  { label: '商户管理', key: 'user', icon: <TeamOutlined />, path: '/user' },
  { label: '订单管理', key: 'order', icon: <FileProtectOutlined />, path: '/order' },
  { label: '地址管理', key: 'address', icon: <TeamOutlined />, path: '/address' },
  { label: '通知管理', key: 'notify', icon: <TeamOutlined />, path: '/notify' },
  { label: '任务管理', key: 'task', icon: <TeamOutlined />, path: '/notify' },
  { label: '日志管理', key: 'logs', icon: <TeamOutlined />, path: '/logs' },
  {
    label: '系统设置',
    key: 'setting',
    icon: <SettingOutlined />,
    path: '/setting',
  },
]
const SiderView = memo((props: Props) => {
  const { className } = props
  const navigate = useNavigate()
  const [defaultKey, setDefaultKey] = useState<string[]>([''])

  const onSelect = (item: any) => {
    const menu = menuData.find(x => x.key === item.key)
    if (menu?.path) {
      setDefaultKey([menu.key])
      navigate(menu.path)
    }
  }

  useEffect(() => {
    const active = menuData.find(x => x.path === location.pathname)
    if (active) {
      setDefaultKey([active.key])
    }
  }, [location.pathname])

  return (
    <Layout.Sider className={className}>
      <Menu
        className="side-menu"
        mode="inline"
        defaultSelectedKeys={defaultKey}
        selectedKeys={defaultKey}
        style={{ height: '100%', borderRight: 0 }}
        items={menuData}
        onSelect={onSelect}
      />
    </Layout.Sider>
  )
})

const Sider = styled(SiderView)`
  .side-menu {
    margin-top: var(--page-header-height);
    height: 100%;
    overflow: auto;
  }
  .ant-menu-item {
    width: 100%;
    border-radius: 0;
    position: relative;
    &.ant-menu-item-selected {
      .ant-menu-title-content {
        &:after {
          content: '';
          position: absolute;
          top: 0;
          right: 4px;
          bottom: 0;
          border-right: 3px solid #1890ff;
        }
      }
    }
  }
`

Sider.displayName = 'Sider'
SiderView.displayName = 'SiderView'

export default Sider
