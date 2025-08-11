import { App } from '@blueskyfish/pierflow/app';
import { StrictMode } from 'react';
import { createRoot } from 'react-dom/client';

import './main.css';

createRoot(document.getElementById('root')!).render(
  <StrictMode>
    <App />
  </StrictMode>,
);
