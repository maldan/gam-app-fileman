import { RestApi } from '@/util/RestApi';
import { createStore } from 'vuex';
import main from './main';
import modal from './modal';
import file from './file';
import tab from './tab';
import extension from './extension';

export default createStore({
  modules: { main, modal, file, tab, extension },
  /*state() {
    return {
      sortBy: 'name',
      path: '/',
      files: [],
      buffer: [],
      lastSelectedFile: null,
      isLoading: false,
      modal: '',
      modalData: {},
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
    SET_BUFFER(state: any, payload: any) {
      state.buffer = payload;
    },
    SET_MODAL(state: any, payload: any) {
      state.modal = payload.name;
      state.modalData = payload.data;
    },
  },
  actions: {
    async getPath({ commit, dispatch }) {
      commit('SET_PATH', await RestApi.file.getPath());
      dispatch('getFiles');
    },
    async changePath({ state, commit, dispatch }, payload: any) {
      commit('SET_PATH', payload.replace(/\/\//g, '/'));
      await RestApi.file.setPath(state.path);
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
      commit('SET_LAST_SELECTED_FILE', null);
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
    copySelectedBuffer({ state, commit }) {
      const buffer = [] as any[];
      for (let i = 0; i < state.files.length; i++) {
        if (state.files[i].isSelected) {
          buffer.push((state.path + '/' + state.files[i].name).replace(/\/\//g, '/'));
        }
      }
      commit('SET_BUFFER', buffer);
    },
    showModal({ commit }, payload) {
      commit('SET_MODAL', payload);
    },
    closeModal({ commit }) {
      commit('SET_MODAL', {});
    },
    modalOk({ commit }) {
      commit('SET_MODAL', {});
    },
  },*/
});
