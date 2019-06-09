import React from 'react';

// reusable wrapper for other components to set up (Columns, Column, etc)
export default ({ children, type, className = '', ...props }) => (
  <div className={`${type} ${className}`} {...props}>
    {children}
  </div>
);
