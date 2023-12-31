/* eslint-disable */
/* prettier-ignore */
// @ts-nocheck
// noinspection JSUnusedGlobalSymbols
// Generated by unplugin-auto-import
export {}
declare global {
  const AppRoute: typeof import('../routers/index')['AppRoute']
  const Link: typeof import('react-router-dom')['Link']
  const NavLink: typeof import('react-router-dom')['NavLink']
  const Navigate: typeof import('react-router-dom')['Navigate']
  const Outlet: typeof import('react-router-dom')['Outlet']
  const Route: typeof import('react-router-dom')['Route']
  const Routes: typeof import('react-router-dom')['Routes']
  const addDays: typeof import('date-fns')['addDays']
  const addMilliseconds: typeof import('date-fns')['addMilliseconds']
  const addMinutes: typeof import('date-fns')['addMinutes']
  const addSeconds: typeof import('date-fns')['addSeconds']
  const addUser: typeof import('../api/user')['addUser']
  const chain: typeof import('lodash')['chain']
  const changePwd: typeof import('../api/system')['changePwd']
  const chunk: typeof import('lodash')['chunk']
  const cloneDeep: typeof import('lodash')['cloneDeep']
  const concat: typeof import('lodash')['concat']
  const config: typeof import('../api/system')['config']
  const createAddress: typeof import('../api/address')['createAddress']
  const createRef: typeof import('react')['createRef']
  const debounce: typeof import('lodash')['debounce']
  const delay: typeof import('lodash')['delay']
  const difference: typeof import('lodash')['difference']
  const differenceBy: typeof import('lodash')['differenceBy']
  const differenceInDays: typeof import('date-fns')['differenceInDays']
  const drop: typeof import('lodash')['drop']
  const editUser: typeof import('../api/user')['editUser']
  const endOfDay: typeof import('date-fns')['endOfDay']
  const filter: typeof import('lodash')['filter']
  const find: typeof import('lodash')['find']
  const findIndex: typeof import('lodash')['findIndex']
  const flatten: typeof import('lodash')['flatten']
  const floor: typeof import('lodash')['floor']
  const forEach: typeof import('lodash')['forEach']
  const format: typeof import('date-fns')['format']
  const forwardRef: typeof import('react')['forwardRef']
  const friendlyFormatTime: typeof import('../utils/helper')['friendlyFormatTime']
  const friendlyTime: typeof import('../utils/helper')['friendlyTime']
  const get: typeof import('lodash')['get']
  const getAddressList: typeof import('../api/address')['getAddressList']
  const getConfig: typeof import('../api/system')['getConfig']
  const getCountData: typeof import('../api/home')['getCountData']
  const getCounts: typeof import('../api/home')['getCounts']
  const getDay: typeof import('date-fns')['getDay']
  const getLogsList: typeof import('../api/logs')['getLogsList']
  const getNotifyList: typeof import('../api/notify')['getNotifyList']
  const getOrderList: typeof import('../api/order')['getOrderList']
  const getRecord: typeof import('../api/home')['getRecord']
  const getRecords: typeof import('../api/home')['getRecords']
  const getTime: typeof import('date-fns')['getTime']
  const getUserList: typeof import('../api/user')['getUserList']
  const groupBy: typeof import('lodash')['groupBy']
  const hasIn: typeof import('lodash')['hasIn']
  const head: typeof import('lodash')['head']
  const http: typeof import('../utils/http')['default']
  const i18n: typeof import('../i18n/index')['default']
  const identity: typeof import('lodash')['identity']
  const init: typeof import('../i18n/index')['init']
  const isAfter: typeof import('date-fns')['isAfter']
  const isBefore: typeof import('date-fns')['isBefore']
  const isEmpty: typeof import('lodash')['isEmpty']
  const isEqual: typeof import('lodash')['isEqual']
  const isNil: typeof import('lodash')['isNil']
  const isNumber: typeof import('lodash')['isNumber']
  const isSameDay: typeof import('date-fns')['isSameDay']
  const isUndefined: typeof import('lodash')['isUndefined']
  const keys: typeof import('lodash')['keys']
  const last: typeof import('lodash')['last']
  const lazy: typeof import('react')['lazy']
  const loginAdmin: typeof import('../api/system')['loginAdmin']
  const logoutAdmin: typeof import('../api/system')['logoutAdmin']
  const makeRequest: typeof import('../utils/http')['makeRequest']
  const map: typeof import('lodash')['map']
  const max: typeof import('lodash')['max']
  const memo: typeof import('react')['memo']
  const negate: typeof import('lodash')['negate']
  const niceModalCreate: typeof import('@ebay/nice-modal-react')['create']
  const niceModalHide: typeof import('@ebay/nice-modal-react')['hide']
  const niceModalRegister: typeof import('@ebay/nice-modal-react')['register']
  const niceModalShow: typeof import('@ebay/nice-modal-react')['show']
  const niceModalUseModal: typeof import('@ebay/nice-modal-react')['useModal']
  const nth: typeof import('lodash')['nth']
  const omit: typeof import('lodash')['omit']
  const omitBy: typeof import('lodash')['omitBy']
  const orderBy: typeof import('lodash')['orderBy']
  const pick: typeof import('lodash')['pick']
  const pickBy: typeof import('lodash')['pickBy']
  const random: typeof import('lodash')['random']
  const range: typeof import('lodash')['range']
  const reduce: typeof import('lodash')['reduce']
  const reduceRight: typeof import('lodash')['reduceRight']
  const riendlyTime: typeof import('../utils/helper')['riendlyTime']
  const router: typeof import('../routers/router')['default']
  const routers: typeof import('../routers/index')['default']
  const routes: typeof import('../routers/router')['routes']
  const saveConfig: typeof import('../api/system')['saveConfig']
  const searchAddressList: typeof import('../api/address')['searchAddressList']
  const searchData: typeof import('../api/user')['searchData']
  const searchOrderList: typeof import('../api/order')['searchOrderList']
  const searchUserData: typeof import('../api/user')['searchUserData']
  const searchUserList: typeof import('../api/user')['searchUserList']
  const shuffle: typeof import('lodash')['shuffle']
  const some: typeof import('lodash')['some']
  const sortBy: typeof import('lodash')['sortBy']
  const split: typeof import('lodash')['split']
  const startOfDay: typeof import('date-fns')['startOfDay']
  const startOfToday: typeof import('date-fns')['startOfToday']
  const startTransition: typeof import('react')['startTransition']
  const styled: typeof import('@emotion/styled')['default']
  const subDays: typeof import('date-fns')['subDays']
  const subMonths: typeof import('date-fns')['subMonths']
  const sumBy: typeof import('lodash')['sumBy']
  const take: typeof import('lodash')['take']
  const toLower: typeof import('lodash')['toLower']
  const union: typeof import('lodash')['union']
  const unionBy: typeof import('lodash')['unionBy']
  const uniq: typeof import('lodash')['uniq']
  const uniqBy: typeof import('lodash')['uniqBy']
  const useCallback: typeof import('react')['useCallback']
  const useContext: typeof import('react')['useContext']
  const useDebugValue: typeof import('react')['useDebugValue']
  const useDeferredValue: typeof import('react')['useDeferredValue']
  const useEffect: typeof import('react')['useEffect']
  const useHref: typeof import('react-router-dom')['useHref']
  const useId: typeof import('react')['useId']
  const useImperativeHandle: typeof import('react')['useImperativeHandle']
  const useInRouterContext: typeof import('react-router-dom')['useInRouterContext']
  const useInsertionEffect: typeof import('react')['useInsertionEffect']
  const useIsLogin: typeof import('../store/user')['useIsLogin']
  const useLayoutEffect: typeof import('react')['useLayoutEffect']
  const useLazy: typeof import('../routers/index')['useLazy']
  const useLinkClickHandler: typeof import('react-router-dom')['useLinkClickHandler']
  const useLocation: typeof import('react-router-dom')['useLocation']
  const useMemo: typeof import('react')['useMemo']
  const useNavigate: typeof import('react-router-dom')['useNavigate']
  const useNavigationType: typeof import('react-router-dom')['useNavigationType']
  const useOutlet: typeof import('react-router-dom')['useOutlet']
  const useOutletContext: typeof import('react-router-dom')['useOutletContext']
  const useParams: typeof import('react-router-dom')['useParams']
  const useReducer: typeof import('react')['useReducer']
  const useRef: typeof import('react')['useRef']
  const useRequest: typeof import('ahooks')['useRequest']
  const useResolvedPath: typeof import('react-router-dom')['useResolvedPath']
  const useRoutes: typeof import('react-router-dom')['useRoutes']
  const useSearchParams: typeof import('react-router-dom')['useSearchParams']
  const useState: typeof import('react')['useState']
  const useSyncExternalStore: typeof import('react')['useSyncExternalStore']
  const useTransition: typeof import('react')['useTransition']
  const useTranslation: typeof import('react-i18next')['useTranslation']
  const useUserStore: typeof import('../store/user')['useUserStore']
  const userAdd: typeof import('../api/user')['userAdd']
  const values: typeof import('lodash')['values']
}
