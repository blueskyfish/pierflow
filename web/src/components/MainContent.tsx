import type { PropsWithChildren } from 'react';
import * as React from 'react';

export const MainContent: React.FC<PropsWithChildren> = ({ children }) => {
  return <main className={'bg-base-100 w-4/5 flex-grow-1'}>{children}</main>;
};
