import React from 'react'
import ReactDOM from 'react-dom/client'
import AppRoute from './routers'
import NiceModal from '@ebay/nice-modal-react'
import './index.css'
// import '@/api/mock'

ReactDOM.createRoot(document.getElementById('root')!).render(
  <React.StrictMode>
    <NiceModal.Provider>
      <AppRoute />
    </NiceModal.Provider>
  </React.StrictMode>,
)
