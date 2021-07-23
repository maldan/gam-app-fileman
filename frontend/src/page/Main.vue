<template>
  <div class="main">
    <div class="header">
      <img
        class="clickable"
        @click="back()"
        src="../asset/back.svg"
        alt=""
        style="margin-right: 10px"
        draggable="false"
      />
      <input type="text" @change="$emit('update:path', $event.target.value)" :value="path" />

      <img class="clickable" src="../asset/mode.svg" alt="" style="" draggable="false" />

      <img
        class="clickable"
        @click="isTileMode = !isTileMode"
        src="../asset/mode.svg"
        alt=""
        style=""
        draggable="false"
      />
      <img
        class="clickable"
        @click="isSplit = !isSplit"
        src="../asset/split.svg"
        alt=""
        style=""
        draggable="false"
      />
      <img class="clickable" src="../asset/settings.svg" alt="" draggable="false" />
    </div>
    <div class="body">
      <!-- <FileList v-model:path="path_1" :mode="isTileMode ? 'tile' : 'list'" />
      <div class="gap" v-if="isSplit">
        <img class="clickable" src="../asset/back.svg" alt="" draggable="false" />
      </div>
      <FileList v-if="isSplit" v-model:path="path_2" :mode="isTileMode ? 'tile' : 'list'" /> -->

      <!-- Tile mode -->
      <div
        v-if="mode === 'tile' && !isLoading"
        ref="file_list"
        :class="[$style.file_tile, $style[`g${maxRows}`]]"
      >
        <div :class="$style.file" v-for="x in files" :key="x.name">
          <img
            @click="go(x)"
            class="clickable"
            :src="iconByFile(x)"
            alt="Folder"
            draggable="false"
          />
          <div :class="$style.name">{{ x.name }}</div>
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import { RestApi } from '../util/RestApi';
import FileList from '../component/FileList.vue';

export default defineComponent({
  components: { FileList },
  async mounted() {
    this.refresh();
  },
  watch: {
    path() {
      this.refresh();
    },
  },
  methods: {
    async refresh() {},
  },
  data: () => {
    return {
      path: 'D:/',
      isSplit: false,
      isTileMode: true,
    };
  },
});
</script>

<style lang="scss" scoped>
.main {
  padding: 10px;
  height: 100%;
  box-sizing: border-box;

  .header {
    display: flex;
    margin-bottom: 15px;
    justify-content: flex-end;

    img {
      margin-left: 15px;
    }

    input {
      flex: 1;
      padding: 4px 8px;
      background: #23232373;
      border: 0;
      border-radius: 2px;
      outline: none;
      color: #a5a5a5;
      font-size: 16px;
    }
  }

  .body {
    display: flex;
    height: calc(100% - 35px);

    .gap {
      width: 26px;
      display: flex;
      align-items: center;
      justify-content: center;
    }
  }
}
</style>
