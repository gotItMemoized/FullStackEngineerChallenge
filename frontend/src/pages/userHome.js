import React from 'react';
import Header from '../containers/header';

export default ({ user, setUser }) => {
  return (
    <div>
      <Header user={user} setUser={setUser} />
      <section className="hero is-light is-bold is-fullheight-with-navbar">
        <div className="section">
          <div className="container">
            <h1 className="title">User Home</h1>
          </div>
        </div>
      </section>
    </div>
  );
};
