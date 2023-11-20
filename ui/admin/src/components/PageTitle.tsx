import { useMatches } from 'react-router-dom'

interface Props extends WithClassName {
  total?: string | number
}

const PageTitleView = memo((props: Props) => {
  const { className, total } = props
  const matches = useMatches()
  const [title, setTitle] = useState('')

  useEffect(() => {
    const title = (matches[1].handle as any)?.title
    if (typeof title === 'string' && title) {
      setTitle(title)
    }
  }, [matches])

  return (
    <h1 className={className}>
      <span>{title}</span>
      {total !== undefined && <span className="ml-4">({total})</span>}
    </h1>
  )
})

const PageTitle = styled(PageTitleView)`
  margin-bottom: 32px;
`

PageTitle.displayName = 'PageTitle'
PageTitleView.displayName = 'PageTitleView'

export default PageTitle
