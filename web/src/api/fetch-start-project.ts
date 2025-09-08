import { errorHandling, fetchingHeaders } from './fetch-helpers.ts';
import axios from 'axios';

export const fetchStartProject = async (projectId: string): Promise<void> => {
  const options = {
    method: 'GET',
    headers: {
      ...fetchingHeaders,
    },
    url: `/api/projects/${projectId}/start`,
  };
  try {
    await axios.request(options);
  } catch (error: any) {
    return errorHandling(error, `/api/projects/${projectId}/start`);
  }
};
