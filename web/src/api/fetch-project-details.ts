import { errorHandling, fetchingHeaders } from './fetch-helpers.ts';
import axios from 'axios';
import type { ProjectDto } from './entities.ts';

export const fetchProjectDetails = async (projectId: string) => {
  const options = {
    method: 'GET',
    headers: {
      ...fetchingHeaders,
    },
    url: `/api/projects/${projectId}`,
  };
  try {
    const { data } = await axios.request<ProjectDto>(options);
    return data;
  } catch (error) {
    return errorHandling(error, `/api/projects/${projectId}`);
  }
};
