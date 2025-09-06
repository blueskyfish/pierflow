import * as React from 'react';
import { useMemo } from 'react';

export interface GitPlaceProps {
  place: string;
}

export const GitPlace: React.FC<GitPlaceProps> = ({ place }) => {
  const placeText = useMemo(() => {
    switch (place) {
      case 'remote':
        return 'Remote';
      case 'local':
        return 'Local';
      default:
        return 'Unknown';
    }
  }, [place]);
  return (
    <div className={'text-center'}>
      <span className={'badge badge-soft badge-info'}>{placeText}</span>
    </div>
  );
};
