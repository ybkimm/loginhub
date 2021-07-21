import React, { DetailedHTMLProps, InputHTMLAttributes, ReactElement } from 'react'
import { buttonClasses } from './Button'

interface FormInputProps extends DetailedHTMLProps<InputHTMLAttributes<HTMLInputElement>, HTMLInputElement> {
  label: string
}

export const FormInput = ({ label, className, ...props}: FormInputProps): ReactElement => {
  return (
		<div className={className}>
			<label className="block px-4 py-2 mb-4 border rounded focus-within:ring-2 ring-primary text-black cursor-text">
				<div className="group relative">
					<input className="block w-full border-0 p-0 pt-6 focus:outline-none focus:ring-0" placeholder=" " {...props} />
					<div className="block absolute top-0 left-0 w-full h-6 text-left text-xs leading-6 transform-gpu origin-left translate-y-3 group-focus-within:translate-y-0 input-not-empty:translate-y-0 scale-125 group-focus-within:scale-100 input-not-empty:scale-100 opacity-50 group-focus-within:opacity-100 transition-all pointer-events-none">
						{label}
					</div>
				</div>
			</label>
		</div>
  )
}

export const FormSubmit = (props: FormInputProps): ReactElement => {
  const { label, className, ...etcProps } = props
  return (
		<input
			className={[buttonClasses, 'block w-full', className].join(' ')}
			type="submit"
			{...etcProps}
			value={label}
		/>
  )
}
