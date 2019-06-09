import React, { useState, useEffect } from 'react';
import PerformanceReviewResponses from '../components/performanceReviewResponses';
import { get } from '../api';

// view submitted performance reviews
export default ({ reviewId }) => {
  const [review, setReview] = useState();

  useEffect(() => {
    const fetchData = async () => {
      const result = await get(`/review/${reviewId}`);
      const converted = result.data;
      converted.user = {
        value: converted.user.id,
        key: converted.user.id,
        label: `${converted.user.name} (${converted.user.username})`,
      };
      converted.feedback = converted.feedback.map(feedback => {
        return {
          ...feedback,
          message:
            feedback.message.String.length > 0
              ? feedback.message.String
              : 'User did not respond to this Performance Review',
        };
      });
      setReview(converted);
    };
    if (reviewId) {
      fetchData();
    }
  }, [reviewId]);

  return <PerformanceReviewResponses review={review} />;
};
