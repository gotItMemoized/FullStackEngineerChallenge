import React from 'react';
import { Redirect } from 'react-router-dom';
import PerformanceReview from '../containers/performanceReview';
import { SectionContent } from '../components';

export default ({ match }) => {
  const { id } = match.params;
  if (!id) {
    return <Redirect to="/performance-reviews" />;
  }
  return (
    <section className="hero is-light is-bold is-fullheight-with-navbar">
      <SectionContent>
        <h1 className="title">Performance Review</h1>
        <PerformanceReview reviewId={id} />
      </SectionContent>
    </section>
  );
};
