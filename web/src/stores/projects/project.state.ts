import { type ProjectDto } from '@blueskyfish/pierflow/api';

export const ProjectFeatureKey = 'projects';

export interface ProjectState {
  map: Record<string, ProjectDto>;
  selectedId: string | null;
}
