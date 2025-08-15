import { MainContent, Sidebar } from '@blueskyfish/pierflow/components';
import { loadProjectList, useAppDispatch } from '@blueskyfish/pierflow/stores';
import * as React from 'react';
import { useEffect } from 'react';
import { Outlet } from 'react-router';

export const RootPage: React.FC = () => {
  const dispatch = useAppDispatch();

  useEffect(() => {
    dispatch(loadProjectList());
  }, [dispatch]);

  return (
    <div className={'flex flex-row align-items-stretch h-screen w-screen bg-base-100'}>
      <Sidebar />
      <MainContent>
        <Outlet />
      </MainContent>
    </div>
  );
};
