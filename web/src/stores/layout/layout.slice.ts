import { ProjectPath } from '@blueskyfish/pierflow/utils';
import { createSlice, type PayloadAction } from '@reduxjs/toolkit';

export const LayoutFeatureKey = 'layout';

export interface LayoutState {
  pageKey: string;
  projectKey: string; // This can be used to track the current project page key if needed
}

export type LayoutFeatureState = {
  [LayoutFeatureKey]: LayoutState;
};

export const layoutSlice = createSlice({
  name: LayoutFeatureKey,
  initialState: {
    pageKey: '',
    projectKey: '',
  } as LayoutState,
  reducers: {
    updatePageKey: (state: LayoutState, action: PayloadAction<string | null>) => {
      return {
        ...state,
        pageKey: action.payload ?? '',
      };
    },
    updateProjectKey: (state: LayoutState, action: PayloadAction<ProjectPath>) => {
      return {
        ...state,
        projectKey: action.payload,
      };
    },
  },
});

export const { updatePageKey, updateProjectKey } = layoutSlice.actions;
export const layoutReducer = layoutSlice.reducer;
