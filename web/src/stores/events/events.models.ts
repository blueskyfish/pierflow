export enum EventStatus {
  Debug = 'debug',
  Info = 'info',
  Warning = 'warning',
  Error = 'error',
  Success = 'success',
}

export interface ServerEvent {
  /**
   * Event status level
   */
  status: EventStatus;

  /**
   * Event message or JSON serialized object
   */
  message: string;

  /**
   * Project unique identifier
   */
  id: string;

  /**
   * ISO 8601 formatted timestamp of the event
   */
  time: string;
}
