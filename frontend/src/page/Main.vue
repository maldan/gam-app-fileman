<template>
  <div class="main">
    <!-- Header -->
    <Header :path="path" @update="path = $event" @refresh="refresh()" />

    <!-- Preload -->
    <div v-if="isLoading" class="body">
      <div style="padding: 10px; display: flex; align-items: center">
        Loading...
        <img
          class="rotating"
          src="../asset/preload.svg"
          alt=""
          style="width: 24px; margin-left: 10px"
        />
      </div>
    </div>

    <!-- File list -->
    <div v-if="!isLoading" class="body">
      <Item
        v-for="(x, i) in list"
        :key="x.name"
        :item="x"
        :isOdd="i % 2 === 1"
        :path="path"
        @update="path = $event"
        @refresh="refresh()"
      />
    </div>

    <div
      class="preview"
      v-for="(x, i) in list.filter((x) => x.isSelected && x.name.match(/\.(png|jpeg|gif|jpg)$/))"
      :key="x.name"
      :style="{ left: i * (164 + 10) + 10 + 'px' }"
    >
      {{ x.name }}
      <img :src="$root.API_URL + `/file/file?path=${path}/${x.name}`" style="width: 100%" />
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import { RestApi } from '../util/RestApi';
import IconButton from '../component/IconButton.vue';
import Path from '../component/Path.vue';
import Item from '../component/Item.vue';
import Header from '../component/Header.vue';

export default defineComponent({
  components: { IconButton, Path, Item, Header },
  async mounted() {
    this.refresh();
  },
  watch: {
    path() {
      this.refresh();
    },
  },
  methods: {
    async refresh() {
      this.isLoading = true;
      try {
        this.list = await RestApi.file.getList(this.path);
      } catch {}
      this.isLoading = false;
    },
  },
  data: () => {
    return {
      path: 'G:',
      list: [],
      isLoading: false,
    };
  },
});
</script>

<style lang="scss" scoped>
.main {
  // padding: 10px;
  height: 100%;
  box-sizing: border-box;

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

  .body {
    display: flex;
    height: calc(100% - 55px);
    flex-direction: column;
    overflow-y: auto;
  }

  .preview {
    position: fixed;
    bottom: 10px;
    left: 0;
    width: 164px;
    height: 164px;
    background: rgba(0, 0, 0, 0.5);
    font-size: 12px;
    color: #cecece;
    overflow: hidden;
  }
}
</style>
