import { configureStore } from '@reduxjs/toolkit';
import { useDispatch, useSelector } from 'react-redux';
import { LayoutFeatureKey, layoutReducer } from './layout';
import { pageTitleMiddleware } from './page-title.middleware.ts';
import { ProjectFeatureKey, projectMiddleware, projectReducer } from './projects';

export const store = configureStore({
  reducer: {
    [ProjectFeatureKey]: projectReducer,
    [LayoutFeatureKey]: layoutReducer,
  },
  middleware: (getDefaultMiddleware) => {
    return getDefaultMiddleware().concat(pageTitleMiddleware, projectMiddleware);
  },
});

export type RootState = ReturnType<typeof store.getState>;
export type AppDispatch = typeof store.dispatch;

export const useAppDispatch = useDispatch.withTypes<AppDispatch>();
export const useAppSelector = useSelector.withTypes<RootState>();
