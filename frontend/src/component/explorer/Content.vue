<template>
  <ui-block :class="$style.content">
    <div :class="$style.header">
      <div :class="$style.path">
        <ui-button
          v-for="(x, i) in pathParts"
          :text="x"
          :key="i"
          :isSelected="i === pathParts.length - 1"
        />
      </div>
      <div :class="$style.sys" class="button_group_round_compact">
        <ui-button icon="control_panel" :isSelected="true" />
        <ui-button icon="list" />
      </div>
    </div>
    <div :class="$style.fileList">
      <explorer-file v-for="x in 24" :key="x" :file="{ name: 'sas' }" />
    </div>
    <div :class="$style.footer" class="button_group_round_compact">
      <ui-button text="By Name" />
      <ui-button text="0 / 20" />
    </div>
  </ui-block>
</template>

<script lang="ts">
import { defineComponent, PropType } from 'vue';

export default defineComponent({
  props: {
    path: String,
  },
  computed: {
    pathParts() {
      const x = '/home/maldan/sasageo/davageo/vrotberageo'.split('/');
      x[0] = '/';
      return x;
    },
  },
  async mounted() {},
  methods: {},
  data: () => {
    return {};
  },
});
</script>

<style lang="scss" module>
.content {
  button {
    padding: 5px;
    font-size: 14px;
    font-weight: normal;
  }

  .footer {
    display: flex;
  }
}

.header {
  display: flex;
  align-items: center;

  .path {
    display: flex;

    button {
      margin-right: 5px;
    }
  }
}

.sys {
  width: max-content;
  margin-left: auto;
}

.fileList {
  display: grid;
  grid-template-columns: repeat(5, 1fr);
  gap: 40px;
  margin-top: 20px;
  margin-bottom: 20px;
}
</style>
