import React, { useState } from 'react';
import { Redirect } from 'react-router-dom';
import { post } from '../api';
import Login from '../components/login';
import jwt from 'jsonwebtoken';

export default ({ user, setUser }) => {
  console.log(user);
  const [error, setError] = useState();

  const submitLogin = (username, password) => {
    // TODO:
    const data = {
      username,
      password,
    };

    post('/user/login', data, false)
      .then(response => {
        console.log('resp', response.data);
        // clear
        setError('');
        if (response.data && response.data.jwt) {
          const newUser = { ...jwt.decode(response.data.jwt), jwt: response.data.jwt };
          console.log('setting user', newUser);
          setUser(newUser);
        } else {
          console.log('nope');
          setError('Could not log you in with that information');
        }
        return response;
      })
      // .then(resp => {
      //   if (error.length === 0) {
      //     setLoggedIn(true);
      //   }
      //   return resp;
      // })
      .catch(err => {
        console.log(err);
        setError('Could not log you in with that information');
      });
  };

  return user.isLoggedIn === true ? (
    <Redirect to="/" />
  ) : (
    <Login error={error} submit={submitLogin} />
  );
};
