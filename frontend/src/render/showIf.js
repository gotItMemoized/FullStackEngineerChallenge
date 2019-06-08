export default (statement, component) => {
  if (statement === true) {
    // function check voodoo
    if (component instanceof Function) {
      return component();
    }
    return component;
  } else if (statement !== false) {
    throw new Error('Expected a true or false statement');
  }
};
