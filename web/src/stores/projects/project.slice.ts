import {
  type BranchDto,
  fetchProjectDetails,
  fetchProjectHistory,
  fetchProjectList,
  type ProjectDto,
  type ProjectHistoryResult,
} from '@blueskyfish/pierflow/api';
import {
  loadProjectDetails,
  loadProjectHistory,
  loadProjectList,
  ProjectFeatureKey,
  type ProjectState,
} from '@blueskyfish/pierflow/stores';
import { createSlice, type MiddlewareAPI, type PayloadAction } from '@reduxjs/toolkit';
import {
  reduceUpdateBranchList,
  reduceUpdateProjectBranch,
  reduceUpdateProjectDetail,
  reduceUpdateProjectHistory,
  reduceUpdateProjectList,
} from './project.reducing.ts';

export const projectSlice = createSlice({
  name: ProjectFeatureKey,
  initialState: {
    map: {},
    selectedId: null,
    historyMap: {},
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

    updateProjectHistory: (
      state: ProjectState,
      action: PayloadAction<{ projectId: string; history: ProjectHistoryResult }>,
    ) => {
      const { projectId, history } = action.payload;
      return reduceUpdateProjectHistory(state, projectId, history);
    },
  },
});

// Exports the actions
export const {
  updateSelectedId,
  updateProjectDetail,
  updateProjectList,
  updateBranchList,
  updateProjectBranch,
  updateProjectHistory,
} = projectSlice.actions;

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
    if (loadProjectHistory.match(action)) {
      const projectId = action.payload;
      const projectHistory = await fetchProjectHistory(projectId);
      dispatch(updateProjectHistory({ projectId, history: projectHistory }));
      return;
    }
    return next(action);
  };
