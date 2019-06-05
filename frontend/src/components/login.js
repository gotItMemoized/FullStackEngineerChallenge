import React, { useState } from 'react';

function useFormInput(initial) {
  const [value, setValue] = useState(initial);

  return {
    value,
    onChange: e => {
      setValue(e.target.value);
    },
  };
}

export default function({ submit, error }) {
  const username = useFormInput('');
  const password = useFormInput('');

  const submitButtonAttribute = {
    disabled: username.value.length === 0 || password.value.length === 0,
  };

  const submitCreds = e => {
    e.preventDefault();
    submit(username.value, password.value);
  };

  const displayError = () => {
    if (error) {
      return <section className="section hero is-danger">{error}</section>;
    }
  };
  return (
    <div>
      {displayError()}
      <section className="hero is-info is-fullheight">
        <div className="section">
          <div className="container">
            <h1 className="title">Welcome</h1>
            <form onSubmit={submitCreds}>
              <div className="field">
                <label className="label" htmlFor="username">
                  Username
                </label>
                <div className="control">
                  <input
                    id="username"
                    type="username"
                    placeholder="myUserName"
                    className="input"
                    {...username}
                  />
                </div>
              </div>
              <div className="field">
                <label className="label" htmlFor="password">
                  Password
                </label>
                <div className="control">
                  <input
                    id="password"
                    type="password"
                    placeholder="kanjiGaSukiDesu"
                    className="input"
                    {...password}
                  />
                </div>
              </div>
              <div className="field">
                <div className="control buttons">
                  <button type="submit" className="button is-primary" {...submitButtonAttribute}>
                    Login
                  </button>
                </div>
              </div>
            </form>
          </div>
        </div>
      </section>
    </div>
  );
}
