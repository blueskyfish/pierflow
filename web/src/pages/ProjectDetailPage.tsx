import { DateFormatter, HeadLine, LabelValue, Panel } from '@blueskyfish/pierflow/components';
import type { ProjectStatusKind } from '@blueskyfish/pierflow/icons';
import { selectSelectProject, updateProjectKey, useAppDispatch, useAppSelector } from '@blueskyfish/pierflow/stores';
import { ProjectPath } from '@blueskyfish/pierflow/utils';
import * as React from 'react';
import { useEffect } from 'react';
import { ProjectStatus } from './project';

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
          <div className={'flex-grow-1 w-1/6'}>
            <ProjectStatus status={project.status as ProjectStatusKind} />
          </div>
          <div className={'flex-shrink-1 w-5/6'}>
            <LabelValue label={'Projekt'} size={'sm'}>
              {project.name}
            </LabelValue>
            <LabelValue label={'Beschreibung'} size={'sm'}>
              {project.description}
            </LabelValue>
            <LabelValue label={'Repository'} size={'sm'}>
              <a className={'link link-primary'} href={project.gitUrl} target={'_blank'}>
                {project.gitUrl}
              </a>
            </LabelValue>
            <LabelValue label={'Branch'} size={'sm'}>
              {project.branch}
            </LabelValue>
            <LabelValue label={'User'} size={'sm'}>
              {project.user}
            </LabelValue>
            <LabelValue label={'Verzeichnis'} size={'sm'}>
              {project.path}
            </LabelValue>
            <LabelValue label={'Erstellt am'} size={'sm'}>
              <DateFormatter date={project.creation} />
            </LabelValue>
            <LabelValue label={'Zuletzt geändert am'} size={'sm'}>
              <DateFormatter date={project.modified} />
            </LabelValue>
          </div>
        </div>
      </Panel>
    </>
  );
};
