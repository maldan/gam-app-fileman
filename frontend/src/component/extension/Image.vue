<template>
  <div class="view_image">
    <img
      class="close clickable"
      src="../../asset/icon/close.svg"
      @click="$store.dispatch('extension/close')"
    />
    <img class="mode clickable" src="../../asset/icon/mode.svg" @click="isWidth = !isWidth" />
    <!-- <IconButton class="close" icon="close" @click="$emit('close')" />
    <IconButton class="mode" icon="mode" @click="isWidth = !isWidth" /> -->
    <!-- <IconButton class="left" icon="left" @click="prev()" />
    <IconButton class="right" icon="right" @click="next()" /> -->

    <div ref="slider" class="swipe">
      <div class="swipe-wrap">
        <!-- <div>1</div>
        <div>
          <img
            class="image"
            :src="$root.API_URL + `/file/file?path=${path + '/' + list[currentId].name}`"
            :style="isWidth ? { width: '100%' } : { height: '100%' }"
            draggable="false"
          />
        </div>
        <div>2</div> -->

        <!-- <div v-for="x in $store.state.file.list" :key="x.name">
          <img
            class="image"
            :src="$root.API_URL + `/file/file?path=${$store.state.main.path + '/' + x.name}`"
            :style="isWidth ? { width: '100%' } : { height: '100%' }"
            draggable="false"
            loading="lazy"
          />
        </div> -->

        <div>
          <img
            class="image"
            :src="$root.API_URL + `/file/file?path=${prev()}`"
            :style="isWidth ? { width: '100%' } : { height: '100%' }"
            draggable="false"
          />
        </div>
        <div>
          <img
            class="image"
            :src="$root.API_URL + `/file/file?path=${current()}`"
            :style="isWidth ? { width: '100%' } : { height: '100%' }"
            draggable="false"
          />
        </div>
        <div>
          <img
            class="image"
            :src="$root.API_URL + `/file/file?path=${next()}`"
            :style="isWidth ? { width: '100%' } : { height: '100%' }"
            draggable="false"
          />
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import Swipe from 'swipejs';

export default defineComponent({
  props: {
    //path: String,
    //list: Array,
    //file: String,
  },
  components: {},
  async mounted() {
    this.currentId = this.$store.state.extension.data.index;

    //const index = (this.$store.file.list as any[]).findIndex(
    //  (x) => this.$store.main.path + '/' + x.name === this.file,
    //);
    //if (index !== -1) {
    //  this.currentId = index;
    //}
    // @ts-ignore
    window.mySwipe = new Swipe(this.$refs['slider'], {
      startSlide: 1, //this.$store.state.extension.data.index,
      speed: 300,
      // auto: 3000,
      draggable: true,
      continuous: true,
      disableScroll: false,
      stopPropagation: false,
      ignore: '.scroller',
      callback: (index, elem, dir) => {
        //this.dir = dir;
        //console.log(dir);
      },
      transitionEnd: (index, elem) => {
        //this.currentId += 1;
        //console.log(this.currentId);
      },
    });
    // @ts-ignore
    //window.mySwipe.slide(index, 10);

    this.keyboardListener = (e: KeyboardEvent) => {
      if (e.key === 'ArrowRight') {
        this.next();
      }
      if (e.key === 'ArrowLeft') {
        this.prev();
      }
    };
    document.addEventListener('keydown', this.keyboardListener);
  },
  beforeUnmount() {
    document.removeEventListener('keydown', this.keyboardListener);
  },
  methods: {
    prev() {
      return (
        this.$store.state.main.path + '/' + this.$store.state.file.list[this.currentId - 1]?.name
      );
    },
    next() {
      return (
        this.$store.state.main.path + '/' + this.$store.state.file.list[this.currentId + 1]?.name
      );
    },
    current() {
      return this.$store.state.main.path + '/' + this.$store.state.file.list[this.currentId]?.name;
    },
  },
  data: () => {
    return {
      currentId: 0,
      isWidth: true,
      keyboardListener: null as any,
      dir: 0,
    };
  },
});
</script>

<style lang="scss" scoped>
.view_image {
  position: fixed;
  width: 100%;
  height: 100%;
  left: 0;
  top: 0;
  background: rgba(0, 0, 0, 0.5);
  overflow-y: auto;

  .image {
    display: block;
    margin: auto;
  }

  .close,
  .mode {
    position: fixed;
    left: 20px;
    top: 10px;
    z-index: 1;
  }

  .mode {
    left: 60px;
  }

  .swipe {
    overflow: hidden;
    visibility: hidden;
    position: relative;
    height: 100%;
    overflow-y: auto;
  }
  .swipe-wrap {
    //overflow: hidden;
    position: relative;
    height: 100%;
  }
  .swipe-wrap > div {
    float: left;
    width: 100%;
    height: 100%;
    position: relative;
    // overflow: hidden;
    overflow-y: auto;
  }

  /*.left {
    right: 100px;
  }
  .right {
    right: 60px;
  }*/
}
</style>
