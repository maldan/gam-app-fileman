<template>
  <div :class="$style.view">
    <img
      class="close clickable"
      src="../../asset/icon/close.svg"
      @click="$store.dispatch('extension/close')"
    />

    <div :class="$style.item" v-for="x in info" :key="x">
      <div>Name: {{ x.name }}</div>
      <div>Kind: {{ x.kind }}</div>
      <div>User: {{ x.user }}</div>
      <div>Size: {{ $root.pretty(x.size) }}</div>
      <div>Created: {{ $root.moment(x.created).format('DD MMM YYYY HH:mm:ss') }}</div>
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
    this.info = [];
    for (let i = 0; i < this.$store.state.file.buffer.length; i++) {
      this.info.push(await RestApi.file.getFullInfo(this.$store.state.file.buffer[i]));
    }
  },
  methods: {},
  data: () => {
    return {
      info: [] as any[],
    };
  },
});
</script>

<style lang="scss" module>
.view {
  .item {
    display: flex;
    flex-direction: column;
    padding: 10px;
    background: #3a3a3a;
    margin-bottom: 10px;

    &:last-child {
      margin-bottom: 0;
    }
  }
}
</style>
