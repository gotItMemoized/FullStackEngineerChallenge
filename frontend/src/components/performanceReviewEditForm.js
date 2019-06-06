import React, { useState } from 'react';
import Select from 'react-select';
import useFormCheckbox from '../hooks/useFormCheckbox';
import useFormSelect from '../hooks/useFormSelect';

export default ({ error, users = [], review = {}, submit = () => {} }) => {
  const user = useFormSelect(review.user);
  const assignedUsers = useFormSelect(review.feedback);
  const isActive = useFormCheckbox(review.isActive || !review.id);

  const displayError = () => {
    if (error) {
      return <section className="section hero is-danger">{error}</section>;
    }
  };

  const userAttributes = {};
  if (review.id) {
    userAttributes.isDisabled = true;
  }

  const buttonAttributes = {};
  if (
    !user.value ||
    !user.value.value ||
    (!assignedUsers.value || assignedUsers.value.length < 1)
  ) {
    buttonAttributes.disabled = true;
  }

  const onSubmit = event => {
    event.preventDefault();
    if (user.value) {
      const data = {
        isActive: isActive.checked,
        user: { id: user.value.value },
        feedback: assignedUsers.value.map(a => {
          return {
            reviewer: {
              id: a.value,
            },
          };
        }),
      };

      submit(data);
    }
  };

  return (
    <div>
      {displayError()}
      <form onSubmit={onSubmit}>
        <div className="field">
          <label className="label" htmlFor="user">
            Review For
          </label>
          <div className="control">
            <Select
              id="user"
              required
              defaultValue={user.value || ''}
              {...user}
              {...userAttributes}
              options={users}
            />
          </div>
        </div>
        <div className="field">
          <label className="checkbox" htmlFor="isActive">
            <input type="checkbox" id="isActive" {...isActive} /> Is Active
          </label>
        </div>
        <div className="field">
          <label className="label" htmlFor="assignedUsers">
            Assigned Users
          </label>
          <div className="control">
            <Select
              id="assignedUsers"
              required
              defaultValue={assignedUsers || ''}
              isMulti
              {...assignedUsers}
              options={users}
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
