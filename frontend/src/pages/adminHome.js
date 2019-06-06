import React from 'react';
import { SectionContent, Columns, Column } from '../components';
import { Link } from 'react-router-dom';

export default () => {
  return (
    <section className="hero is-info is-fullheight-with-navbar">
      <SectionContent>
        <h1 className="title">Admin Home</h1>
        <Columns className="is-hidden-desktop">
          <Column>
            <Link className="button is-primary" to="/users">
              Manage Users
            </Link>
          </Column>
          <Column>
            <Link className="button is-primary" to="/performance">
              Performance Reviews
            </Link>
          </Column>
        </Columns>
      </SectionContent>
    </section>
  );
};
