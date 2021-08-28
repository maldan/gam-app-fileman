<template>
  <div class="bottom">
    <input ref="file_input" type="file" multiple style="position: fixed; top: -100px" />

    <!-- Preload -->
    <img
      v-if="isLoading"
      class="rotating"
      src="../asset/preload.svg"
      style="width: 16px; margin-right: 20px"
    />

    <!-- Tabs -->
    <div style="margin-right: auto; display: flex">
      <div
        v-for="(x, i) in tabs"
        @click="$emit('tab', i)"
        :key="i"
        class="tab clickable"
        :class="x === tab ? 'selected' : ''"
      >
        {{ i + 1 }}
      </div>
      <div class="tab clickable" @click="$emit('open', 'download')">Download</div>
    </div>

    <div @click="$emit('changeView', !isList)" class="selected_files clickable">
      {{ isList ? 'List' : 'Block' }}
    </div>

    <!-- Sort mode -->
    <div
      v-if="list.filter((x) => x.isSelected).length"
      @click="$emit('copy')"
      class="selected_files clickable"
    >
      COPY
    </div>
    <div
      v-if="list.filter((x) => x.isSelected).length"
      @click="$emit('cut')"
      class="selected_files clickable"
    >
      CUT
    </div>
    <div v-if="buffer.length" @click="$emit('paste')" class="selected_files clickable">PASTE</div>

    <!-- Sort mode -->
    <div @click="du()" class="selected_files clickable">DU</div>

    <!-- Sort mode -->
    <div @click="changeSortBy()" class="selected_files clickable">By {{ sort }}</div>

    <div @click="unselect()" class="selected_files clickable">
      {{ list.filter((x) => x.isSelected).length }} / {{ list.length }}
    </div>

    <IconButton @click="selectFiles()" icon="upload" style="margin-right: 15px" />
    <IconButton @click="addDir()" icon="add_folder" style="margin-right: 15px" />
    <IconButton @click="addFile()" icon="add_file" />
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import { RestApi } from '../util/RestApi';
import IconButton from './IconButton.vue';

export default defineComponent({
  props: {
    isLoading: Boolean,
    path: String,
    list: Array,
    sort: String,
    buffer: Array,
    tabs: Array,
    tab: Object,
    isList: Boolean,
  },
  components: { IconButton },
  async mounted() {},
  methods: {
    async addFile() {
      const name = prompt('Enter file name');
      if (name) {
        await RestApi.file.createFile(this.path + '/' + name);
        this.$emit('refresh');
      }
    },
    async addDir() {
      const name = prompt('Enter dir name');
      if (name) {
        await RestApi.file.createDir(this.path + '/' + name);
        this.$emit('refresh');
      }
    },
    unselect() {
      this.list?.forEach((x: any) => {
        x.isSelected = false;
        return x;
      });
    },
    du() {
      this.$emit('du');
    },
    changeSortBy() {
      if (this.sort === 'name') {
        this.$emit('changeSort', 'size');
      }
      if (this.sort === 'size') {
        this.$emit('changeSort', 'created');
      }
      if (this.sort === 'created') {
        this.$emit('changeSort', 'name');
      }
      this.$emit('refresh');
    },
    selectFiles() {
      const fileInput = this.$refs['file_input'] as HTMLInputElement;

      if (fileInput) {
        fileInput.onchange = (e: any) => {
          this.upload(e.target.files);
        };
        fileInput.click();
      }
    },
    async upload(files: FileList) {
      const items = Array.from(files).filter(
        (x: File) => x.type.match(/^image\//) || x.name.toLowerCase().match('.heic'),
      );
      let amount = items.length;

      for (let i = 0; i < items.length; i++) {
        (async () => {
          try {
            await RestApi.file.uploadFile(
              this.path + '/' + Math.random() + '_' + items[i].name,
              items[i],
              (e) => {
                // this.progress[i] = e.loaded / items[i].size;
              },
            );
            // this.done[i] = true;
            this.$emit('refresh');
          } catch {}

          amount -= 1;
          if (amount <= 0) {
          }
        })();
      }
    },
  },
  data: () => {
    return {
      isEditMode: false,
    };
  },
});
</script>

<style lang="scss" scoped>
.bottom {
  display: flex;
  justify-content: flex-end;
  background: #1d1d1d;
  padding: 10px;
  box-sizing: border-box;

  .selected_files {
    padding: 0 10px;
    background: #2c589e;
    color: #b5b5b5;
    border-radius: 3px;
    margin-right: 15px;
    font-size: 12px;
    display: flex;
    align-items: center;
    font-weight: bold;
  }

  .tab {
    padding: 0 10px;
    background: #686868;
    color: #b5b5b5;
    border-radius: 3px;
    margin-right: 15px;
    font-size: 12px;
    display: flex;
    align-items: center;
    font-weight: bold;

    &.selected {
      background: #2c589e;
    }
  }
}
</style>
