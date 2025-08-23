import { createSlice } from '@reduxjs/toolkit';

export type EventStatus = 'disconnected' | 'connecting' | 'connected' | 'error';

export interface EventState {
  status: EventStatus;
  messages: string[];
  error: string | null;
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
    addMessage: (state, action) => {
      state.messages.push(action.payload);
    },
    setError: (state, action) => {
      state.error = action.payload;
    },
    clearMessages: (state) => {
      state.messages = [];
    },
  },
});

export const { updateStatus, addMessage, setError, clearMessages } = eventStore.actions;
export const eventReducer = eventStore.reducer;
