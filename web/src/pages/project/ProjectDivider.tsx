import React from 'react';

export const ProjectDivider: React.FC = () => {
  return (
    <div className={'w-[36px] flex flex-col flex-shrink-1 items-center justify-center px-1 pt-10'}>
      <span className={'mdi mdi-chevron-double-right mdi-36px text-gray-400'} />
    </div>
  );
};
