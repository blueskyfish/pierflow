import { HeadLine, ScrollBar, ScrollingDirection } from '@blueskyfish/pierflow/components';
import {
  selectSelectProject,
  setError,
  updateProjectDetail,
  updateProjectKey,
  useAppDispatch,
  useAppSelector,
} from '@blueskyfish/pierflow/stores';
import { ProjectPath } from '@blueskyfish/pierflow/utils';
import * as React from 'react';
import { useCallback, useEffect, useRef, useState } from 'react';
import { ProjectDetail } from './project';
import { type ErrorDto, fetchChangeTaskfile, fetchGetTaskFileList, type ProjectDto } from '@blueskyfish/pierflow/api';

export const ProjectDetailPage: React.FC = () => {
  const [taskfiles, setTaskfiles] = useState<string[]>([]);
  const [loading, setLoading] = useState(false);
  const dispatch = useAppDispatch();
  const project = useAppSelector(selectSelectProject)!;

  const branchDialogRef = useRef<HTMLDialogElement>(null);

  // request the list of taskfiles and open the dialog "choose taskfile"
  const chooseTaskfile = useCallback(() => {
    setLoading(true);
    fetchGetTaskFileList(project.id)
      .then((list) => {
        setTaskfiles(list);
        setLoading(false);
        if (branchDialogRef.current) {
          branchDialogRef.current.showModal();
        }
      })
      .catch(() => {
        setLoading(false);
      });
  }, [project.id]);

  // choose command from ProjectDetail component
  // currently only "taskfile" is supported
  // which will open the dialog to choose a taskfile
  const chooseCommand = useCallback(
    (command: string) => {
      if (command === 'taskfile') {
        chooseTaskfile();
      }
    },
    [chooseTaskfile],
  );

  // update the taskfile of the project and close the dialog
  const updateTaskfile = useCallback(
    (taskfile: string) => {
      setLoading(true);

      fetchChangeTaskfile(project.id, taskfile)
        .then((p: ProjectDto) => {
          dispatch(updateProjectDetail(p));
        })
        .catch((error: ErrorDto) => {
          dispatch(setError(error));
        })
        .finally(() => setLoading(false));

      if (branchDialogRef.current) {
        branchDialogRef.current.close();
      }
    },
    [dispatch, project.id],
  );

  // Update the project key to Detail when this component is mounted
  useEffect(() => {
    dispatch(updateProjectKey(ProjectPath.Detail));
  }, [dispatch]);

  return (
    <>
      <HeadLine
        title={`Detail: ${project.name}`}
        as={'h2'}
        icon={'mdi mdi-factory'}
        className={'mb-4 px-3 pt-3'}
        loading={loading}
      />
      <ScrollBar direction={ScrollingDirection.Vertical} className={'p-3'}>
        <ProjectDetail project={project} onChange={chooseCommand} />
      </ScrollBar>
      <dialog ref={branchDialogRef} className={'modal'}>
        <div className='modal-box'>
          <form method='dialog'>
            <h3 className='font-bold text-lg'>Change Taskfile</h3>
            <table className={'table table-zebra table-pin-rows table-sm bg-base-100'}>
              <colgroup>
                <col width={'5%'} />
                <col width={'*'} />
                <col width={'15%'} />
              </colgroup>
              <thead>
                <tr>
                  <th className={'text-right'}>No</th>
                  <th>Taskfile</th>
                  <th className={'text-right'}>Action</th>
                </tr>
              </thead>
              <tbody>
                {taskfiles.map((taskfile, index) => (
                  <tr
                    key={index}
                    className={`${project.taskfile === taskfile ? 'bg-primary text-white' : 'hover:bg-base-300'}`}
                  >
                    <td className={'text-right'}>{index + 1}</td>
                    <td>{taskfile}</td>
                    <td className={'text-right'}>
                      {project.taskfile === taskfile && (
                        <span className={`badge badge-soft badge-sm badge-primary`}>Current Taskfile</span>
                      )}
                      {project.taskfile !== taskfile && (
                        <button
                          type={'button'}
                          className={'btn btn-xs btn-soft btn-primary'}
                          onClick={() => updateTaskfile(taskfile)}
                        >
                          <span className={'mr-2'}>Select</span>
                          <span className={'mdi mdi-chevron-right'} />
                        </button>
                      )}
                    </td>
                  </tr>
                ))}
              </tbody>
            </table>
            <div className='modal-action'>
              <button className='btn'>Cancel</button>
            </div>
          </form>
        </div>
      </dialog>
    </>
  );
};
