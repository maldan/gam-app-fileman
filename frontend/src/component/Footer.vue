<template>
  <div :class="$style.footer">
    <!-- Preload -->
    <img
      v-if="$store.state.main.isLoading"
      class="animation_rotation"
      src="../asset/preload.svg"
      style="width: 16px; margin-right: 20px"
    />

    <!-- Buttons -->
    <button v-for="x in extensions" @click="x.onClick()" :key="x.name" class="clickable">
      {{ x.name }}
    </button>

    <!-- Info -->
    <div :class="$style.info">
      <div @click="$store.dispatch('file/switchSort')" class="clickable" :class="$style.item">
        By {{ $store.state.file.sortBy }}
      </div>
      <div :class="$style.item">
        {{ $store.state.file.list.filter((x) => x.isSelected).length }}
        /
        {{ $store.state.file.list.length }}
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';

export default defineComponent({
  props: {},
  components: {},
  async mounted() {
    this.extensions = [
      {
        name: 'Download Manager',
        onClick: () => {
          this.$store.dispatch(`extension/show`, {
            name: 'download',
            data: {},
          });
        },
      },
      {
        name: 'Disk Usage',
        onClick: () => {
          this.$store.dispatch(`extension/show`, {
            name: 'usage',
            data: {},
          });
        },
      },
    ];
  },
  methods: {},
  data: () => {
    return {
      extensions: [] as any[],
    };
  },
});
</script>

<style lang="scss" module>
@import '../gam_sdk_ui/vue/main.scss';

.footer {
  display: flex;
  flex: 1;
  background: #1b1b1b;
  padding: 10px;

  button {
    @include base_button(#2d88ff, #fefefe, 5px 10px, 2px);
    margin-right: 10px;
  }

  .info {
    margin-left: auto;
    align-items: center;
    display: flex;
    justify-content: center;

    .item {
      background: #2d88ff;
      padding: 2px 10px;
      color: #fefefe;
      border-radius: 2px;
      font-size: 14px;
      margin-left: 10px;
    }
  }
}
</style>
