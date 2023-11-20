import { Button, Divider, Form, Input, Radio, Spin, message } from 'antd'

interface Props extends WithClassName {
  data?: ApiType.UserListItem
  close: () => void
}

const EditFormView = memo((props: Props) => {
  const { className, data, close } = props
  const [form] = Form.useForm()
  const isEdit = !!data?.id

  const { run: runSubmit, loading } = useRequest(data => (isEdit ? editUser(data) : addUser(data)), {
    manual: true,
    onSuccess(res) {
      if (res?.success) {
        close()
        message.success('操作成功')
      }
    },
  })

  return (
    <div className={className}>
      <Spin spinning={loading}>
        <Form
          form={form}
          layout="vertical"
          labelAlign="left"
          // size="large"
          wrapperCol={{ span: 24 }}
          style={{ maxWidth: 600 }}
          initialValues={data}
          onFinish={runSubmit}
        >
          <Form.Item label="商户名称" name="app_name" rules={[{ required: true, message: '请输入商户名称' }]}>
            <Input placeholder="请输入商户名称" />
          </Form.Item>
          <Form.Item label="回调地址" name="web_hook" rules={[{ required: true, message: '请输入回调地址' }]}>
            <Input placeholder="请输入回调地址" />
          </Form.Item>
          {isEdit && (
            <>
              <Form.Item name="id" rules={[{ required: true }]} style={{ display: 'none' }}>
                <Input />
              </Form.Item>

              <Form.Item label="商户密钥" name="app_secret" rules={[{ required: true, message: '请输入商户密钥' }]}>
                <Input.Password placeholder="请输入商户密钥" readOnly />
              </Form.Item>
            </>
          )}

          <Form.Item label="商户状态" name="app_status" rules={[{ required: true, message: '请选择商户状态' }]}>
            <Radio.Group onChange={() => {}} value={1}>
              <Radio value={0}>禁用</Radio>
              <Radio value={1}>启用</Radio>
            </Radio.Group>
          </Form.Item>
          <Form.Item label="结算权限" name="payment_status" rules={[{ required: true, message: '请选择结算权限' }]}>
            <Radio.Group onChange={() => {}} value={1}>
              <Radio value={0}>关闭</Radio>
              <Radio value={1}>开启</Radio>
            </Radio.Group>
          </Form.Item>
          <Divider />
          <div className="flex justify-center space-x-8">
            <Button onClick={() => close()}>取消</Button>
            <Button type="primary" htmlType="submit">
              保存
            </Button>
          </div>
        </Form>
      </Spin>
    </div>
  )
})

const EditForm = styled(EditFormView)``

EditForm.displayName = 'EditForm'
EditFormView.displayName = 'EditFormView'
export default EditForm
