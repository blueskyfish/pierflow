import * as React from 'react';
import { useMemo } from 'react';

export interface MessageLabelProps {
  status: string;
}

export const MessageLabel: React.FC<MessageLabelProps> = ({ status }) => {
  const badgeColor = useMemo(() => {
    switch (status.toLowerCase()) {
      case 'info':
        return 'badge-info';
      case 'success':
        return 'badge-success';
      case 'warning':
        return 'badge-warning';
      case 'error':
        return 'badge-error';
      default:
        return 'badge-neutral';
    }
  }, [status]);
  return (
    <div className='text-center'>
      <span className={`badge badge-soft ${badgeColor}`}>{status}</span>
    </div>
  );
};
