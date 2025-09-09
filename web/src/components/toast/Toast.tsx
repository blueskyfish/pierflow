import React, { useCallback, useEffect } from 'react';
import { ToastIcon, type ToastMessage, ToastStyle } from './toast.models';

export interface ToastProps {
  toast: ToastMessage;
  onRemove: (id: string) => void;
}

export const Toast: React.FC<ToastProps> = ({ toast, onRemove }) => {
  const [cleanupTimer, setCleanupTimer] = React.useState<any | null>(null);
  const closeToast = useCallback(
    (id: string) => {
      onRemove(id);
      clearTimeout(cleanupTimer);
      setCleanupTimer(null);
    },
    [onRemove, cleanupTimer],
  );

  useEffect(() => {
    const clearTimer = setTimeout(() => {
      clearTimeout(clearTimer);
      onRemove(toast.id);
    }, toast.timeout ?? 2_000);

    setCleanupTimer(clearTimer);
  }, [toast, onRemove]);
  return (
    <div key={toast.id} className={`alert ${ToastStyle[toast.state]} shadow flex flex-col items-stretch`}>
      <h4 className={`text-lg flex items-center gap-2`}>
        <span className={`mr-2 mdi ${ToastIcon[toast.state]}`} />
        <span className={'flex-grow-1'}>{toast.title}</span>
        <button className='btn btn-sm btn-circle btn-ghost absolute right-2 top-2' onClick={() => closeToast(toast.id)}>
          ✕
        </button>
      </h4>
      <p className={`text-${toast.state}-content`}>{toast.message}</p>
    </div>
  );
};
