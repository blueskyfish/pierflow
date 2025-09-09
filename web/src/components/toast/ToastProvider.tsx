import React, { useCallback, useState } from 'react';
import type { ToastMessage } from './toast.models';
import { ToastContext } from './toast.hooks.ts';
import { ToastContainer } from './ToastContainer.tsx';

export const ToastProvider: React.FC<{ children: React.ReactNode }> = ({ children }) => {
  const [toasts, setToasts] = useState<ToastMessage[]>([]);

  const add = useCallback((msg: Omit<ToastMessage, 'id'>) => {
    const id = crypto.randomUUID() as string;
    setToasts((prev) => [...prev, { id, ...msg }]);
  }, []);

  const remove = useCallback((id: string) => {
    setToasts((prev) => prev.filter((t) => t.id !== id));
  }, []);

  return (
    <ToastContext.Provider value={{ add }}>
      {children}
      <ToastContainer toasts={toasts} onRemove={remove} />
    </ToastContext.Provider>
  );
};
