import wrapper from './wrapper';

export default ({ className, children = '', ...props }) =>
  wrapper({ children, type: 'column', className, ...props });
