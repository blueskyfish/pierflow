import { createSlice, type MiddlewareAPI, type PayloadAction } from '@reduxjs/toolkit';
import { loadProjectList } from './project.actions.ts';
import { type ProjectDto } from './project.models.ts';
import type { BranchDto } from '../fetching';
import { fetchProjectList } from '../fetching';

export const ProjectFeatureKey = 'projects';

export interface ProjectState {
  map: Record<string, ProjectDto>;
  selectedId: string | null;
}

export const projectSlice = createSlice({
  name: ProjectFeatureKey,
  initialState: {
    map: {},
    selectedId: null,
  } as ProjectState,
  reducers: {
    updateSelectedId: (state: ProjectState, action: PayloadAction<string | null>) => {
      return {
        ...state,
        selectedId: action.payload,
      };
    },
    updateProjectList: (state: ProjectState, action: PayloadAction<ProjectDto[]>) => {
      const projects = action.payload;
      const map: Record<string, ProjectDto> = {};
      projects.forEach((project) => {
        map[project.id] = project;
      });
      return {
        ...state,
        selectedId: state.selectedId ?? null,
        map,
      };
    },
    updateBranchList: (state: ProjectState, action: PayloadAction<{ projectId: string; branchList: BranchDto[] }>) => {
      const { projectId, branchList } = action.payload;
      let project = state.map[projectId];
      if (project) {
        project = {
          ...project,
          branchList,
        };
      }
      return {
        ...state,
        map: {
          ...state.map,
          [projectId]: project,
        },
      };
    },
  },
});

export const { updateSelectedId, updateProjectList, updateBranchList } = projectSlice.actions;
export const projectReducer = projectSlice.reducer;

/**
 * Middleware to handle project-related actions.
 */
export const projectMiddleware =
  ({ dispatch }: MiddlewareAPI) =>
  (next: (action: unknown) => unknown) =>
  async (action: unknown) => {
    if (loadProjectList.match(action)) {
      const list = await fetchProjectList();
      dispatch(updateProjectList(list));
      return;
    }
    return next(action);
  };
