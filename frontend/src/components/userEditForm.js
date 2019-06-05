import React from 'react';
import useFormInput from '../hooks/useFormInput';
import useFormCheckbox from '../hooks/useFormCheckbox';

export default ({ error, user = {}, submit = () => {} }) => {
  const username = useFormInput(user.username);
  const name = useFormInput(user.name);
  const newPassword = useFormInput();
  const isAdmin = useFormCheckbox(user.isAdmin);

  const displayError = () => {
    if (error) {
      return <section className="section hero is-danger">{error}</section>;
    }
  };

  const buttonAttributes = {};
  if ((!user.id && newPassword.value.length === 0) || username.value.length === 0) {
    buttonAttributes.disabled = true;
  }

  const onSubmit = event => {
    event.preventDefault();
    if ((newPassword.value.length !== 0 || user.id) && username.value.length !== 0) {
      submit({
        name: name.value,
        username: username.value,
        newPassword: newPassword.value,
        isAdmin: isAdmin.value,
      });
    }
  };

  const passwordWarning =
    !user.id || newPassword.value <= 0 ? (
      React.Fragment
    ) : (
      <article className="message is-warning">
        <div className="message-body">
          This is a password for a current user, be careful when editing
        </div>
      </article>
    );

  return (
    <div>
      {displayError()}
      <form onSubmit={onSubmit}>
        <div className="field">
          <label className="label" htmlFor="name">
            Name
          </label>
          <div className="control">
            <input id="name" placeholder="James" className="input" {...name} />
          </div>
        </div>
        <div className="field">
          <label className="label" htmlFor="username">
            Username
          </label>
          <div className="control">
            <input
              id="username"
              type="text"
              placeholder="jcrisman"
              className="input"
              {...username}
              required
            />
          </div>
        </div>
        <div className="field">
          <label className="label" htmlFor="password">
            New Password
          </label>
          <div className="control">
            <input
              id="newPassword"
              type="password"
              placeholder="kanjiGaSukiDesuka"
              className="input"
              minLength="10"
              {...newPassword}
            />
          </div>
        </div>
        {passwordWarning}
        <div className="field">
          <label className="checkbox" htmlFor="isAdmin">
            <input type="checkbox" id="isAdmin" {...isAdmin} /> Is Admin User
          </label>
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
