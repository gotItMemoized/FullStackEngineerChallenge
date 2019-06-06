import React from 'react';
import PerformanceReviewEditor from '../containers/performanceReviewEditor';
import { SectionContent } from '../components';

export default () => {
  return (
    <section className="hero is-light is-bold is-fullheight-with-navbar">
      <SectionContent>
        <h1 className="title">New Performance Review</h1>
        <PerformanceReviewEditor />
      </SectionContent>
    </section>
  );
};
