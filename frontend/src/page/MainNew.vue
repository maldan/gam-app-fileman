<template>
  <div class="main">
    <div class="header">
      <Header />
      <Path />
    </div>

    <div class="body">
      <div v-if="!$store.state.main.isLoading" class="files">
        <File
          @click.stop="selectFile($event, x)"
          v-for="x in $store.state.file.list"
          :key="x.name"
          :file="x"
        />
      </div>
    </div>

    <Footer />

    <ui-modal v-if="$store.state.modal.name">
      <component :is="$store.state.modal.name" />
    </ui-modal>
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import Header from '../component/Header.vue';
import Footer from '../component/Footer.vue';
import File from '../component/File.vue';
import Path from '../component/Path.vue';
import ModalName from '../component/modal/ModalName.vue';

export default defineComponent({
  components: { Header, Footer, File, Path, ModalName },
  async mounted() {
    this.$store.dispatch('main/getPath');

    document.addEventListener('click', () => {
      this.$store.dispatch('file/clearSelection');
    });
  },

  methods: {
    selectFile(e: MouseEvent, x: any) {
      if (e.ctrlKey) {
        this.$store.dispatch('file/select', x);
      } else if (e.shiftKey && this.$store.state.file.lastSelected) {
        const lastIndex = this.$store.state.file.lastSelected.index;
        let direction = x.index - lastIndex;
        if (direction > 0) {
          for (let i = 0; i <= direction; i++) {
            this.$store.dispatch('file/select', this.$store.state.file.list[i + lastIndex]);
          }
        } else {
          for (let i = 0; i <= -direction; i++) {
            this.$store.dispatch('file/select', this.$store.state.file.list[lastIndex - i]);
          }
        }
      } else {
        this.$store.dispatch('file/clearSelection');
        this.$store.dispatch('file/select', x);
      }
    },
  },
  data: () => {
    return {};
  },
});
</script>

<style lang="scss" scoped>
.main {
  height: 100%;
  box-sizing: border-box;

  .header {
    display: flex;
    flex-direction: column;
  }

  .body {
    height: calc(100% - 164px);
    overflow-y: auto;
    margin-top: 10px;

    .files {
      display: grid;
      grid-template-columns: 1fr 1fr 1fr 1fr 1fr 1fr 1fr;
      grid-auto-rows: max-content;
    }
  }
}
</style>
