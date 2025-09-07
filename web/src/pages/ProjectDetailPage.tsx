import { HeadLine, ScrollBar, ScrollingDirection } from '@blueskyfish/pierflow/components';
import { selectSelectProject, updateProjectKey, useAppDispatch, useAppSelector } from '@blueskyfish/pierflow/stores';
import { ProjectPath } from '@blueskyfish/pierflow/utils';
import * as React from 'react';
import { useEffect } from 'react';
import { ProjectDetail } from './project';

export const ProjectDetailPage: React.FC = () => {
  const dispatch = useAppDispatch();
  // Update the project key to Detail when this component is mounted
  useEffect(() => {
    dispatch(updateProjectKey(ProjectPath.Detail));
  }, [dispatch]);

  const project = useAppSelector(selectSelectProject)!;
  return (
    <>
      <HeadLine title={`Detail: ${project.name}`} as={'h2'} icon={'mdi mdi-factory'} className={'mb-4 px-3 pt-3'} />
      <ScrollBar direction={ScrollingDirection.Vertical} className={'p-3'}>
        <ProjectDetail project={project} />
      </ScrollBar>
    </>
  );
};
