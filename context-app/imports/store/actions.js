import { REQUEST_DATA, REQUEST_LIST_DATA, REQUEST_ID_DATA, RECIEVE_DATA, RECIEVE_LIST_DATA, RECIEVE_ID_DATA, CACHE_DATA } from "./action-types";

const defaultRequestAction = (type, key, url, params, data) => ({
  type,
  key,
  url,
  params,
  data
});

const defaultReceiveAction = (type, key, url, data) => ({
  type,
  key,
  url,
  data
});

export const requestData = (key, url, params, data) => defaultRequestAction(REQUEST_DATA, key, url, params, data);


export const requestListData = (key, url, params, data) => defaultRequestAction(REQUEST_LIST_DATA, key, url, params, data);

export const requestIDData = (key, url, params, data, id) => {
  let action = defaultRequestAction(REQUEST_ID_DATA, key, url, params, data);
  return {
    ...action,
    id
  };
}

export const receiveData = (key, url, data) => defaultReceiveAction(RECIEVE_DATA, key, url, data);

export const receiveListData = (key, url, data) => defaultReceiveAction(RECIEVE_LIST_DATA, key, url, data);

export const receiveIDData = (key, url, data, id) => {
  let action = defaultReceiveAction(RECIEVE_ID_DATA, key, url, data);
  return {
    ...action,
    id
  };
}

export const cacheData = (key, data) => ({
  type: CACHE_DATA,
  key,
  data
});
