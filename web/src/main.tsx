import { App } from '@blueskyfish/pierflow/app';
import { ToastProvider } from '@blueskyfish/pierflow/components';
import { EventsProvider, StoreProvider } from '@blueskyfish/pierflow/stores';
import { StrictMode } from 'react';
import { HashRouter as Router } from 'react-router';
import { createRoot } from 'react-dom/client';

import './main.css';

const strictMode = false;

const renderNode = strictMode ? (
  <StrictMode>
    <ToastProvider>
      <StoreProvider>
        <Router>
          <App />
        </Router>
      </StoreProvider>
    </ToastProvider>
  </StrictMode>
) : (
  <ToastProvider>
    <StoreProvider>
      <EventsProvider>
        <Router>
          <App />
        </Router>
      </EventsProvider>
    </StoreProvider>
  </ToastProvider>
);

createRoot(document.getElementById('root')!).render(renderNode);
