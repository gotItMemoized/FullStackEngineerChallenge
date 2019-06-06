import React from 'react';
import UserList from '../containers/userList';
import { SectionContent } from '../components';

export default ({ user }) => {
  return (
    <section className="hero is-light is-bold is-fullheight-with-navbar">
      <SectionContent>
        <h1 className="title">Users</h1>
        <UserList currentUser={user} />
      </SectionContent>
    </section>
  );
};
