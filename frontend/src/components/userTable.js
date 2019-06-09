import React, { useState } from 'react';
import { Link } from 'react-router-dom';
import { showIf } from '../render';

// display table for users
export default ({ currentUser, users = [], deleteAction = () => {} }) => {
  const [resetToggle, setResetToggle] = useState({});
  const headerFooter = (
    <tr>
      <th>Name</th>
      <th>Username</th>
      <th>Edit</th>
      <th>Delete</th>
    </tr>
  );
  const rows = users.map(user => {
    const deleteDisabled = user.id === currentUser.id;
    const deleteButton = (
      <button
        type="button"
        className="button is-small"
        disabled={deleteDisabled}
        onClick={() => setResetToggle({ [user.id]: true })}
      >
        Delete
      </button>
    );
    const confirmButton = showIf(
      resetToggle[user.id] === true,
      <button
        type="button"
        className="button is-danger is-small"
        onClick={() => deleteAction(user.id)}
      >
        Confirm
      </button>,
    );
    const adminTag = showIf(user.isAdmin === true, <span className="tag is-primary">Admin</span>);
    return (
      <tr key={user.id}>
        <td>
          {user.name + ' '}
          {adminTag}
        </td>
        <td>{user.username}</td>
        <td>
          <Link className="button is-small" to={`/users/${user.id}/edit`}>
            Edit
          </Link>
        </td>
        <td>{confirmButton ? confirmButton : deleteButton}</td>
      </tr>
    );
  });

  return (
    <table className="table is-fullwidth is-hoverable is-striped">
      <thead>{headerFooter}</thead>
      <tfoot>{headerFooter}</tfoot>
      <tbody>{rows}</tbody>
    </table>
  );
};
