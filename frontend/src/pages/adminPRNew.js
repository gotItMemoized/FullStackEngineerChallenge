import React from 'react';
import Header from '../containers/header';
import PerformanceReviewEditor from '../containers/performanceReviewEditor';

export default ({ currentUser, setCurrentUser }) => {
  return (
    <div>
      <Header user={currentUser} setUser={setCurrentUser} />
      <section className="hero is-light is-bold is-fullheight-with-navbar">
        <div className="section">
          <div className="container">
            <h1 className="title">New Performance Review</h1>
            <PerformanceReviewEditor />
          </div>
        </div>
      </section>
    </div>
  );
};
