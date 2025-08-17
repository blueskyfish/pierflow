import type { PropsWithChildren } from 'react';
import * as React from 'react';

export interface LabelValueProps {
  label: string;
  className?: string;
  size?: 'xs' | 'sm' | 'md' | 'lg';
}

export const LabelValue: React.FC<PropsWithChildren<LabelValueProps>> = ({ label, className, size, children }) => {
  return (
    <div className={`${className ?? ''} flex flex-row align-items-center`}>
      <div className={`font-bold text-${size ?? 'md'} text-base-content/60 w-1/4 py-0.5 text-truncate`}>{label}</div>
      <div className={`text-${size ?? 'md'} text-base-content w-3/4 py-0.5`}>{children}</div>
    </div>
  );
};
