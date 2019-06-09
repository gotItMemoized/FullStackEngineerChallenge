import React, { useState, useEffect } from 'react';
import UserEditForm from '../components/userEditForm';
import { Redirect } from 'react-router-dom';
import { get, post, put } from '../api';

// update a single user
export default ({ userId }) => {
  const [user, setUser] = useState();
  const [submitted, setSubmitted] = useState(false);
  const [error, setError] = useState();

  useEffect(() => {
    const fetchData = async () => {
      const result = await get(`/user/${userId}`);
      setUser(result.data);
    };
    if (userId) {
      fetchData();
    }
  }, [userId]);

  const submit = userData => {
    if (userId) {
      // update
      put(`/user/${userId}`, userData)
        .then(() => {
          setSubmitted(true);
          // then redirect back to list
        })
        .catch(err => {
          setError(err.response ? err.response.data : 'Error saving your data');
          // set the error
        });
    } else {
      // create
      post('/user', userData)
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
    <Redirect to="/users" />
  ) : (
    <UserEditForm user={user} error={error} submit={submit} />
  );
};
