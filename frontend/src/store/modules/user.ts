import * as types from "../types";

const initPageState = () => {
  return {
    userInfo: {
      id: 0,
      name: "",
      avatar: ""
    }
  };
};
const user = {
  state: initPageState(),
  mutations: {
    [types.SAVE_USER](state: object | any, pageState: object | any) {
      for (const prop in pageState) {
        state[prop] = pageState[prop];
      }
    }
  },
  actions: {}
};

export default user;
