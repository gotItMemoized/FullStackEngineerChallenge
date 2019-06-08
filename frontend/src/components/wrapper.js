import React from 'react';

export default ({ children, type, className = '', ...props }) => (
  <div className={`${type} ${className}`} {...props}>
    {children}
  </div>
);
