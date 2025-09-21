import type { ProjectDto } from '@blueskyfish/pierflow/api';
import React from 'react';
import { Section } from '../../components/common';

export interface ProjectPanelProps {
  project: ProjectDto;
  onAction: (command: string) => void;
}

export const ProjectPanel: React.FC<ProjectPanelProps> = ({ project, onAction }) => {
  return (
    <Section className={'flex flex-col items-stretch mx-4 mt-4 p-4 gap-2 flex-shrink-1'}>
      <h1 className={'text-2xl font-bold'}>{project.name}</h1>
      <p className={'text-sm text-gray-500'}>{project.description}</p>
      <p className={'text-base'}>
        <a className={'decoration-1 underline decoration-primary text-primary'} href={project.gitUrl} target={'_blank'}>
          {project.gitUrl}
        </a>
      </p>
      <p className={'text-base'}>{project.branch}</p>
      <div className={'flex flex-row items-center justify-end'}>
        <ul className={'menu menu-horizontal bg-gray-100 rounded-box'}>
          <li>
            <button type={'button'} className={'btn'} onClick={() => onAction('branch')}>
              <span className={'mdi mdi-source-branch'} />
              <span className={'ml-2'}>Branch</span>
            </button>
          </li>
          <li>
            <button type={'button'} className={'btn'} onClick={() => onAction('git')}>
              <span className={'mdi mdi-source-repository'} />
              <span className={'ml-2'}>Git</span>
            </button>
          </li>
          <li>
            <button type={'button'} className={'btn'} onClick={() => onAction('edit')}>
              <span className={'mdi mdi-pencil'} />
              <span className={'ml-2'}>Edit</span>
            </button>
          </li>
        </ul>
      </div>
    </Section>
  );
};
