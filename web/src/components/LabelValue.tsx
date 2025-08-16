import type { PropsWithChildren } from 'react';
import * as React from 'react';

export interface LabelValueProps {
  label: string;
  className?: string;
}

export const LabelValue: React.FC<PropsWithChildren<LabelValueProps>> = ({ label, className, children }) => {
  return (
    <div className={`${className ?? ''} flex flex-row align-items-center`}>
      <div className={'font-bold text-base-content/60 w-1/4 py-1 text-truncate'}>{label}</div>
      <div className={'text-base-content w-3/4 py-1'}>{children}</div>
    </div>
  );
};
