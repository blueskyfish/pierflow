import type { PropsWithChildren } from 'react';
import * as React from 'react';
import { Provider } from 'react-redux';
import { store } from './stores.ts';

export const StoreProvider: React.FC<PropsWithChildren> = ({ children }) => {
  return <Provider store={store}>{children}</Provider>;
};
