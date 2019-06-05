import React from 'react';
import Header from '../components/header';

export default ({ user, setUser }) => {
  const logoutAction = () => {
    if (user.loggedIn) {
      setUser({ loggedIn: false, isAdmin: false });
    }
  };

  return <Header logoutAction={logoutAction} currentUser={user} />;
};
