import { ProjectCommand } from '@blueskyfish/pierflow/api';
import React, { useMemo } from 'react';
import { commandIconMap, commandTextMap } from './Project.utils.ts';

export interface ProjectButtonProps {
  command: ProjectCommand;
  onAction: () => void;
}

export const ProjectButton: React.FC<ProjectButtonProps> = ({ command, onAction }) => {
  const iconClassName = useMemo(() => {
    return commandIconMap[command];
  }, [command]);

  const title = useMemo(() => {
    return commandTextMap[command];
  }, [command]);

  const buttonVariant = useMemo(() => {
    switch (command) {
      case ProjectCommand.DeleteProject:
        return 'error';
      case ProjectCommand.StopProject:
        return 'warning';
      case ProjectCommand.StartProject:
      case ProjectCommand.PullRepository:
      case ProjectCommand.CheckoutRepository:
      case ProjectCommand.BuildProject:
        return 'secondary';
      default:
        return 'neutral';
    }
  }, [command]);

  return (
    <button
      type={'button'}
      className={`btn-${buttonVariant} btn btn-soft btn-sm flex flex-row items-center`}
      onClick={() => onAction()}
    >
      <span className={`${iconClassName} flex-shrink-1`} />
      <span className={`ml-2 flex-grow-1 truncate text-left`}>{title}</span>
      <span className={'mdi mdi-chevron-right flex-shrink-1'} />
    </button>
  );
};
