import { ProjectCommand, selectSelectProject, useAppSelector } from '@blueskyfish/pierflow/stores';
import { ProjectPath } from '@blueskyfish/pierflow/utils';
import * as React from 'react';
import { ProjectAllow, ProjectStop } from './project';

export const ProjectStopPage: React.FC = () => {
  const project = useAppSelector(selectSelectProject);
  return (
    <ProjectAllow command={ProjectCommand.StopProject} project={project} projectKey={ProjectPath.Stop}>
      <ProjectStop project={project!} />
    </ProjectAllow>
  );
};
