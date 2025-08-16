import type { PropsWithChildren } from 'react';
import * as React from 'react';

export type ParagraphSize = 'xs' | 'sm' | 'base' | 'md' | 'lg' | 'xl' | '2xl';

export interface ParagraphProps {
  size?: ParagraphSize;
  className?: string;
}

export const Paragraph: React.FC<PropsWithChildren<ParagraphProps>> = ({ size = 'base', className, children }) => {
  const sizeClass = {
    xs: 'text-xs',
    sm: 'text-sm',
    md: 'text-base',
    base: 'text-base',
    lg: 'text-lg',
    xl: 'text-xl',
    '2xl': 'text-2xl',
  }[size];

  return <p className={`${sizeClass} ${className ?? 'mb-1'}`}>{children}</p>;
};
