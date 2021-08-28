<template>
  <div class="view_paste">
    <IconButton class="close" icon="close" @click="$emit('close'), $emit('refresh')" />

    <div class="window">
      <img :src="data[0]" alt="" style="max-height: 640px" />
      <Button @click="upload" text="Add" />
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import IconButton from '../IconButton.vue';
import Input from '../Input.vue';
import Button from '../Button.vue';
import { RestApi } from '../../util/RestApi';

export default defineComponent({
  props: {
    path: String,
    data: Array,
  },
  components: { IconButton, Input, Button },
  async mounted() {},
  methods: {
    async upload() {
      const file = this.data?.[1] as File;

      await RestApi.file.uploadFile(
        this.path + '/' + Math.random() + '_' + file.name,
        file,
        (e) => {
          // this.progress[i] = e.loaded / items[i].size;
          this.$emit('close');
        },
      );
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
