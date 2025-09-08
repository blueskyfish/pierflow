import type { ProjectDto } from './entities.ts';
import { errorHandling, fetchingHeaders } from './fetch-helpers.ts';
import axios from 'axios';

export const fetchChangeTaskfile = async (projectId: string, taskfile: string): Promise<ProjectDto> => {
  const options = {
    method: 'PUT',
    headers: {
      ...fetchingHeaders,
    },
    url: `/api/projects/${projectId}/tasks`,
    data: { taskfile },
  };
  try {
    const { data } = await axios.request(options);
    return data as ProjectDto;
  } catch (error) {
    return errorHandling(error, `/api/projects/${projectId}/tasks`);
  }
};
