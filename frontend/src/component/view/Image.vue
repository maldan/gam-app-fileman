<template>
  <div class="view_image">
    <IconButton class="close" icon="close" @click="$emit('close')" />
    <IconButton class="left" icon="left" @click="prev()" />
    <IconButton class="right" icon="right" @click="next()" />

    <img
      :src="$root.API_URL + `/file/file?path=${path + '/' + list[currentId].name}`"
      style="width: 100%"
    />
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import IconButton from '../IconButton.vue';

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
  },
  methods: {
    prev() {
      this.currentId -= 1;
      if (this.currentId <= 0) {
        this.currentId = 0;
      }
    },
    next() {
      this.currentId += 1;
      if (this.currentId > (this.list as any[]).length - 1) {
        this.currentId = (this.list as any[]).length - 1;
      }
    },
  },
  data: () => {
    return {
      currentId: 0,
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

  .close,
  .left,
  .right {
    position: fixed;
    right: 20px;
    top: 10px;
  }

  .left {
    right: 100px;
  }
  .right {
    right: 60px;
  }
}
</style>
