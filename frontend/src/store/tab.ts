import { ActionContext } from 'vuex';
import { MainTree } from '@/store/index';

export interface TabStore {
  tabs: string[];
  tab: { label: string; value: string };
  tabIndex: number;
}
export type TabActionContext = ActionContext<TabStore, MainTree>;

export default {
  namespaced: true,
  state(): TabStore {
    return {
      tabs: ['/', '/home'],
      tab: { label: '/', value: '/' },
      tabIndex: 0,
    };
  },
  mutations: {
    SET_PATH(state: TabStore, path: string): void {
      state.tabs[state.tabIndex] = path;
      state.tab.label = path;
      state.tab.value = path;
    },
    SET_INDEX(state: TabStore, index: number): void {
      state.tabIndex = index;
    },
  },
  actions: {
    async changePath(action: TabActionContext, path: string): Promise<void> {
      action.commit('SET_PATH', path.replace(/\/\//g, '/'));
    },

    async setIndex(action: TabActionContext, index: number): Promise<void> {
      action.commit('SET_INDEX', index);
    },
  },
};
