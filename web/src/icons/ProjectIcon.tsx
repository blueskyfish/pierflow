import type { ReactNode } from 'react';
import * as React from 'react';

import builtIcon from './assets/built.svg';
import checkedOutIcon from './assets/checked-out.svg';
import clonedIcon from './assets/cloned.svg';
import createdIcon from './assets/created.svg';
import deletedIcon from './assets/deleted.svg';
import pulledIcon from './assets/pulled.svg';
import runIcon from './assets/run.svg';
import stoppedIcon from './assets/stopped.svg';
import taggedIcon from './assets/tagged.svg';

export type ProjectStatusKind =
  | 'created'
  | 'cloned'
  | 'checked-out'
  | 'built'
  | 'run'
  | 'pulled'
  | 'stopped'
  | 'deleted'
  | 'tagged';

export interface ProjectIconProps {
  kind: ProjectStatusKind;
}

const iconFactory: Record<ProjectStatusKind, string> = {
  created: createdIcon,
  cloned: clonedIcon,
  'checked-out': checkedOutIcon,
  built: builtIcon,
  run: runIcon,
  pulled: pulledIcon,
  stopped: stoppedIcon,
  deleted: deletedIcon,
  tagged: taggedIcon,
};

const getProjectIcon = (kind: ProjectStatusKind): ReactNode => {
  if (!iconFactory[kind]) {
    return (
      <span className='w-16 h-16 text-md rounded-full bg-red-200 text-red-700 border-1 border-red-700 flex flex-col items-center justify-center'>
        <span>??</span>
      </span>
    );
  }
  return <img src={iconFactory[kind]!} alt={kind} className='w-8 h-8' />;
};

export const ProjectIcon: React.FC<ProjectIconProps> = ({ kind }) => {
  return <div className='w-16 h-16 flex items-center justify-center'>{getProjectIcon(kind)}</div>;
};
