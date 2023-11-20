import { create } from 'zustand'

interface IUser {
  token: string
  setLogin: (v: string) => void
  setLogout: () => void
}

export const useUserStore = create<IUser>()(set => ({
  token: localStorage.getItem('x-token') || '',
  setLogin: (token: string) => {
    localStorage.setItem('x-token', token)
    set(() => ({ token: token }))
  },
  setLogout: () => {
    localStorage.removeItem('x-token')
    set(() => ({ token: '' }))
  },
}))

export const useIsLogin = () => useUserStore(state => !!state.token)
