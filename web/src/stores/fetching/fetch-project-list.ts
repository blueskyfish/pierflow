import axios from 'axios';
import { fetchingHeaders } from './fetch-user';
import type { BranchDto } from './fetch-branch-list.ts';

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

export const fetchProjectList = async (): Promise<ProjectDto[]> => {
  const options = {
    method: 'GET',
    headers: {
      ...fetchingHeaders,
    },
    url: '/api/projects',
  };
  try {
    const { data } = await axios.request(options);
    return data as ProjectDto[];
  } catch (e) {
    console.error('> Failed to load project list', e);
    return [];
  }
};
