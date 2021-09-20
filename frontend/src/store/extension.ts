export default {
  namespaced: true,
  state() {
    return {
      name: '',
      data: {},
      func: () => {},
    };
  },
  mutations: {
    SET_EXTENSION(state: any, payload: any) {
      state.name = payload.name;
      state.data = payload.data;
      state.func = payload.func || function () {};
    },
  },
  actions: {
    show({ commit }: any, payload: any) {
      commit('SET_EXTENSION', payload);
    },
    close({ commit }: any) {
      commit('SET_EXTENSION', {});
    },
    ok({ state, commit }: any) {
      state?.func();
      commit('SET_EXTENSION', {});
    },
  },
};
