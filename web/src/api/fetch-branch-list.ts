import axios from 'axios';
import { fetchingHeaders } from './fetch-user.ts';

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
  } catch (e) {
    console.error('> Failed to load branch list', e);
    return;
  }
};
