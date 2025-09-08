import { ProjectCommand, selectSelectProject, useAppSelector } from '@blueskyfish/pierflow/stores';
import { ProjectPath } from '@blueskyfish/pierflow/utils';
import * as React from 'react';
import { ProjectAllow, ProjectStart } from './project';

export const ProjectStartPage: React.FC = () => {
  const project = useAppSelector(selectSelectProject);
  return (
    <ProjectAllow command={ProjectCommand.StartProject} project={project} projectKey={ProjectPath.Start}>
      <ProjectStart project={project!} />
    </ProjectAllow>
  );
};
