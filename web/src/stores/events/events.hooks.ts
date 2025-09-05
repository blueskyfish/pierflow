import { EventsContext } from './events.provider';
import { useContext } from 'react';

export const useEventSource = () => {
  const { eventSource } = useContext(EventsContext)!;
  return eventSource!;
};
