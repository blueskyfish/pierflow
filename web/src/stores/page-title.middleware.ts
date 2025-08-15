import type { Action, Middleware, MiddlewareAPI } from '@reduxjs/toolkit';

export const pageTitleMiddleware: Middleware =
  // eslint-disable-next-line @typescript-eslint/no-unused-vars
  (_: MiddlewareAPI) => (next: (action: unknown) => unknown) => (action: unknown) => {
    console.log('> Dispatching %s', (action as Action).type);
    return next(action);
  };
