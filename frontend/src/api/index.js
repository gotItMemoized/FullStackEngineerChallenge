import axios from 'axios';

// Wrappers for the get/post

export const post = (url, data, useAuth) => {
  const headers = {};
  if (useAuth !== false) {
    const { jwt } = JSON.parse(localStorage.getItem('jwt_token') || {});
    headers.Authorization = `BEARER ${jwt}`;
  }
  return axios.post(url, data, { headers });
};

// could add a toggle for useAuth on get, but most data will require a user context
export const get = url => {
  const { jwt } = JSON.parse(localStorage.getItem('jwt_token') || {});
  return axios.get(url, {
    headers: {
      Authorization: `BEARER ${jwt}`,
    },
  });
};

export const deleteUser = id => {
  const { jwt } = JSON.parse(localStorage.getItem('jwt_token') || {});
  return axios.delete(`/user/${id}`, {
    headers: {
      Authorization: `BEARER ${jwt}`,
    },
  });
};

axios.interceptors.response.use(
  response => {
    return response;
  },
  error => {
    const {
      response: { status },
    } = error;

    if (status === 401) {
      localStorage.removeItem('jwt_token');
    }
    return Promise.reject(error);
  },
);