import type { PropsWithChildren } from 'react';
import * as React from 'react';

export interface CardImageProps {
  fillColor: string;
  height: number;
}

export const CardImage: React.FC<PropsWithChildren<CardImageProps>> = ({ fillColor, height, children }) => {
  return (
    <div className={'width-100'} style={{ height: `${height}px`, position: 'relative' }}>
      <svg width='100%' height={`${height}px`}>
        <rect width='100%' height={`${height}px`} fill={fillColor} />
      </svg>
      <div
        className={'flex flex-row align-items-center justify-center z-10'}
        style={{
          position: 'absolute',
          left: 0,
          top: 0,
          width: '100%',
          height: `${height}px`,
        }}
      >
        {children}
      </div>
    </div>
  );
};
