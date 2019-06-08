import React, { useState, useEffect } from 'react';
import PerformanceReviewEditForm from '../components/performanceReviewEditForm';
import { Redirect } from 'react-router-dom';
import { get, post, put } from '../api';

export default ({ currentUser, reviewId }) => {
  const [review, setReview] = useState();
  const [users, setUsers] = useState([]);
  const [submitted, setSubmitted] = useState(false);
  const [error, setError] = useState();

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
          feedback,
          key: feedback.reviewer.id,
          value: feedback.reviewer.id,
          label: `${feedback.reviewer.name} (${feedback.reviewer.username})`,
        };
      });
      setReview(converted);
    };
    if (reviewId) {
      fetchData();
    }
  }, [reviewId]);

  useEffect(() => {
    const fetchData = async () => {
      const result = await get(`/user/all`);
      setUsers(
        result.data.map(user => {
          return {
            feedback: user,
            label: `${user.name} (${user.username})`,
            key: user.id,
            value: user.id,
          };
        }),
      );
    };
    fetchData();
  }, []);

  const submit = reviewData => {
    if (reviewId) {
      put(`/review/${reviewId}`, reviewData)
        .then(() => {
          setSubmitted(true);
          // then redirect back to list
        })
        .catch(err => {
          setError(err.response ? err.response.data : 'Error saving your data');
          // set the error
        });
    } else {
      post('/review', reviewData)
        .then(() => {
          setSubmitted(true);
          // then redirect back to list
        })
        .catch(err => {
          setError(err.response ? err.response.data : 'Error saving your data');
          // set the error
        });
    }
  };

  return submitted ? (
    <Redirect to="/performance-manager" />
  ) : (
    <PerformanceReviewEditForm
      currentUser={currentUser}
      users={users}
      review={review}
      error={error}
      submit={submit}
    />
  );
};
