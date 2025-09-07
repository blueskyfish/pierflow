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
  user: string;
  creation: string;
  modified: string;
  status: string;
  commandMap: Record<string, boolean>;

  /**
   * Optional the list of branches for the project repository
   */
  branchList?: BranchDto[];
}

export interface CheckoutPayload {
  branch: string;
  place: string;
  message: string;
}

export interface ErrorDto {
  type: string;
  status: number;
  message: string;
}
