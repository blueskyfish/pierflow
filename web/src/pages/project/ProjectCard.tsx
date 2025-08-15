import { CardImage, HeadLine, Paragraph } from '@blueskyfish/pierflow/components';
import type { ProjectDto } from '@blueskyfish/pierflow/stores';
import * as React from 'react';
import { Link } from 'react-router';
import { RouteBuilder } from '../../utils/routing/route-names.ts';

export interface ProjectCardProps {
  project: ProjectDto;
}

export const ProjectCard: React.FC<ProjectCardProps> = ({ project }) => {
  return (
    <div className={'card w-1/5 bg-base-100 shadow-md'}>
      <figure>
        <CardImage fillColor={'oklch(92% 0 0)'} height={65}>
          <span className={'mdi mdi-factory mdi-48px text-base'}></span>
        </CardImage>
      </figure>
      <div className={'card-body p-3'}>
        <HeadLine title={project.name} as={'h4'} className={'card-title'} />
        <Paragraph size={'sm'} className={'text-base-content/70'}>
          {project.description}
        </Paragraph>
        <div className={'card-actions justify-end'}>
          <Link to={RouteBuilder.buildProjectDetailPath(project.id)} className={'btn btn-xs btn-neutral'}>
            Projekt Details
          </Link>
        </div>
      </div>
    </div>
  );
};
