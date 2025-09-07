import type { BuildPayload } from './entities.ts';
import { errorHandling, fetchingHeaders } from './fetch-helpers.ts';
import axios from 'axios';

export const fetchBuildProject = async (projectId: string, payload: BuildPayload): Promise<void> => {
  const options = {
    method: 'PUT',
    headers: {
      ...fetchingHeaders,
    },
    url: `/api/projects/${projectId}/build`,
    data: payload,
  };
  try {
    await axios.request(options);
  } catch (error: any) {
    return errorHandling(error, `/api/projects/${projectId}/build`);
  }
};
