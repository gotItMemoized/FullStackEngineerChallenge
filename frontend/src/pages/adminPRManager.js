import React from 'react';
import Header from '../containers/header';
import ManagePerformanceReviewList from '../containers/managePerformanceReviewList';

export default ({ currentUser, setCurrentUser }) => {
  return (
    <div>
      <Header user={currentUser} setUser={setCurrentUser} />
      <section className="hero is-light is-bold is-fullheight-with-navbar">
        <div className="section">
          <div className="container">
            <h1 className="title">Manage Performance Reviews</h1>
            <ManagePerformanceReviewList />
          </div>
        </div>
      </section>
    </div>
  );
};
