<template>
  <div :class="$style.main">
    <div :class="$style.x">
      <tree-list />
      <explorer-content style="flex: 1" />
      <explorer-content style="flex: 1" />
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';

export default defineComponent({
  components: {},
  async mounted() {},

  methods: {},
  data: () => {
    return {};
  },
});
</script>

<style lang="scss" module>
.main {
  height: 100%;
  box-sizing: border-box;
  padding: 10px;

  .x {
    display: grid;
    grid-template-columns: 280px 1fr 1fr;
    gap: 10px;
  }
}
</style>
