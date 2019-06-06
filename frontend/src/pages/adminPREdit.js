import React from 'react';
import { Redirect } from 'react-router-dom';
import Header from '../containers/header';
import PerformanceReviewEditor from '../containers/performanceReviewEditor';

export default ({ currentUser, setCurrentUser, match }) => {
  const { id } = match.params;
  if (!id) {
    return <Redirect to="/performance-manager" />;
  }
  return (
    <div>
      <Header user={currentUser} setUser={setCurrentUser} />
      <section className="hero is-light is-bold is-fullheight-with-navbar">
        <div className="section">
          <div className="container">
            <h1 className="title">Edit Performance Review</h1>
            <PerformanceReviewEditor reviewId={id} />
          </div>
        </div>
      </section>
    </div>
  );
};
