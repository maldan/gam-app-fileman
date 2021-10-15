<template>
  <div :class="$style.tabs">
    <ui-select
      @change="
        $store.dispatch('main/changePath', $event.value),
          $store.dispatch('tab/setIndex', $event.index)
      "
      icon="rename"
      :items="$store.state.tab.tabs.map((x) => ({ label: x, value: x }))"
      v-model="$store.state.tab.tab"
    />
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';

export default defineComponent({
  props: {},
  components: {},
  async mounted() {},
  methods: {},
  data: () => {
    return {
      path: '',
    };
  },
});
</script>

<style lang="scss" module>
.tabs {
}
</style>
