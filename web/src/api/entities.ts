export interface BranchDto {
  branch: string;
  place: number;
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
