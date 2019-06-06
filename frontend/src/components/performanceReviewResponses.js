import React from 'react';
import Message from './message';
import { showIf } from '../render';

export default ({ review = {} }) => {
  if (!review.feedback) {
    return <div>Loading...</div>;
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
  return (
    <div>
      <h2 className="is-size-5 title">Review of {review.user.label}</h2>
      <div>{responses}</div>
    </div>
  );
};
