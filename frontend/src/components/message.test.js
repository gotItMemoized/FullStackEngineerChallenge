import React from 'react';
import { render, cleanup } from '@testing-library/react';
import Message from './Message';
import 'jest-dom/extend-expect';

afterEach(cleanup);

it('renders without crashing', () => {
  const { container } = render(<Message />);
  expect(container).toMatchSnapshot();
});

it('renders with classname', () => {
  const { container } = render(<Message className="test" />);
  expect(container).toMatchSnapshot();
});

it('renders children', () => {
  const { container } = render(
    <Message>
      <div id="something">hello</div>
    </Message>,
  );
  expect(container).toMatchSnapshot();
  const innerDiv = document.getElementById('something');
  expect(innerDiv).toMatchSnapshot();
  expect(innerDiv).toHaveTextContent('hello');
});
