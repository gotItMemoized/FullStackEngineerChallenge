import React from 'react';
import { Link } from 'react-router-dom';
import { Columns, Column, SectionContent } from '../components';

export default () => {
  return (
    <section className="hero is-success is-bold is-fullheight-with-navbar">
      <SectionContent>
        <h1 className="title">Performance Reviews</h1>
        <Columns>
          <Column>
            <Link className="button" to="/performance-manager">
              Performance Review Management
            </Link>
          </Column>
        </Columns>
        <Columns>
          <Column>
            <Link className="button" to="/performance-reviews">
              Reviews Assigned to Me
            </Link>
          </Column>
        </Columns>
      </SectionContent>
    </section>
  );
};
