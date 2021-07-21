import { ReactNode } from 'react'

const FAILBACK_LANG = 'en'

class I18nString {
  private strings: {[lang: string]: string}
  public static prefferLang: string

  constructor (strings: {[lang: string]: string}) {
    this.strings = strings
  }

  public toString (): string | undefined {
    const prefferLang = I18nString.prefferLang
    if (prefferLang in this.strings) {
      return this.strings[I18nString.prefferLang]
    }
    if (FAILBACK_LANG in this.strings) {
      return this.strings[FAILBACK_LANG]
    }
    return undefined
  }

  // Render implements ReactNode.
  public render (): ReactNode {
    return this.toString()
  }
}
export default I18nString
