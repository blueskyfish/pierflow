import type { ProjectHistoryDto } from '@blueskyfish/pierflow/api';
import {
  loadProjectHistory,
  selectProjectProcessing,
  selectSelectProject,
  useAppDispatch,
  useAppSelector,
} from '@blueskyfish/pierflow/stores';
import React, { useCallback, useEffect, useMemo } from 'react';
import { ProjectPanel, ProjectPipeline } from '../components/projects';

export const ProjectProcessPage: React.FC = () => {
  const project = useAppSelector(selectSelectProject)!;
  const { loading, current, previous } = useAppSelector(selectProjectProcessing);
  const dispatch = useAppDispatch();

  const executeCommand = useCallback((command: string) => {
    console.log(`Execute command: ${command}`);
  }, []);

  const realCurrent = useMemo(() => {
    return (
      current ??
      ({
        id: '',
        event: '',
        group: 'project',
        project: project,
        timestamp: Date.now(),
      } as ProjectHistoryDto)
    );
  }, [current, project]);

  const actionCommand = useCallback((projectId: string, command: string) => {
    console.log(`Action command: ${command} on project ${projectId}`);
  }, []);

  useEffect(() => {
    if (loading && !!current && !!previous) {
      return;
    }
    dispatch(loadProjectHistory(project.id));
    return;
  }, [dispatch, loading, project.id, current, previous]);

  return (
    <div className={`flex flex-col items-stretch height-100 overflow-auto gap-4`}>
      <ProjectPanel project={project} onAction={executeCommand} />
      {(current || project) && <ProjectPipeline current={realCurrent} previous={previous} onExecute={actionCommand} />}
    </div>
  );
};
