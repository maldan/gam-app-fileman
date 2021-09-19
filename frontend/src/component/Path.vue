<template>
  <div class="path">
    <img
      @click="isEditMode = !isEditMode"
      style="margin-right: 15px"
      src="../asset/icon/pencil.svg"
      alt=""
      draggable="false"
    />
    <input
      v-if="isEditMode"
      type="text"
      :value="$store.state.main.path"
      @change="$store.dispatch('changePath', $event.target.value), (isEditMode = false)"
    />
    <div v-if="!isEditMode" style="display: flex">
      <div @click="$store.dispatch('main/changePath', '/')" class="part clickable">/</div>
      <div
        @click="go(i)"
        class="part clickable"
        v-for="(x, i) in $store.state.main.path.split('/').filter((x) => Boolean(x))"
        :key="x"
      >
        {{ x }}
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import { Helper } from '../util/Helper';

export default defineComponent({
  props: {},
  components: {},
  async mounted() {},
  methods: {
    go(i: number) {
      const tt = this.$store.state.main.path.split('/').filter((x: any) => Boolean(x)) || [];
      this.$store.dispatch('main/changePath', '/' + tt.slice(0, i + 1).join('/'));
    },
  },
  data: () => {
    return {
      isEditMode: false,
    };
  },
});
</script>

<style lang="scss" scoped>
.path {
  display: flex;
  align-items: center;
  padding: 10px;
  box-sizing: border-box;
  background: #363636;

  input {
    flex: 1;
    padding: 4px 8px;
    background: #3b3b3b;
    border: 0;
    border-radius: 2px;
    outline: none;
    color: #a5a5a5;
    font-size: 16px;
  }

  .part {
    padding: 3px 10px;
    background: #565656;
    border-radius: 4px;
    margin-right: 10px;
    color: #ababab;
  }
}
</style>
