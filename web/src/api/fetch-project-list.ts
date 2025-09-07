import axios from 'axios';
import { errorHandling, fetchingHeaders } from './fetch-helpers.ts';
import type { ProjectDto } from './entities.ts';

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
  } catch (error: any) {
    return errorHandling(error, '/api/projects');
  }
};
