import {
  addEventMessager,
  addMessage,
  EventStatus,
  ProjectCommand,
  type ServerEvent,
  toEventType,
  updateBranchList,
  useAppDispatch,
  useEventSource,
} from '@blueskyfish/pierflow/stores';
import { HeadLine } from '@blueskyfish/pierflow/components';
import * as React from 'react';
import { useEffect, useState } from 'react';
import { type BranchDto, fetchBranchList, type ProjectDto } from '@blueskyfish/pierflow/api';
import { ProjectMessage } from './ProjectMessage.tsx';

export interface CheckoutProps {
  project: ProjectDto;
}

export const ProjectCheckout: React.FC<CheckoutProps> = ({ project }) => {
  const [refresh, setRefresh] = useState<boolean>(false);
  const [loading, setLoading] = useState<boolean>(false);
  const dispatch = useAppDispatch();
  const eventSource = useEventSource();

  // load the branch list if not already loaded
  const loadBranchList = React.useCallback(() => {
    setLoading(true);
    const removeListener = addEventMessager(
      eventSource,
      toEventType(ProjectCommand.BranchList),
      (event: ServerEvent) => {
        switch (event.status) {
          case EventStatus.Success:
            console.log('> BranchList event:', event);
            if (event.id === project.id && event.message) {
              const branchList = JSON.parse(event.message) as BranchDto[];
              dispatch(updateBranchList({ projectId: project.id, branchList }));
            }
            removeListener();
            setLoading(false);
            return;
          case EventStatus.Error:
            dispatch(updateBranchList({ projectId: project.id, branchList: [] }));
            dispatch(addMessage(event));
            removeListener();
            setLoading(false);
            return;
          default:
            dispatch(addMessage(event));
            return;
        }
      }, // end of event handler
    );

    // start to fetch the branch list. The branch list will be delivered via server side events
    fetchBranchList(project.id, refresh).catch(() => removeListener());

    return removeListener;
  }, [dispatch, eventSource, project, refresh]);

  useEffect(() => {
    if (Array.isArray(project.branchList)) {
      return;
    }
    const removeListener = loadBranchList();
    return () => {
      if (removeListener) removeListener();
    };
  }, [loadBranchList, project.branchList]);

  // Callback für Reload
  const handleReload = () => {
    loadBranchList();
  };

  const handleRefreshChange = () => {
    setRefresh(!refresh);
  };

  return (
    <div className={'flex flex-col items-stretch h-full'}>
      <HeadLine
        title={`Checkout ${project!.name}`}
        as={'h2'}
        icon={'mdi mdi-factory'}
        className={'mb-4 flex-shrink-1'}
        loading={loading}
      />
      <ul className={'menu menu-horizontal rounded-box mb-4 flex-shrink-1'}>
        <li>
          <button
            className={'tooltip'}
            data-tip='Reload the list of branches'
            disabled={loading}
            onClick={handleReload}
          >
            <span className={'mdi mdi-refresh'} />
            Refresh
          </button>
        </li>
        <li>
          <button
            className={'tooltip'}
            disabled={loading}
            data-tip='Force refresh from remote repository'
            onClick={handleRefreshChange}
          >
            <span
              className={`mdi ${refresh ? 'mdi-checkbox-marked-circle-outline' : 'mdi-checkbox-blank-circle-outline'}`}
            />
            Remote
          </button>
        </li>
      </ul>
      <ul className={'flex-grow-1 overflow-auto'}>
        {project.branchList && project.branchList.map((item: BranchDto) => <li key={item.branch}>{item.branch}</li>)}
      </ul>
      <ProjectMessage filterId={project.id} />
    </div>
  );
};
