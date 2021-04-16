import i18n from 'i18next';
import { initReactI18next } from 'react-i18next';

import { resources } from './locales/resources';

i18n
  .use(initReactI18next)
  .init({
    resources,
    fallbackLng: 'en_US',  // language to use if translations in user language are not available.
    lng: 'zh_Hans',
    debug: true,
    interpolation: {
      escapeValue: false, // not needed for react as it escapes by default
    }
  });

export default i18n;