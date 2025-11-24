import { StrictMode } from 'react';
import { createRoot } from 'react-dom/client';
import { App } from './app';

import '@fontsource/inter';
import '@/styles/main.css';

createRoot(document.getElementById('root') as HTMLElement).render(
  <StrictMode>
    <App />
  </StrictMode>,
);
