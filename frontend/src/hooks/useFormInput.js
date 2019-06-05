import { useState, useEffect } from 'react';

export default function useFormInput(initial) {
  const defaultValue = initial || '';
  const [value, setValue] = useState(defaultValue);

  useEffect(() => {
    setValue(initial || '');
  }, [initial]);

  return {
    value,
    onChange: event => {
      setValue(event.target.value);
    },
  };
}
