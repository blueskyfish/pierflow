import * as React from 'react';
import { useCallback, useEffect, useState } from 'react';
import { type ErrorDto, fetchBuildProject, fetchGetTaskFileList, type ProjectDto } from '@blueskyfish/pierflow/api';
import { HeadLine, ScrollBar, ScrollingDirection } from '@blueskyfish/pierflow/components';
import { ProjectDetail } from './ProjectDetail.tsx';
import {
  addEventMessager,
  addMessage,
  EventStatus,
  ProjectCommand,
  type ServerEvent,
  setError,
  toEventType,
  updateProjectTaskfileList,
  useAppDispatch,
  useEventSource,
} from '@blueskyfish/pierflow/stores';
import { ProjectTaskfileList } from './ProjectTaskfileList.tsx';
import { Duration } from 'luxon';

export interface ProjectBuildProps {
  project: ProjectDto;
}

export const ProjectBuild: React.FC<ProjectBuildProps> = ({ project }) => {
  const [loading, setLoading] = useState(false);
  const [taskfile, setTaskfile] = useState<string>('');
  const [startTime, setStartTime] = useState<number>(-1);
  const [duration, setDuration] = useState<number>(-1);
  const dispatch = useAppDispatch();
  const eventSource = useEventSource();

  const loadTaskfileList = useCallback(() => {
    setLoading(true);

    fetchGetTaskFileList(project.id)
      .then((taskfileList: string[]) => {
        dispatch(updateProjectTaskfileList({ projectId: project.id, taskfileList }));
      })
      .catch((error: ErrorDto) => {
        dispatch(setError(error));
      })
      .finally(() => setLoading(false));
  }, [dispatch, project.id]);

  const buildProject = useCallback(
    (taskfile: string) => {
      // Implement the build project logic here
      console.log(`Building project ${project.name} with taskfile ${taskfile}`);

      setTaskfile(taskfile);
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
              if (event.id === project.id) {
                setTaskfile('');
              }
              setLoading(false);
              setStartTime(-1);
              setDuration(-1);
              clearInterval(timerHandle);
              removeListener();
              return;
            case EventStatus.Error:
              dispatch(addMessage(event));
              setLoading(false);
              setStartTime(-1);
              setDuration(-1);
              clearInterval(timerHandle);
              removeListener();
              return;
            default:
              dispatch(addMessage(event));
              return;
          }
        }, // end of event handler
      );

      fetchBuildProject(project.id, { taskfile, message: 'TODO' }).catch((error: ErrorDto) => {
        dispatch(setError(error));
        setLoading(false);
        setStartTime(-1);
        setDuration(-1);
        clearInterval(timerHandle);
        removeListener();
      });
    },
    [dispatch, eventSource, project.id, project.name, startTime],
  );

  useEffect(() => {
    if (Array.isArray(project.taskfileList)) {
      return;
    }
    loadTaskfileList();
  }, [loadTaskfileList, project]);

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
          <button type={'button'} className={'btn btn-soft btn-primary'} onClick={loadTaskfileList}>
            <span className={'mdi mdi-refresh mr-2'} />
            Reload Taskfiles
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
        <ProjectTaskfileList taskFiles={project.taskfileList ?? []} onBuild={buildProject} activeTaskfile={taskfile} />
      </ScrollBar>
    </>
  );
};
