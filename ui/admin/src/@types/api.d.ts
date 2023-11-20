declare namespace ApiUser {
  interface LoginParam {
    username: string
    password: string
    authcode: string
  }
}

declare namespace ApiType {
  interface ChangePwd {
    old_password: string
    new_password: string
    authcode: string
  }

  interface TokenItem {
    id: number
    title: string
    symbol: string
    status: number
    min_confirm: number
    min_address: number
    interval_time: number
    withdraw_fee_type: number
    withdraw_fee: number
    min_withdraw: number
  }

  interface HomeCounts {
    amount: {
      total: number
      curr: number
    }
    order: {
      total: number
      curr: number
    }
    user: {
      total: number
      curr: number
    }
  }

  interface HomeRecordItem {
    id: number
    tx: string
    from: string
    to: string
    amount: number
    type: number
    created_at: number
  }

  interface UserListItem {
    id: number
    key: number
    app_name: string
    web_hook: string
    app_secret: string
    app_status: number
    payment_status: number
    created_at: number
  }

  interface OrderItem {
    id: number
    key: number
    from_address: string
    to_address: string
    tx_id: string
    product_id: string
    symbol: string
    amount: number
    handle_status: number
    handle_time: number
    created_at: number
  }

  interface AddressItem {
    id: number
    key: number
    symbol: string
    address: string
    status: number
    created_at: number
  }

  interface LogsItem {
    id: number
    ip: string
    username: string
    action_type: number
    action_func: string
    params: string
    created_at: number
  }

  interface NotifyItem {
    id: number
    nonce: string
    product_id: number
    product_name: string
    item_type: string
    item_id: number
    notify_type: number
    send_url: string
    send_body: string
    handle_status: number
    handle_time: number
    handle_msg: string
    created_at: number
  }
}
