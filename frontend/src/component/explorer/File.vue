<template>
  <div :class="$style.file">
    <ui-icon name="folder" :width="48" color="#feb000" />
    <div :class="$style.name">{{ file.name }}</div>
  </div>
</template>

<script lang="ts">
import { defineComponent, PropType } from 'vue';
import { IFile } from '@/util/Types';

export default defineComponent({
  props: {
    file: { type: Object as PropType<IFile> },
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
  flex-direction: column;

  svg {
    &:hover {
      opacity: 0.7;
    }
  }

  svg,
  .name {
    margin: 0 auto;
    cursor: pointer;
  }
}
</style>
