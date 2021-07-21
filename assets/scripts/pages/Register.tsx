import React, { ReactElement, FormEvent, useCallback, useEffect, useRef } from 'react'
import { SectionLeading, SectionItem } from '../components/Section'
import { FormInput, FormSubmit } from '../components/Form'
import { Link, Switch, Route, useRouteMatch, useLocation, useHistory } from 'react-router-dom'
import { TransitionGroup, CSSTransition } from 'react-transition-group'
import { Button } from '../components/Button'

interface PageRequiredFormProps {
	baseURL: string
}

const PageRequiredForm = ({ baseURL }: PageRequiredFormProps): ReactElement => {
	const history = useHistory()

	const handleRequiredFormSubmit = useCallback((e: FormEvent<HTMLFormElement>) => {
		history.push(`${baseURL}/email_verification`)
		e.preventDefault()
	}, [])

	return (
		<div className="flex-none w-md">
			<div className="text-center">
				<div className="pt-6">[BRAND LOGO HERE]</div>
				<SectionLeading title="계정 만들기" className="-mt-6" />
			</div>
			<SectionItem>
				<form className="px-4" onSubmit={handleRequiredFormSubmit}>
					<div className="py-4">
						<FormInput type="text" name="name" label="이름" />
						<FormInput type="text" name="display_name" label="닉네임" />
						<FormInput type="email" name="email" label="이메일" />
						<div className="flex gap-4">
							<FormInput className="flex-1" type="password" name="password" label="비밀번호" />
							<FormInput className="flex-1" type="password" name="password_confirm" label="확인" />
						</div>
					</div>
					<div className="flex flex-1 mt-12 gap-4 items-center">
            <div className="block w-full text-center">
              <Link to="/login" className="inline-block px-2 py-1 hover:underline opacity-50">대신 로그인하기</Link>
            </div>
            <FormSubmit label="다음" />
          </div>
        </form>
      </SectionItem>
    </div>
	)
}

interface PageEmailVerificationProps {
	baseURL: string
}

const PageEmailVerification = ({ baseURL }: PageEmailVerificationProps): ReactElement => {
	const history = useHistory()

	useEffect(() => {
		// TODO: Use API
		const timeout = setTimeout(() => { history.push(`${baseURL}/optional_info`) }, 2000)

		return () => {
			clearTimeout(timeout)
		}
	})

	return (
		<div className="flex-none w-md">
			<div className="text-center">
				<SectionLeading title="메일을 확인해주세요" />
			</div>
			<SectionItem>
				<p className="pb-2">
					admin@sekiramen.studio로 인증 메일을 보냈어요.
				</p>
				<p>
					30분 안에 확인하지 않으면 메일이 만료되니 주의하세요!
				</p>
				<div className="flex py-12 justify-center items-center">
					<svg className="animate-spin w-5 h-5 mr-2" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24">
						<circle
							className="opacity-25"
							cx="12" cy="12" r="10"
							stroke="currentColor"
							strokeWidth="4"
						/>
						<path
							className="opacity-75"
							fill="currentColor"
							d="M4 12a8 8 0 018-8V0A12 12 0 000 12h4zm2 5.3A7.96 7.96 0 014 12H0c0 3.04 1.14 5.82 3 7.94l3-2.65z"
						/>
					</svg>
					<p>메일이 확인되길 기다리는 중...</p>
				</div>
				<p><a>혹시 메일이 도착하지 않았나요?</a></p>
			</SectionItem>
		</div>
	)
}

interface PageOptionalFormProps {
	baseURL: string
}

const PageOptionalForm = ({ baseURL }: PageOptionalFormProps): ReactElement => {
	const history = useHistory()

	const handleOptionalFormSubmit = useCallback((e: FormEvent<HTMLFormElement>) => {
		// TODO: Use API
		history.push(`${baseURL}/welcome`)
		e.preventDefault()
	}, [])

	return (
		<div className="flex-none w-md">
			<div className="text-center">
				<SectionLeading title="추가 정보를 입력해주세요" />
			</div>
			<SectionItem>
				<p className="pb-2">서비스에 필요한 추가적인 정보를 입력해주세요.</p>
				<p className="pb-8">모든 정보를 다 입력하실 필요는 없습니다. 원하는 만큼만 입력해주세요!</p>
				<form className="px-4" onSubmit={handleOptionalFormSubmit}>
					<div className="py-4">
						<FormInput type="text" name="name" label="생일" />
					</div>
					<div className="block w-1/2 pl-2 mt-12 ml-auto">
            <FormSubmit label="다음" />
					</div>
        </form>
			</SectionItem>
		</div>
	)
}

const PageWelcome = (): ReactElement => {
	return (
		<div className="flex-none w-md">
			<div className="text-center">
				<SectionLeading title="어서오세요!" />
			</div>
			<SectionItem>
				<p>가입해주셔서 감사합니다. 이제 아래 버튼을 눌러 시작해보세요!</p>
				<div className="w-1/2 pl-2 ml-auto mt-24">
					<Button>시작하기</Button>
				</div>
			</SectionItem>
		</div>
	)
}
const PageRegister = (): ReactElement => {
	const location = useLocation()
	const { path, url } = useRouteMatch()
	const innerDiv = useRef<HTMLDivElement>()
	const outerDiv = useRef<HTMLDivElement>()

	const handleTransitionDone = useCallback(() => {
		setTimeout(() => {
			outerDiv.current.style.height = `${innerDiv.current.scrollHeight}px`
		}, 1)
	}, [])

	useEffect(() => {
		handleTransitionDone()
	}, [])

  return (
    <div className="flex-1 py-8">
			<div className="w-full max-w-md mx-auto rounded border overflow-hidden transition-all" ref={outerDiv}>
				<div ref={innerDiv}>
					<TransitionGroup className="flex overflow-hidden pt-6 pb-3">
						<CSSTransition
							key={location.key}
							timeout={400}
							classNames="page-transition-register"
							onExited={handleTransitionDone}
						>
							<Switch location={location}>
								<Route exact path={`${path}`}>
									<PageRequiredForm baseURL={url} />
								</Route>
								<Route exact path={`${path}/email_verification`}>
									<PageEmailVerification baseURL={url} />
								</Route>
								<Route exact path={`${path}/optional_info`}>
									<PageOptionalForm baseURL={url} />
								</Route>
								<Route exact path={`${path}/welcome`}>
									<PageWelcome />
								</Route>
							</Switch>
						</CSSTransition>
					</TransitionGroup>
				</div>
			</div>
    </div>
  )
}
export default PageRegister
