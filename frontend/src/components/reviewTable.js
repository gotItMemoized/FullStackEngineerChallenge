import React from 'react';
import { Link } from 'react-router-dom';
import { showIf } from '../render';

// view all reviews
export default ({ reviews = [] }) => {
  const headerFooter = (
    <tr>
      <th>Review For</th>
      <th>Responses</th>
      <th>Edit</th>
    </tr>
  );
  const rows = reviews.map(review => {
    return (
      <tr key={review.id}>
        <td>{review.user.name}</td>
        <td>
          {showIf(
            !review.isActive,
            <Link
              className="button is-primary is-outlined is-small"
              to={`/performance-manager/${review.id}/view`}
            >
              View Responses
            </Link>,
          )}
          {showIf(
            review.isActive,
            <button className="button is-outlined is-small" disabled>
              Open For Responses
            </button>,
          )}
        </td>
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
