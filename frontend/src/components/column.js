import React from 'react';

export default ({ className, children, ...props }) => (
  <div className={'column ' + (className || '')} {...props}>
    {children}
  </div>
);
