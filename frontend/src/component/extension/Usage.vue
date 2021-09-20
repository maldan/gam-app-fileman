<template>
  <div class="view_usage">
    <img
      class="close clickable"
      src="../../asset/icon/close.svg"
      @click="$store.dispatch('extension/close')"
    />

    <div class="window">
      <div class="item" v-for="x in list" :key="x.mount">
        <div style="display: flex">
          <div style="flex: 1">{{ x.fs }}</div>
          <div style="flex: 1">{{ x.mount }}</div>
          <div style="display: flex; flex: 1; justify-content: flex-end">
            <div style="color: #fedf00">{{ $root.pretty(x.used * 1024) }}</div>
            <div style="width: 20px; text-align: center">/</div>
            <div style="color: #fea400">{{ $root.pretty(x.total * 1024) }}</div>
          </div>
        </div>
        <div class="progress" style="flex: 1">
          <div class="fill" :style="{ width: (x.used / x.total) * 100 + '%' }"></div>
        </div>
      </div>

      <div class="item">
        <div style="display: flex">
          <div style="flex: 1">Total</div>
          <div style="flex: 1"></div>
          <div style="display: flex; flex: 1; justify-content: flex-end">
            <div style="color: #fedf00">{{ $root.pretty(usedTotal * 1024) }}</div>
            <div style="width: 20px; text-align: center">/</div>
            <div style="color: #fea400">{{ $root.pretty(total * 1024) }}</div>
          </div>
        </div>
        <div class="progress" style="flex: 1">
          <div class="fill" :style="{ width: (usedTotal / total) * 100 + '%' }"></div>
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import { RestApi } from '../../util/RestApi';

export default defineComponent({
  props: {},
  components: {},
  async mounted() {
    this.list = await RestApi.disk.getUsage();
    for (let i = 0; i < this.list.length; i++) {
      this.total += this.list[i].total;
      this.usedTotal += this.list[i].used;
    }
  },
  methods: {},
  data: () => {
    return {
      list: [] as any[],
      total: 0,
      usedTotal: 0,
    };
  },
});
</script>

<style lang="scss" scoped>
.view_usage {
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
    height: 400px;
    overflow-y: auto;

    .item {
      color: #bebebe;
      margin-bottom: 10px;
      background: #1d1d1d;
      padding: 10px;

      &:last-child {
        margin-bottom: 0;
      }

      .progress {
        border-radius: 3px;
        margin-right: 10px;
        background: #0f0f0f;
        overflow: hidden;
        margin-top: 15px;

        .fill {
          width: 0%;
          height: 100%;
          background: #4360a0;
          transition: width 0.3s ease;
          color: #cecece;
          font-size: 14px;
          display: flex;
          align-items: center;
          height: 4px;
          padding-left: 10px;
        }
      }
    }
  }
}
</style>
