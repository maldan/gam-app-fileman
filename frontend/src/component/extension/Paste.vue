<template>
  <div class="view_paste">
    <img
      class="close clickable"
      src="../../asset/icon/close.svg"
      @click="$store.dispatch('extension/close')"
    />

    <div class="window">
      <img :src="$store.state.extension.data.image" alt="" style="max-height: 640px" />
      <ui-button @click="upload" text="Add" />
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import { RestApi } from '../../util/RestApi';

export default defineComponent({
  props: {},
  components: {},
  async mounted() {},
  methods: {
    async upload() {
      await RestApi.file.uploadFile(
        this.$store.state.main.path +
          '/' +
          Math.random() +
          '_' +
          this.$store.state.extension.data.file.name,
        this.$store.state.extension.data.file,
      );

      this.$store.dispatch('extension/close');
      this.$store.dispatch('file/getList');
    },
  },
  data: () => {
    return {};
  },
});
</script>

<style lang="scss" scoped>
.view_paste {
  position: fixed;
  width: 100%;
  height: 100%;
  left: 0;
  top: 0;
  background: rgba(0, 0, 0, 0.5);
  overflow-y: auto;
  display: flex;
  align-items: center;
  justify-content: center;

  .close {
    position: fixed;
    right: 20px;
    top: 10px;
  }

  .window {
    background: #fefefe;
    padding: 10px;
    background: #2b2b2b;
    border-radius: 3px;
    display: flex;
    flex-direction: column;

    .top {
      display: flex;
      flex-direction: row;

      button {
        margin-left: 10px;
      }
    }

    .list {
      display: flex;
      flex-direction: column;
      margin-top: 10px;

      .item {
        display: flex;
        background: #1b1b1b;
        color: #999999;
        margin-bottom: 10px;

        > div {
          flex: 1;
          white-space: nowrap;
          overflow: hidden;
          text-overflow: ellipsis;
          padding: 10px;
        }

        &:last-child {
          margin-bottom: 0;
        }
      }
    }
  }
}
</style>
