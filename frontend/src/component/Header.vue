<template>
  <div :class="$style.header">
    <button
      @click="x.onClick()"
      v-for="x in buttons"
      :key="x.name"
      class="clickable"
      :style="{
        opacity:
          (x.isNeedFile && !$store.state.file.lastSelected) ||
          (x.isNeedBuffer && !$store.state.file.buffer.length)
            ? 0.5
            : 1,
      }"
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
  async mounted() {
    this.buttons = [
      {
        icon: 'add_folder',
        name: 'New Dir',
        onClick: () => {
          this.$store.dispatch('modal/show', {
            name: 'modalName',
            data: { title: 'Add new directory', name: '' },
            func: () => {
              this.$store.dispatch('file/createDir');
            },
          });
        },
      },
      {
        icon: 'add_file',
        name: 'New File',
        onClick: () => {
          this.$store.dispatch('modal/show', {
            name: 'modalName',
            data: { title: 'Add new file', name: '' },
            func: () => {
              this.$store.dispatch('file/createFile');
            },
          });
        },
      },
      {
        icon: 'add_file',
        name: 'Rename',
        isNeedFile: true,
        onClick: () => {
          this.$store.dispatch('file/copySelectedToBuffer');

          this.$store.dispatch('modal/show', {
            name: 'modalName',
            data: {
              title: 'Rename file/folder',
              name: this.$store.state.file.lastSelected.name,
            },
            func: () => {
              this.$store.dispatch('file/rename');
            },
          });
        },
      },
      {
        icon: 'add_file',
        name: 'Delete',
        isNeedFile: true,
        onClick: () => {
          this.$store.dispatch('main/copySelectedBuffer');
        },
      },
      {
        icon: 'add_file',
        name: 'Copy',
        isNeedFile: true,
        onClick: () => {
          this.$store.dispatch('main/copySelectedBuffer');
        },
      },
      {
        icon: 'add_file',
        name: 'Cut',
        isNeedFile: true,
        onClick: () => {
          this.$store.dispatch('main/copySelectedBuffer');
        },
      },
      {
        icon: 'add_file',
        name: 'Paste',
        isNeedBuffer: true,
      },
      {
        icon: 'upload',
        name: 'Upload',
      },
    ];
  },
  methods: {},
  data: () => {
    return {
      buttons: [] as any[],
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
