export enum ProjectCommand {
  CreateProject = 'create-project',
  CloneRepository = 'clone-repository',
  CheckoutRepository = 'checkout-repository',
  PullRepository = 'pull-repository',
  BuildProject = 'build-project',
  StartProject = 'start-project',
  StopProject = 'stop-project',
  DeleteProject = 'delete-project',

  BranchList = 'branch-list',
  TaskList = 'task-list',
}

export interface BranchDto {
  branch: string;
  place: string;
  active: boolean;
}

export interface ProjectDto {
  id: string;
  name: string;
  description: string;
  path: string;
  gitUrl: string;
  branch: string;
  taskfile: string;
  user: string;
  creation: string;
  modified: string;
  status: string;
  commandMap: Record<string, boolean>;
}

export interface ProjectHistoryDto {
  id: string;
  group: string;
  event: string;
  project: ProjectDto;
  timestamp: number;
}

export interface ProjectHistoryResult {
  current: ProjectHistoryDto;
  previous: ProjectHistoryDto[];
}

export interface CheckoutPayload {
  branch: string;
  place: string;
  message: string; // TODO remove later
}

/**
 * @deprecated
 */
export interface BuildPayload {
  taskfile: string;
  message: string; // TODO remove later
}

export interface ErrorDto {
  type: string;
  status: number;
  message: string;
}

export interface CreateProjectPayload {
  name: string;
  description: string;
  path: string;
  giturl: string;
  user: string;
  token: string;
}
