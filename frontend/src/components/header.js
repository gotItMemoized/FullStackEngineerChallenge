import React from 'react';
import { NavLink } from 'react-router-dom';

export default function({ currentUser, logoutAction }) {
  const sections = () => {
    const links = [];
    if (currentUser.isAdmin) {
      links.push(
        <NavLink key="user" to="/users" className="navbar-item">
          Users
        </NavLink>,
      );
    }

    links.push(
      <NavLink key="performance" to="/performance" className="navbar-item">
        Performance Reviews
      </NavLink>,
    );
    return <div className="navbar-start">{links}</div>;
  };

  const userControls = () => {
    return (
      <div className="navbar-end">
        <div className="navbar-item">
          <div className="buttons">
            <button type="button" className="button is-light" onClick={logoutAction}>
              Log out
            </button>
          </div>
        </div>
      </div>
    );
  };

  return (
    <nav className="navbar" role="navigation" aria-label="main navigation">
      <div className="navbar-brand">
        <NavLink className="navbar-item" to="/">
          <strong>Coworker Yelp</strong>
        </NavLink>

        <button
          type="button"
          className="navbar-burger burger"
          aria-label="menu"
          aria-expanded="false"
          data-target="navbarBasicExample"
        >
          <span aria-hidden="true" />
          <span aria-hidden="true" />
          <span aria-hidden="true" />
        </button>
      </div>

      <div id="navbarBasic" className="navbar-menu">
        {sections()}

        {userControls()}
      </div>
    </nav>
  );
}
