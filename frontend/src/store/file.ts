import { RestApi } from '@/util/RestApi';

export default {
  namespaced: true,
  state() {
    return {
      list: [],
      buffer: [],
      lastSelected: null,
    };
  },
  mutations: {
    SET_FILES(state: any, payload: any) {
      state.list = payload;
    },
    SET_LAST_SELECTED(state: any, payload: any) {
      state.lastSelected = payload;
    },
    SET_BUFFER(state: any, payload: any) {
      state.buffer = payload;
    },
  },
  actions: {
    async getList({ commit, dispatch, rootState }: any) {
      commit('main/SET_LOADING', true, { root: true });
      const files = await RestApi.file.getList(rootState.main.path);
      commit(
        'SET_FILES',
        files.map((x: any, i: number) => {
          x.index = i;
          return x;
        }),
      );
      commit('main/SET_LOADING', false, { root: true });

      dispatch('sort');
    },
    async getListSilent({ commit, dispatch, rootState }: any) {
      const files = await RestApi.file.getList(rootState.main.path);
      commit(
        'SET_FILES',
        files.map((x: any, i: number) => {
          x.index = i;
          return x;
        }),
      );
      dispatch('sort');
    },
    sort({ state, commit }: any) {
      let folders = state.list.filter((x: any) => x.kind === 'dir');
      let files = state.list.filter((x: any) => x.kind !== 'dir');

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
    clearSelection({ state, commit }: any) {
      const files = state.list;
      for (let i = 0; i < files.length; i++) {
        files[i].isSelected = false;
      }
      commit('SET_FILES', files);
      commit('SET_LAST_SELECTED', null);
    },
    select({ state, commit }: any, payload: any) {
      const files = state.list;
      for (let i = 0; i < files.length; i++) {
        if (files[i].name === payload.name) {
          files[i].isSelected = true;
          commit('SET_LAST_SELECTED', files[i]);
          return;
        }
      }
      commit('SET_FILES', files);
    },
    copySelectedToBuffer({ rootState, state, commit }: any) {
      const buffer = [] as any[];
      for (let i = 0; i < state.list.length; i++) {
        if (state.list[i].isSelected) {
          buffer.push((rootState.main.path + '/' + state.list[i].name).replace(/\/\//g, '/'));
        }
      }
      commit('SET_BUFFER', buffer);
    },
    clearBuffer({ commit }: any) {
      commit('SET_BUFFER', []);
    },
    async createFile({ dispatch, rootState }: any) {
      await RestApi.file.createFile(rootState.main.path + '/' + rootState.modal.data.name);
      dispatch('getListSilent');
    },
    async createDir({ dispatch, rootState }: any) {
      await RestApi.file.createDir(rootState.main.path + '/' + rootState.modal.data.name);
      dispatch('getListSilent');
    },
    async rename({ dispatch, rootState }: any) {
      await RestApi.file.rename(
        rootState.file.buffer[0],
        rootState.main.path + '/' + rootState.modal.data.name,
      );
      dispatch('clearBuffer');
      dispatch('getListSilent');
    },
  },
};
