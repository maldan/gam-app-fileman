<template>
  <router-view />
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import Moment from 'moment';
import PrettyBytes from 'pretty-bytes';

export default defineComponent({
  components: {},
  async mounted() {},
  methods: {},
  data: () => {
    return {
      moment: Moment,
      pretty: PrettyBytes,
      // @ts-ignore
      API_URL: process.env.VUE_APP_API_URL || `${window.location.origin}/api`,
    };
  },
});
</script>

<style lang="scss" scoped>
.main {
}
</style>
