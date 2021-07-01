<template>
  <div class="main">
    <div class="header">
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
      <img class="clickable" @click="back()" src="../asset/settings.svg" alt="" draggable="false" />
    </div>
    <div class="body">
      <FileList v-model:path="path_1" :mode="isTileMode ? 'tile' : 'list'" />
      <div v-if="isSplit" style="width: 10px"></div>
      <FileList v-if="isSplit" v-model:path="path_2" :mode="isTileMode ? 'tile' : 'list'" />
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
      path_1: 'D:/',
      path_2: 'D:/',
      isSplit: false,
      isTileMode: true,
    };
  },
});
</script>

<style lang="scss" scoped>
.main {
  padding: 20px;
  height: 100%;
  box-sizing: border-box;

  .header {
    display: flex;
    margin-bottom: 10px;
    justify-content: flex-end;

    img {
      margin-left: 15px;
    }
  }

  .body {
    display: flex;
    height: calc(100% - 35px);
  }
}
</style>
