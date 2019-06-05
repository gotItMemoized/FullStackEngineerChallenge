import React from 'react';

export default ({ children, ...props }) => (
  <div className="columns" {...props}>
    {children}
  </div>
);
