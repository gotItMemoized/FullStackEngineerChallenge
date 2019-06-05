import React from 'react';
import { Redirect } from 'react-router-dom';
import Header from '../containers/header';
import UserEditor from '../containers/userEditor';

export default ({ currentUser, setCurrentUser, match }) => {
  const { id } = match.params;
  if (!id) {
    return <Redirect to="/users" />;
  }
  return (
    <div>
      <Header user={currentUser} setUser={setCurrentUser} />
      <section className="hero is-primary is-bold is-fullheight-with-navbar">
        <div className="section">
          <div className="container">
            <h1 className="title">Edit User</h1>
            <UserEditor currentUser={currentUser} userId={id} />
          </div>
        </div>
      </section>
    </div>
  );
};
