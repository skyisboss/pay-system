import PageTitle from '@/components/PageTitle'
import { Avatar, Button, Card, Radio, Segmented, Spin, Statistic } from 'antd'
import MyDatePicker from '@/components/date-picker'
import zh_CN from 'antd/es/date-picker/locale/zh_CN'
import { ArrowUpOutlined, UserOutlined } from '@ant-design/icons'
import Empty from '@/components/empty'

interface Props extends WithClassName {}

const today = startOfToday().getTime()
const todayEnd = today + 86400 * 1000 - 1
const rangeDays = [
  { label: '今日', key: 0, value: [today, todayEnd] },
  { label: '昨日', key: 1, value: [subDays(today, 1).getTime(), todayEnd] },
  { label: '本周', key: 2, value: [subDays(today, 7).getTime(), todayEnd] },
  { label: '本月', key: 3, value: [subDays(today, 30).getTime(), todayEnd] },
  { label: '其他', key: 4, value: [] },
]
const options = rangeDays.map(x => {
  return { label: x.label, value: x.key }
})
const HomeView = memo((props: Props) => {
  const { className } = props
  const [counts, setCounts] = useState<ApiType.HomeCounts>()
  const [recordList, setRecordList] = useState<ApiType.HomeRecordItem[]>([])
  const [selectDay, setSelectDay] = useState(0)
  const [showDateInput, setShowDateInput] = useState(false)

  const { run, loading } = useRequest(data => getCounts(data), {
    manual: true,
    onSuccess(res) {
      if (res?.success) {
        setCounts(res?.data)
      }
    },
  })

  const { run: runRecords, loading: loading2 } = useRequest((param: { type: number }) => getRecords(param), {
    manual: true,
    onSuccess(res) {
      if (res?.success) {
        setRecordList(res?.data?.rows ?? [])
      }
    },
  })

  const handleFilterDate = (day: number) => {
    const date = rangeDays.find((_x, i) => i === day)

    if (date) {
      setSelectDay(day)
      const [start, end] = date.value
      run({ start, end })
    }
  }

  useEffect(() => {
    handleFilterDate(0)
    runRecords({ type: 0 })
  }, [])

  return (
    <div className={className}>
      <div>
        <PageTitle />
        <div className="flex justify-between mb-4">
          <h3>数据统计</h3>
          <div className="flex space-x-4">
            <Radio.Group
              options={options}
              onChange={e => {
                const x = rangeDays.find(x => x.key === e.target.value)
                if (x) {
                  if (x.key === 4) {
                    setShowDateInput(true)
                  } else {
                    setShowDateInput(false)
                    handleFilterDate(x.key)
                  }
                }
              }}
              defaultValue={0}
              optionType="button"
              className="flex-none"
            />
            {showDateInput && (
              <MyDatePicker.RangePicker
                locale={zh_CN as any}
                allowClear={false}
                inputReadOnly={true}
                defaultValue={rangeDays[selectDay].value as any}
                disabledDate={current => current && current > startOfToday()}
                onChange={e => {
                  if (e) {
                    let [start, end] = e
                    start = new Date(start as any).getTime() as any
                    end = (new Date(end as any).getTime() + 86400 * 1000 - 1) as any
                    run({ start: start, end: end })
                  }
                }}
              />
            )}
          </div>
        </div>

        <div className="flex space-x-8">
          <Card className="flex flex-1 statistic-card">
            <Statistic
              loading={loading}
              title={
                <div className="flex justify-between">
                  <div>交易金额</div>
                  <div>总计: {counts?.amount?.total}</div>
                </div>
              }
              value={counts?.amount?.curr}
              precision={2}
              prefix="$"
            />
          </Card>
          <Card className="flex flex-1 statistic-card">
            <Statistic
              loading={loading}
              value={counts?.order?.curr}
              title={
                <div className="flex justify-between">
                  <div>订单数量</div>
                  <div>总计: {counts?.order?.total}</div>
                </div>
              }
            />
          </Card>
          <Card className="flex flex-1 statistic-card">
            <Statistic
              loading={loading}
              value={counts?.user?.curr}
              title={
                <div className="flex justify-between">
                  <div>商户数量</div>
                  <div>总计: {counts?.user?.total}</div>
                </div>
              }
            />
          </Card>
        </div>
      </div>

      <div className="mt-12">
        <div className="flex items-center justify-between">
          <h3>最近记录</h3>
          <Segmented
            options={['全部', '转入', '转出']}
            onChange={val => {
              ;['全部', '转入', '转出'].forEach((x, index) => {
                if (x === val) {
                  runRecords({ type: index })
                }
              })
            }}
          />
        </div>
        <Card className="mt-4">
          {loading2 ? (
            <Spin />
          ) : recordList.length === 0 ? (
            <Empty />
          ) : (
            recordList.map((item, index) => (
              <div className="flex r-item hover:bg-gray-50" key={index}>
                <div className="flex flex-1 items-center space-x-4">
                  <Avatar icon={<UserOutlined />} />
                  <div className="flex flex-col">
                    <h4>{item.tx}</h4>
                    <div className="text-gray-400">{friendlyTime(item.created_at)}</div>
                  </div>
                </div>
                <div className="flex flex-col flex-1 ml-14">
                  <div>
                    从 <span className="text-blue-500">{item.from}</span>
                  </div>
                  <div>
                    到 <span className="text-blue-500">{item.to}</span>
                  </div>
                </div>
                <div className="flex justify-end-items-center flex-1">
                  <div className="flex flex-col">
                    <div>Eth</div>
                    <div>{item.amount}</div>
                  </div>
                </div>
              </div>
            ))
          )}
        </Card>
      </div>
    </div>
  )
})
const Home = styled(HomeView)`
  .r-item {
    padding: 16px 8px;
    border-bottom: 1px solid rgba(5, 5, 5, 0.06);
    &:last-of-type {
      border-bottom: 0;
    }
  }
  .statistic-card {
    .ant-card-body {
      width: 100%;
    }
    .ant-statistic {
      height: 100%;
      display: flex;
      flex-direction: column;
      justify-content: space-between;
    }
  }
`

Home.displayName = 'Home'
HomeView.displayName = 'HomeView'

export default Home
