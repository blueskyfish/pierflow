import type { BranchDto, ProjectDto } from '@blueskyfish/pierflow/api';

export interface ProjectInfo extends ProjectDto {
  /**
   * Optional the list of branches for the project repository
   */
  branchList?: BranchDto[];

  /**
   * Optional the list of tasks for the project from taskfiles
   */
  taskfileList?: string[];
}
