<template>
  <ui-block>
    <div :class="$style.header">
      <div>/home/maldan/sasageo/davageo/vrotberageo</div>
      <div :class="$style.sys" class="button_group_round_compact">
        <ui-button icon="pencil" />
        <ui-button icon="pencil" />
      </div>
    </div>
    <div :class="$style.fileList">
      <explorer-file v-for="x in 24" :key="x" :file="{ name: 'sas' }" />
    </div>
    <div>Stat: 20 selected, sort by name</div>
  </ui-block>
</template>

<script lang="ts">
import { defineComponent, PropType } from 'vue';

export default defineComponent({
  props: {
    path: String,
  },
  async mounted() {},
  methods: {},
  data: () => {
    return {};
  },
});
</script>

<style lang="scss" module>
.header {
  display: flex;
  align-items: center;
}

.sys {
  width: max-content;
  margin-left: auto;

  button {
    padding: 5px;
  }
}

.fileList {
  display: grid;
  grid-template-columns: repeat(5, 1fr);
  gap: 40px;
  margin-top: 20px;
  margin-bottom: 20px;
}
</style>
