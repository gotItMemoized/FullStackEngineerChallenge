import React from 'react';
import Header from '../components/header';
import { SectionContent } from '../components';

export default ({ user, setUser }) => {
  return (
    <div>
      <Header currentUser={user} setCurrentUser={setUser} />
      <section className="hero is-light is-bold is-fullheight-with-navbar">
        <SectionContent>
          <h1 className="title">Loading Your Content</h1>
        </SectionContent>
      </section>
    </div>
  );
};
