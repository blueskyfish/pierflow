import { selectPageKey, selectProjectList, useAppSelector } from '@blueskyfish/pierflow/stores';
import * as React from 'react';
import { BrandHead } from '../brand';
import { SidebarMenu } from './SidebarMenu.tsx';

export const Sidebar: React.FC = () => {
  const list = useAppSelector(selectProjectList);
  const pageKey = useAppSelector(selectPageKey);
  return (
    <aside className={'bg-base-200 w-1/5 flex-shrink-1 py-2 px-3 border-r border-base-300'}>
      <BrandHead />
      <div className={'card card-border border-base-300 bg-base-100 overflow-y-auto'}>
        <SidebarMenu menu={list} selected={pageKey} />
      </div>
    </aside>
  );
};
