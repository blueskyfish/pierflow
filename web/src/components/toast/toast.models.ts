export type ToastStatus = 'info' | 'warn' | 'success' | 'error';

export interface ToastMessage {
  id: string;
  title: string;
  message: string;
  state: ToastStatus;
  timeout?: number;
}

export const ToastStyle: Record<ToastStatus, string> = {
  info: 'alert-info',
  warn: 'alert-warning',
  success: 'alert-success',
  error: 'alert-error',
};

export const ToastIcon: Record<ToastStatus, string> = {
  info: 'mdi-information-outline',
  warn: 'mdi-alert-circle-outline',
  success: 'mdi-check-circle-outline',
  error: 'mdi-alert-circle',
};
