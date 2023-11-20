import { useMatches } from 'react-router-dom'
import { Layout as LayoutBox } from 'antd'
import Header from './header'
import Sider from './sider'

interface Props extends WithClassName {}

const LayoutView = memo((props: Props) => {
  const { className } = props
  const matches = useMatches()
  const navigate = useNavigate()
  const isLogin = useIsLogin()

  useEffect(() => {
    const title = (matches[1].handle as any)?.title
    if (typeof title === 'string' && title) {
      document.title = title
    }

    if (!isLogin && matches[1].pathname !== '/login') {
      navigate('/login', { replace: true })
    }
  }, [matches])

  return (
    <div className={className}>
      <LayoutBox>
        <Sider />
        <LayoutBox>
          <Header className="hedaer" />
          <LayoutBox.Content className="content">
            <div className="main">
              <Outlet />
            </div>
          </LayoutBox.Content>
        </LayoutBox>
      </LayoutBox>
    </div>
  )
})

const Layout = styled(LayoutView)`
  display: flex;
  flex-direction: column;
  width: 100%;
  min-height: 100%;
  --page-header-height: 48px;
  --page-sider-width: 200px;
  .hedaer {
    padding: 0;
    position: fixed;
    top: 0;
    z-index: 100;
    width: 100%;
    height: var(--page-header-height);
    line-height: var(--page-header-height);
    display: flex;
    .logo {
      width: var(--page-sider-width);
      padding: 0 16px;
      color: #fff;
    }
    .topbar {
      flex: 1;
      padding: 0 16px;
      color: #fff;
    }
  }
  .ant-layout-sider {
    position: fixed;
    top: 0;
    left: 0;
    z-index: 100;
    height: 100%;
    background: #fff;
    border-right: 1px solid #ececec;
  }

  .content {
    margin-left: var(--page-sider-width);
    margin-top: var(--page-header-height);
    background-color: #fbfdfe;
    .main {
      padding: 32px;
    }
  }
`

Layout.displayName = 'Layout'
LayoutView.displayName = 'LayoutView'

export default Layout
