import { errorHandling, fetchingHeaders } from './fetch-helpers.ts';
import axios from 'axios';

/**
 * Fetch to get the list of taskfiles for the given project
 * @param projectId - the project id
 * @returns a promise that resolves to an array of taskfile names
 */
export const fetchGetTaskFileList = async (projectId: string): Promise<string[]> => {
  const options = {
    method: 'GET',
    headers: {
      ...fetchingHeaders,
    },
    url: `/api/projects/${projectId}/tasks`,
  };
  try {
    const { data } = await axios.request(options);
    return data as string[];
  } catch (error: any) {
    return errorHandling(error, `/api/projects/${projectId}/tasks`);
  }
};
