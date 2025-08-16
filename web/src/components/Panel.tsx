import type { PropsWithChildren } from 'react';
import * as React from 'react';

export const Panel: React.FC<PropsWithChildren> = ({ children }) => {
  return <section className={'card bg-base-200 card-border border-base-300 p-3'}>{children}</section>;
};
