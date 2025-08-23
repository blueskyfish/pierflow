import { HeadLine } from '@blueskyfish/pierflow/components';
import { selectMessageList, updatePageKey, useAppDispatch, useAppSelector } from '@blueskyfish/pierflow/stores';
import * as React from 'react';
import { useEffect } from 'react';

export const SsePage: React.FC = () => {
  const messages = useAppSelector(selectMessageList);

  const dispatch = useAppDispatch();
  // Update the project key to Detail when this component is mounted
  useEffect(() => {
    dispatch(updatePageKey('sse'));
  }, [dispatch]);

  return (
    <div className={'flex flex-col align-items-stretch height-100 overflow-auto p-3'}>
      <HeadLine title={'SSE Test Pages'} icon={'mdi mdi-server'} className={'mb-4'} />
      <ul>
        {messages.map((msg, index) => (
          <li key={index}>{msg}</li>
        ))}
      </ul>
    </div>
  );
};
