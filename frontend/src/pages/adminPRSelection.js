import React from 'react';
import Header from '../containers/header';
import { Link } from 'react-router-dom';
import {Columns, Column} from '../components';

export default ({ currentUser, setCurrentUser }) => {
  return (
    <div>
      <Header user={currentUser} setUser={setCurrentUser} />
      <section className="hero is-success is-bold is-fullheight-with-navbar">
        <div className="section">
          <div className="container">
            <h1 className="title">Performance Reviews</h1>
            <Columns>
              <Column>
                <Link className="button" to="/performance-manager">Performance Review Management</Link>
              </Column>
            </Columns>
            <Columns>
              <Column>
                <Link className="button" to="/performance-view">Reviews Assigned to Me</Link>
              </Column>
            </Columns>
          </div>
        </div>
      </section>
    </div>
  );
};
