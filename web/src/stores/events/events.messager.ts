import type { ServerEvent } from './events.models.ts';
import { ProjectCommand } from '@blueskyfish/pierflow/stores';

export type CleanupFunction = () => void;

/**
 * Add an event listener to the EventSource for a specific event type.
 * The handler will be called with the parsed ServerEvent when an event is received.
 * Returns a cleanup function to remove the event listener.
 * @param es EventSource
 * @param eventType string
 * @param handler (event: ServerEvent) => void
 * @returns CleanupFunction
 */
export const addEventMessager = (
  es: EventSource,
  eventType: string,
  handler: (event: ServerEvent) => void,
): CleanupFunction => {
  // Wraps the handler to parse the event data
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
