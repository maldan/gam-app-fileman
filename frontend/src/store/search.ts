import { ActionContext } from 'vuex';
import { MainTree } from '@/store/index';
import { FileInfo } from '@/store/file';
import Axios from 'axios';

export interface SearchStore {
  query: string;
}
export type SearchActionContext = ActionContext<SearchStore, MainTree>;

export default {
  namespaced: true,
  state(): SearchStore {
    return {
      query: '',
    };
  },
  mutations: {},
  actions: {
    async search(action: SearchActionContext): Promise<void> {
      const list = (
        await Axios.get(`${action.rootState.main.API_URL}/file/search?query=${action.state.query}`)
      ).data.response;
      action.commit(
        'file/SET_FILES',
        list.map((x: FileInfo, index: number) => {
          x.index = index;
          return x;
        }),
        { root: true },
      );
    },
  },
};
