import { HeadLine } from '@blueskyfish/pierflow/components';
import { ProjectCommand, selectSelectProject, useAppSelector } from '@blueskyfish/pierflow/stores';
import { ProjectPath } from '@blueskyfish/pierflow/utils';
import * as React from 'react';
import { ProjectAllow } from './project';

export const ProjectClonePage: React.FC = () => {
  const project = useAppSelector(selectSelectProject);
  return (
    <ProjectAllow command={ProjectCommand.CloneRepository} project={project} projectKey={ProjectPath.Clone}>
      <HeadLine title={`Clone: ${project!.name}`} as={'h2'} icon={'mdi mdi-factory'} className={'mb-4'} />
    </ProjectAllow>
  );
};
