import React, { useState, useEffect } from 'react';
import { Columns, Column } from '../components';
import { Link } from 'react-router-dom';
import { get } from '../api';
import { showIf } from '../render';

const FeedbackTable = ({ reviews = [] }) => {
  const headerFooter = (
    <tr>
      <th>Reviewing</th>
      <th className="has-text-right">Action</th>
    </tr>
  );
  const rows = reviews.map(review => {
    return (
      <tr key={review.id}>
        <td>
          {`${review.name} (${review.username}) `}
          {showIf(
            review.message.String.length > 0,
            <span className="tag is-success">Complete</span>,
          )}
        </td>
        <td className="has-text-right">
          <Link className="button is-primary is-small" to={`/performance-reviews/${review.id}`}>
            Start
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

export default () => {
  const [reviews, setReviews] = useState([]);
  const [loaded, setLoaded] = useState(false);

  useEffect(() => {
    const fetchData = async () => {
      const result = await get('/feedback/all');
      setLoaded(true);
      setReviews(result.data);
    };
    fetchData();
  }, []);

  return (
    <Columns>
      <Column>
        {showIf(reviews.length > 0, <FeedbackTable reviews={reviews} />)}
        {showIf(
          reviews.length === 0 && loaded,
          <div>No performance reviews to give feedback on.</div>,
        )}
        {showIf(!loaded, <div>Loading</div>)}
      </Column>
    </Columns>
  );
};
