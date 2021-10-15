import { RestApi } from '@/util/RestApi';

export default {
  namespaced: true,
  state() {
    return {
      tabs: ['/', '/home'],
      tab: { label: '/', value: '/' },
      tabIndex: 0,
    };
  },
  mutations: {
    SET_PATH(state: any, path: string) {
      state.tabs[state.tabIndex] = path;
      state.tab.label = path;
      state.tab.value = path;
    },
    SET_INDEX(state: any, index: number) {
      state.tabIndex = index;
    },
  },
  actions: {
    async changePath({ commit }: any, path: string) {
      commit('SET_PATH', path.replace(/\/\//g, '/'));
    },

    async setIndex({ commit }: any, index: number) {
      commit('SET_INDEX', index);
    },
  },
};
