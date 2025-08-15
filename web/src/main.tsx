import { App } from '@blueskyfish/pierflow/app';
import { StoreProvider } from '@blueskyfish/pierflow/stores';
import { StrictMode } from 'react';
import { createRoot } from 'react-dom/client';
import { HashRouter as Router } from 'react-router';

import './main.css';

const strictMode = false;

const renderNode = strictMode ? (
  <StrictMode>
    <StoreProvider>
      <Router>
        <App />
      </Router>
    </StoreProvider>
  </StrictMode>
) : (
  <StoreProvider>
    <Router>
      <App />
    </Router>
  </StoreProvider>
);

createRoot(document.getElementById('root')!).render(renderNode);
