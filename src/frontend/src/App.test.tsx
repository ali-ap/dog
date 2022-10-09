import React from 'react';
import { render, screen } from '@testing-library/react';
import App from './App';
import { HashRouter } from 'react-router-dom';

test('renders learn react link', () => {
  render(
    <HashRouter>
      <React.StrictMode>
        <App />
      </React.StrictMode>
    </HashRouter>
  );
  const element = screen.getAllByText('Dog')
  expect(element.length).toBeGreaterThan(0)
});
