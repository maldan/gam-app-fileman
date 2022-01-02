<template>
  <div class="main">
    <div class="header">
      <ui-button text="Back" @click="$router.push('/')" :style="{ flex: 0 }" />
      <ui-input
        placeholder="Search..."
        @keyup.enter="$store.dispatch('search/search')"
        v-model="$store.state.search.query"
      />
      <ui-button text="Search" @click="$store.dispatch('search/search')" :style="{ flex: 0 }" />
    </div>
    <div class="body">
      <div v-if="!$store.state.main.isLoading" class="files">
        <File
          @click.stop="selectFile($event, x)"
          @contextmenu="
            selectFile($event, x);
            showContextMenu($event);
          "
          v-for="x in $store.state.file.list"
          :key="x.name"
          :file="x"
        />
      </div>
    </div>

    <ui-modal v-if="$store.state.extension.name">
      <component :is="'ext-' + $store.state.extension.name" />
    </ui-modal>

    <ui-modal v-if="$store.state.modal.name">
      <component :is="'modal-' + $store.state.modal.name" />
    </ui-modal>

    <ui-context-menu
      v-if="isShowContextMenu"
      @close="isShowContextMenu = false"
      :items="[
        {
          label: 'Set Tags',
          icon: 'pencil',
          async onClick() {
            const tags = [];
            for (x in $store.state.file.lastSelected.tags)
              tags.push({ key: x, value: $store.state.file.lastSelected.tags[x] });

            await $store.dispatch('modal/show', {
              name: 'tags',
              data: {
                tags,
              },
              onSuccess: () => {
                $store.dispatch('file/setTags');
              },
            });
          },
        },
      ]"
      :style="contextPosition"
    />
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
    document.addEventListener('click', () => {
      this.$store.dispatch('file/clearSelection');
      // this.isShowContextMenu = false;
    });

    await this.$store.dispatch('file/clearList');
  },

  methods: {
    showContextMenu(e: MouseEvent) {
      e.preventDefault();
      this.isShowContextMenu = true;
      this.contextPosition.left = e.pageX + 20 + 'px';
      this.contextPosition.top = e.pageY + 20 + 'px';
    },
    selectFile(e: MouseEvent, x: any) {
      this.isShowContextMenu = false;

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
  },
  data: () => {
    return {
      isDrag: false,
      isShowContextMenu: false,
      pasteData: {
        image: null as any,
        imageFile: null as any,
      },
      contextPosition: {
        left: '0',
        top: '0',
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
  }

  .body {
    height: calc(100% - 175px);
    overflow-y: auto;
    position: relative;

    .files {
      padding-top: 10px;
      display: grid;
      gap: 10px;
      grid-template-columns: repeat(1, 1fr);
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

@for $i from 0 through 12 {
  @media (min-width: #{575 + $i * 256}px) {
    .main {
      .body {
        .files {
          grid-template-columns: repeat(#{$i + 2}, 1fr);
        }
      }
    }
  }
}
</style>
