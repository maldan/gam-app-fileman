<template>
  <img class="clickable" :src="require(`../asset/icon/${icon}.svg`)" alt="" draggable="false" />
</template>

<script lang="ts">
import { defineComponent } from 'vue';

export default defineComponent({
  props: {
    icon: String,
  },
  components: {},
  async mounted() {},
  methods: {},
  data: () => {
    return {};
  },
});
</script>

<style lang="scss" scoped>
.main {
}
</style>
