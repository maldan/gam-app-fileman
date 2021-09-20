import { RestApi } from '@/util/RestApi';

export default {
  namespaced: true,
  state() {
    return {
      path: '/',
      isLoading: false,
    };
  },
  mutations: {
    SET_PATH(state: any, payload: any) {
      state.path = payload;
    },

    SET_LOADING(state: any, payload: any) {
      state.isLoading = payload;
    },
  },
  actions: {
    async getPath({ commit, dispatch }: any) {
      commit('SET_PATH', await RestApi.file.getPath());
      dispatch('file/getList', null, { root: true });
    },
    async changePath({ state, commit, dispatch }: any, payload: any) {
      commit('SET_PATH', payload.replace(/\/\//g, '/'));
      await RestApi.file.setPath(state.path);
      dispatch('file/getList', null, { root: true });
    },
    setLoading({ commit }: any, payload: any) {
      commit('SET_LOADING', payload);
    },
  },
};
