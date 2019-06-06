import React from 'react';
import UserEditor from '../containers/userEditor';

export default () => {
  return (
    <section className="hero is-primary is-bold is-fullheight-with-navbar">
      <div className="section">
        <div className="container">
          <h1 className="title">New User</h1>
          <UserEditor />
        </div>
      </div>
    </section>
  );
};
