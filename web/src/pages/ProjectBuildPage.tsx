import { ProjectCommand, selectSelectProject, useAppSelector } from '@blueskyfish/pierflow/stores';
import { ProjectPath } from '@blueskyfish/pierflow/utils';
import * as React from 'react';
import { ProjectAllow, ProjectBuild } from './project';

export const ProjectBuildPage: React.FC = () => {
  const project = useAppSelector(selectSelectProject);
  return (
    <ProjectAllow command={ProjectCommand.BuildProject} project={project} projectKey={ProjectPath.Build}>
      <ProjectBuild project={project!} />
    </ProjectAllow>
  );
};
