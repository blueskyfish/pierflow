import { createSelector } from '@reduxjs/toolkit';
import { ProjectFeatureKey, type ProjectState } from './project.state.ts';

export const selectProjectState = (state: { [ProjectFeatureKey]: ProjectState }) => state[ProjectFeatureKey];

export const selectProjectList = createSelector([selectProjectState], (projectState: ProjectState) =>
  Object.values(projectState.map),
);

export const selectSelectProject = createSelector([selectProjectState], (projectState: ProjectState) => {
  if (projectState.selectedId) {
    return projectState.map[projectState.selectedId] || null;
  }
  return null;
});
