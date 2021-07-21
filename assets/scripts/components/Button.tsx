import React, { DetailedHTMLProps, ButtonHTMLAttributes, ReactElement } from 'react'

export const buttonClasses = 'block w-full bg-white hover:bg-primary-dark hover:bg-opacity-10 border rounded py-2 cursor-pointer'

export interface ButtonProps extends DetailedHTMLProps<ButtonHTMLAttributes<HTMLButtonElement>, HTMLButtonElement> { }

export const Button = ({ className, children, ...props }: ButtonProps): ReactElement => {
	return (
		<button className={[buttonClasses, className].join(' ')} {...props}>
			{children}
		</button>
	)
}
