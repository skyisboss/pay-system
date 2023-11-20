import Layout from '@/layout'
import { Skeleton } from 'antd'
import React from 'react'
import { RouteObject } from 'react-router-dom'

export const routes: RouteObject[] = [
  {
    path: '/',
    element: <Layout />,
    children: [
      {
        path: '',
        handle: { title: '控制台' },
        element: lazy(() => import('@/views/home')),
      },
      {
        path: 'user',
        handle: { title: '商户管理' },
        element: lazy(() => import('@/views/user')),
      },
      {
        path: 'order',
        handle: { title: '订单管理' },
        element: lazy(() => import('@/views/order')),
      },
      {
        path: 'address',
        handle: { title: '地址管理' },
        element: lazy(() => import('@/views/address')),
      },
      {
        path: 'logs',
        handle: { title: '日志管理' },
        element: lazy(() => import('@/views/logs')),
      },
      {
        path: 'notify',
        handle: { title: '通知管理' },
        element: lazy(() => import('@/views/notify')),
      },
      {
        path: 'setting',
        handle: { title: '系统管理' },
        element: lazy(() => import('@/views/setting')),
        children: [
          {
            path: 'token',
            index: true,
            handle: { title: '币种设置' },
            element: lazy(() => import('@/views/setting/Toekn')),
          },
          {
            path: 'password',
            index: true,
            handle: { title: '账号设置' },
            element: lazy(() => import('@/views/setting/Password')),
          },
        ],
      },
    ],
  },
  {
    path: '/login',
    handle: { title: '登录后台' },
    element: lazy(() => import('@/views/login')),
  },
]

function lazy(callback: () => Promise<{ default: React.ComponentType<any> }>) {
  const LazyComp = React.lazy(callback)

  return (
    <React.Suspense fallback={<Skeleton active />}>
      <LazyComp />
    </React.Suspense>
  )
}

export default routes
