import React from 'react';
import { render } from '@testing-library/react';
import showIf from './showIf';
import 'jest-dom/extend-expect';

it('should not render on false', () => {
  const { container } = render(showIf(false, <div id="test">hello</div>));
  expect(container).toMatchSnapshot();
});

it('should render on true', () => {
  const { container } = render(showIf(true, <div id="test">hello</div>));
  expect(container).toMatchSnapshot();
  const innerDiv = document.getElementById('test');
  expect(innerDiv).toMatchSnapshot();
  expect(innerDiv).toHaveTextContent('hello');
});

it('should explode on not true', () => {
  const renderComponent = () => render(showIf({}, <div id="test">hello</div>));
  expect(renderComponent).toThrow('Expected a true or false statement');
});

it('should execute function if sent in', () => {
  const result = showIf(true, () => 'hello');
  expect(result).toEqual('hello');
});
