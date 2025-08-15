import { createSlice } from '@reduxjs/toolkit';

export const LayoutFeatureKey = 'layout';

export interface LayoutState {
  pageKey: string;
}

export type LayoutFeatureState = {
  [LayoutFeatureKey]: LayoutState;
};

export const layoutSlice = createSlice({
  name: LayoutFeatureKey,
  initialState: {
    pageKey: '',
  } as LayoutState,
  reducers: {
    updatePageKey: (state: LayoutState, action) => {
      return {
        ...state,
        pageKey: action.payload,
      };
    },
  },
});

export const { updatePageKey } = layoutSlice.actions;
export const layoutReducer = layoutSlice.reducer;
