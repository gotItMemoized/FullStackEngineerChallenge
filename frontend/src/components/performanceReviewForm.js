import React, { useState, useEffect } from 'react';

export default ({ error, review = {}, submit = () => {} }) => {
  const [message, setMessage] = useState(review.message);
  useEffect(() => {
    setMessage(review.message);
  }, [review.message]);

  const displayError = () => {
    if (error) {
      return <section className="section hero is-danger">{error}</section>;
    }
  };

  const buttonAttributes = {};
  if (!message || message.length < 1) {
    buttonAttributes.disabled = true;
  }

  const onSubmit = event => {
    event.preventDefault();
    if (message && message.length > 1) {
      const data = {
        message: { String: message, valid: true },
      };

      submit(data);
    }
  };

  return (
    <div>
      {displayError()}
      <form onSubmit={onSubmit}>
        <div className="field">
          <label className="label" htmlFor="message">
            Feedback for {review.name}
          </label>
          <div className="control">
            <textarea
              id="message"
              className="textarea"
              required
              value={message}
              onChange={e => {
                setMessage(e.target.value);
              }}
            />
          </div>
        </div>
        <div className="control">
          <button type="submit" className="button is-primary" {...buttonAttributes}>
            Save
          </button>
        </div>
      </form>
    </div>
  );
};
