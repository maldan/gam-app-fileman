import { RestApi } from '@/util/RestApi';
import { createStore } from 'vuex';

export default createStore({
  state() {
    return {
      sortBy: 'name',
      path: '/',
      files: [],
      lastSelectedFile: null,
      isLoading: false,
    };
  },
  mutations: {
    SET_PATH(state: any, payload: any) {
      state.path = payload;
    },
    SET_FILES(state: any, payload: any) {
      state.files = payload;
    },
    SET_LOADING(state: any, payload: any) {
      state.isLoading = payload;
    },
    SET_LAST_SELECTED_FILE(state: any, payload: any) {
      state.lastSelectedFile = payload;
    },
  },
  actions: {
    changePath({ commit, dispatch }, payload: any) {
      commit('SET_PATH', payload.replace(/\/\//g, '/'));
      dispatch('getFiles');
    },
    async getFiles({ state, commit, dispatch }) {
      commit('SET_LOADING', true);
      const files = await RestApi.file.getList(state.path);
      commit(
        'SET_FILES',
        files.map((x: any, i: number) => {
          x.index = i;
          return x;
        }),
      );
      commit('SET_LOADING', false);

      dispatch('sortFiles');
    },
    sortFiles({ state, commit }) {
      let folders = state.files.filter((x: any) => x.kind === 'dir');
      let files = state.files.filter((x: any) => x.kind !== 'dir');

      if (state.sortBy === 'name') {
        folders = folders.sort((a: any, b: any) => a.name.localeCompare(b.name));
        files = files.sort((a: any, b: any) => a.name.localeCompare(b.name));
      }
      if (state.sortBy === 'size') {
        folders = folders.sort((a: any, b: any) => a.size - b.size);
        files = files.sort((a: any, b: any) => a.size - b.size);
      }
      if (state.sortBy === 'created') {
        folders = folders.sort(
          (a: any, b: any) => new Date(a.created).getTime() - new Date(b.created).getTime(),
        );
        files = files.sort(
          (a: any, b: any) => new Date(a.created).getTime() - new Date(b.created).getTime(),
        );
      }

      commit(
        'SET_FILES',
        [...folders, ...files].map((x: any, i: number) => {
          x.index = i;
          return x;
        }),
      );
    },
    clearSelection({ state, commit }) {
      const files = state.files;
      for (let i = 0; i < files.length; i++) {
        files[i].isSelected = false;
      }
      commit('SET_FILES', files);
    },
    selectFile({ state, commit }, payload) {
      const files = state.files;
      for (let i = 0; i < files.length; i++) {
        if (files[i].name === payload.name) {
          files[i].isSelected = true;
          commit('SET_LAST_SELECTED_FILE', files[i]);
          return;
        }
      }
      commit('SET_FILES', files);
    },
  },
});
