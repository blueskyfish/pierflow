import * as React from 'react';

// Import the logo image
import brandLogo from './BrandLogo.64x64.png';

export const BrandLogo: React.FC = () => {
  return (
    <img src={brandLogo} alt={'Pierflow Logo'} className={'m-0 flex-shrink-1 rounded-sm'} width={36} height={36} />
  );
};
