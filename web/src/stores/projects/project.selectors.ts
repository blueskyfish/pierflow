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

export const selectProjectProcessing = createSelector([selectProjectState], (projectState: ProjectState) => {
  if (projectState.selectedId) {
    const history = projectState.historyMap[projectState.selectedId];
    if (history) {
      return {
        loading: history.previous.length > 0,
        current: history.current,
        previous: history.previous[0] || undefined,
      };
    }
    const project = projectState.map[projectState.selectedId];
    if (project) {
      return {
        loading: false,
        current: {
          id: '',
          group: 'project',
          event: '',
          project: project,
          timestamp: Date.now(),
        },
        previous: undefined,
      };
    }
  }
  return {
    loading: false,
    current: undefined,
    previous: undefined,
  };
});
