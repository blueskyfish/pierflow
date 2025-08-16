export enum ProjectCommand {
  CreateProject = 'create-project',
  CheckoutRepository = 'checkout-repository',
  BuildProject = 'build-project',
  CloneRepository = 'clone-repository',
  PullRepository = 'pull-repository',
  StartProject = 'start-project',
  StopProject = 'stop-project',
  DeleteProject = 'delete-project',
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
}
