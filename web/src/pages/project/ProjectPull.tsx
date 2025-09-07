import * as React from 'react';
import { useCallback } from 'react';
import { type ErrorDto, fetchPullRepository, type ProjectDto } from '@blueskyfish/pierflow/api';
import { HeadLine, ScrollBar, ScrollingDirection } from '@blueskyfish/pierflow/components';
import { ProjectDetail } from './ProjectDetail.tsx';
import {
  addEventMessager,
  addMessage,
  EventStatus,
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

  const pullRepository = useCallback(() => {
    setLoading(true);

    const remoteListener = addEventMessager(eventSource, toEventType(ProjectCommand.PullRepository), (event) => {
      switch (event.status) {
        case EventStatus.Success:
          setLoading(false);
          if (event.id === project.id) {
            const { branch } = JSON.parse(event.message) as { branch: string };
            dispatch(updateProjectBranch({ projectId: project.id, branch }));
          }
          remoteListener();
          return;
        case EventStatus.Error:
          setLoading(false);
          dispatch(addMessage(event));
          remoteListener();
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
  }, [dispatch, eventSource, project.id]);

  return (
    <>
      <HeadLine
        title={`Pull ${project!.name}`}
        as={'h2'}
        icon={'mdi mdi-source-pull'}
        className={'mb-4 px-3 pt-3'}
        loading={loading}
      />
      <ul className={'menu menu-horizontal mb-4'}>
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
