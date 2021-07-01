<template>
  <div :class="$style.body">
    <div :class="$style.header">
      <img
        class="clickable"
        @click="back()"
        src="../asset/back.svg"
        alt=""
        style="margin-right: 20px"
        draggable="false"
      />
      <input type="text" @change="$emit('update:path', $event.target.value)" :value="path" />
    </div>

    <!-- Tile mode -->
    <div
      v-if="mode === 'tile' && !isLoading"
      ref="file_list"
      :class="[$style.file_tile, $style[`g${maxRows}`]]"
    >
      <div :class="$style.file" v-for="x in files" :key="x.name">
        <img @click="go(x)" class="clickable" :src="iconByFile(x)" alt="Folder" draggable="false" />
        <div :class="$style.name">{{ x.name }}</div>
      </div>
    </div>

    <!-- List mode -->
    <div v-if="mode === 'list' && !isLoading" ref="file_list" :class="[$style.file_list]">
      <div @click="go(x)" class="clickable" :class="$style.file" v-for="x in files" :key="x.name">
        <img :src="iconByFile(x)" alt="Folder" draggable="false" />
        <div :class="$style.name">{{ x.name }}</div>
        <div :class="$style.size">{{ x.size }}</div>
        <div :class="$style.created">{{ $root.moment(x.created).fromNow() }}</div>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import { RestApi } from '../util/RestApi';

export default defineComponent({
  props: {
    path: String,
    mode: String,
  },
  components: {},
  async mounted() {
    this.refresh();

    this.intervalId = setInterval(() => {
      if (!(this.$refs['file_list'] as any)) {
        return;
      }
      const w = (this.$refs['file_list'] as any).getBoundingClientRect().width;
      if (w < 600) {
        this.maxRows = 4;
      } else if (w < 400) {
        this.maxRows = 3;
      } else {
        this.maxRows = 5;
      }
    }, 16);
  },
  beforeUnmount() {
    clearInterval(this.intervalId);
  },
  watch: {
    path() {
      this.refresh();
    },
  },
  methods: {
    async refresh() {
      this.isLoading = true;
      const all = await RestApi.file.getList(this.path + '');
      const files = all
        .filter((x: any) => x.kind !== 'dir')
        .sort((a: any, b: any) => {
          return a.name.localeCompare(b.name);
        });
      const folder = all
        .filter((x: any) => x.kind === 'dir')
        .sort((a: any, b: any) => {
          return a.name.localeCompare(b.name);
        });

      this.files = [...folder, ...files];
      this.isLoading = false;
    },
    back() {
      const t = this.path?.split('/').slice(0, -1) || [];

      if (t.length === 1) {
        this.$emit('update:path', t.join('/') + '/');
      } else {
        this.$emit('update:path', t.join('/'));
      }
    },
    go(file: any) {
      if (file.kind !== 'dir') {
        return;
      }
      this.$emit('update:path', (this.path + '/' + file.name).replace(/\/\//g, '/'));
    },
    iconByFile(file: any) {
      if (file.kind === 'file') return require('../asset/file.svg');
      return require('../asset/folder.svg');
    },
  },
  data: () => {
    return {
      files: [] as any[],
      maxRows: 4,
      intervalId: 0,
      isLoading: false,
    };
  },
});
</script>

<style lang="scss" module>
.body {
  background: rgba(0, 0, 0, 0.3);
  padding: 10px;
  flex: 1;

  .header {
    display: flex;
    align-items: center;
    padding-left: 5px;

    input {
      flex: 1;
      padding: 5px 10px;
      background: rgba(255, 255, 255, 0.15);
      border: 0;
      border-radius: 4px;
      outline: none;
      color: #a5a5a5;
      font-size: 16px;
    }
  }

  .file_list {
    margin-top: 10px;
    height: calc(100% - 35px - 10px);
    overflow-y: auto;
    color: #dbdbdb;

    .file {
      display: flex;
      padding: 10px;
      align-items: center;

      img {
        width: 32px;
        display: block;
        margin-right: 10px;
      }

      .name {
        text-overflow: ellipsis;
        overflow: hidden;
        white-space: nowrap;
        word-wrap: break-word;
        // flex-grow: 0;
        max-width: 200px;
        text-align: center;
        font-size: 14px;
        margin-top: 5px;
        flex: 2;
        text-align: left;
      }

      .size {
        flex: 1;
        text-align: left;
      }
      .created {
        flex: 1;
        text-align: left;
      }
    }
  }

  .file_tile {
    margin-top: 10px;
    height: calc(100% - 35px - 10px);
    overflow-y: auto;
    color: #dbdbdb;
    display: grid;
    grid-template-columns: repeat(4, 1fr);
    gap: 10px;
    grid-auto-rows: 100px;

    &.g6 {
      grid-template-columns: repeat(6, 1fr);
    }
    &.g5 {
      grid-template-columns: repeat(5, 1fr);
    }
    &.g4 {
      grid-template-columns: repeat(4, 1fr);
    }
    &.g3 {
      grid-template-columns: repeat(3, 1fr);
    }
    &.g2 {
      grid-template-columns: repeat(2, 1fr);
    }

    .file {
      display: flex;
      flex-direction: column;
      align-items: center;
      width: 100%;
      height: 100%;

      img {
        display: block;
        margin-top: auto;
      }

      .name {
        text-overflow: ellipsis;
        overflow: hidden;
        white-space: nowrap;
        word-wrap: break-word;
        max-width: 100px;
        text-align: center;
        font-size: 14px;
        margin-top: 15px;
        margin-bottom: auto;
      }
    }
  }
}
</style>
