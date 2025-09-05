import type { BranchDto } from '../fetching';

export enum ProjectCommand {
  CreateProject = 'create-project',
  CheckoutRepository = 'checkout-repository',
  BuildProject = 'build-project',
  CloneRepository = 'clone-repository',
  PullRepository = 'pull-repository',
  StartProject = 'start-project',
  StopProject = 'stop-project',
  DeleteProject = 'delete-project',

  BranchList = 'branch-list',
  TaskList = 'task-list',
}

export interface ProjectDto {
  id: string;
  name: string;
  description: string;
  path: string;
  gitUrl: string;
  branch: string;
  user: string;
  creation: string;
  modified: string;
  status: string;
  commandMap: Record<ProjectCommand, boolean>;

  /**
   * Optional the list of branches for the project repository
   */
  branchList?: BranchDto[];
}
