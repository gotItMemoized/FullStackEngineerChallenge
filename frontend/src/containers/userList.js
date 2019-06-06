import React, { useState, useEffect } from 'react';
import { Columns, Column } from '../components';
import { Link } from 'react-router-dom';
import { get, deleteUser } from '../api';

const UserTable = ({ currentUser, users = [], deleteAction = () => {} }) => {
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
    let adminTag;
    let deleteButton = (
      <button
        type="button"
        className="button is-small"
        onClick={() => setResetToggle({ [user.id]: true })}
      >
        Delete
      </button>
    );
    let confirmButton =
      resetToggle[user.id] === true ? (
        <button
          type="button"
          className="button is-danger is-small"
          onClick={() => deleteAction(user.id)}
        >
          Confirm
        </button>
      ) : (
        undefined
      );
    if (user.isAdmin) {
      adminTag = <span className="tag is-primary">Admin</span>;
    }
    if (user.id === currentUser.id) {
      deleteButton = (
        <button type="button" className="button is-small" disabled>
          Delete
        </button>
      );
    }
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

export default ({ currentUser }) => {
  const [users, setUsers] = useState([]);

  useEffect(() => {
    const fetchData = async () => {
      const result = await get('/user/all');
      setUsers(result.data);
    };
    fetchData();
  }, []);

  const deleteAction = userId => {
    deleteUser(userId)
      .then(resp => {
        setUsers(users.filter(u => u.id !== userId));
      })
      .catch(err => {
        // TODO: error view?
      });
  };

  return (
    <div>
      <Columns>
        <Column className="is-offset-10 has-text-right">
          <Link className="button" to="/users/new">
            New User
          </Link>
        </Column>
      </Columns>
      <Columns>
        <UserTable users={users} deleteAction={deleteAction} currentUser={currentUser} />
      </Columns>
    </div>
  );
};
