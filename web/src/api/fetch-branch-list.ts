import axios from 'axios';
import { errorHandling, fetchingHeaders } from './fetch-helpers.ts';

export const fetchBranchList = async (projectId: string, refresh: boolean = false): Promise<void> => {
  const url = refresh ? `/api/projects/${projectId}/branches?refresh=true` : `/api/projects/${projectId}/branches`;
  const options = {
    method: 'GET',
    headers: {
      ...fetchingHeaders,
    },
    url,
  };
  try {
    await axios.request(options);
    return;
  } catch (error: any) {
    return errorHandling(error, `/api/projects/${projectId}/branches`);
  }
};
