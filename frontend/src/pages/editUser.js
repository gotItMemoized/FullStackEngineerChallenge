import React from 'react';
import { Redirect } from 'react-router-dom';
import UserEditor from '../containers/userEditor';
import { SectionContent } from '../components';

export default ({ match }) => {
  const { id } = match.params;
  if (!id) {
    return <Redirect to="/users" />;
  }
  return (
    <section className="hero is-primary is-bold is-fullheight-with-navbar">
      <SectionContent>
        <h1 className="title">Edit User</h1>
        <UserEditor userId={id} />
      </SectionContent>
    </section>
  );
};
