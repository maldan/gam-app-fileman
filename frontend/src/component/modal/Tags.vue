<template>
  <div :class="$style.modal">
    <div style="margin-bottom: 10px">File Tags</div>
    <div :class="$style.group">
      <div :class="$style.field" v-for="(x, i) in $store.state.modal.data.tags" :key="i">
        <ui-input placeholder="Name" v-model="x.key" />
        <ui-input placeholder="Value" v-model="x.value" />
        <ui-button
          @click="$store.state.modal.data.tags.splice(i, 1)"
          icon="trash"
          style="padding: 5px"
        />
      </div>
      <ui-button @click="$store.state.modal.data.tags.push({ key: '', value: '' })" text="Add" />
    </div>
    <div style="display: flex; margin-top: 10px">
      <ui-button @click.stop="$store.dispatch('modal/close')" text="Cancel" />
      <ui-button @click.stop="$store.dispatch('modal/ok')" text="OK" />
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

  .group {
    display: flex;
    flex-direction: column;

    .field {
      display: flex;
      margin-bottom: 10px;
    }
  }
}
</style>
