<template>
  <div
    v-doubleclick="
      () => {
        open();
      }
    "
    class="file clickable"
    :class="file.isSelected ? 'selected' : ''"
  >
    <img :class="'icon'" :src="icon(file)" alt="File" loading="lazy" draggable="false" />
    <div class="name">{{ file.name }}</div>
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';

export default defineComponent({
  props: {
    file: {
      type: Object,
      required: true,
    },
  },
  components: {},
  async mounted() {},
  methods: {
    icon(file: any): any {
      if (file.kind === 'file') {
        if (file.name.match(/\.(png|jpeg|gif|jpg|webp)$/)) {
          return `${(this.$root as any).API_URL}/file/imageThumbnail?path=${
            this.$store.state.main.path
          }/${file.name}`;
        }
        if (file.name.match(/\.(mp4|avi)$/)) {
          return `${(this.$root as any).API_URL}/file/videoThumbnail?path=${
            this.$store.state.main.path
          }/${file.name}`;
        }
        // @ts-ignore
        return require('../asset/file/file.svg');
      }
      // @ts-ignore
      return require('../asset/file/folder.svg');
    },
    open() {
      if (this.file.kind === 'dir') {
        this.$store.dispatch('main/changePath', this.$store.state.main.path + '/' + this.file.name);
      }
      if (this.file.kind === 'file') {
        if (this.file.name.match(/\.(png|jpeg|gif|jpg|webp)$/)) {
          this.$store.dispatch('extension/show', {
            name: 'image',
            data: {
              index: this.file.index,
              fullPath: this.$store.state.main.path + '/' + this.file.name,
            },
          });
        }
      }
    },
  },
  data: () => {
    return {};
  },
});
</script>

<style lang="scss">
.file {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  padding: 10px;
  box-sizing: border-box;

  .name {
    margin-top: 10px;
    text-align: center;
    word-break: break-all;
  }

  &.selected {
    background: #2d4b65;
  }

  > img {
    max-width: 100%;
  }
}
</style>
