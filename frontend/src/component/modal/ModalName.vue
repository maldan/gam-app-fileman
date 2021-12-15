<template>
  <div :class="$style.modal">
    <div style="margin-bottom: 10px">{{ $store.state.modal.data.title }}</div>
    <ui-input
      ref="input"
      @keydown.enter="$store.dispatch('modal/ok')"
      placeholder="Name..."
      v-model="$store.state.modal.data.name"
    />
    <div style="display: flex; margin-top: 10px">
      <ui-button @click="$store.dispatch('modal/close')" text="Cancel" />
      <ui-button @click="$store.dispatch('modal/ok')" text="OK" />
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';

export default defineComponent({
  props: {},
  components: {},
  async mounted() {
    (this.$refs['input'] as any).focus();
  },
  methods: {},
  data: () => {
    return {};
  },
});
</script>

<style lang="scss" module>
.modal {
  min-width: 360px; //
}
</style>
