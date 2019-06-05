import React, { useState, useEffect } from 'react';
import UserEditForm from '../components/userEditForm';
import { Redirect } from 'react-router-dom';
import { get, post } from '../api';

export default ({ currentUser, userId }) => {
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
    let url = userId ? `/user/${userId}` : '/user';
    post(url, userData)
      .then(response => {
        console.log('user submit resp', response);
        setSubmitted(true);
        // then redirect back to list
      })
      .catch(err => {
        console.log('error', err.response);
        setError(err.response ? err.response.data : 'Error saving your data');
        // set the error
      });
  };

  return submitted ? (
    <Redirect to="/users" />
  ) : (
    <UserEditForm currentUser={currentUser} user={user} error={error} submit={submit} />
  );
};
