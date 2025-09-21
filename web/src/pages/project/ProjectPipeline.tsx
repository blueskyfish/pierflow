import { ProjectCommand, type ProjectHistoryDto } from '@blueskyfish/pierflow/api';
import React, { useCallback, useMemo } from 'react';
import { Section } from '../../components/common';
import { Fieldset } from '../../components/common/Fieldset.tsx';
import { ProjectButton } from './ProjectButton.tsx';
import { ProjectDivider } from './ProjectDivider.tsx';
import { ProjectHistoryPanel } from './ProjectHistoryPanel.tsx';

export interface ProjectHistoryProps {
  current: ProjectHistoryDto;
  previous?: ProjectHistoryDto;
  onExecute: (projectId: string, command: ProjectCommand) => void;
}

const COMMAND_LIST = [
  ProjectCommand.CloneRepository,
  ProjectCommand.CheckoutRepository,
  ProjectCommand.PullRepository,
  ProjectCommand.BuildProject,
  ProjectCommand.StartProject,
  ProjectCommand.StopProject,
  ProjectCommand.DeleteProject,
];

export const ProjectPipeline: React.FC<ProjectHistoryProps> = ({ current, previous, onExecute }) => {
  // Build the command list from current
  const commandList = useMemo(() => {
    return COMMAND_LIST.filter((cmd) => {
      return current.project.commandMap[cmd] ?? false;
    });
  }, [current]);

  const commandExecute = useCallback(
    (command: ProjectCommand) => {
      onExecute(current.project.id, command);
    },
    [current.project.id, onExecute],
  );

  return (
    <Section className={'flex flex-row items-stretch flex-grow-1 mx-4 mb-4 p-0 gap-0'}>
      <Fieldset title={'Previous Status'} titleAlign={'center'}>
        <div className={`h-full flex flex-col items-center justify-center pag-2`}>
          {previous ? (
            <>
              <ProjectHistoryPanel item={previous} />
            </>
          ) : (
            <span className={'text-gray-400'}>
              No previous
              <br />
              status available
            </span>
          )}
        </div>
      </Fieldset>
      <ProjectDivider />
      <Fieldset title={'Current Status'} titleAlign={'center'}>
        <div className={`h-full flex flex-col items-center justify-center pag-2`}>
          <ProjectHistoryPanel item={current} />
        </div>
      </Fieldset>
      <ProjectDivider />
      <Fieldset title={'Next Action'} titleAlign={'center'}>
        <div className={`h-full flex flex-col items-center justify-center gap-2`}>
          <div className={'flex flex-col gap-2 items-stretch'}>
            {commandList.map((cmd) => (
              <ProjectButton key={cmd} command={cmd} onAction={() => commandExecute(cmd)} />
            ))}
          </div>
        </div>
      </Fieldset>
    </Section>
  );
};
