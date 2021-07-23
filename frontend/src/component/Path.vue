<template>
  <div class="path">
    <input v-if="mode === 'string'" type="text" :value="value" />
    <div v-if="mode !== 'string'" style="display: flex">
      <div
        @click="go(i)"
        class="part clickable"
        v-for="(x, i) in value.split('/').filter((x) => Boolean(x))"
        :key="x"
      >
        {{ x }}
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';

export default defineComponent({
  props: {
    icon: String,
    value: String,
  },
  components: {},
  async mounted() {},
  methods: {
    go(i: number) {
      const tt = this.value?.split('/').filter((x) => Boolean(x)) || [];
      this.$emit('update', tt.slice(0, i + 1).join('/'));
    },
  },
  data: () => {
    return {
      mode: 's',
    };
  },
});
</script>

<style lang="scss" scoped>
.path {
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
