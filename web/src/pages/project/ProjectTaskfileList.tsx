import * as React from 'react';

export interface ProjectTaskfileListProps {
  taskFiles: string[];
  activeTaskfile: string;
  onBuild: (taskfile: string) => void;
}

export const ProjectTaskfileList: React.FC<ProjectTaskfileListProps> = ({ taskFiles, activeTaskfile, onBuild }) => {
  return (
    <table className={'table table-zebra'}>
      <colgroup>
        <col width={'5%'} />
        <col width={'*'} />
        <col width={'15%'} />
      </colgroup>
      <thead>
        <tr>
          <td className={'text-right'}>No</td>
          <th>Taskfile</th>
          <td>&nbsp;</td>
        </tr>
      </thead>
      <tbody>
        {taskFiles.map((taskfile, index) => (
          <tr key={index} className={`${activeTaskfile === taskfile ? 'bg-primary text-white' : 'hover:bg-base-300'}`}>
            <td className={'text-right'}>{index + 1}</td>
            <td>{taskfile}</td>
            <td>
              {activeTaskfile === taskfile && <span className={`text-sm`}>Current build with</span>}
              {activeTaskfile !== taskfile && (
                <button type={'button'} className={'btn btn-xs btn-soft btn-primary'} onClick={() => onBuild(taskfile)}>
                  Build
                  <span className={'ml-1 mdi mdi-chevron-right'}></span>
                </button>
              )}
            </td>
          </tr>
        ))}
      </tbody>
    </table>
  );
};
