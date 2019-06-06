import { useState, useEffect } from 'react';

export default function useFormSelect(initial) {
  const defaultValue = initial || '';
  const [value, setValue] = useState(defaultValue);

  useEffect(() => {
    setValue(initial || '');
  }, [initial]);

  return {
    value,
    onChange: event => {
      setValue(event);
    },
  };
}
