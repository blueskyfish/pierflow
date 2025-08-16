import { ProjectCommand } from '@blueskyfish/pierflow/stores';
import { ProjectPath, RouteBuilder } from '@blueskyfish/pierflow/utils';
import * as React from 'react';
import { Link, useNavigate } from 'react-router';

export interface ProjectDockProps {
  commandMap: Record<ProjectCommand, boolean>;
  selectKey: string;
  projectId: string;
}

const ProjectDockMaps: Record<ProjectCommand, { icon: string; label: string; pageKey: ProjectPath }> = {
  [ProjectCommand.CloneRepository]: { icon: 'mdi mdi-store-plus-outline', label: 'Clone', pageKey: ProjectPath.Clone },
  [ProjectCommand.CheckoutRepository]: {
    icon: 'mdi mdi-store-check-outline',
    label: 'Checkout',
    pageKey: ProjectPath.Checkout,
  },
  [ProjectCommand.BuildProject]: { icon: 'mdi mdi-cog-outline', label: 'Build Project', pageKey: ProjectPath.Build },
  [ProjectCommand.StartProject]: {
    icon: 'mdi mdi-play-circle-outline',
    label: 'Start Project',
    pageKey: ProjectPath.Start,
  },
  [ProjectCommand.StopProject]: {
    icon: 'mdi mdi-stop-circle-outline',
    label: 'Stop Project',
    pageKey: ProjectPath.Stop,
  },
  [ProjectCommand.PullRepository]: { icon: 'mdi mdi-refresh', label: 'Pull', pageKey: ProjectPath.Pull },
  [ProjectCommand.DeleteProject]: {
    icon: 'mdi mdi-delete-outline',
    label: 'Delete Project',
    pageKey: ProjectPath.Delete,
  },
  [ProjectCommand.CreateProject]: {
    icon: 'mdi mdi-plus-circle-outline',
    label: 'Create Project',
    pageKey: ProjectPath.Create,
  },
};

const ProjectCommandList: ProjectCommand[] = [
  ProjectCommand.CloneRepository,
  ProjectCommand.CheckoutRepository,
  ProjectCommand.BuildProject,
  ProjectCommand.StartProject,
  ProjectCommand.StopProject,
  ProjectCommand.PullRepository,
  ProjectCommand.DeleteProject,
];

export const ProjectDock: React.FC<ProjectDockProps> = ({ commandMap, projectId, selectKey }) => {
  const navigate = useNavigate();
  // TODO add Command is executing and disable all buttons

  const handleCommand = (command: ProjectCommand) => {
    const commandPath = ProjectDockMaps[command].pageKey;
    navigate(RouteBuilder.buildProjectCommandPath(projectId, commandPath));
  };

  return (
    <div className={'dock dock-md absolute bg-base-200 text-base-200-content'}>
      <Link
        to={RouteBuilder.buildProjectHomePath(projectId)}
        className={`${selectKey === 'detail' ? 'dock-active' : ''} dock-item`}
        key={'detail'}
      >
        <span className={'app-icon size-[1.2em] mdi mdi-information-outline'} />
        <span className={'app-title dock-label'}>Details</span>
      </Link>
      {ProjectCommandList.map((command) => {
        if (typeof commandMap[command] === 'undefined') {
          return null;
        }
        return (
          <button
            type={'button'}
            className={`${selectKey === ProjectDockMaps[command].pageKey ? 'dock-active' : ''} dock-item`}
            onClick={() => handleCommand(command)}
            key={ProjectDockMaps[command].pageKey}
            disabled={!commandMap[command]}
          >
            <span className={`app-icon size-[1.2em] ${ProjectDockMaps[command].icon}`} />
            <span className={'app-title dock-label'}>{ProjectDockMaps[command].label}</span>
          </button>
        );
      })}
    </div>
  );
};
