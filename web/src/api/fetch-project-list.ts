import axios from 'axios';
import { fetchingHeaders } from './fetch-user.ts';
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
  } catch (e) {
    console.error('> Failed to load project list', e);
    return [];
  }
};
