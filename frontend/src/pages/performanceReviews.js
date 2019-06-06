import React from 'react';
import PerformanceReviewList from '../containers/performanceReviewList';
import { SectionContent } from '../components';

export default () => {
  return (
    <section className="hero is-light is-bold is-fullheight-with-navbar">
      <SectionContent>
        <h1 className="title">Performance Reviews</h1>
        <PerformanceReviewList />
      </SectionContent>
    </section>
  );
};
