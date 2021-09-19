<template>
  <div :class="$style.header">
    <button
      v-for="x in buttons"
      :key="x.name"
      class="clickable"
      :style="{ opacity: x.isNeedFile ? 0.5 : 1 }"
    >
      <img :src="require(`../asset/icon/${x.icon}.svg`)" alt="" />
      {{ x.name }}
    </button>
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
      buttons: [
        {
          icon: 'add_folder',
          name: 'New Folder',
        },
        {
          icon: 'add_file',
          name: 'New File',
        },
        {
          icon: 'add_file',
          name: 'Rename',
          isNeedFile: true,
        },
        {
          icon: 'add_file',
          name: 'Delete',
        },
        {
          icon: 'add_file',
          name: 'Copy',
        },
        {
          icon: 'add_file',
          name: 'Cut',
        },
        {
          icon: 'add_file',
          name: 'Paste',
        },
        {
          icon: 'upload',
          name: 'Upload',
        },
      ],
    };
  },
});
</script>

<style lang="scss" module>
.header {
  display: flex;
  flex: 1;

  button {
    background: none;
    color: #dadada;
    display: flex;
    flex-direction: column;
    outline: none;
    border: 0;
    align-items: center;
    justify-content: center;
    padding: 10px;
    background: #1b1b1b;
    flex: 1;
    text-transform: uppercase;
    font-size: 12px;

    img {
      display: block;
      margin-bottom: 12px;
    }
  }
}
</style>
