import React from 'react';
import { Redirect } from 'react-router-dom';
import PerformanceReviewViewer from '../containers/performanceReviewViewer';
import { SectionContent } from '../components';

export default ({ match }) => {
  const { id } = match.params;
  if (!id) {
    return <Redirect to="/performance-manager" />;
  }
  return (
    <section className="hero is-light is-bold is-fullheight-with-navbar">
      <SectionContent>
        <h1 className="title">View Performance Review</h1>
        <PerformanceReviewViewer reviewId={id} />
      </SectionContent>
    </section>
  );
};
