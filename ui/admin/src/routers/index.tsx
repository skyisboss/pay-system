import { RouterProvider, createBrowserRouter } from 'react-router-dom'
import routes from './router'

const AppRoute = () => {
  return <RouterProvider router={createBrowserRouter(routes)} />
}

export default AppRoute
