import type { ServerEvent } from './events.models.ts';
import { ProjectCommand } from '@blueskyfish/pierflow/stores';

export const addEventMessager = (es: EventSource, eventType: string, handler: (event: ServerEvent) => void) => {
  const serverEventHandler = (ev: MessageEvent) => {
    const sEvent = JSON.parse(ev.data) as ServerEvent;
    if (!sEvent) {
      return;
    }
    handler(sEvent);
  };

  es.addEventListener(eventType, serverEventHandler);
  return () => {
    es.removeEventListener(eventType, serverEventHandler);
  };
};

/**
 * Convert a ProjectCommand to the corresponding event type string.
 * @param command ProjectCommand
 * @returns event type string
 */
export const toEventType = (command: ProjectCommand): string => {
  return `message-${command}`;
};
