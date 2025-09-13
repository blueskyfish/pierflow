import type { CreateProjectPayload, ProjectDto } from './entities.ts';
import { errorHandling, fetchingHeaders } from './fetch-helpers.ts';
import axios from 'axios';

export const fetchCreateProject = async (data: CreateProjectPayload): Promise<ProjectDto> => {
  const options = {
    method: 'POST',
    headers: {
      ...fetchingHeaders,
    },
    url: `/api/projects`,
    data,
  }
  try {
    const { data } = await axios.request<ProjectDto>(options);
    return data as ProjectDto;
  } catch (error: any) {
    return errorHandling(error, `/api/projects`);
  }
}
