<template>
  <div class="view_image">
    <IconButton class="close" icon="close" @click="$emit('close')" />
    <IconButton class="mode" icon="mode" @click="isWidth = !isWidth" />
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

        <div v-for="x in list" :key="x.name">
          <img
            class="image"
            :src="$root.API_URL + `/file/file?path=${path + '/' + x.name}`"
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
import IconButton from '../IconButton.vue';
import Swipe from 'swipejs';

export default defineComponent({
  props: {
    path: String,
    list: Array,
    file: String,
  },
  components: { IconButton },
  async mounted() {
    const index = (this.list as any[]).findIndex((x) => this.path + '/' + x.name === this.file);
    if (index !== -1) {
      this.currentId = index;
    }
    // @ts-ignore
    window.mySwipe = new Swipe(this.$refs['slider'], {
      startSlide: index,
      speed: 600,
      // auto: 3000,
      draggable: true,
      continuous: true,
      disableScroll: false,
      stopPropagation: false,
      ignore: '.scroller',
      callback: function (index, elem, dir) {},
      transitionEnd: function (index, elem) {},
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
      this.currentId -= 1;
      if (this.currentId <= 0) {
        this.currentId = 0;
      }
      // @ts-ignore
      window.mySwipe.prev();
    },
    next() {
      this.currentId += 1;
      if (this.currentId > (this.list as any[]).length - 1) {
        this.currentId = (this.list as any[]).length - 1;
      }
      // @ts-ignore
      window.mySwipe.next();
    },
  },
  data: () => {
    return {
      currentId: 0,
      isWidth: true,
      keyboardListener: null as any,
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
