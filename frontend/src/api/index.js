import axios from 'axios';

// Auth Wrappers for the get/post/put/delete

export const post = (url, data, useAuth) => {
  const headers = {};
  if (useAuth !== false) {
    const { jwt } = JSON.parse(localStorage.getItem('jwt_token') || {});
    headers.Authorization = `BEARER ${jwt}`;
  }
  return axios.post(`/api${url}`, data, { headers });
};

export const put = (url, data, useAuth) => {
  const headers = {};
  if (useAuth !== false) {
    const { jwt } = JSON.parse(localStorage.getItem('jwt_token') || {});
    headers.Authorization = `BEARER ${jwt}`;
  }
  return axios.put(`/api${url}`, data, { headers });
};

// could add a toggle for useAuth on get, but most data will require a user context
export const get = url => {
  const { jwt } = JSON.parse(localStorage.getItem('jwt_token') || {});
  return axios.get(`/api${url}`, {
    headers: {
      Authorization: `BEARER ${jwt}`,
    },
  });
};

export const deleteUser = id => {
  const { jwt } = JSON.parse(localStorage.getItem('jwt_token') || {});
  return axios.delete(`/api/user/${id}`, {
    headers: {
      Authorization: `BEARER ${jwt}`,
    },
  });
};
