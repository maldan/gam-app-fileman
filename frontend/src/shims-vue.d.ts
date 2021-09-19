import type { DefineComponent } from 'vue';

/* eslint-disable */
declare module '*.vue' {
  const component: DefineComponent<{}, {}, any>;
  export default component;
}

/* eslint-disable */
declare module '@vue/runtime-core' {
  import { ComponentCustomProperties } from 'vue';
  import { Store } from 'vuex';

  interface ComponentCustomProperties {
    $store: Store<any>;
  }
}
