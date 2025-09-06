import { HeadLine, Paragraph } from '@blueskyfish/pierflow/components';
import {
  selectProjectKey,
  selectSelectProject,
  updatePageKey,
  updateSelectedId,
  useAppDispatch,
  useAppSelector,
} from '@blueskyfish/pierflow/stores';
import * as React from 'react';
import { useEffect } from 'react';
import { Outlet, useParams } from 'react-router';
import { ProjectDock, ProjectMessage } from './project';

export const ProjectHomePage: React.FC = () => {
  const projectId = useParams().projectId ?? null;
  const dispatch = useAppDispatch();

  // Update pageKey to project id and update also selected project id
  // when this component is mounted
  useEffect(() => {
    dispatch(updateSelectedId(projectId));
    dispatch(updatePageKey(projectId));
  }, [dispatch, projectId]);

  const project = useAppSelector(selectSelectProject);
  const selectKey = useAppSelector(selectProjectKey);

  if (!project) {
    return (
      <div className={'flex flex-col align-items-stretch height-100 overflow-auto p-3'}>
        <HeadLine title={'Fehler'} icon={'mdi mdi-alert'} className={'mb-4'} />
        <div className={'alert alert-error w-full'}>
          <Paragraph size={'md'}>Projekt ist nicht vorhanden.</Paragraph>
        </div>
      </div>
    );
  }

  return (
    <div className={'flex flex-col align-items-stretch height-100 overflow-auto relative'}>
      <div className={'flex flex-col items-stretch flex-grow-1 overflow-auto'}>
        <Outlet />
      </div>
      <ProjectMessage filterId={project.id} />
      <ProjectDock commandMap={project.commandMap} projectId={project.id} selectKey={selectKey} />
    </div>
  );
};
