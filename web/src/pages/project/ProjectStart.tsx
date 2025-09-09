import * as React from 'react';
import { useCallback, useState } from 'react';
import { type ErrorDto, fetchStartProject, type ProjectDto } from '@blueskyfish/pierflow/api';
import { HeadLine, ScrollBar, ScrollingDirection, useToast } from '@blueskyfish/pierflow/components';
import { ProjectDetail } from './ProjectDetail.tsx';
import {
  addEventMessager,
  addMessage,
  EventStatus,
  ProjectCommand,
  type ServerEvent,
  toEventType,
  useAppDispatch,
  useEventSource,
} from '@blueskyfish/pierflow/stores';

export interface ProjectStartProps {
  project: ProjectDto;
}

export const ProjectStart: React.FC<ProjectStartProps> = ({ project }) => {
  const [loading, setLoading] = useState(false);
  const dispatch = useAppDispatch();
  const eventSource = useEventSource();
  const toaster = useToast();

  const startProject = useCallback(() => {
    setLoading(true);
    const removeListener = addEventMessager(
      eventSource,
      toEventType(ProjectCommand.StartProject),
      (event: ServerEvent) => {
        switch (event.status) {
          case EventStatus.Success:
            console.log('> StartProject event:', event);
            removeListener();
            setLoading(false);
            toaster.add({
              state: 'success',
              title: 'Start',
              message: `Start operation completed successfully for project ${project.name}.`,
            });
            return;
          case EventStatus.Error:
            dispatch(addMessage(event));
            setLoading(false);
            removeListener();
            toaster.add({
              state: 'error',
              title: 'Start',
              message: `Start operation failed for project ${project.name}.`,
            });
            return;
          default:
            dispatch(addMessage(event));
            return;
        }
      },
    );

    fetchStartProject(project.id).catch((error: ErrorDto) => {
      console.error('Failed to start project:', error);
      setLoading(false);
      removeListener();
    });
  }, [dispatch, eventSource, project.id, project.name, toaster]);

  return (
    <>
      <HeadLine
        title={`Start ${project!.name}`}
        as={'h2'}
        icon={`mdi mdi-factory`}
        className={'mb-4 px-3 pt-3'}
        loading={loading}
      />
      <ul className={'menu menu-horizontal mb-4 flex-shrink-1 border-t border-b border-gray-200 w-full'}>
        <li>
          <button type={'button'} className={'btn btn-soft btn-primary'} disabled={loading} onClick={startProject}>
            <span className={'mr-1'}>Start Project</span>
            <span className={'mdi mdi-chevron-right'} />
          </button>
        </li>
      </ul>
      <ScrollBar direction={ScrollingDirection.Vertical} className={'p-3'}>
        <ProjectDetail project={project} />
      </ScrollBar>
    </>
  );
};
