import axios from 'axios';
import type { ErrorDto, ProjectHistoryDto, ProjectHistoryResult } from './entities.ts';
import { errorHandling, fetchingHeaders } from './fetch-helpers.ts';

export const fetchProjectHistory = async (projectId: string): Promise<ProjectHistoryResult> => {
  const url = `/api/projects/${projectId}/history`;
  const options = {
    method: 'GET',
    headers: {
      ...fetchingHeaders,
    },
    url,
  };
  try {
    const { data } = await axios<ProjectHistoryDto[]>(options);
    const [current, ...previous] = data;
    if (!current) {
      return Promise.reject({
        status: 404,
        type: url,
        message: 'No current project data found',
      } as ErrorDto);
    }
    if (!previous) {
      return {
        current,
        previous: [],
      };
    }
    return {
      current,
      previous,
    };
  } catch (error: any) {
    return errorHandling(error, url);
  }
};
