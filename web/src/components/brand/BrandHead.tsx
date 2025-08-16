import * as React from 'react';
import { BrandLogo } from './BrandLogo.tsx';
import { BrandName } from './BrandName.tsx';

export const BrandHead: React.FC = () => {
  return (
    <div className={'flex flex-row align-items-center px-2 mb-3'}>
      <BrandLogo />
      <BrandName />
    </div>
  );
};
