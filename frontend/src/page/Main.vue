<template>
  <div class="main">
    <!-- Header -->
    <Header :path="path" @update="path = fixPath($event)" @refresh="refresh()" />

    <!-- Empty body -->
    <div v-show="isLoading" class="body"></div>

    <!-- File list -->
    <div v-show="!isLoading" :class="isList ? 'body' : 'body_grid'">
      <Item
        v-for="(x, i) in list"
        :key="x.name"
        :item="x"
        :isOdd="i % 2 === 1"
        :path="path"
        :isList="isList"
        @update="path = fixPath($event)"
        @refresh="refresh()"
        @open="openFile"
      />
    </div>

    <!-- Preview -->
    <!--<div
      class="preview"
      v-for="(x, i) in list.filter(
        (x) => x.isSelected && x.name.match(/\.(png|jpeg|gif|jpg|mp4)$/),
      )"
      :key="x.name"
      :style="{ left: i * (164 + 10) + 10 + 'px' }"
    >
      {{ x.name }}
      <img
        v-if="x.name.match(/\.(png|jpeg|gif|jpg)$/)"
        :src="$root.API_URL + `/file/file?path=${path}/${x.name}`"
        style="width: 100%"
      />
      <img
        v-if="x.name.match(/\.(mp4)$/)"
        :src="$root.API_URL + `/file/videoThumbnail?path=${path}/${x.name}`"
        style="width: 100%"
      />
    </div> -->

    <!-- Bottom -->
    <Bottom
      :isLoading="isLoading"
      :path="path"
      :list="list"
      :sort="sortBy"
      :buffer="buffer"
      :tab="tab"
      :tabs="tabs"
      :isList="isList"
      @changeView="isList = $event"
      @changeSort="sortBy = $event"
      @update="path = fixPath($event)"
      @refresh="refresh()"
      @paste="paste()"
      @du="du()"
      @copy="(bufferPath = path), (copyMode = 'copy'), (buffer = list.filter((x) => x.isSelected))"
      @cut="(bufferPath = path), (copyMode = 'cut'), (buffer = list.filter((x) => x.isSelected))"
      @tab="selectTab($event)"
      @open="view = $event"
    />

    <!-- Views -->
    <component
      v-if="view"
      :is="`view-${view}`"
      @refresh="refresh()"
      @close="view = ''"
      :list="list"
      :path="path"
      :file="selectedFile"
      :data="[pasteData.image, pasteData.imageFile]"
    />
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import { RestApi } from '../util/RestApi';
import { Helper } from '../util/Helper';
import IconButton from '../component/IconButton.vue';
import Path from '../component/Path.vue';
import Item from '../component/Item.vue';
import Header from '../component/Header.vue';
import Bottom from '../component/Bottom.vue';
import ViewImage from '../component/view/Image.vue';
import ViewVideo from '../component/view/Video.vue';
import ViewDownload from '../component/view/Download.vue';
import ViewPaste from '../component/view/Paste.vue';
import ViewUsage from '../component/view/Usage.vue';

export default defineComponent({
  components: {
    IconButton,
    Path,
    Item,
    Header,
    Bottom,
    ViewImage,
    ViewVideo,
    ViewDownload,
    ViewPaste,
    ViewUsage,
  },
  async mounted() {
    this.selectTab(0);
    this.path = await RestApi.file.getPath();
    this.refresh();

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
            this.view = 'paste';
          };
          this.pasteData.imageFile = blob as File;
          reader.readAsDataURL(blob as any);
        }
      }
    };
  },
  watch: {
    async path() {
      this.tab.path = this.path;
      this.refresh();
      await RestApi.file.setPath(this.path);
    },
  },
  methods: {
    fixPath: Helper.fixPath,
    async refresh() {
      this.isLoading = true;
      try {
        const list = await RestApi.file.getList(this.path);
        this.list = this.sort(list);
      } catch {}
      this.isLoading = false;
    },
    selectTab(id: number) {
      this.tab = this.tabs[id];
      this.path = this.tab.path;
      this.refresh();
    },
    du() {
      for (let i = 0; i < this.list.length; i++) {
        if (this.list[i].kind === 'dir') {
          this.list[i].size = 0;
          this.getDirSize(this.list[i]);
        }
      }
    },
    unselectAll() {
      for (let i = 0; i < this.list.length; i++) {
        this.list[i].isSelected = false;
      }
    },
    async getDirSize(dir: any) {
      dir.size = await RestApi.file.getDirSize(this.fixPath(this.path + '/' + dir.name));
      this.list = this.sort(this.list);
    },
    sort(list: any[]): any[] {
      let folders = list.filter((x) => x.kind === 'dir');
      let files = list.filter((x) => x.kind !== 'dir');

      if (this.sortBy === 'name') {
        folders = folders.sort((a, b) => a.name.localeCompare(b.name));
        files = files.sort((a, b) => a.name.localeCompare(b.name));
      }
      if (this.sortBy === 'size') {
        folders = folders.sort((a, b) => a.size - b.size);
        files = files.sort((a, b) => a.size - b.size);
      }
      if (this.sortBy === 'created') {
        folders = folders.sort(
          (a, b) => new Date(a.created).getTime() - new Date(b.created).getTime(),
        );
        files = files.sort((a, b) => new Date(a.created).getTime() - new Date(b.created).getTime());
      }

      return [...folders, ...files];
    },
    openFile(path: string) {
      this.selectedFile = path;
      if (path.match(/\.(png|jpeg|gif|jpg)$/)) {
        this.view = 'image';
      }
      if (path.match(/\.(mp4)$/)) {
        this.view = 'video';
      }

      this.unselectAll();
    },
    async paste() {
      this.isLoading = true;
      for (let i = 0; i < this.buffer.length; i++) {
        if (this.copyMode === 'copy') {
          await RestApi.file.copy(
            this.bufferPath + '/' + this.buffer[i].name,
            this.path + '/' + this.buffer[i].name,
          );
        } else {
          await RestApi.file.move(
            this.bufferPath + '/' + this.buffer[i].name,
            this.path + '/' + this.buffer[i].name,
          );
        }
      }
      this.isLoading = false;
      this.buffer = [];
      this.refresh();
    },
  },
  data: () => {
    return {
      sortBy: 'name',
      path: '/',
      list: [] as any[],
      isLoading: false,
      view: '',
      selectedFile: '',
      copyMode: 'copy',
      bufferPath: '',
      buffer: [] as any[],
      tab: null as any,
      tabs: [{ path: '/' }, { path: '/' }, { path: '/' }],
      isList: true,

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

  .body {
    display: flex;
    height: calc(100% - 45px - 40px);
    flex-direction: column;
    overflow-y: auto;
    position: relative;
  }

  .body_grid {
    display: grid;
    grid-template-columns: repeat(14, 1fr);
    height: calc(100% - 45px - 40px);
    overflow-y: auto;
    position: relative;
    grid-auto-rows: max-content;
  }

  .preview {
    position: fixed;
    bottom: 50px;
    left: 0;
    width: 164px;
    height: 164px;
    background: rgba(0, 0, 0, 0.5);
    font-size: 12px;
    color: #cecece;
    overflow: hidden;
  }
}

@for $i from 0 through 11 {
  @media (max-width: (2550px-$i*140px)) {
    .main {
      .body_grid {
        grid-template-columns: repeat(13-$i, 1fr);
      }
    }
  }
}

/*@media (max-width: 1990px) {
  .main {
    .body_grid {
      grid-template-columns: repeat(10, 1fr);
    }
  }
}*/
</style>
