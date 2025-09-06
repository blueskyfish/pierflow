import * as React from 'react';
import { selectMessageList, useAppSelector } from '@blueskyfish/pierflow/stores';
import { useMemo } from 'react';
import { MessageLabel } from './MessageLabel.tsx';

export interface ProjectMessageProps {
  filterId: string;
}

const MAX_MESSAGES = 10;

export const ProjectMessage: React.FC<ProjectMessageProps> = ({ filterId }) => {
  const messageList = useAppSelector(selectMessageList);

  const filteredMessageList = useMemo(() => {
    const list = messageList.filter((message) => message.id === filterId);
    if (list.length > MAX_MESSAGES) {
      return list.slice(-MAX_MESSAGES).reverse();
    }
    return list.reverse();
  }, [messageList, filterId]);

  const hasMessages = useMemo(() => {
    return filteredMessageList.length > 0;
  }, [filteredMessageList]);

  const messageOffset = useMemo(() => {
    const count = messageList.length;
    if (count > MAX_MESSAGES) {
      return count - MAX_MESSAGES;
    }
    return 0;
  }, [messageList]);

  return (
    <div className={'overflow-x-auto !h-48 bg-base-200 mb-13 border-t border-t-solid border-t-base-300'}>
      <table className={'table table-xs table-zebra table-pin-rows'}>
        <thead>
          <tr className={'bg-base-200'}>
            <th>No</th>
            <th>
              <div className={'text-center'}>Status</div>
            </th>
            <td>Message</td>
            <td>Time</td>
          </tr>
        </thead>
        <tbody>
          {!hasMessages && (
            <tr>
              <td colSpan={4}>No Data</td>
            </tr>
          )}
          {hasMessages &&
            filteredMessageList.map((message, index, list) => (
              <tr key={index}>
                <td>{messageOffset + (list.length - index)}</td>
                <td>
                  <MessageLabel status={message.status} />
                </td>
                <td>{message.message}</td>
                <td>{message.time}</td>
              </tr>
            ))}
        </tbody>
      </table>
    </div>
  );
};
