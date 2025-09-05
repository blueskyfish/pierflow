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
import { useEffect } from 'react';
import { type BranchDto, fetchBranchList, type ProjectDto } from '@blueskyfish/pierflow/api';

export interface CheckoutProps {
  project: ProjectDto;
}

export const ProjectCheckout: React.FC<CheckoutProps> = ({ project }) => {
  const [refresh, setRefresh] = React.useState<boolean>(false);
  const dispatch = useAppDispatch();
  const eventSource = useEventSource();

  // load the branch list if not already loaded
  const loadBranchList = React.useCallback(() => {
    if (Array.isArray(project.branchList)) {
      return;
    }

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
            return;
          case EventStatus.Error:
            dispatch(updateBranchList({ projectId: project.id, branchList: [] }));
            dispatch(addMessage(event));
            removeListener();
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
    const removeListener = loadBranchList();
    return () => {
      if (removeListener) removeListener();
    };
  }, [loadBranchList]);

  // Callback für Reload
  const handleReload = () => {
    setRefresh((prev) => !prev);
    loadBranchList();
  };

  return (
    <>
      <HeadLine title={`Checkout ${project!.name}`} as={'h2'} icon={'mdi mdi-factory'} className={'mb-4'} />
      <button onClick={handleReload}>Reload</button>
      <ul>
        {project.branchList && project.branchList.map((item: BranchDto) => <li key={item.branch}>{item.branch}</li>)}
      </ul>
    </>
  );
};
