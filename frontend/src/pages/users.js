import React from 'react';
import Header from '../containers/header';
import UserList from '../containers/userList';

export default ({ user, setUser }) => {
  return (
    <div>
      <Header user={user} setUser={setUser} />
      <section className="hero is-light is-bold is-fullheight-with-navbar">
        <div className="section">
          <div className="container">
            <h1 className="title">Users</h1>
            <UserList currentUser={user} />
          </div>
        </div>
      </section>
    </div>
  );
};
