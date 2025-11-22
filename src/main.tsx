import { StrictMode } from 'react';
import { createRoot } from 'react-dom/client';
import './styles/main.css';

createRoot(document.getElementById('root') as HTMLElement).render(
  <StrictMode>
    <h1>Hello, World</h1>
  </StrictMode>,
);
