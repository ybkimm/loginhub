import { assert } from 'chai'
import I18nString from './strings'

test('I18nString#toString', () => {
  const src = new I18nString({
    en: 'Hello!',
    kr: '안녕하세요!',
    jp: 'こんにちは！'
  })

  assert.equal(src.toString(), 'Hello!')

  I18nString.prefferLang = 'kr'
  assert.equal(src.toString(), '안녕하세요!')

  I18nString.prefferLang = 'jp'
  assert.equal(src.toString(), 'こんにちは！')

  I18nString.prefferLang = 'id'
  assert.equal(src.toString(), 'Hello!')
})
