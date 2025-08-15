import axios from 'axios';
import type { ProjectDto } from './project.models.ts';

export const fetchProjectList = async (): Promise<ProjectDto[]> => {
  const options = {
    method: 'GET',
    headers: {
      'Content-Type': 'application/json',
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
