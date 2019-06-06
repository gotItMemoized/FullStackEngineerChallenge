export default (statement, component) => {
  if (statement === true) {
    return component;
  } else if (statement !== false) {
    throw new Error('Expected a true or false statement');
  }
};
