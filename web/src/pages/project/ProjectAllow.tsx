import type { ProjectDto } from '@blueskyfish/pierflow/api';
import { ProjectCommand, updateProjectKey, useAppDispatch } from '@blueskyfish/pierflow/stores';
import { ProjectPath, RouteBuilder, RoutePath } from '@blueskyfish/pierflow/utils';
import * as React from 'react';
import { type PropsWithChildren, useEffect } from 'react';
import { Navigate, useParams } from 'react-router';

export interface ProjectAllowProps {
  command: ProjectCommand;
  project: ProjectDto | null;
  projectKey: ProjectPath;
}

/**
 * ProjectAllow component checks if the user has permission to access a specific project command.
 * If the user does not have permission, it redirects them to the project home page or home page.
 * @param command the project command to check permissions for
 * @param project the project to check permissions against
 * @param projectKey the key of the project to update in the layout state
 * @param children the children components to render if permission is granted
 * @constructor
 */
export const ProjectAllow: React.FC<PropsWithChildren<ProjectAllowProps>> = ({
  command,
  project,
  projectKey,
  children,
}) => {
  const projectId = useParams<{ projectId?: string }>().projectId;
  const dispatch = useAppDispatch();

  useEffect(() => {
    dispatch(updateProjectKey(projectKey));
  }, [dispatch, projectKey]);

  if (!project && projectId) {
    return <Navigate to={RouteBuilder.buildProjectHomePath(projectId)} />;
  } else if (!projectId || !project) {
    return <Navigate to={RoutePath.HomePath} />;
  }

  if (!project.commandMap[command]) {
    return <Navigate to={RouteBuilder.buildProjectHomePath(projectId)} />;
  }
  return <>{children}</>;
};
