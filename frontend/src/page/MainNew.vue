<template>
  <div class="main">
    <div class="header">
      <Header />
      <Path />
    </div>

    <div class="body">
      <div v-if="!$store.state.isLoading" class="files">
        <File
          @click.stop="selectFile($event, x)"
          v-for="x in $store.state.files"
          :key="x.name"
          :file="x"
        />
      </div>
    </div>

    <Footer />
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import Header from '../component/Header.vue';
import Footer from '../component/Footer.vue';
import File from '../component/File.vue';
import Path from '../component/Path.vue';

export default defineComponent({
  components: { Header, Footer, File, Path },
  async mounted() {
    this.$store.dispatch('getFiles');

    document.addEventListener('click', () => {
      this.$store.dispatch('clearSelection');
    });
  },

  methods: {
    selectFile(e: MouseEvent, x: any) {
      // Hold control
      /*if (e.ctrlKey) {
        x.isSelected = !x.isSelected;
        this.lastSelectedFile = x;
        return;
      }

      // Hold shift
      if (e.shiftKey && this.lastSelectedFile) {
        const lastIndex = this.lastSelectedFile.index;
        this.clearSelection();
        let direction = x.index - lastIndex;
        if (direction > 0) {
          for (let i = 0; i <= direction; i++) {
            // this.files[i + lastIndex].isSelected = true;
          }
        } else {
          for (let i = 0; i <= -direction; i++) {
            //this.files[lastIndex - i].isSelected = true;
          }
        }
        return;
      }
      this.clearSelection();*/
      //x.isSelected = true;
      //this.lastSelectedFile = x;

      if (e.ctrlKey) {
        this.$store.dispatch('selectFile', x);
      } else if (e.shiftKey && this.$store.state.lastSelectedFile) {
        const lastIndex = this.$store.state.lastSelectedFile.index;
        let direction = x.index - lastIndex;
        if (direction > 0) {
          for (let i = 0; i <= direction; i++) {
            this.$store.dispatch('selectFile', this.$store.state.files[i + lastIndex]);
          }
        } else {
          for (let i = 0; i <= -direction; i++) {
            this.$store.dispatch('selectFile', this.$store.state.files[lastIndex - i]);
          }
        }
      } else {
        this.$store.dispatch('clearSelection');
        this.$store.dispatch('selectFile', x);
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
