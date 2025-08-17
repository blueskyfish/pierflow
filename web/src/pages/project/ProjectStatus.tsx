import { ProjectIcon, type ProjectStatusKind } from '@blueskyfish/pierflow/icons';
import * as React from 'react';

export interface ProjectStatusProps {
  status: ProjectStatusKind;
}

export const ProjectStatus: React.FC<ProjectStatusProps> = ({ status }) => {
  return (
    <div className={'flex flex-col items-center justify-center w-full h-full'}>
      <ProjectIcon kind={status} />
      <span className={'text-sm text-base-content'}>{status}</span>
    </div>
  );
};
