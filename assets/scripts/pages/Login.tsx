import React, { ReactElement } from 'react'
import { SectionLeading, SectionItem } from '../components/Section'
import { FormInput, FormSubmit } from '../components/Form'
import { Link } from 'react-router-dom'

const PageLogin = (): ReactElement => {
  return (
    <div className="flex flex-1 flex-col py-8 items-center justify-center">
      <div className="w-full max-w-md mx-2 mb-6 rounded border pt-6 pb-2 text-center">
        <SectionLeading title="PLATFORM 계정으로 로그인" />
        <SectionItem>
          <form className="px-4">
            <div className="py-4">
              <FormInput type="email" name="email" label="이메일" />
              <FormInput type="password" name="password" label="비밀번호" />
            </div>
            <div className="flex flex-1 items-center">
              <div className="block w-full">
                <Link to="/register" className="inline-block px-2 py-1">계정 만들기</Link>
              </div>
              <FormSubmit label="로그인" />
            </div>
          </form>
        </SectionItem>
      </div>
    </div>
  )
}
export default PageLogin
