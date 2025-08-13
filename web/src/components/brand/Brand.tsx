import { BrandLogo } from './BrandLogo.tsx';
import { BrandName } from './BrandName.tsx';

export const Brand: React.FC = () => {
  return (
    <div className={'flex flex-row align-items-center p-3 mb-3'}>
      <BrandLogo />
      <BrandName />
    </div>
  );
};
