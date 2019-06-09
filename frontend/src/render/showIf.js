export default (statement, component) => {
  if (statement === true) {
    // if it's a function, we'll run it and return
    // this allows us to have things that need to execute later
    //  and not when they're written into the parameter
    if (component instanceof Function) {
      return component();
    }
    return component;
  } else if (statement !== false) {
    throw new Error('Expected a true or false statement');
  }
};
