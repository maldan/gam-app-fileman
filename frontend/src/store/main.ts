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

    SET_LOADING(state: any, payload: boolean) {
      state.isLoading = payload;
    },
  },
  actions: {
    async getPath({ commit, dispatch }: any) {
      commit('SET_PATH', await RestApi.file.getPath());
      dispatch('file/getList', null, { root: true });
    },
    async changePath({ state, commit, dispatch }: any, path: string) {
      commit('SET_PATH', path.replace(/\/\//g, '/'));
      commit('SET_LOADING', true);
      await RestApi.file.setPath(state.path);
      commit('SET_LOADING', false);
      dispatch('file/getList', null, { root: true });
      dispatch('file/clearSelection', null, { root: true });
      dispatch('tab/changePath', path, { root: true });
    },
    setLoading({ commit }: any, payload: any) {
      commit('SET_LOADING', payload);
    },
  },
};
