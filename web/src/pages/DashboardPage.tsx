import { updatePageKey, useAppDispatch } from '@blueskyfish/pierflow/stores';
import * as React from 'react';
import { useEffect } from 'react';

export const DashboardPage: React.FC = () => {
  // Update pageKey to 'dashboard' when this component is mounted
  const dispatch = useAppDispatch();
  useEffect(() => {
    dispatch(updatePageKey('dashboard'));
  }, [dispatch]);
  return (
    <div className={'flex flex-col align-items-stretch height-100 overflow-auto'}>
      <h1>Dashboard</h1>
      <p>This is the dashboard page.</p>
    </div>
  );
};
