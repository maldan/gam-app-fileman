import { RestApi } from '@/util/RestApi';
import { ActionContext } from 'vuex';
import { MainTree } from '.';
import Axios from 'axios';

export interface MainStore {
  API_URL: string;
  path: string;
  isLoading: boolean;
}
export type MainActionContext = ActionContext<MainStore, MainTree>;

export default {
  namespaced: true,
  state(): MainStore {
    return {
      API_URL: process.env.VUE_APP_API_URL || `${window.location.origin}/api`,
      path: '/',
      isLoading: false,
    };
  },
  mutations: {
    SET_PATH(state: MainStore, path: string): void {
      state.path = path;
    },

    SET_LOADING(state: MainStore, status: boolean): void {
      state.isLoading = status;
    },
  },
  actions: {
    async getPath(action: MainActionContext): Promise<void> {
      action.commit(
        'SET_PATH',
        (await Axios.get(`${action.rootState.main.API_URL}/file/path`)).data.response,
      );
      await action.dispatch('file/getList', null, { root: true });
    },
    async changePath(action: MainActionContext, path: string): Promise<void> {
      action.commit('SET_PATH', path.replace(/\/\//g, '/'));
      action.commit('SET_LOADING', true);
      await RestApi.file.setPath(action.state.path);
      action.commit('SET_LOADING', false);
      await action.dispatch('file/getList', null, { root: true });
      await action.dispatch('file/clearSelection', null, { root: true });
      await action.dispatch('tab/changePath', path, { root: true });
    },
    setLoading(action: MainActionContext, status: boolean): void {
      action.commit('SET_LOADING', status);
    },
  },
};
