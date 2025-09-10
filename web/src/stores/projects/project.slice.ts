import { createSlice, type MiddlewareAPI, type PayloadAction } from '@reduxjs/toolkit';
import { type BranchDto, fetchProjectDetails, fetchProjectList, type ProjectDto } from '@blueskyfish/pierflow/api';
import {
  loadProjectDetails,
  loadProjectList,
  ProjectFeatureKey,
  type ProjectState,
} from '@blueskyfish/pierflow/stores';
import {
  reduceUpdateBranchList,
  reduceUpdateProjectBranch,
  reduceUpdateProjectDetail,
  reduceUpdateProjectList,
} from './project.reducing.ts';

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
      const projectList = action.payload;
      return reduceUpdateProjectList(state, projectList);
    },
    updateProjectDetail: (state: ProjectState, action: PayloadAction<ProjectDto>) => {
      return reduceUpdateProjectDetail(state, action.payload);
    },
    updateBranchList: (state: ProjectState, action: PayloadAction<{ projectId: string; branchList: BranchDto[] }>) => {
      const { projectId, branchList } = action.payload;
      return reduceUpdateBranchList(state, projectId, branchList);
    },
    /**
     * Update the current branch of a project
     */
    updateProjectBranch: (state: ProjectState, action: PayloadAction<{ projectId: string; branch: string }>) => {
      const { projectId, branch } = action.payload;
      return reduceUpdateProjectBranch(state, projectId, branch);
    },
  },
});

// Exports the actions
export const { updateSelectedId, updateProjectDetail, updateProjectList, updateBranchList, updateProjectBranch } =
  projectSlice.actions;

// Exports the reducer
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
    if (loadProjectDetails.match(action)) {
      const projectId = action.payload;
      const project = await fetchProjectDetails(projectId);
      dispatch(updateProjectDetail(project));
      return;
    }
    return next(action);
  };
