import React from 'react';

// common reusable from Bulma
export default ({ className, children, ...props }) => (
  <article className={'message ' + (className || '')} {...props}>
    <div className="message-body">{children}</div>
  </article>
);
