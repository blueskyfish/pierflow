import { DateFormatter, HeadLine, LabelValue, Panel } from '@blueskyfish/pierflow/components';
import { selectSelectProject, updateProjectKey, useAppDispatch, useAppSelector } from '@blueskyfish/pierflow/stores';
import { ProjectPath } from '@blueskyfish/pierflow/utils';
import * as React from 'react';
import { useEffect } from 'react';

export const ProjectDetailPage: React.FC = () => {
  const dispatch = useAppDispatch();
  // Update the project key to Detail when this component is mounted
  useEffect(() => {
    dispatch(updateProjectKey(ProjectPath.Detail));
  }, [dispatch]);

  const project = useAppSelector(selectSelectProject)!;
  return (
    <>
      <HeadLine title={`Detail: ${project.name}`} as={'h2'} icon={'mdi mdi-factory'} className={'mb-4'} />
      <Panel>
        <div className={'flex flex-row align-items-stretch'}>
          <div className={'flex-grow-1 w-1/5'}>
            <div>{project.status}</div>
          </div>
          <div className={'flex-shrink-1 w-4/5'}>
            <LabelValue label={'Projekt'}>{project.name}</LabelValue>
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
    </>
  );
};
