import * as React from 'react';
import { useMemo } from 'react';

export interface GitPlaceProps {
  place: number;
}

export const GitPlace: React.FC<GitPlaceProps> = ({ place }) => {
  const placeText = useMemo(() => {
    switch (place) {
      case 1:
        return 'Remote';
      case 0:
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
