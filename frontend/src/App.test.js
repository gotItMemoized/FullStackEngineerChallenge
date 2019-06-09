import React from 'react';
import ReactDOM from 'react-dom';
import { render, cleanup } from '@testing-library/react';
import App from './App';

afterEach(cleanup);

it('renders without crashing and directs to login', () => {
  const div = document.createElement('div');
  ReactDOM.render(<App />, div);
  ReactDOM.unmountComponentAtNode(div);
  const { container } = render(<App />);
  expect(container).toMatchSnapshot();
  const usernameField = document.getElementById('username');
  expect(usernameField).toMatchSnapshot();
  const passwordField = document.getElementById('password');
  expect(passwordField).toMatchSnapshot();
});
