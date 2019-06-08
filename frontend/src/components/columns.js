import wrapper from './wrapper';

export default ({ children, type, className = '', ...props }) =>
  wrapper({ children, type: 'columns', className, ...props });
