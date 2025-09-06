import { HeadLine } from '@blueskyfish/pierflow/components';
import { ProjectCommand, selectSelectProject, useAppSelector } from '@blueskyfish/pierflow/stores';
import { ProjectPath } from '@blueskyfish/pierflow/utils';
import * as React from 'react';
import { ProjectAllow } from './project';

export const ProjectStartPage: React.FC = () => {
  const project = useAppSelector(selectSelectProject);
  return (
    <ProjectAllow command={ProjectCommand.StartProject} project={project} projectKey={ProjectPath.Start}>
      <HeadLine title={`Start ${project!.name}`} as={'h2'} icon={`mdi mdi-factory`} className={'mb-4 px-3 pt-3'} />
    </ProjectAllow>
  );
};
