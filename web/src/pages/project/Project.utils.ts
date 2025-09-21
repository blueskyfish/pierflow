import { ProjectCommand } from '@blueskyfish/pierflow/api';

export const iconMap: Record<string, string> = {
  created: 'mdi mdi-moon-new',
  cloned: 'mdi mdi-git',
  'checked-out': 'mdi mdi-source-branch',
  pulled: 'mdi mdi-source-pull',
  built: 'mdi mdi-cog-outline',
  run: 'mdi mdi-play-circle-outline',
  stopped: 'mdi mdi-stop-circle-outline',
  deleted: 'mdi mdi-delete-outline',
  unknown: 'mdi mdi-help-circle-outline',
};

export const textMap: Record<string, string> = {
  created: 'Created',
  cloned: 'Cloned',
  'checked-out': 'Checked Out',
  pulled: 'Pulled',
  built: 'Built',
  run: 'Running',
  stopped: 'Stopped',
  deleted: 'Deleted',
  unknown: 'Unknown',
};

export const commandIconMap: Record<ProjectCommand, string> = {
  [ProjectCommand.CreateProject]: 'mdi mdi-new-box',
  [ProjectCommand.CloneRepository]: 'mdi mdi-git',
  [ProjectCommand.CheckoutRepository]: 'mdi mdi-source-branch',
  [ProjectCommand.PullRepository]: 'mdi mdi-source-pull',
  [ProjectCommand.BuildProject]: 'mdi mdi-cog-outline',
  [ProjectCommand.StartProject]: 'mdi mdi-play-circle-outline',
  [ProjectCommand.StopProject]: 'mdi mdi-stop-circle-outline',
  [ProjectCommand.DeleteProject]: 'mdi mdi-delete-outline',

  [ProjectCommand.BranchList]: 'mdi mdi-format-list-bulleted',
  [ProjectCommand.TaskList]: 'mdi mdi-format-list-bulleted-type',
};

export const commandTextMap: Record<ProjectCommand, string> = {
  [ProjectCommand.CreateProject]: 'Create',
  [ProjectCommand.CloneRepository]: 'Clone',
  [ProjectCommand.CheckoutRepository]: 'Checkout',
  [ProjectCommand.PullRepository]: 'Pull',
  [ProjectCommand.BuildProject]: 'Build',
  [ProjectCommand.StartProject]: 'Start',
  [ProjectCommand.StopProject]: 'Stop',
  [ProjectCommand.DeleteProject]: 'Delete',

  [ProjectCommand.BranchList]: 'Branches',
  [ProjectCommand.TaskList]: 'Tasks',
};
