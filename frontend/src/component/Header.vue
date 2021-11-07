<template>
  <div :class="$style.header">
    <input ref="file_input" type="file" multiple style="position: fixed; left: -1000px" />
    <!-- <button
      @click="x.onClick()"
      v-for="x in buttons"
      :key="x.name"
      class="clickable"
      :disabled="isDisabled(x)"
    >
      <ui-icon :size="17" :name="x.icon" />
      {{ x.name }}
    </button> -->
    <ui-button
      @click="x.onClick()"
      v-for="x in buttons"
      :class="$style.button"
      :key="x.name"
      :icon="x.icon"
      :text="x.name"
      iconPosition="left"
      :iconColor="x.color"
      :disabled="isDisabled(x)"
    />
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
        icon: 'folder',
        name: 'New Dir',
        onClick: () => {
          this.$store.dispatch('modal/show', {
            name: 'name',
            data: { title: 'Add new directory', name: '' },
            func: () => {
              this.$store.dispatch('file/createDir');
            },
          });
        },
      },
      {
        icon: 'file',
        name: 'New File',
        color: '#b9b9b9',
        onClick: () => {
          this.$store.dispatch('modal/show', {
            name: 'name',
            data: { title: 'Add new file', name: '' },
            func: () => {
              this.$store.dispatch('file/createFile');
            },
          });
        },
      },
      {
        icon: 'rename',
        name: 'Rename',
        color: '#45a1ff',
        isNeedFile: true,
        onClick: () => {
          this.$store.dispatch('file/copySelectedToBuffer');

          this.$store.dispatch('modal/show', {
            name: 'name',
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
        icon: 'trash',
        name: 'Delete',
        color: '#cf5144',
        isNeedFile: true,
        onClick: () => {
          this.$store.dispatch('file/copySelectedToBuffer');
          this.$store.dispatch('modal/show', {
            name: 'approve',
            data: {
              title: `Are you sure you want to delete these ${this.$store.state.file.buffer.length} files?`,
            },
            func: () => {
              this.$store.dispatch('file/delete');
            },
          });
        },
      },
      {
        icon: 'copy',
        name: 'Copy',
        color: '#fed300',
        isNeedFile: true,
        onClick: () => {
          this.$store.dispatch('file/copySelectedToBuffer', 'copy');
        },
      },
      {
        icon: 'scissors',
        name: 'Cut',
        color: '#50b11f',
        isNeedFile: true,
        onClick: () => {
          this.$store.dispatch('file/copySelectedToBuffer', 'cut');
        },
      },
      {
        icon: 'paste',
        name: 'Paste',
        color: '#fed300',
        isNeedBuffer: true,
        onClick: () => {
          this.$store.dispatch('file/paste');
        },
      },
      {
        icon: 'upload',
        name: 'Upload',
        onClick: () => {
          const fileInput = this.$refs['file_input'] as HTMLInputElement;

          if (fileInput) {
            fileInput.onchange = (e: any) => {
              this.$store.dispatch('file/upload', e.target.files);
            };
            fileInput.click();
          }
        },
      },
      {
        icon: 'info',
        name: 'Info',
        color: '#579aff',
        isNeedFile: true,
        onClick: () => {
          this.$store.dispatch('file/copySelectedToBuffer');
          this.$store.dispatch('extension/show', {
            name: 'info',
            data: {},
          });
        },
      },
    ];
  },
  methods: {
    isDisabled(x: any) {
      return (
        (x.isNeedFile && !this.$store.state.file.lastSelected) ||
        (x.isNeedBuffer && !this.$store.state.file.buffer.length)
      );
    },
  },
  data: () => {
    return {
      buttons: [] as any[],
    };
  },
});
</script>

<style lang="scss" module>
@import '../gam_sdk_ui/vue/style/color.scss';
@import '../gam_sdk_ui/vue/style/size.scss';

.header {
  display: flex;
  flex: 1;

  .button {
    text-transform: uppercase;
    font-size: 14px;
    border-radius: 0;
    background: $gray-dark;
  }

  /*button {
    background: none;
    color: #dadada;
    display: flex;
    flex-direction: column;
    outline: none;
    border: 0;
    align-items: center;
    justify-content: center;
    padding: 7px;
    background: #1b1b1b;
    flex: 1;
    text-transform: uppercase;
    font-size: 12px;
  }*/
}
</style>
