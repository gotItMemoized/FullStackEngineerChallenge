import React from 'react';
import { Redirect } from 'react-router-dom';
import PerformanceReviewEditor from '../containers/performanceReviewEditor';
import { SectionContent } from '../components';

export default ({ match }) => {
  const { id } = match.params;
  if (!id) {
    return <Redirect to="/performance-manager" />;
  }
  return (
    <section className="hero is-light is-bold is-fullheight-with-navbar">
      <SectionContent>
        <h1 className="title">Edit Performance Review</h1>
        <PerformanceReviewEditor reviewId={id} />
      </SectionContent>
    </section>
  );
};
