import React from 'react';
import { render, cleanup } from '@testing-library/react';
import { MemoryRouter } from 'react-router-dom';
import Header from './Header';
import 'jest-dom/extend-expect';

afterEach(cleanup);

it('renders without crashing', () => {
  const { container } = render(
    <MemoryRouter>
      <Header />
    </MemoryRouter>,
  );
  expect(container).toMatchSnapshot();
});

it('should not nav buttons if not logged in', () => {
  const { container } = render(
    <MemoryRouter>
      <Header />
    </MemoryRouter>,
  );
  expect(container).toMatchSnapshot();
  const innerDiv = document.getElementById('logout-btn');
  expect(innerDiv).toEqual(null);
  const innerDiv2 = document.getElementById('performance-menu-btn');
  expect(innerDiv2).toEqual(null);
});

it('should have some buttons if logged in as regular user', () => {
  const { container } = render(
    <MemoryRouter>
      <Header currentUser={{ loggedIn: true }} />
    </MemoryRouter>,
  );
  expect(container).toMatchSnapshot();
  const innerDiv = document.getElementById('logout-btn');
  expect(innerDiv).toHaveTextContent('Log out');
  const innerDiv2 = document.getElementById('performance-menu-btn');
  expect(innerDiv2).toHaveTextContent('Performance Reviews');
  const innerDiv3 = document.getElementById('user-menu-btn');
  expect(innerDiv3).toEqual(null);
});

it('should have more buttons if logged in as admin user', () => {
  const { container } = render(
    <MemoryRouter>
      <Header currentUser={{ loggedIn: true, isAdmin: true }} />
    </MemoryRouter>,
  );
  expect(container).toMatchSnapshot();
  const innerDiv = document.getElementById('logout-btn');
  expect(innerDiv).toHaveTextContent('Log out');
  const innerDiv2 = document.getElementById('performance-menu-btn');
  expect(innerDiv2).toHaveTextContent('Performance Reviews');
  const innerDiv3 = document.getElementById('users-menu-btn');
  expect(innerDiv3).toHaveTextContent('Users');
});
