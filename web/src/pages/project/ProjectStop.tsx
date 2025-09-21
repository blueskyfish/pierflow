import { type ErrorDto, fetchStopProject, ProjectCommand } from '@blueskyfish/pierflow/api';
import { HeadLine, ScrollBar, ScrollingDirection, useToast } from '@blueskyfish/pierflow/components';
import {
  addEventMessager,
  addMessage,
  EventStatus,
  loadProjectDetails,
  type ProjectInfo,
  type ServerEvent,
  setError,
  toEventType,
  useAppDispatch,
  useEventSource,
} from '@blueskyfish/pierflow/stores';
import * as React from 'react';
import { useCallback, useState } from 'react';
import { ProjectDetail } from './ProjectDetail.tsx';

export interface ProjectStopProps {
  project: ProjectInfo;
}

export const ProjectStop: React.FC<ProjectStopProps> = ({ project }) => {
  const [loading, setLoading] = useState(false);
  const dispatch = useAppDispatch();
  const eventSource = useEventSource();
  const toaster = useToast();

  const stopProject = useCallback(() => {
    setLoading(true);

    const remoteListener = addEventMessager(
      eventSource,
      toEventType(ProjectCommand.StopProject),
      (event: ServerEvent) => {
        switch (event.status) {
          case EventStatus.Success:
            console.log('> StopProject event:', event);
            remoteListener();
            setLoading(false);
            toaster.add({
              state: 'success',
              title: 'Stop',
              message: `Stop operation completed successfully for project ${project.name}.`,
            });
            dispatch(loadProjectDetails(project.id)); // reload project details to reset the branch if necessary
            return;
          case EventStatus.Error:
            dispatch(addMessage(event));
            setLoading(false);
            remoteListener();
            toaster.add({
              state: 'error',
              title: 'Stop',
              message: `Stop operation failed for project ${project.name}.`,
            });
            return;
          default:
            dispatch(addMessage(event));
            return;
        }
      },
    );

    fetchStopProject(project.id).catch((error: ErrorDto) => {
      dispatch(setError(error));
      setLoading(false);
      remoteListener();
    });
  }, [dispatch, eventSource, project.id, project.name, toaster]);

  return (
    <>
      <HeadLine
        title={`Stop ${project!.name}`}
        as={'h2'}
        icon={`mdi mdi-factory`}
        className={'mb-4 px-3 pt-3'}
        loading={loading}
      />
      <ul className={'menu menu-horizontal border-t border-b border-gray-200 w-full'}>
        <li>
          <button type={'button'} className={'btn btn-soft btn-primary'} onClick={stopProject}>
            <span className={'mr-1'}>Stop Project</span>
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
