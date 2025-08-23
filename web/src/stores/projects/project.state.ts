import { createSlice, type MiddlewareAPI, type PayloadAction } from '@reduxjs/toolkit';
import { loadProjectList } from './project.actions.ts';
import { fetchProjectList } from './project.backend.ts';
import type { ProjectDto } from './project.models.ts';

export const ProjectFeatureKey = 'projects';

export interface ProjectState {
  userId: string | null;
  map: Record<string, ProjectDto>;
  selectedId: string | null;
}

export const projectSlice = createSlice({
  name: ProjectFeatureKey,
  initialState: {
    userId: null,
    map: {},
    selectedId: null,
  } as ProjectState,
  reducers: {
    updateUserId: (state: ProjectState, action: PayloadAction<string | null>) => {
      localStorage.setItem('blueskyfish.pierflow.userId', action.payload ?? '');
      return {
        ...state,
        userId: action.payload,
      };
    },
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
  },
});

export const { updateUserId, updateSelectedId, updateProjectList } = projectSlice.actions;
export const projectReducer = projectSlice.reducer;

/**
 * Middleware to handle project-related actions.
 */
export const projectMiddleware =
  ({ dispatch }: MiddlewareAPI) =>
  (next: (action: unknown) => unknown) =>
  async (action: unknown) => {
    if (loadProjectList.match(action)) {
      // read or perhaps update the user id from local storage
      let userId = localStorage.getItem('blueskyfish.pierflow.userId');
      if (!userId) {
        userId = crypto.randomUUID() as string;
      }
      dispatch(updateUserId(userId));

      const list = await fetchProjectList(userId);
      dispatch(updateProjectList(list));
      return;
    }
    return next(action);
  };
