import axios from 'axios';
import type { ProjectDto } from './project.models.ts';

export const fetchProjectList = async (userId: string): Promise<ProjectDto[]> => {
  const options = {
    method: 'GET',
    headers: {
      'Content-Type': 'application/json',
      'x-pierflow-user': userId,
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
