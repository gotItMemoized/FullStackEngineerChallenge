import React, { useState } from 'react';
import { Redirect } from 'react-router-dom';
import { post } from '../api';
import Login from '../components/login';
import jwt from 'jsonwebtoken';

// login page
export default ({ user, setUser }) => {
  const [error, setError] = useState();

  const submitLogin = (username, password) => {
    const data = {
      username,
      password,
    };

    post('/user/login', data, false)
      .then(response => {
        // clear
        setError('');
        if (response.data && response.data.jwt) {
          const newUser = { ...jwt.decode(response.data.jwt), jwt: response.data.jwt };
          setUser(newUser);
        } else {
          setError('Could not log you in with that information');
        }
        return response;
      })
      .catch(() => {
        setError('Could not log you in with that information');
      });
  };

  return user.isLoggedIn === true ? (
    <Redirect to="/" />
  ) : (
    <Login error={error} submit={submitLogin} />
  );
};
