<template>
  <div class="main">
    <div class="header">
      <Header />
      <Path />
      <Tabs />
    </div>

    <div class="body" @dragover="dragOver" @dragleave="dragLeave" @drop="drop">
      <div v-if="!$store.state.main.isLoading" class="files">
        <File
          @click.stop="selectFile($event, x)"
          v-for="x in $store.state.file.list"
          :key="x.name"
          :file="x"
        />
      </div>
    </div>

    <Footer />

    <ui-modal v-if="$store.state.extension.name">
      <component :is="'ext-' + $store.state.extension.name" />
    </ui-modal>

    <ui-modal v-if="$store.state.modal.name">
      <component :is="'modal-' + $store.state.modal.name" />
    </ui-modal>
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import Header from '../component/Header.vue';
import Footer from '../component/Footer.vue';
import File from '../component/File.vue';
import Path from '../component/Path.vue';
import Tabs from '../component/Tabs.vue';
import ModalName from '../component/modal/ModalName.vue';
import ModalApprove from '../component/modal/ModalApprove.vue';
import ExtImage from '../component/extension/Image.vue';
import ExtVideo from '../component/extension/Video.vue';
import ExtUsage from '../component/extension/Usage.vue';
import ExtDownload from '../component/extension/Download.vue';
import ExtInfo from '../component/extension/Info.vue';
import ExtPaste from '../component/extension/Paste.vue';
import { RestApi } from '../util/RestApi';

export default defineComponent({
  components: {
    Header,
    Footer,
    File,
    Path,
    Tabs,
    ModalName,
    ModalApprove,
    ExtImage,
    ExtVideo,
    ExtUsage,
    ExtDownload,
    ExtInfo,
    ExtPaste,
  },
  async mounted() {
    this.$store.dispatch('main/getPath');

    document.addEventListener('click', () => {
      this.$store.dispatch('file/clearSelection');
    });

    // Paste image
    document.onpaste = (event) => {
      var items = event.clipboardData?.items || [];
      for (const index in items) {
        var item = items[index];
        if (item.kind === 'file') {
          var blob = item.getAsFile();
          var reader = new FileReader();
          reader.onload = (event) => {
            // @ts-ignore
            this.pasteData.image = event.target?.result || '';
            // this.view = 'paste';
            this.$store.dispatch('extension/show', {
              name: 'paste',
              data: {
                image: this.pasteData.image,
                file: this.pasteData.imageFile,
              },
            });
          };
          this.pasteData.imageFile = blob as File;
          reader.readAsDataURL(blob as any);
        }
      }
    };
  },

  methods: {
    selectFile(e: MouseEvent, x: any) {
      if (e.ctrlKey) {
        this.$store.dispatch('file/select', x);
      } else if (e.shiftKey && this.$store.state.file.lastSelected) {
        const lastIndex = this.$store.state.file.lastSelected.index;
        let direction = x.index - lastIndex;
        if (direction > 0) {
          for (let i = 0; i <= direction; i++) {
            this.$store.dispatch('file/select', this.$store.state.file.list[i + lastIndex]);
          }
        } else {
          for (let i = 0; i <= -direction; i++) {
            this.$store.dispatch('file/select', this.$store.state.file.list[lastIndex - i]);
          }
        }
      } else {
        this.$store.dispatch('file/clearSelection');
        this.$store.dispatch('file/select', x);
      }
    },
    dragOver(e: any) {
      e.stopPropagation();
      e.preventDefault();
      this.isDrag = true;
    },

    dragLeave() {
      this.isDrag = false;
    },

    async drop(e: any) {
      e.stopPropagation();
      e.preventDefault();
      this.isDrag = false;

      const files = e.dataTransfer.files;

      for (let i = 0; i < files.length; i++) {
        const file = files[i];
        const name = ~~(new Date().getTime() / 1000) + '_' + file.name;
        await RestApi.file.uploadFile(this.$store.state.main.path + '/' + name, file);
      }

      this.$store.dispatch('file/getList');
    },
  },
  data: () => {
    return {
      isDrag: false,
      pasteData: {
        image: null as any,
        imageFile: null as any,
      },
    };
  },
});
</script>

<style lang="scss" scoped>
.main {
  height: 100%;
  box-sizing: border-box;

  .header {
    display: flex;
    flex-direction: column;
  }

  .body {
    height: calc(100% - 188px);
    overflow-y: auto;
    // margin-top: 10px;
    position: relative;

    .files {
      padding-top: 10px;
      display: grid;
      grid-template-columns: 1fr 1fr 1fr 1fr 1fr 1fr 1fr;
      grid-auto-rows: max-content;
    }
  }

  .drag_file {
    position: absolute;
    left: 0;
    top: 0;
    width: 100%;
    height: 100%;
    background: rgba(0, 0, 0, 0.4);
    display: flex;
    align-items: center;
    justify-content: center;
    z-index: 1;
  }
}
</style>
