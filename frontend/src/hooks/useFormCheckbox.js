import { useState, useEffect } from 'react';

export default function useFormCheckbox(initial) {
  const [checked, setChecked] = useState(initial || false);

  useEffect(() => {
    setChecked(initial || false);
  }, [initial]);

  return {
    checked,
    onChange: event => {
      setChecked(event.target.checked);
    },
  };
}
