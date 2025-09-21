import { DateTime } from 'luxon';
import React, { useMemo } from 'react';
import { iconMap, textMap } from './Project.utils.ts';

export type ProjectHistoryItem = {
  event: string;
  project: {
    status: string;
  };
  timestamp: number;
};

export interface ProjectStateProps {
  item: ProjectHistoryItem;
}

export const ProjectHistoryPanel: React.FC<ProjectStateProps> = ({ item }) => {
  const iconClassName = useMemo(() => {
    return iconMap[item.project.status] || iconMap['unknown'];
  }, [item.project.status]);
  const textState = useMemo(() => {
    return textMap[item.project.status] || textMap['unknown'];
  }, [item.project.status]);
  const formattedDate = useMemo(() => {
    return DateTime.fromSeconds(item.timestamp).toFormat('dd.LL.yyyy HH:mm:ss');
  }, [item.timestamp]);
  return (
    <dl className='grid grid-cols-[1fr_10fr] gap-x-4 gap-y-2 items-center text-sm'>
      <dt className='font-medium text-gray-700 flex items-center'>
        <span className={iconClassName} />
      </dt>
      <dd className='text-gray-900'>{textState}</dd>
      <dt className='font-medium text-gray-700 flex items-center'>
        <span className={'mdi mdi-calendar-check-outline'} />
      </dt>
      <dd className='text-gray-900'>{item.event}</dd>
      <dt className='font-medium text-gray-700 flex items-center'>
        <span className={'mdi mdi-clock-outline'} />
      </dt>
      <dd className='text-gray-500'>{formattedDate}</dd>
    </dl>
  );
};
