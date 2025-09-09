import { useToast } from '@blueskyfish/pierflow/components';
import { updatePageKey, useAppDispatch } from '@blueskyfish/pierflow/stores';
import * as React from 'react';
import { useEffect } from 'react';

export const DashboardPage: React.FC = () => {
  const toaster = useToast();

  const addToast = () => {
    toaster.add({ state: 'info', title: 'Info', message: 'This is an info toast message.' });
  };

  const addSuccessToast = () => {
    toaster.add({ state: 'success', title: 'Success', message: 'This is a success toast message.', timeout: 20_000 });
  };

  const addErrorToast = () => {
    toaster.add({ state: 'error', title: 'Error', message: 'This is an error toast message.' });
  };

  // Update pageKey to 'dashboard' when this component is mounted
  const dispatch = useAppDispatch();
  useEffect(() => {
    dispatch(updatePageKey('dashboard'));
  }, [dispatch]);
  return (
    <div className={'flex flex-col align-items-stretch height-100 overflow-auto'}>
      <h1>Dashboard</h1>
      <p>This is the dashboard page.</p>
      <button type='button' className={'btn btn-soft'} onClick={addToast}>
        Add Toast
      </button>
      <button type='button' className={'btn btn-soft'} onClick={addSuccessToast}>
        Success
      </button>
      <button type='button' className={'btn btn-soft'} onClick={addErrorToast}>
        Error
      </button>
    </div>
  );
};
