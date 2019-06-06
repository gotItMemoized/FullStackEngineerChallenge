import React, { useState, useEffect } from 'react';
import { Columns, Column } from '../components';
import { Link } from 'react-router-dom';
import { get, deleteUser } from '../api';

const ReviewTable = ({ reviews = [] }) => {
  const headerFooter = (
    <tr>
      <th>Name</th>
      <th>Complete</th>
      <th>Edit</th>
    </tr>
  );
  const rows = reviews.map(review => {
    return (
      <tr key={review.id}>
        <td>{review.user.name + ' '}</td>
        <td>{!review.isActive ? 'Completed' : ''}</td>
        <td>
          <Link className="button is-small" to={`/performance-manager/${review.id}/edit`}>
            Edit
          </Link>
        </td>
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

export default ({}) => {
  const [reviews, setReviews] = useState([]);

  useEffect(() => {
    const fetchData = async () => {
      const result = await get('/review/all');
      setReviews(result.data);
    };
    fetchData();
  }, []);

  return (
    <div>
      <Columns>
        <Column className="is-offset-10 has-text-right">
          <Link className="button" to="/performance-manager/new">
            New Review
          </Link>
        </Column>
      </Columns>
      <Columns>
        <ReviewTable reviews={reviews} />
      </Columns>
    </div>
  );
};
