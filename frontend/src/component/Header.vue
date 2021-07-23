<template>
  <div class="header">
    <IconButton @click="back()" icon="back" style="margin-right: 15px" />

    <Path @update="$emit('update', $event)" :value="path" style="flex: 1" />

    <!-- <div class="selected_files" v-if="list.filter((x) => x.isSelected).length">
      {{ list.filter((x) => x.isSelected).length }}
    </div> -->

    <input ref="file_input" type="file" multiple style="position: fixed; top: -100px" />

    <IconButton @click="selectFiles()" icon="upload" style="margin-right: 15px" />
    <IconButton @click="addFolder()" icon="add_folder" style="margin-right: 15px" />
    <IconButton @click="addFile()" icon="add" />
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import { RestApi } from '../util/RestApi';
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
    async addFile() {
      const name = prompt('Enter file name');
      if (name) {
        await RestApi.file.createFolder(this.path + '/' + name);
        this.$emit('refresh');
      }
    },
    async addFolder() {
      const name = prompt('Enter folder name');
      if (name) {
        await RestApi.file.createFile(this.path + '/' + name);
        this.$emit('refresh');
      }
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
    return {};
  },
});
</script>

<style lang="scss" scoped>
.header {
  display: flex;
  // margin-bottom: 15px;
  justify-content: flex-end;
  background: #1d1d1d;
  padding: 10px;
  box-sizing: border-box;

  .selected_files {
    padding: 3px 10px;
    background: #2c589e;
    color: #b5b5b5;
    border-radius: 3px;
    margin-right: 15px;
  }
}
</style>
