import React, { useState, useEffect } from 'react';
import PerformanceReviewForm from '../components/performanceReviewForm';
import { Redirect } from 'react-router-dom';
import { get, put } from '../api';

// submit a performance review
export default ({ reviewId }) => {
  const [review, setReview] = useState();
  const [submitted, setSubmitted] = useState(false);
  const [error, setError] = useState();

  useEffect(() => {
    const fetchData = async () => {
      const result = await get(`/feedback/${reviewId}`);
      const converted = {
        ...result.data,
        message: result.data.message.String,
      };
      setReview(converted);
    };
    if (reviewId) {
      fetchData();
    }
  }, [reviewId]);

  const submit = reviewData => {
    put(`/feedback/${reviewId}`, reviewData)
      .then(() => {
        setSubmitted(true);
        // then redirect back to list
      })
      .catch(err => {
        setError(err.response ? err.response.data : 'Error saving your data');
        // set the error
      });
  };

  return submitted ? (
    <Redirect to="/performance-reviews" />
  ) : (
    <PerformanceReviewForm review={review} error={error} submit={submit} />
  );
};
