import { createSelector } from '@reduxjs/toolkit';
import type { LayoutFeatureState, LayoutState } from './layout.slice.ts';

export const selectLayoutState = (state: LayoutFeatureState) => state.layout;

export const selectPageKey = createSelector([selectLayoutState], (layoutState: LayoutState) => layoutState.pageKey);
export const selectProjectKey = createSelector(
  [selectLayoutState],
  (layoutState: LayoutState) => layoutState.projectKey,
);
