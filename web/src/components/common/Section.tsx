import React, { type PropsWithChildren } from 'react';

export interface PanelProps {
  className?: string;
}

export const Section: React.FC<PropsWithChildren<PanelProps>> = ({ className, children }) => {
  return <section className={`rounded border border-gray-300 bg-white ${className ?? ''}`}>{children}</section>;
};
