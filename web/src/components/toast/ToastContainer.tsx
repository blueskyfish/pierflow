import React from 'react';
import type { ToastMessage } from './toast.models.ts';
import { Toast } from './Toast.tsx';

export interface ToastContainerProps {
  toasts: ToastMessage[];
  onRemove: (id: string) => void;
}

export const ToastContainer: React.FC<ToastContainerProps> = ({ toasts, onRemove }) => {
  return (
    <div className={'toast'}>
      {(toasts ?? []).map((toast) => (
        <Toast key={toast.id} toast={toast} onRemove={onRemove} />
      ))}
    </div>
  );
};
