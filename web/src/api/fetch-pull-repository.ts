import { errorHandling, fetchingHeaders } from './fetch-helpers.ts';
import axios from 'axios';

export const fetchPullRepository = async (projectId: string): Promise<void> => {
  const options = {
    url: `/api/projects/${projectId}/branches/pull`,
    method: 'GET',
    headers: {
      // Assuming fetchingHeaders is defined elsewhere and imported here
      ...fetchingHeaders,
    },
  };
  try {
    await axios.request(options);
  } catch (error: any) {
    return errorHandling(error, `/api/projects/${projectId}/branches/pull`);
  }
};
