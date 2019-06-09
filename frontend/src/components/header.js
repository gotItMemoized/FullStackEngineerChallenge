import React, { useState } from 'react';
import { NavLink } from 'react-router-dom';
import { showIf } from '../render';
import { userDefault } from '../App';

// the page header
export default function({ currentUser = userDefault, logoutAction }) {
  const [showMenu, setShowMenu] = useState(false);

  const sections = () => {
    if (!currentUser.loggedIn) return;
    const links = [];
    if (currentUser.isAdmin) {
      links.push(
        <NavLink key="user" id="users-menu-btn" to="/users" className="navbar-item">
          Users
        </NavLink>,
      );
    }

    links.push(
      <NavLink
        key="performance"
        id="performance-menu-btn"
        to="/performance"
        className="navbar-item"
      >
        Performance Reviews
      </NavLink>,
    );
    return <div className="navbar-start">{links}</div>;
  };

  const userControls = () => {
    if (!currentUser.loggedIn) return;
    return (
      <div className="navbar-end">
        <div className="navbar-item">
          <div className="buttons">
            <button
              type="button"
              id="logout-btn"
              className="button is-light"
              onClick={logoutAction}
            >
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
          onClick={() => {
            setShowMenu(!showMenu);
          }}
        >
          <span aria-hidden="true" />
          <span aria-hidden="true" />
          <span aria-hidden="true" />
        </button>
      </div>

      <div id="navbarBasic" className={`navbar-menu ${showIf(showMenu, 'is-active')}`}>
        {sections()}

        {userControls()}
      </div>
    </nav>
  );
}
