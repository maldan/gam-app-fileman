import { RestApi } from '@/util/RestApi';
import { ActionContext } from 'vuex';
import { MainTree } from '@/store/index';
import Axios from 'axios';
import { Helper } from '@/util/Helper';

export interface FileInfo {
  user: string;
  created: string;
  index: number;
  isSelected: boolean;
  kind: string;
  name: string;
  path: string;
  size: number;
  tags: Record<string, string | number | boolean>;
}

export interface FileStore {
  list: FileInfo[];
  buffer: any[];
  lastSelected: unknown;
  bufferMode: string;
  sortBy: string;
}
export type FileActionContext = ActionContext<FileStore, MainTree>;

export default {
  namespaced: true,
  state(): FileStore {
    return {
      list: [] as FileInfo[],
      buffer: [],
      lastSelected: null,
      bufferMode: 'copy',
      sortBy: 'name',
    };
  },
  mutations: {
    SET_FILES(state: FileStore, payload: any): void {
      state.list = payload;
    },
    SET_LAST_SELECTED(state: FileStore, payload: any): void {
      state.lastSelected = payload;
    },
    SET_BUFFER(state: FileStore, payload: any): void {
      state.buffer = payload.data;
      state.bufferMode = payload.mode;
    },
    SET_SORT(state: FileStore, sortBy: string): void {
      state.sortBy = sortBy;
    },
    SET_TAGS(state: FileStore, tags: any): void {
      // state.tags = tags;
    },
  },
  actions: {
    async clearList(action: FileActionContext): Promise<void> {
      action.commit('SET_FILES', []);
    },
    async getList({ commit, dispatch, rootState }: any) {
      dispatch('main/setLoading', true, { root: true });
      const files = await RestApi.file.getList(rootState.main.path);
      commit(
        'SET_FILES',
        files.map((x: any, i: number) => {
          x.index = i;
          return x;
        }),
      );
      dispatch('main/setLoading', false, { root: true });

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
    switchSort({ commit, state, dispatch }: any) {
      if (state.sortBy === 'name') {
        commit('SET_SORT', 'size');
        dispatch('sort');
      } else if (state.sortBy === 'size') {
        commit('SET_SORT', 'created');
        dispatch('sort');
      } else if (state.sortBy === 'created') {
        commit('SET_SORT', 'name');
        dispatch('sort');
      }
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
    copySelectedToBuffer({ rootState, state, commit }: any, payload: string) {
      const buffer = [] as any[];
      for (let i = 0; i < state.list.length; i++) {
        if (state.list[i].isSelected) {
          buffer.push((rootState.main.path + '/' + state.list[i].name).replace(/\/\//g, '/'));
        }
      }
      commit('SET_BUFFER', { data: buffer, mode: payload });
    },
    clearBuffer({ commit }: any) {
      commit('SET_BUFFER', { data: [], mode: 'copy' });
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
    async delete({ dispatch, state }: any) {
      for (let i = 0; i < state.buffer.length; i++) {
        await RestApi.file.deleteAny(state.buffer[i]);
      }
      dispatch('clearBuffer');
      dispatch('getListSilent');
    },
    async paste({ dispatch, state, rootState }: any) {
      dispatch('main/setLoading', true, { root: true });
      for (let i = 0; i < state.buffer.length; i++) {
        if (state.bufferMode === 'copy') {
          await RestApi.file.copy(state.buffer[i], rootState.main.path);
        } else {
          await RestApi.file.move(state.buffer[i], rootState.main.path);
        }
      }
      dispatch('clearBuffer');
      dispatch('getListSilent');
      dispatch('main/setLoading', false, { root: true });
    },
    async move(action: FileActionContext, path: string): Promise<void> {
      await action.dispatch('main/setLoading', true, { root: true });
      for (let i = 0; i < action.state.buffer.length; i++) {
        await RestApi.file.move(action.state.buffer[i], path);
      }
      await action.dispatch('clearBuffer');
      await action.dispatch('getListSilent');
      await action.dispatch('main/setLoading', false, { root: true });
    },
    async upload({ dispatch, rootState }: any, payload: any) {
      for (let i = 0; i < payload.length; i++) {
        const path = rootState.main.path + '/' + new Date().getTime() + '_' + payload[i].name;
        await RestApi.file.uploadFile(path, payload[i]);
      }
      dispatch('getListSilent');
    },
    async setVideoPreview({ dispatch }: any, payload: any) {
      await RestApi.file.setVideoPreview(payload.path, payload.time);
      dispatch('getListSilent');
    },
    async setFileHashName(action: FileActionContext): Promise<void> {
      const selectedList = action.state.list.filter((file) => file.isSelected);
      for (const file of selectedList) {
        try {
          await Axios.post(`${action.rootState.main.API_URL}/file/setHashName`, {
            path: Helper.fixPath(`${action.rootState.main.path}/${file.name}`),
          });
        } catch (e) {
          //
        }
      }
      await action.dispatch('getListSilent');
    },
    async setTags(action: FileActionContext): Promise<void> {
      const selectedList = action.state.list.filter((file) => file.isSelected);
      const file = selectedList[0];
      if (file) {
        const tags: Record<string, unknown> = {};
        for (const x of action.rootState.modal.data.tags) {
          if (x.value === 'true') tags[x.key] = true;
          else if (x.value === 'false') tags[x.key] = false;
          else if (typeof x.value === 'string' && x.value.match(/^[+-]?\d+(\.\d+)?$/)) {
            tags[x.key] = Number.parseFloat(x.value);
          } else tags[x.key] = x.value;
        }
        await Axios.post(`${action.rootState.main.API_URL}/file/tags`, {
          path: file.path,
          tags: tags,
        });
      }

      await action.dispatch('getListSilent');
    },
  },
};
