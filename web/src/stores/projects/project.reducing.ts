import type { BranchDto, ProjectDto, ProjectHistoryResult } from '@blueskyfish/pierflow/api';
import type { ProjectState } from './project.state';

export const reduceUpdateProjectList = (state: ProjectState, projectList: ProjectDto[]): ProjectState => {
  const map: Record<string, ProjectDto> = {};
  projectList.forEach((project) => {
    map[project.id] = project;
  });
  return {
    ...state,
    selectedId: state.selectedId ?? null,
    map,
  };
};

export const reduceUpdateProjectDetail = (state: ProjectState, project: ProjectDto): ProjectState => {
  return {
    ...state,
    map: {
      ...state.map,
      [project.id]: {
        ...state.map[project.id],
        ...project,
      },
    },
  };
};

export const reduceUpdateBranchList = (
  state: ProjectState,
  projectId: string,
  branchList: BranchDto[],
): ProjectState => {
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
};

export const reduceUpdateProjectBranch = (state: ProjectState, projectId: string, branch: string): ProjectState => {
  let project = state.map[projectId];
  if (project) {
    project = {
      ...project,
      branch,
      branchList: project.branchList
        ? project.branchList.map((b) => ({ ...b, active: b.branch === branch }))
        : undefined,
    };
  }
  return {
    ...state,
    map: {
      ...state.map,
      [projectId]: project,
    },
  };
};

export const reduceUpdateProjectHistory = (
  state: ProjectState,
  projectId: string,
  history: ProjectHistoryResult,
): ProjectState => {
  return {
    ...state,
    historyMap: {
      ...state.historyMap,
      [projectId]: history,
    },
  };
};
