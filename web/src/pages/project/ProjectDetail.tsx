import type { ProjectDto } from '@blueskyfish/pierflow/api';
import { DateFormatter, LabelValue, Panel } from '@blueskyfish/pierflow/components';
import { ProjectStatus } from './ProjectStatus.tsx';
import type { ProjectStatusKind } from '@blueskyfish/pierflow/icons';
import * as React from 'react';

export interface ProjectDetailProps {
  project: ProjectDto;
}

export const ProjectDetail: React.FC<ProjectDetailProps> = ({ project }) => {
  return (
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
  );
};
