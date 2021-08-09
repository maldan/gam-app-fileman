<template>
  <div @click="click()" class="item clickable" :class="isOdd ? 'odd' : ''">
    <img class="icon" :src="icon(item)" alt="File" />
    <div class="name">{{ item.name }}</div>
    <div v-if="!item.isSelected" class="user">{{ item.user }}</div>
    <div v-if="!item.isSelected" class="size">{{ $root.pretty(item.size) }}</div>
    <div v-if="!item.isSelected" class="created">
      {{ $root.moment(item.created).format('YYYY-MM-DD HH:mm:ss') }}
    </div>

    <IconButton
      @click.stop="download()"
      v-if="item.isSelected"
      icon="download"
      style="margin-left: auto"
    />
    <IconButton
      @click.stop="rename()"
      v-if="item.isSelected"
      icon="pencil"
      style="margin-left: 15px"
    />
    <IconButton
      @click.stop="remove()"
      v-if="item.isSelected"
      icon="delete"
      style="margin-left: 15px"
    />
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import IconButton from '../component/IconButton.vue';
import { RestApi } from '../util/RestApi';

export default defineComponent({
  props: {
    item: Object,
    isOdd: Boolean,
    path: String,
  },
  components: { IconButton },
  async mounted() {},
  methods: {
    icon(file: any) {
      if (file.kind === 'file') {
        if (file.name.match(/\.(png|jpeg|gif|jpg)$/)) {
          return `${(this.$root as any).API_URL}/file/imageThumbnail?path=${this.path}/${
            file.name
          }`;
        }
        if (file.name.match(/\.(mp4|avi)$/)) {
          return `${(this.$root as any).API_URL}/file/videoThumbnail?path=${this.path}/${
            file.name
          }`;
        }
        // @ts-ignore
        return require('../asset/file/file.svg');
      }
      // @ts-ignore
      return require('../asset/file/folder.svg');
    },
    download() {
      window.open(
        `${(this.$root as any).API_URL}/file/file?path=${this.path + '/' + this.item?.name}`,
        '_blank',
      );
    },
    async rename() {
      const to = prompt('Enter new name', this.item?.name);
      if (to) {
        await RestApi.file.rename(this.path + '/' + this.item?.name, this.path + '/' + to);
      }
      this.$emit('refresh');
    },
    async remove() {
      if (confirm(`Are you sure? Delete ${this.path}/${this.item?.name}`)) {
        if (this.item?.kind === 'dir') {
          if (prompt('Enter "sasageo" to approve action') === 'sasageo') {
            await RestApi.file.deleteDir(this.path + '/' + this.item?.name);
          }
        } else {
          await RestApi.file.deleteFile(this.path + '/' + this.item?.name);
        }
        this.$emit('refresh');
      }
    },
    click() {
      if (new Date().getTime() - this.delay.getTime() > 300) {
        this.clickTime = 0;
      }
      this.clickTime += 1;
      if (this.item) {
        this.item.isSelected = !this.item.isSelected;
      }
      if (this.clickTime > 1) {
        this.clickTime = 0;
        this.doubleClick();
      }
      this.delay = new Date();
    },
    doubleClick() {
      if (this.item?.kind === 'dir') {
        this.$emit('update', this.path + '/' + this.item?.name);
      } else {
        this.$emit('open', this.path + '/' + this.item?.name);
      }
    },
  },
  data: () => {
    return {
      clickTime: 0,
      delay: new Date(),
    };
  },
});
</script>

<style lang="scss" scoped>
.item {
  display: flex;
  align-items: center;
  padding: 10px;
  color: #ababab;
  font-size: 14px;

  .icon {
    width: 80px;
  }

  img {
    display: block;
  }

  .name {
    margin-left: 10px;
    flex: 1;
    overflow: hidden;
    text-overflow: ellipsis;
    white-space: nowrap;
  }

  .user {
    color: #8ab0f1;
    flex: 0.2;
    text-align: right;
    padding-right: 15px;
  }

  .size {
    color: #5eafa1;
    font-weight: bold;
    text-align: right;
    padding-right: 15px;
    flex: none;
    width: 65px;
  }

  .created {
    color: #fec700;
    flex: none;
  }
}

.item.odd {
  background: #333333;
}
</style>
