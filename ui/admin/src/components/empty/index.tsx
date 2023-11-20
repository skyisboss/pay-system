import { Empty as MyEmpty } from 'antd'

interface Props extends WithClassName {
  title?: string
}

const EmptyView = memo((props: Props) => {
  const { className, title } = props

  return <MyEmpty className={className} description={title || '暂无数据'} style={{ margin: '32px' }} />
})

const Empty = styled(EmptyView)``

Empty.displayName = 'Empty'
EmptyView.displayName = 'EmptyView'

export default Empty
