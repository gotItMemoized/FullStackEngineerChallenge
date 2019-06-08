import React from 'react';
import Message from './message';
import { showIf } from '../render';

export default ({ review = {} }) => {
  if (!review.feedback) {
    return <div id="performance-reviews-loading">Loading...</div>;
  }

  const responses = review.feedback.map(fb => {
    return (
      <Message key={fb.id} className="message">
        {fb.message}
        <hr /> - {fb.reviewer.name} ({fb.reviewer.username}){' '}
        {showIf(
          fb.reviewer.id === review.user.key,
          <span className="tag is-info">Self Review</span>,
        )}
      </Message>
    );
  });
  const { user } = review;
  const title = showIf(!!user && !!user.label, () => (
    <h2 id="performance-review-response-title" className="is-size-5 title">
      Review of {user.label}
    </h2>
  ));
  return (
    <div>
      {title}
      <div>{responses}</div>
    </div>
  );
};
