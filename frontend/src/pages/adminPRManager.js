import React from 'react';
import ManagePerformanceReviewList from '../containers/managePerformanceReviewList';
import { SectionContent } from '../components';

export default () => {
  return (
    <section className="hero is-light is-bold is-fullheight-with-navbar">
      <SectionContent>
        <h1 className="title">Manage Performance Reviews</h1>
        <ManagePerformanceReviewList />
      </SectionContent>
    </section>
  );
};
