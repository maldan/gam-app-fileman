<template>
  <div
    v-doubleclick="
      () => {
        open();
      }
    "
    :class="[$style.file, 'clickable', file.isSelected ? $style.selected : '']"
  >
    <ui-icon
      :class="$style.icon"
      v-if="!isUseImage && file.kind === 'dir'"
      name="folder"
      color="#dd9c0b"
    />
    <ui-icon
      :class="$style.icon"
      v-if="!isUseImage && file.kind === 'file'"
      name="file"
      color="#9d9d9d"
    />

    <div :class="$style.image" v-if="isUseImage">
      <div :class="$style.info">
        <div v-if="file.name.split('.')[0].length === 40" :class="$style.hash">H</div>
        <div v-if="file.tags?.path" :class="$style.tag">T</div>
        <div :class="$style.size">{{ $root.pretty(file.size) }}</div>
      </div>
      <img :src="icon(file)" alt="File" loading="lazy" draggable="false" />
    </div>
    <div :class="$style.name">{{ file.tags?.title || file.name }}</div>
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import { FileInfo } from '@/store/file';

export default defineComponent({
  props: {
    file: {
      type: Object,
      required: true,
    },
  },
  computed: {
    isUseImage() {
      if (this.file.kind === 'file') {
        if (this.file.name.match(/\.(png|jpeg|gif|jpg|webp)$/)) return true;
        if (this.file.name.match(/\.(mp4|avi)$/)) return true;
      }
      return false;
    },
  },
  components: {},
  async mounted() {},
  methods: {
    icon(file: FileInfo): any {
      if (file.kind === 'file') {
        if (file.name.match(/\.(png|jpeg|gif|jpg|webp)$/)) {
          return `${(this.$root as any).API_URL}/file/imageThumbnail?path=${file.path}`;
        }
        if (file.name.match(/\.(mp4|avi)$/)) {
          return `${(this.$root as any).API_URL}/file/videoThumbnail?path=${file.path}&time=${
            file.tags?.videoPreviewTime || '00:00:00'
          }`;
        }
        // @ts-ignore
        return require('../gam_sdk_ui/vue/asset/icon/file.svg');
      }
      // @ts-ignore
      return require('../gam_sdk_ui/vue/asset/icon/folder.svg');
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
        } else if (this.file.name.match(/\.(mp4|avi)$/)) {
          this.$store.dispatch('extension/show', {
            name: 'video',
            data: {
              index: this.file.index,
              fullPath: this.$store.state.main.path + '/' + this.file.name,
            },
          });
        } else {
          window.open(
            `${(this.$root as any).API_URL}/file/file?path=${
              this.$store.state.main.path + '/' + this.file.name
            }`,
            '_blank',
          );
        }
      }
    },
  },
  data: () => {
    return {};
  },
});
</script>

<style module lang="scss">
.file {
  display: flex;
  flex-direction: column;
  align-items: center;
  // justify-content: center;
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

  .icon {
    height: 64px;
  }

  .image {
    position: relative;

    .info {
      position: absolute;
      right: 5px;
      bottom: 8px;
      display: flex;

      .tag,
      .size,
      .hash {
        background-color: #19aa00;
        padding: 2px 5px;
        border-radius: 4px;
        font-size: 12px;
        margin-left: 5px;
      }

      .tag {
        background-color: #e75a25;
      }

      .hash {
        background-color: #bd25e7;
      }
    }

    img {
      max-width: 100%;
    }
  }
}
</style>
