import React from 'react';
import { render, cleanup } from '@testing-library/react';
import Columns from './Columns';
import 'jest-dom/extend-expect';

afterEach(cleanup);

it('renders without crashing', () => {
  const { container } = render(<Columns />);
  expect(container).toMatchSnapshot();
});

it('renders with classname', () => {
  const { container } = render(<Columns className="test" />);
  expect(container).toMatchSnapshot();
});

it('renders children', () => {
  const { container } = render(
    <Columns>
      <div id="something">hello</div>
    </Columns>,
  );
  expect(container).toMatchSnapshot();
  const innerDiv = document.getElementById('something');
  expect(innerDiv).toMatchSnapshot();
  expect(innerDiv).toHaveTextContent('hello');
});
