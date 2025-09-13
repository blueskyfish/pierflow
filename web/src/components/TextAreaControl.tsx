import React from 'react';

export interface TextAreaControlProps {
  name: string;
  value: string;
  label: string;
  placeholder?: string;
  required?: boolean;
  helpText?: string;
  rows?: number;
  onChange: (e: React.ChangeEvent<HTMLTextAreaElement>) => void;
}

export const TextAreaControl: React.FC<TextAreaControlProps> = ({
  name,
  value,
  label,
  placeholder,
  required,
  helpText,
  rows,
  onChange,
}) => {
  return (
    <fieldset className={'fieldset w-full h-full flex flex-col items-stretch'}>
      <legend className={'fieldset-legend flex-shrink-1'}>
        {label}
        {required && <span className={'text-red-600 font-bold'}>*</span>}
      </legend>
      <textarea
        className={`textarea flex-grow-1 w-full h-full {required ? 'textarea-primary' : 'textarea-secondary'}`}
        name={name}
        id={name}
        value={value}
        onChange={onChange}
        placeholder={placeholder}
        required={required}
        rows={rows || 3}
      />
      {helpText && <p className={'label text-wrap flex-shrink-1'}>{helpText}</p>}
    </fieldset>
  );
};
