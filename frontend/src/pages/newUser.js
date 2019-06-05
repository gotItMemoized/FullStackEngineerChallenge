import React from 'react';
import Header from '../containers/header';
import UserEditor from '../containers/userEditor';

export default ({ currentUser, setCurrentUser }) => {
  return (
    <div>
      <Header user={currentUser} setUser={setCurrentUser} />
      <section className="hero is-primary is-bold is-fullheight-with-navbar">
        <div className="section">
          <div className="container">
            <h1 className="title">New User</h1>
            <UserEditor currentUser={currentUser} />
          </div>
        </div>
      </section>
    </div>
  );
};
