import React, { type ChangeEvent } from 'react';

export interface InputControlProps {
  type: 'text' | 'password' | 'email';
  name: string;
  value: string;
  label: string;
  placeholder?: string;
  required?: boolean;
  helpText?: string;
  onChange: (e: ChangeEvent<HTMLInputElement>) => void;
}

export const InputControl: React.FC<InputControlProps> = ({
  type,
  name,
  value,
  label,
  placeholder,
  required,
  helpText,
  onChange,
}) => {
  return (
    <fieldset className={'fieldset w-full'}>
      <legend className={'fieldset-legend'}>
        {label}
        {required && <span className={'text-red-600 font-bold'}>*</span>}
      </legend>
      <input
        type={type}
        name={name}
        id={name}
        value={value}
        className={`input w-full ${required ? 'input-primary' : 'input-secondary'}`}
        onChange={onChange}
        placeholder={placeholder}
        required={required}
      />
      {helpText && <p className={'label text-wrap'}>{helpText}</p>}
    </fieldset>
  );
};
