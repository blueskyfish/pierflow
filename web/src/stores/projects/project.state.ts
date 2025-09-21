import type { ProjectHistoryResult } from '@blueskyfish/pierflow/api';
import type { ProjectInfo } from './project.models.ts';

export const ProjectFeatureKey = 'projects';

export interface ProjectState {
  map: Record<string, ProjectInfo>;
  selectedId: string | null;
  historyMap: Record<string, ProjectHistoryResult>;
}
