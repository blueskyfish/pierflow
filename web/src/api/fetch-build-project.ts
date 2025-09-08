import { errorHandling, fetchingHeaders } from './fetch-helpers.ts';
import axios from 'axios';

export const fetchBuildProject = async (projectId: string): Promise<void> => {
  const options = {
    method: 'GET',
    headers: {
      ...fetchingHeaders,
    },
    url: `/api/projects/${projectId}/build`,
  };
  try {
    await axios.request(options);
  } catch (error: any) {
    return errorHandling(error, `/api/projects/${projectId}/build`);
  }
};
