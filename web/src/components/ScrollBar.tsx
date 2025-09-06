import * as React from 'react';
import type { PropsWithChildren } from 'react';

// eslint-disable-next-line react-refresh/only-export-components
export enum ScrollingDirection {
  Horizontal = 'overflow-x-auto',
  Vertical = 'overflow-y-auto',
  Both = 'overflow-auto',
}

export interface ScrollingProps {
  direction?: ScrollingDirection;
  className?: string;
}

export const ScrollBar: React.FC<PropsWithChildren<ScrollingProps>> = ({ direction, className, children }) => {
  return <div className={`${direction ?? 'overflow'} ${className ?? ''}`}>{children}</div>;
};
