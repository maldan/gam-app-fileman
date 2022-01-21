<template>
  <ui-block title="File" :scrollY="true">
    <div v-for="x in 103" :key="x" :class="$style.file">
      <ui-icon name="arrow_right_small" />
      <ui-icon name="folder" color="#feb000" />
      <div :class="$style.name">home</div>
    </div>
  </ui-block>
</template>

<script lang="ts">
import { defineComponent, PropType } from 'vue';

export default defineComponent({
  props: {
    path: String,
  },
  async mounted() {},
  methods: {},
  data: () => {
    return {};
  },
});
</script>

<style lang="scss" module>
.file {
  display: flex;
  align-items: center;
  cursor: pointer;

  svg {
    &:hover {
      opacity: 0.7;
    }
  }

  .name {
    margin-left: 5px;
    &:hover {
      opacity: 0.7;
    }
  }
}
</style>
