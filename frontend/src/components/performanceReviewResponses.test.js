import React from 'react';
import { render, cleanup } from '@testing-library/react';
import PerformanceReviewResponses from './performanceReviewResponses';
import 'jest-dom/extend-expect';

afterEach(cleanup);

it('should render loading', () => {
  const { container } = render(<PerformanceReviewResponses />);
  expect(container).toMatchSnapshot();
  const innerDiv = document.getElementById('performance-reviews-loading');
  expect(innerDiv).toMatchSnapshot();
  expect(innerDiv).toHaveTextContent('Loading');
});

it('should render not render loading', () => {
  const { container } = render(<PerformanceReviewResponses review={{ feedback: [] }} />);
  expect(container).toMatchSnapshot();
  const innerDiv = document.getElementById('performance-reviews-loading');
  expect(innerDiv).toEqual(null);
  const innerDiv2 = document.getElementById('performance-review-response-title');
  expect(innerDiv2).toEqual(null);
});

it('should render title', () => {
  const { container } = render(
    <PerformanceReviewResponses review={{ feedback: [], user: { label: 'hi' } }} />,
  );
  expect(container).toMatchSnapshot();
  const innerDiv = document.getElementById('performance-reviews-loading');
  expect(innerDiv).toEqual(null);
  const innerDiv2 = document.getElementById('performance-review-response-title');
  expect(innerDiv2).toHaveTextContent('Review of hi');
});

it('should render reviews', () => {
  const { container } = render(
    <PerformanceReviewResponses
      review={{
        feedback: [
          { id: '1', message: 'test', reviewer: { name: 'j', username: 'c', id: '1' } },
          { id: '2', message: 'test2', reviewer: { name: 'j', username: 'c', id: '2' } },
        ],
        user: { label: 'hi', key: '1' },
      }}
    />,
  );
  expect(container).toMatchSnapshot();
  const innerDiv = document.getElementById('performance-reviews-loading');
  expect(innerDiv).toEqual(null);
  const innerDiv2 = document.getElementById('performance-review-response-title');
  expect(innerDiv2).toHaveTextContent('Review of hi');
});
