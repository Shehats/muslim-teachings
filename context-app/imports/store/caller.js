import axios from 'axios';

export const getData = (url) => axios.get(url);

export const createData = (url, data) => axios.post(url, data);

export const deleteData = (url) => axios.delete(url);

export const updateData = (url, data) => axios.put(url, data);
