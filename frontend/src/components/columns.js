import React from 'react';

export default ({ children, className = '', ...props }) => (
  <div className={`columns ${className}`} {...props}>
    {children}
  </div>
);
