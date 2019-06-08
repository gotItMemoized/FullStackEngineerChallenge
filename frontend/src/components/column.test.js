import React from 'react';
import { render, cleanup } from '@testing-library/react';
import Column from './Column';
import 'jest-dom/extend-expect';

afterEach(cleanup);

it('renders without crashing', () => {
  const { container } = render(<Column />);
  expect(container).toMatchSnapshot();
});

it('renders with classname', () => {
  const { container } = render(<Column className="test" />);
  expect(container).toMatchSnapshot();
});

it('renders children', () => {
  const { container } = render(
    <Column>
      <div id="something">hello</div>
    </Column>,
  );
  expect(container).toMatchSnapshot();
  const innerDiv = document.getElementById('something');
  expect(innerDiv).toMatchSnapshot();
  expect(innerDiv).toHaveTextContent('hello');
});
