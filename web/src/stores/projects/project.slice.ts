import { createSlice, type MiddlewareAPI, type PayloadAction } from '@reduxjs/toolkit';
import { type BranchDto, fetchProjectList, type ProjectDto } from '@blueskyfish/pierflow/api';
import { loadProjectList, ProjectFeatureKey, type ProjectState } from '@blueskyfish/pierflow/stores';
import { reduceUpdateBranchList, reduceUpdateProjectBranch, reduceUpdateProjectList } from './project.reducing.ts';

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

export const { updateSelectedId, updateProjectList, updateBranchList, updateProjectBranch } = projectSlice.actions;
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
