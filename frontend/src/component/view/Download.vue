<template>
  <div class="view_download">
    <IconButton class="close" icon="close" @click="$emit('close'), $emit('refresh')" />

    <div class="window">
      <div class="top">
        <TextArea placeholder="Url" v-model="url" />
        <Button text="Add" @click="add()" />
      </div>
      <div class="list_container">
        <div v-if="list.length > 0" class="list">
          <div class="item" v-for="x in list" :key="x.name">
            <div>{{ x.url }}</div>
            <div>{{ x.path }}{{ x.name }}</div>
            <div style="display: flex">
              <div class="progress" style="flex: 1">
                <div class="fill" :style="{ width: x.progress * 100 + '%' }">
                  {{ ~~(x.progress * 100) }}%
                </div>
              </div>
              <div>{{ $root.moment(x.created).format('HH:mm:ss') }}</div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import IconButton from '../IconButton.vue';
import Input from '../Input.vue';
import TextArea from '../TextArea.vue';
import Button from '../Button.vue';
import { RestApi } from '../../util/RestApi';

export default defineComponent({
  props: {
    path: String,
    file: String,
  },
  components: { IconButton, Input, Button, TextArea },
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
      const list = this.url.split(' ');
      for (let i = 0; i < list.length; i++) {
        await RestApi.download.add(list[i], (this.path + '/').replace(/\/\//g, '/'));
      }

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
    width: 600px;

    .top {
      display: flex;
      flex-direction: row;

      button {
        margin-left: 10px;
        flex: none;
      }
    }

    .list_container {
      max-height: 400px;
      overflow-y: auto;

      .list {
        display: flex;
        flex-direction: column;
        margin-top: 10px;

        .item {
          display: flex;
          background: #1b1b1b;
          color: #999999;
          margin-bottom: 10px;
          flex-direction: column;

          > div {
            flex: 1;
            white-space: nowrap;
            overflow: hidden;
            text-overflow: ellipsis;
            padding: 10px;
          }

          .progress {
            border-radius: 3px;
            margin-right: 10px;
            background: #202020;
            overflow: hidden;

            .fill {
              width: 0%;
              height: 100%;
              background: #4360a0;
              transition: width 0.3s ease;
              color: #cecece;
              font-size: 14px;
              display: flex;
              align-items: center;
              padding-left: 5px;
            }
          }

          &:last-child {
            margin-bottom: 0;
          }
        }
      }
    }
  }
}
</style>
