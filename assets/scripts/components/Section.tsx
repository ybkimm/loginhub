import React, { HTMLAttributes, ReactElement } from 'react'

const SectionWrapper = (props: HTMLAttributes<HTMLDivElement>): ReactElement => {
  const { children, className, ...etcProps } = props
  return (
    <div className={['border rounded mb-6 pt-4', className].join(' ')} {...etcProps}>
      {children}
    </div>
  )
}
export default SectionWrapper

interface SectionLeadingProps extends Omit<HTMLAttributes<HTMLDivElement>, 'children'> {
  title: string
  subtitle?: string
}

export const SectionItem = (props: HTMLAttributes<HTMLDivElement>): ReactElement => {
  const { children, className, ...etcProps } = props
  return (
    <div className={['p-4 pt-8', className].join(' ')} {...etcProps}>
      {children}
    </div>
  )
}

export const SectionLeading = (props: SectionLeadingProps): ReactElement => {
  const { title, subtitle, ...etcProps } = props
  return (
    <SectionItem {...etcProps}>
      <div className="text-xl">{title}</div>
      {subtitle == null ? (<div className="text-sm opacity-50">{subtitle}</div>) : null}
    </SectionItem>
  )
}
