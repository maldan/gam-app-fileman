<template>
  <div class="main">
    <div>Header</div>
    <div>Files</div>
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import { RestApi } from '../util/RestApi';

export default defineComponent({
  components: {},
  async mounted() {},
  watch: {},
  methods: {},
  data: () => {
    return {
      files: [{}],
    };
  },
});
</script>

<style lang="scss" scoped>
.main {
  height: 100%;
  box-sizing: border-box;
}
</style>
