<template>
  <div :class="$style.main" @contextmenu="showContextMenu($event)">
    <div :class="$style.content">
      <tree-list />
      <explorer-content style="flex: 1" />
      <explorer-content style="flex: 1" />
    </div>
    <ui-block style="margin-top: 10px; display: flex">
      <ui-button-group
        :class="$style.z"
        :items="[
          { id: 'fs', text: 'FS' },
          { id: 'dm', text: 'DM' },
        ]"
        v-model="tab"
      />

      <ui-button-group
        v-if="tab === 'fs'"
        :class="$style.y"
        :items="[
          { id: 'max', text: 'Max' },
          { id: 'med', text: 'Med' },
          { id: 'min', text: 'Min' },
        ]"
        v-model="displayMode"
      />
    </ui-block>

    <ui-context-menu
      v-if="isShowContextMenu"
      :items="[{ label: 'Sasageeoo' }]"
      :style="contextPosition"
    />
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';

export default defineComponent({
  components: {},
  async mounted() {
    document.addEventListener('click', () => {
      this.isShowContextMenu = false;
    });
  },
  methods: {
    showContextMenu(e: MouseEvent) {
      e.preventDefault();
      this.isShowContextMenu = true;
      this.contextPosition.left = e.pageX + 'px';
      this.contextPosition.top = e.pageY + 'px';
    },
  },
  data: () => {
    return {
      tab: 'fs',
      displayMode: 'max',
      isShowContextMenu: false,
      contextPosition: {
        left: '0',
        top: '0',
      },
    };
  },
});
</script>

<style lang="scss" module>
.main {
  height: 100%;
  box-sizing: border-box;
  padding: 10px;

  .content {
    display: grid;
    grid-template-columns: 280px 1fr 1fr;
    gap: 10px;
    height: calc(100% - 65px);
  }

  button {
    padding: 5px;
    font-weight: normal;
  }

  .z {
    width: max-content;
  }

  .y {
    width: max-content;
    margin-left: auto;
  }
}
</style>
