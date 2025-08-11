import type { PropsWithChildren } from 'react';
import * as React from 'react';

export const RootPage: React.FC<PropsWithChildren> = ({ children }) => {
  return <div className={'flex flex-row align-items-stretch h-screen w-screen bg-base-100'}>{children}</div>;
};
