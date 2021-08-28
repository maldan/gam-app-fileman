<template>
  <div class="view_download">
    <IconButton class="close" icon="close" @click="$emit('close'), $emit('refresh')" />

    <div class="window">
      <div class="top">
        <Input placeholder="Url" v-model="url" />
        <Button text="Add" @click="add()" />
      </div>
      <div class="list">
        <div class="item" v-for="x in list" :key="x">
          <div>{{ x.url }}</div>
          <div>{{ x.path }}{{ x.name }}</div>
          <div>{{ ~~(x.progress * 100) }}%</div>
          <div>{{ $root.moment(x.created).fromNow() }}</div>
        </div>
      </div>
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
    file: String,
  },
  components: { IconButton, Input, Button },
  async mounted() {
    await this.refresh();

    this.id = setInterval(() => {
      this.refresh();
    }, 1000);
  },
  beforeUnmount() {
    clearInterval(this.id);
  },
  methods: {
    async add() {
      await RestApi.download.add(this.url, (this.path + '/').replace(/\/\//g, '/'));
      this.url = '';
      await this.refresh();
    },
    async refresh() {
      this.list = await RestApi.download.getList();
    },
  },
  data: () => {
    return {
      url: '',
      list: [],
      id: 0,
    };
  },
});
</script>

<style lang="scss" scoped>
.view_download {
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
