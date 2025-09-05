import { useAppDispatch } from '@blueskyfish/pierflow/stores';
import * as React from 'react';
import { createContext, type PropsWithChildren, useEffect, useRef, useState } from 'react';
import { addMessage, setError, updateStatus } from './events.slice';
import { type ServerEvent } from './events.models.ts';

const PROJECT_CONNECT_PATH = '/api/projects/connect/{userId}';

export interface EventsContextValue {
  eventSource: EventSource | null;
}

// eslint-disable-next-line react-refresh/only-export-components
export const EventsContext = createContext<EventsContextValue | undefined>(undefined);

/**
 * Provider component to manage and provide EventSource connection via context.
 *
 * @param children
 */
export const EventsProvider: React.FC<PropsWithChildren> = ({ children }) => {
  const [eventSource, setEventSource] = useState<EventSource | null>(null);
  const userRef = useRef<string>(null);

  const dispatch = useAppDispatch();

  useEffect(() => {
    if (!userRef.current) {
      let userId = localStorage.getItem('blueskyfish.pierflow.userId');
      if (!userId) {
        userId = crypto.randomUUID();
        localStorage.setItem('blueskyfish.pierflow.userId', userId);
      }
      userRef.current = userId;
    }
  }, [userRef]);

  useEffect(() => {
    // timer handle for reconnection attempts
    let reconnectTimer: number | undefined = undefined;

    const connect = () => {
      if (reconnectTimer) {
        clearTimeout(reconnectTimer);
        reconnectTimer = undefined;
      }

      if (!eventSource && !!userRef.current) {
        dispatch(updateStatus('connecting'));
        dispatch(setError(null));
        const es = new EventSource(PROJECT_CONNECT_PATH.replace('{userId}', encodeURIComponent(userRef.current)));
        es.onopen = () => {
          dispatch(updateStatus('connected'));
          dispatch(setError(null));
        };
        es.onerror = (error) => {
          console.error('EventSource failed:', error);
          dispatch(updateStatus('error'));
          dispatch(setError('EventSource failed'));
          es.close();
          setEventSource(null);

          // eslint-disable-next-line @typescript-eslint/ban-ts-comment
          // @ts-expect-error
          reconnectTimer = setTimeout(() => connect(), 5_000);
        };

        es.addEventListener('message', (event) => {
          const sEvent = JSON.parse(event.data) as ServerEvent;
          dispatch(addMessage(sEvent));
        });

        es.addEventListener('heartbeat', (event) => {
          console.log('Received heartbeat =>', event.data);
        });

        // store the event source in state
        setEventSource(es);
      }
    };

    // connect initially
    connect();

    // Cleanup on unmount
    return () => {
      if (reconnectTimer) {
        clearTimeout(reconnectTimer);
      }
      if (eventSource) {
        eventSource.close();
        setEventSource(null);
      }
    };
  }, [dispatch, eventSource, userRef]);

  const contextValue: EventsContextValue = {
    eventSource,
  };

  return <EventsContext.Provider value={contextValue}>{children}</EventsContext.Provider>;
};
