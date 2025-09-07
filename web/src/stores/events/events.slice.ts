import { createSlice, type PayloadAction } from '@reduxjs/toolkit';
import type { ServerEvent } from './events.models.ts';
import type { ErrorDto } from '@blueskyfish/pierflow/api';

export type EventConnectStatus = 'disconnected' | 'connecting' | 'connected' | 'error';

export interface EventState {
  status: EventConnectStatus;
  messages: ServerEvent[];
  error: ErrorDto | null;
}

export const EventFeatureKey = 'events';

/**
 * Initial state for the event slice, excluding userId which is set dynamically.
 */
export const initialEventState: EventState = {
  status: 'disconnected',
  messages: [],
  error: null,
};

const eventStore = createSlice({
  name: EventFeatureKey,
  initialState: initialEventState,
  reducers: {
    updateStatus: (state, action) => {
      state.status = action.payload;
    },
    addMessage: (state, action: PayloadAction<ServerEvent>) => {
      state.messages.push(action.payload);
    },
    setError: (state, action: PayloadAction<ErrorDto | null>) => {
      state.error = action.payload;
    },
    clearMessages: (state) => {
      state.messages = [];
    },
  },
});

export const { updateStatus, addMessage, setError, clearMessages } = eventStore.actions;
export const eventReducer = eventStore.reducer;
