import { createStore, Store } from 'vuex';
import main, { MainStore } from './main';
import modal, { ModalStore } from '../gam_sdk_ui/vue/store/modal';
import { InjectionKey } from 'vue';

// define your typings for the store state
export interface MainTree {
  main: MainStore;
  modal: ModalStore;
}

// define injection key
export const key: InjectionKey<Store<MainTree>> = Symbol();

export default createStore<MainTree>({
  modules: { main, modal },
});
