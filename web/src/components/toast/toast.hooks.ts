import { createContext, useContext } from 'react';
import type { ToastMessage } from './toast.models.ts';

interface ToastContextType {
  // Add a new toast
  add: (msg: Omit<ToastMessage, 'id'>) => void;
}

export const ToastContext = createContext<ToastContextType | undefined>(undefined);

export function useToast() {
  const ctx = useContext(ToastContext);
  if (!ctx) throw new Error('useToast must be used within a ToastProvider');
  return ctx;
}
