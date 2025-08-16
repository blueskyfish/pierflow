import { DateFormatter, HeadLine, LabelValue, Paragraph } from '@blueskyfish/pierflow/components';
import {
  selectSelectProject,
  updatePageKey,
  updateSelectedId,
  useAppDispatch,
  useAppSelector
} from '@blueskyfish/pierflow/stores';
import * as React from 'react';
import { useEffect } from 'react';
import { useParams } from 'react-router';
import { Panel } from '../components/Panel.tsx';

export const ProjectDetailPage: React.FC = () => {
  const projectId = useParams().projectId ?? '??';
  const dispatch = useAppDispatch();

  // Update pageKey to project id and update also selected project id
  // when this component is mounted
  useEffect(() => {
    dispatch(updateSelectedId(projectId));
    dispatch(updatePageKey(projectId));
  }, [dispatch, projectId]);

  const project = useAppSelector(selectSelectProject);
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
    <div className={'flex flex-col align-items-stretch height-100 overflow-auto p-3'}>
      <HeadLine title={project.name} as={'h2'} icon={'mdi mdi-factory'} className={'mb-4'} />
      <Panel>
        <div className={'flex flex-row align-items-stretch'}>
          <div className={'flex-grow-1 w-1/5'}>
            <div>{project.status}</div>
          </div>
          <div className={'flex-shrink-1 w-4/5'}>
            <LabelValue label={'Project'}>{project.name}</LabelValue>
            <LabelValue label={'Beschreibung'}>{project.description}</LabelValue>
            <LabelValue label={'Git Url'}>
              <a className={'link link-primary'} href={project.gitUrl} target={'_blank'}>
                {project.gitUrl}
              </a>
            </LabelValue>
            <LabelValue label={'Git Branch'}>{project.branch}</LabelValue>
            <LabelValue label={'Git User'}>{project.user}</LabelValue>
            <LabelValue label={'Verzeichnis'}>{project.path}</LabelValue>
            <LabelValue label={'Erstellt am'}>
              <DateFormatter date={project.creation} />
            </LabelValue>
            <LabelValue label={'Zuletzt geändert am'}>
              <DateFormatter date={project.modified} />
            </LabelValue>
          </div>
        </div>
      </Panel>
    </div>
  );
};
