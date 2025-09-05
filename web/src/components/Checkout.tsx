import {
  addMessage,
  ProjectCommand,
  type ProjectDto,
  updateBranchList,
  useAppDispatch,
  useEventSource,
} from '@blueskyfish/pierflow/stores';
import { HeadLine } from './typograhie';
import * as React from 'react';
import { useEffect } from 'react';
import { type BranchDto, fetchBranchList } from '../stores/fetching';
import { addEventMessager, toEventType } from '../stores/events/events.messager.ts';
import { EventStatus, type ServerEvent } from '../stores/events/events.models.ts';

export interface CheckoutProps {
  project: ProjectDto;
}

export const Checkout: React.FC<CheckoutProps> = ({ project }) => {
  // eslint-disable-next-line @typescript-eslint/no-unused-vars
  const [refresh, _] = React.useState<boolean>(false);
  const dispatch = useAppDispatch();
  const eventSource = useEventSource();

  useEffect(() => {
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

    fetchBranchList(project.id, refresh).catch(() => removeListener());

    return () => {
      removeListener();
    };
  }, [dispatch, eventSource, project, refresh]);
  return (
    <>
      <HeadLine title={`Checkout ${project!.name}`} as={'h2'} icon={'mdi mdi-factory'} className={'mb-4'} />
      <ul>
        {project.branchList && project.branchList.map((item: BranchDto) => <li key={item.branch}>{item.branch}</li>)}
      </ul>
    </>
  );
};
