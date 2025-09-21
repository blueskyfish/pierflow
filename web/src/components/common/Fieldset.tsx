import React, { type PropsWithChildren, useMemo } from 'react';

export interface FieldsetProps {
  title: string;
  titleAlign?: 'left' | 'center' | 'right';
  border?: boolean;
}

export const Fieldset: React.FC<PropsWithChildren<FieldsetProps>> = ({ title, titleAlign, border, children }) => {
  const alignClass = useMemo(() => {
    switch (titleAlign) {
      case 'center':
        return 'text-center w-full';
      case 'right':
        return 'text-right w-full';
      case 'left':
      default:
        return 'text-left';
    }
  }, [titleAlign]);
  return (
    <fieldset
      className={`h-full w-full flex-grow-1 rounded border ${border ? 'border border-gray-300 bg-white rounded' : 'border-none'}`}
    >
      <legend className={`p-2 font-semibold text-base-content ${alignClass} ${border ? 'bg-white' : ''}`}>
        {title}
      </legend>
      {children}
    </fieldset>
  );
};
