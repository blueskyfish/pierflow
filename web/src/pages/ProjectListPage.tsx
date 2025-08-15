import { HeadLine, Paragraph } from '@blueskyfish/pierflow/components';
import { selectProjectList, useAppDispatch, useAppSelector } from '@blueskyfish/pierflow/stores';
import * as React from 'react';
import { useEffect } from 'react';
import { updatePageKey } from '../stores/layout';
import { updateSelectedId } from '../stores/projects';
import { ProjectCard } from './project';

export const ProjectListPage: React.FC = () => {
  const dispatch = useAppDispatch();
  const list = useAppSelector(selectProjectList);

  const isEmpty = list.length === 0;

  // Reset the selected project id and update pageKey to 'project-list' when this component is mounted
  useEffect(() => {
    dispatch(updateSelectedId(null));
    dispatch(updatePageKey('project-list'));
  }, [dispatch]);

  return (
    <div className={'flex flex-col align-items-stretch height-100 overflow-auto p-3'}>
      <HeadLine as={'h2'} title={'Project Overview'} icon={'mdi mdi-list-box-outline'} />
      {isEmpty && <Paragraph>Leider kein Projek definiert.</Paragraph>}
      {!isEmpty && (
        <div className={'flex flex-row flex-wrap gap-4'}>
          {list.map((project) => (
            <ProjectCard key={project.id} project={project} />
          ))}
        </div>
      )}
    </div>
  );
};
