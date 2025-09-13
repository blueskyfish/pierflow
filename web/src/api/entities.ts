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

  /**
   * Optional the list of branches for the project repository
   */
  branchList?: BranchDto[];

  /**
   * Optional the list of tasks for the project from taskfiles
   */
  taskfileList?: string[];
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
