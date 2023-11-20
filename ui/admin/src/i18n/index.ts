import i18n from 'i18next'
import { initReactI18next } from 'react-i18next'

const defaultLang: string = localStorage.getItem('lang') || 'zh'

export const init = async () => {
  let all = import.meta.glob('./*.json')
  let translation: any = await all[`./${defaultLang}.json`]()
  let nsRules: any = await all[`./${defaultLang}_rules.json`]()
  const resources = {
    translation,
    nsRules,
  }
  i18n.use(initReactI18next).init({
    resources: { [defaultLang]: resources },
    lng: defaultLang,
    fallbackLng: defaultLang,
    interpolation: {
      escapeValue: false,
      prefix: '{',
      suffix: '}',
    },
  })
}

export default i18n
