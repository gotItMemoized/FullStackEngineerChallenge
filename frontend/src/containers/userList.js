import React, { useState, useEffect } from 'react';
import { Columns, Column } from '../components';
import UserTable from '../components/userTable';
import { Link } from 'react-router-dom';
import { get, deleteUser } from '../api';

// view edit delete users
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
    deleteUser(userId).then(() => {
      // when we get the response from the backend,
      // we'll just manually remove it so we don't need to reload
      // we'd want to reload once we put in pagination however.
      setUsers(users.filter(u => u.id !== userId));
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
