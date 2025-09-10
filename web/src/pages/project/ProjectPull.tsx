import * as React from 'react';
import { useCallback } from 'react';
import { type ErrorDto, fetchPullRepository, type ProjectDto } from '@blueskyfish/pierflow/api';
import { HeadLine, ScrollBar, ScrollingDirection, useToast } from '@blueskyfish/pierflow/components';
import { ProjectDetail } from './ProjectDetail.tsx';
import {
  addEventMessager,
  addMessage,
  EventStatus,
  loadProjectDetails,
  ProjectCommand,
  setError,
  toEventType,
  updateProjectBranch,
  useAppDispatch,
  useEventSource,
} from '@blueskyfish/pierflow/stores';

export interface ProjectPullProps {
  project: ProjectDto;
}

export const ProjectPull: React.FC<ProjectPullProps> = ({ project }) => {
  const [loading, setLoading] = React.useState(false);
  const dispatch = useAppDispatch();
  const eventSource = useEventSource();
  const toaster = useToast();

  const pullRepository = useCallback(() => {
    setLoading(true);

    const remoteListener = addEventMessager(eventSource, toEventType(ProjectCommand.PullRepository), (event) => {
      switch (event.status) {
        case EventStatus.Success:
          setLoading(false);
          if (event.id === project.id) {
            const { branch } = JSON.parse(event.message) as { branch: string };
            dispatch(updateProjectBranch({ projectId: project.id, branch }));
            toaster.add({
              state: 'success',
              title: 'Pull',
              message: `Pull operation completed successfully for project ${project.name} and ${branch}.`,
            });
            dispatch(loadProjectDetails(project.id)); // reload project details to reset the branch if necessary
          }
          remoteListener();
          return;
        case EventStatus.Error:
          setLoading(false);
          dispatch(addMessage(event));
          remoteListener();
          toaster.add({
            state: 'error',
            title: 'Pull',
            message: `Pull operation failed for project ${project.name}.`,
          });
          return;
        default:
          dispatch(addMessage(event));
          return;
      }
    });

    fetchPullRepository(project.id).catch((error: ErrorDto) => {
      dispatch(setError(error));
      setLoading(false);
      remoteListener();
    });

    return () => {
      setLoading(false);
      if (remoteListener) {
        remoteListener();
      }
    };
  }, [dispatch, eventSource, project.id, project.name, toaster]);

  return (
    <>
      <HeadLine
        title={`Pull ${project!.name}`}
        as={'h2'}
        icon={'mdi mdi-source-pull'}
        className={'mb-4 px-3 pt-3'}
        loading={loading}
      />
      <ul className={'menu menu-horizontal border-t border-b border-gray-200 w-full'}>
        <li>
          <button type={'button'} className={'btn btn-soft btn-primary'} onClick={() => pullRepository()}>
            <span className={'mdi mdi-source-pull mr-2'} />
            <span>Pull Repository</span>
          </button>
        </li>
      </ul>
      <ScrollBar direction={ScrollingDirection.Vertical} className={'p-3'}>
        <ProjectDetail project={project} />
      </ScrollBar>
    </>
  );
};
