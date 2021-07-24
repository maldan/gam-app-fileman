<template>
  <div class="header">
    <IconButton @click="back()" icon="back" style="margin-right: 15px" />
    <Path @update="$emit('update', $event)" :value="path" style="flex: 1" />
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import IconButton from '../component/IconButton.vue';
import Path from '../component/Path.vue';
import Item from '../component/Item.vue';

export default defineComponent({
  props: {
    path: String,
  },
  components: { IconButton, Path, Item },
  async mounted() {},

  methods: {
    back() {
      const t = this.path?.split('/').slice(0, -1) || [];

      if (t.length === 1) {
        this.$emit('update', t.join('/') + '/');
      } else {
        this.$emit('update', t.join('/'));
      }
    },
  },
  data: () => {
    return {};
  },
});
</script>

<style lang="scss" scoped>
.header {
  display: flex;
  justify-content: flex-end;
  background: #1d1d1d;
  padding: 10px;
  box-sizing: border-box;
}
</style>
