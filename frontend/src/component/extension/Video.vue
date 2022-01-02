<template>
  <div class="view_video">
    <!-- Video -->
    <video
      @click="changeState"
      v-doubleclick="changeOffset"
      ref="video"
      :src="`${$root.API_URL}/file/file?path=${$store.state.extension.data.fullPath}`"
    ></video>

    <img
      class="close clickable"
      src="../../asset/icon/close.svg"
      @click="$store.dispatch('extension/close')"
    />

    <!-- Bar -->
    <div @click="clickOnBar" ref="bar" class="bar">
      <div class="current">{{ current }}</div>
      <div class="duration">{{ duration }}</div>
      <div class="fill" :style="{ width: progress * 100 + '%' }"></div>
      <div
        v-if="loopStart && loopEnd"
        class="loop_fill"
        :style="{
          left: (loopStart / $refs['video'].duration) * 100 + '%',
          width: ((loopEnd - loopStart) / $refs['video'].duration) * 100 + '%',
        }"
      ></div>

      <div
        v-for="x in bookmark"
        class="bookmark"
        :key="x[0]"
        :style="{
          left: (x[0] / $refs['video'].duration) * 100 + '%',
          width: ((x[1] - x[0]) / $refs['video'].duration) * 100 + '%',
        }"
      ></div>
    </div>

    <!-- Control -->
    <div v-if="$refs['video']" class="control">
      <button class="clickable" @click="offset(-$refs['video'].playbackRate)">&lt;</button>
      <button class="clickable" @click="offset($refs['video'].playbackRate)">&gt;</button>

      <button class="clickable" @click="changeSpeed(-0.1)">Slow</button>
      <div style="width: 50px; text-align: center">
        {{ $refs['video'].playbackRate.toFixed(1) }}
      </div>
      <button class="clickable" @click="changeSpeed(0.1)">Fast</button>
      <button class="clickable" @click="loopStart = $refs['video'].currentTime">Loop Start</button>
      <button class="clickable" @click="loopEnd = $refs['video'].currentTime">Loop End</button>
      <button class="clickable" @click="(loopStart = 0), (loopEnd = 0)">Loop Clear</button>
      <button
        class="clickable"
        @click="bookmark.push([loopStart, loopEnd]), (loopStart = 0), (loopEnd = 0), saveInfo()"
      >
        Save Loop
      </button>

      <button class="clickable" @click="changeVolume(-0.1)">Low</button>
      <div style="width: 50px; text-align: center">
        {{ $refs['video'].volume.toFixed(1) }}
      </div>
      <button class="clickable" @click="changeVolume(0.1)">High</button>

      <button
        class="clickable"
        @click="
          $store.dispatch('file/setVideoPreview', {
            path: $store.state.extension.data.fullPath,
            time: current,
          })
        "
      >
        Set Preview
      </button>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import Moment from 'moment';

export default defineComponent({
  props: {},
  components: {},
  async mounted() {
    setInterval(() => {
      const video = this.$refs['video'] as HTMLVideoElement;

      if (!video) return 0;

      this.progress = video.currentTime / video.duration;

      this.current = Moment.utc(
        Moment.duration(video.currentTime, 'seconds').asMilliseconds(),
      ).format('HH:mm:ss.S');

      this.duration = Moment.utc(
        Moment.duration(video.duration, 'seconds').asMilliseconds(),
      ).format('HH:mm:ss');
    }, 16);

    document.addEventListener('keydown', this.keyEvents);
    (this.$refs['video'] as HTMLVideoElement).addEventListener('timeupdate', () => {
      if (this.loopStart && this.loopEnd) {
        if ((this.$refs['video'] as HTMLVideoElement).currentTime >= this.loopEnd) {
          (this.$refs['video'] as HTMLVideoElement).currentTime = this.loopStart;
        }
      }
    });
  },
  beforeUnmount() {
    document.removeEventListener('keydown', this.keyEvents);
  },
  methods: {
    changeOffset(x: number) {
      if (x < 0.5) {
        this.offset(-(this.$refs['video'] as any).playbackRate * 5);
      } else {
        this.offset((this.$refs['video'] as any).playbackRate * 5);
      }
    },

    clickOnBar(e: any) {
      const video = this.$refs['video'] as HTMLVideoElement;
      const bar = this.$refs['bar'] as HTMLElement;

      const offset = (e.pageX - bar.getBoundingClientRect().x) / bar.getBoundingClientRect().width;
      video.currentTime = offset * video.duration;
    },
    changeState() {
      const video = this.$refs['video'] as HTMLVideoElement;

      if (video.paused) {
        video.play();
      } else {
        video.pause();
      }

      this.isPlay = !video.paused;
    },
    keyEvents(e: any) {
      const video = this.$refs['video'] as HTMLVideoElement;
      if (!video) {
        return;
      }

      if (e.key === 'ArrowRight') {
        video.currentTime += video.playbackRate;
      }
      if (e.key === 'ArrowLeft') {
        video.currentTime -= video.playbackRate;
      }
      if (e.key === ' ') {
        if (video.paused) {
          video.play();
        } else {
          video.pause();
        }
      }
    },
    offset(p: number) {
      const video = this.$refs['video'] as HTMLVideoElement;
      video.currentTime += p;
    },
    changeSpeed(speed: number) {
      const video = this.$refs['video'] as HTMLVideoElement;
      video.playbackRate += speed;
    },
    changeVolume(amount: number) {
      const video = this.$refs['video'] as HTMLVideoElement;
      video.volume += amount;
    },
  },
  data: () => {
    return {
      currentId: 0,

      isPlay: false,
      progress: 0,
      current: '',
      duration: '',

      loopStart: 0,
      loopEnd: 0,
      bookmark: [],

      info: {} as any,
    };
  },
});
</script>

<style lang="scss" scoped>
.view_video {
  position: fixed;
  width: 100%;
  height: 100%;
  left: 0;
  top: 0;
  background: rgba(0, 0, 0, 0.99);
  overflow-y: auto;

  .close,
  .left,
  .right {
    position: fixed;
    right: 10px;
    top: 10px;
  }

  .left {
    right: 90px;
  }
  .right {
    right: 50px;
  }

  video {
    width: 100%;
    height: calc(100% - 72px);
  }

  .control {
    display: flex;
    align-items: center;
    margin-top: 8px;

    button {
      background: #2c72e2;
      color: #fefefe;
      border: 0;
      padding: 10px 20px;
    }
  }

  .bar {
    display: flex;
    position: relative;
    bottom: 0px;
    width: 100%;
    font-size: 17px;
    height: 24px;
    box-sizing: border-box;
    color: #d4d4d4;
    // border-top: 1px solid #19191985;

    .fill,
    .loop_fill {
      width: 0%;
      background: #b4b4b45e;
    }
    .loop_fill {
      position: absolute;
      left: 0;
      top: 0;
      height: 50%;
      background: #03c513ce;
    }
    .bookmark {
      position: absolute;
      left: 0;
      top: 19px;
      height: 25%;
      background: #e0ad03e8;
    }

    .duration,
    .current {
      position: absolute;
    }

    .current {
      left: 10px;
      top: 2px;
    }

    .duration {
      right: 10px;
      top: 2px;
    }
  }
}
</style>
