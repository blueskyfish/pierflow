import type { PropsWithChildren } from 'react';
import * as React from 'react';

export type HeadAs = 'h1' | 'h2' | 'h3' | 'h4' | 'h5' | 'h6';

export interface HeadLineProps {
  as?: HeadAs;
  className?: string;
  title: string;
  icon?: string;
}

const Caption: React.FC<PropsWithChildren<{ as: HeadAs; className: string }>> = ({ as, className, children }) => {
  switch (as) {
    case 'h1':
      return <h1 className={`${className} text-2xl font-bold flex flex-row align-items-center`}>{children}</h1>;
    case 'h2':
      return <h2 className={`${className} text-xl font-semibold flex flex-row align-items-center`}>{children}</h2>;
    case 'h3':
      return <h3 className={`${className} text-lg font-medium flex flex-row align-items-center`}>{children}</h3>;
    case 'h4':
      return <h4 className={`${className} text-base font-normal flex flex-row align-items-center`}>{children}</h4>;
    case 'h5':
      return <h5 className={`${className} text-sm font-light flex flex-row align-items-center`}>{children}</h5>;
    case 'h6':
      return <h6 className={`${className} text-xs font-thin flex flex-row align-items-center`}>{children}</h6>;
    default:
      return <span className={'flex flex-row align-items-center'}>{children}</span>;
  }
};

export const HeadLine: React.FC<HeadLineProps> = ({ as, icon, title, className }) => {
  return (
    <Caption as={as ?? 'h2'} className={className ?? 'mb-2'}>
      {icon && <span className={`app-icon ${icon} mr-2 flex-shrink-1`}></span>}
      <span className='app-title flex-grow-1 text-truncate'>{title}</span>
    </Caption>
  );
};
