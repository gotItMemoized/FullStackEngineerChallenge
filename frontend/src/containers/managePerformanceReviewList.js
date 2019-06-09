import React, { useState, useEffect } from 'react';
import { Columns, Column } from '../components';
import { Link } from 'react-router-dom';
import { get } from '../api';
import { showIf } from '../render';
import ReviewTable from '../components/reviewTable';

export default () => {
  const [reviews, setReviews] = useState([]);
  const [loaded, setLoaded] = useState(false);

  useEffect(() => {
    const fetchData = async () => {
      const result = await get('/review/all');
      setReviews(result.data);
      setLoaded(true);
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
        <Column>
          {showIf(reviews.length > 0, <ReviewTable reviews={reviews} />)}
          {showIf(reviews.length === 0 && loaded, <div>No performance reviews created.</div>)}
          {showIf(!loaded, <div>Loading</div>)}
        </Column>
      </Columns>
    </div>
  );
};
