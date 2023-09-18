

import { useI18n } from 'vue-i18n'
type I18nGlobalTranslation = ['t'];
/**
 * 国际化
 * @param data 国际化语句
 * @returns 
 */
export const uselocals = (data: string): { t: I18nGlobalTranslation } => {
    const { t } = useI18n()
    return t(data)
}