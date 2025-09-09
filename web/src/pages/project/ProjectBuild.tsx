import * as React from 'react';
import { useCallback, useState } from 'react';
import { type ErrorDto, fetchBuildProject, type ProjectDto } from '@blueskyfish/pierflow/api';
import { HeadLine, ScrollBar, ScrollingDirection, useToast } from '@blueskyfish/pierflow/components';
import { ProjectDetail } from './ProjectDetail.tsx';
import {
  addEventMessager,
  addMessage,
  EventStatus,
  ProjectCommand,
  type ServerEvent,
  setError,
  toEventType,
  useAppDispatch,
  useEventSource,
} from '@blueskyfish/pierflow/stores';
import { Duration } from 'luxon';

export interface ProjectBuildProps {
  project: ProjectDto;
}

export const ProjectBuild: React.FC<ProjectBuildProps> = ({ project }) => {
  const [loading, setLoading] = useState(false);
  const [startTime, setStartTime] = useState<number>(-1);
  const [duration, setDuration] = useState<number>(-1);
  const dispatch = useAppDispatch();
  const eventSource = useEventSource();
  const toaster = useToast();

  const buildProject = useCallback(() => {
    // Implement the build project logic here
    console.log(`Building project ${project.name} with taskfile ${project.taskfile}`);

    setStartTime(Date.now());

    const timerHandle = setInterval(() => {
      setDuration(Date.now() - startTime);
    }, 1_000);

    const removeListener = addEventMessager(
      eventSource,
      toEventType(ProjectCommand.BuildProject),
      (event: ServerEvent) => {
        switch (event.status) {
          case EventStatus.Success:
            setLoading(false);
            setStartTime(-1);
            setDuration(-1);
            clearInterval(timerHandle);
            removeListener();
            toaster.add({
              state: 'success',
              title: 'Build',
              message: `Build operation completed successfully for project ${project.name}.`,
              timeout: 3_000,
            });
            // TODO reload project details
            return;
          case EventStatus.Error:
            dispatch(addMessage(event));
            setLoading(false);
            setStartTime(-1);
            setDuration(-1);
            clearInterval(timerHandle);
            removeListener();
            toaster.add({
              state: 'error',
              title: 'Build',
              message: `Build operation failed for project ${project.name}.`,
              timeout: 3_000,
            });
            return;
          default:
            dispatch(addMessage(event));
            return;
        }
      }, // end of event handler
    );

    fetchBuildProject(project.id).catch((error: ErrorDto) => {
      dispatch(setError(error));
      setLoading(false);
      setStartTime(-1);
      setDuration(-1);
      clearInterval(timerHandle);
      removeListener();
    });
  }, [dispatch, eventSource, project.id, project.name, project.taskfile, startTime, toaster]);

  return (
    <>
      <HeadLine
        title={`Build ${project!.name}`}
        as={'h2'}
        icon={`mdi mdi-application-brackets-outline`}
        className={'mb-4 px-3 pt-3'}
        loading={loading}
      />
      <ul className={'menu menu-horizontal'}>
        <li>
          <button type={'button'} className={'btn btn-soft btn-primary'} disabled={loading} onClick={buildProject}>
            <span className={'mdi mdi-application-brackets-outline mr-2'} />
            Build Project
          </button>
        </li>
        {duration > 0 && (
          <li>
            <span className={'text-sm'}>Duration: {Duration.fromMillis(duration).toFormat('mm:ss:SSS')}</span>
          </li>
        )}
        <li></li>
      </ul>
      <ScrollBar direction={ScrollingDirection.Vertical} className={'p-3'}>
        <ProjectDetail project={project} />
      </ScrollBar>
    </>
  );
};
