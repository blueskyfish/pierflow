import {
  selectSelectProject,
  updatePageKey,
  updateSelectedId,
  useAppDispatch,
  useAppSelector,
} from '@blueskyfish/pierflow/stores';
import * as React from 'react';
import { useEffect } from 'react';
import { useParams } from 'react-router';

export const ProjectDetailPage: React.FC = () => {
  const projectId = useParams().projectId ?? '??';
  const dispatch = useAppDispatch();

  // Update pageKey to project id and update also selected project id
  // when this component is mounted
  useEffect(() => {
    dispatch(updateSelectedId(projectId));
    dispatch(updatePageKey(projectId));
  }, [dispatch, projectId]);

  const project = useAppSelector(selectSelectProject);
  if (!project) {
    return (
      <div className={'flex flex-col align-items-stretch h-100'}>
        <h1>Project Detail</h1>
        <p>Project not found.</p>
      </div>
    );
  }

  return (
    <div className={'flex flex-col align-items-stretch height-100 overflow-auto'}>
      <h1 className={'text-3xl'}>Project Detail {project.name}</h1>
      <p>
        This is the project detail page <b>{projectId}</b>
      </p>
    </div>
  );
};
