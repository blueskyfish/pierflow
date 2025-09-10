import { createAction } from '@reduxjs/toolkit';

export const loadProjectList = createAction('project/loadProjectList');
export const loadProjectDetails = createAction('project/loadProjectDetails', (projectId: string) => ({
  payload: projectId,
}));
