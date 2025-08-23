import { createSelector } from '@reduxjs/toolkit';
import { EventFeatureKey, type EventState } from './events.slice.ts';

export const selectEventState = (state: { [EventFeatureKey]: EventState }) => state[EventFeatureKey];
export const selectEvent = createSelector([selectEventState], (eventState: EventState) => eventState);
export const selectMessageList = createSelector([selectEventState], (eventState: EventState) => eventState.messages);
