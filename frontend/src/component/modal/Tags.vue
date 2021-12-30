<template>
  <div :class="$style.modal">
    <div style="margin-bottom: 10px">File Tags</div>
    <div :class="$style.field">
      <ui-input placeholder="Name" />
      <ui-input placeholder="Value" />
      <ui-select placeholder="Type" />
    </div>
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
  async mounted() {},
  methods: {},
  data: () => {
    return {};
  },
});
</script>

<style lang="scss" module>
.modal {
  min-width: 360px;

  .field {
    display: flex;
  }
}
</style>
