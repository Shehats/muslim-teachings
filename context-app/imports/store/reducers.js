import { REQUEST_DATA, REQUEST_LIST_DATA, REQUEST_ID_DATA, RECIEVE_DATA, RECIEVE_LIST_DATA, RECIEVE_ID_DATA, CACHE_DATA } from "./action-types";
import { combineReducers } from 'redux';

let initialState = {
};

const getObjectOrCreate = (state, key) => (state[key] || {
  isFeching: false,
  data: null,
  listData: null,
  IDData: null
})

const actionsReducer = (state = initialState, action) => {
  switch (action.type) {
    case REQUEST_DATA: case REQUEST_LIST_DATA: case REQUEST_ID_DATA: {
      let curnData = getObjectOrCreate(state, action.key)
      return {
        ...state,
        [action.key]: {
          ...curnData,
          isFeching: true
        }
      };
    }

    case RECIEVE_DATA: {
      let curnData = getObjectOrCreate(state, action.key);
      return {
        ...state,
        [action.key]: {
          ...curnData,
          isFeching: false,
          data: action.data
        }
      };
    }

    case RECIEVE_LIST_DATA: {
      let curnData = getObjectOrCreate(state, action.key);
      return {
        ...state,
        [action.key]: {
          ...curnData,
          isFeching: false,
          listData: action.data
        }
      };
    }

    case RECIEVE_ID_DATA: {
      let curnData = getObjectOrCreate(state, action.key);
      let curnIDData = curnData.IDData
      return {
        ...state,
        [action.key]: {
          ...curnData,
          isFeching: false,
          IDData: {
            ...curnIDData,
            [action.id]: action.data
          }
        }
      };
    }

    default:
      return state;
  }
}

const rootReducer = combineReducers({
  actionsReducer
});

export default rootReducer;
