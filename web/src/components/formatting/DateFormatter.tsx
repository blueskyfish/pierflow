import { DateTime } from 'luxon';
import * as React from 'react';
import { useMemo } from 'react';

export interface DateTimeProps {
  date: string;
  format?: string;
}

/**
 * DateFormatter component formats a date string using Luxon's DateTime.
 * It takes a date string and an optional format string as props.
 * If no format is provided, it defaults to 'dd.LL.yyyy HH:mm:ss
 * @param date the date string to format, in ISO format
 * @param format the optional format string, using Luxon's DateTime formatting
 * @returns a React component that displays the formatted date
 */
export const DateFormatter: React.FC<DateTimeProps> = ({ date, format }) => {
  const value = useMemo(() => {
    return DateTime.fromISO(date).toFormat(format ?? 'dd.LL.yyyy HH:mm:ss');
  }, [date, format]);
  return <span className={'text-base-content/60'}>{value}</span>;
};
