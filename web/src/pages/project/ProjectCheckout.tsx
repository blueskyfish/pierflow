import {
  addEventMessager,
  addMessage,
  EventStatus,
  ProjectCommand,
  type ServerEvent,
  setError,
  toEventType,
  updateBranchList,
  updateProjectBranch,
  useAppDispatch,
  useEventSource,
} from '@blueskyfish/pierflow/stores';
import { HeadLine, ScrollBar, ScrollingDirection, useToast } from '@blueskyfish/pierflow/components';
import * as React from 'react';
import { useCallback, useEffect, useState } from 'react';
import {
  type BranchDto,
  type ErrorDto,
  fetchBranchList,
  fetchCheckoutRepository,
  type ProjectDto,
} from '@blueskyfish/pierflow/api';
import { GitPlace } from './GitPlace.tsx';

export interface CheckoutProps {
  project: ProjectDto;
}

export const ProjectCheckout: React.FC<CheckoutProps> = ({ project }) => {
  const [refresh, setRefresh] = useState<boolean>(false);
  const [loading, setLoading] = useState<boolean>(false);
  const dispatch = useAppDispatch();
  const eventSource = useEventSource();
  const toaster = useToast();

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
            toaster.add({
              state: 'error',
              title: 'Checkout',
              message: `Failed to load branch list from project ${project.name} repository`,
              timeout: 3_000,
            });
            return;
          default:
            dispatch(addMessage(event));
            return;
        }
      }, // end of event handler
    );

    // start to fetch the branch list. The branch list will be delivered via server side events
    fetchBranchList(project.id, refresh).catch((error: ErrorDto) => {
      dispatch(setError(error));
      setLoading(false);
      removeListener();
    });

    return removeListener;
  }, [dispatch, eventSource, project.id, project.name, refresh, toaster]);

  const checkoutBranchRepositor = useCallback(
    (branch: string, place: string) => {
      setLoading(true);
      const removeListener = addEventMessager(
        eventSource,
        toEventType(ProjectCommand.CheckoutRepository),
        (event: ServerEvent) => {
          switch (event.status) {
            case EventStatus.Success:
              console.log('> CheckoutRepository event:', event);
              if (event.id === project.id && event.message) {
                const { branch } = JSON.parse(event.message) as { branch: string };
                dispatch(updateProjectBranch({ projectId: project.id, branch }));
              }
              removeListener();
              setLoading(false);
              toaster.add({
                state: 'success',
                title: 'Checkout',
                message: `Branch ${branch} is checked out from project ${project.name} repository`,
                timeout: 3_000,
              });
              return;
            case EventStatus.Error:
              dispatch(addMessage(event));
              removeListener();
              setLoading(false);
              toaster.add({
                state: 'error',
                title: 'Checkout',
                message: `Failed to checkout branch ${branch} from project ${project.name} repository`,
                timeout: 3_000,
              });
              return;
            default:
              dispatch(addMessage(event));
              return;
          }
        }, // end of event handler
      );

      // start to check out the branch. The result will be delivered via server side events
      fetchCheckoutRepository(project.id, { branch, place, message: 'TODO' }).catch((error: ErrorDto) => {
        dispatch(setError(error));
        setLoading(false);
        removeListener();
      });

      return () => {
        setLoading(false);
        if (removeListener) removeListener();
      };
    },
    [dispatch, eventSource, project.id, project.name, toaster],
  );

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
    <div className={'flex flex-col items-stretch h-full overflow-auto'}>
      <HeadLine
        title={`Checkout ${project!.name}`}
        as={'h2'}
        icon={'mdi mdi-source-branch'}
        className={'mb-4 flex-shrink-1 px-3 pt-3'}
        loading={loading}
      />
      <ul className={'menu menu-horizontal mb-4 flex-shrink-1 border-t border-b border-gray-200 w-full'}>
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
      <ScrollBar direction={ScrollingDirection.Vertical} className={'p-3 flex-grow-1'}>
        <div className={'overflow-auto'}>
          <table className={'table table-sm'}>
            <colgroup>
              <col width={'*'} />
              <col width={'10%'} />
              <col width={'15%'} />
            </colgroup>
            <thead>
              <tr>
                <th>Branch</th>
                <th className={'text-center'}>Place</th>
                <th>&nbsp;</th>
              </tr>
            </thead>
            <tbody>
              {project.branchList &&
                project.branchList.map((item: BranchDto) => {
                  return (
                    <tr className={`${item.active ? 'bg-primary text-white' : 'hover:bg-base-300'}`} key={item.branch}>
                      <td>{item.branch}</td>
                      <td>
                        <GitPlace place={item.place} />
                      </td>
                      <td>
                        {item.active && (
                          <>
                            <span className={'mdi mdi-check mr-3'} />
                            Active
                          </>
                        )}
                        {!item.active && (
                          <>
                            <button
                              type={'button'}
                              className={'btn btn-xs btn-soft btn-secondary'}
                              onClick={() => checkoutBranchRepositor(item.branch, item.place)}
                            >
                              Checkout
                              <span className={'ml-1 mdi mdi-chevron-right'}></span>
                            </button>
                          </>
                        )}
                      </td>
                    </tr>
                  );
                })}
            </tbody>
          </table>
        </div>
      </ScrollBar>
    </div>
  );
};
