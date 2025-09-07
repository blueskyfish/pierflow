import { ProjectCommand, selectSelectProject, useAppSelector } from '@blueskyfish/pierflow/stores';
import { ProjectPath } from '@blueskyfish/pierflow/utils';
import * as React from 'react';
import { ProjectAllow, ProjectPull } from './project';

export const ProjectPullPage: React.FC = () => {
  const project = useAppSelector(selectSelectProject);
  return (
    <ProjectAllow command={ProjectCommand.PullRepository} project={project} projectKey={ProjectPath.Pull}>
      <ProjectPull project={project!} />
    </ProjectAllow>
  );
};
