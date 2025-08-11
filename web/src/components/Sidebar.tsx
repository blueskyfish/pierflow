import type { PropsWithChildren } from 'react';
import * as React from 'react';

export const Sidebar: React.FC<PropsWithChildren> = ({ children }) => {
  return <aside className={'bg-base-300 w-1/5 flex-shrink-1'}>{children}</aside>;
};
