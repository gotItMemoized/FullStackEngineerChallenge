import React from 'react';

export default ({ className, children, ...props }) => (
  <article className={'message ' + (className || '')} {...props}>
    <div className="message-body">{children}</div>
  </article>
);
