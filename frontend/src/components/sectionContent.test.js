import React from 'react';
import { render, cleanup } from '@testing-library/react';
import SectionContent from './SectionContent';
import 'jest-dom/extend-expect';

afterEach(cleanup);

it('renders without crashing', () => {
  const { container } = render(<SectionContent />);
  expect(container).toMatchSnapshot();
});

it('renders with classname', () => {
  const { container } = render(<SectionContent className="test" />);
  expect(container).toMatchSnapshot();
});

it('renders children', () => {
  const { container } = render(
    <SectionContent>
      <div id="something">hello</div>
    </SectionContent>,
  );
  expect(container).toMatchSnapshot();
  const innerDiv = document.getElementById('something');
  expect(innerDiv).toMatchSnapshot();
  expect(innerDiv).toHaveTextContent('hello');
});
