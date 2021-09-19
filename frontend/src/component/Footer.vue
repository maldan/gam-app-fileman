<template>
  <div :class="$style.footer">
    <!-- Preload -->
    <img
      v-if="$store.state.isLoading"
      class="rotating"
      src="../asset/preload.svg"
      style="width: 16px; margin-right: 20px"
    />

    <button v-for="x in extensons" :key="x.name" class="clickable">
      {{ x.name }}
    </button>
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';

export default defineComponent({
  props: {},
  components: {},
  async mounted() {},
  methods: {},
  data: () => {
    return {
      extensons: [{ name: 'Download Manager' }, { name: 'Disk Usage' }],
    };
  },
});
</script>

<style lang="scss" module>
.footer {
  display: flex;
  flex: 1;
  background: #1b1b1b;
  padding: 10px;

  button {
    background: #2d88ff;
    outline: none;
    border: none;
    padding: 5px 10px;
    color: #fefefe;
    margin-right: 10px;
    border-radius: 2px;
  }
}
</style>
