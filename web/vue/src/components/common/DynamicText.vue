<template>
  <div
    class="dynamic-text"
    :style="{ fontFamily: fontInfo.fontFamily, fontSize: fontInfo.fontSize }"
  >
    <span>{{ showContent }}</span>
    <span v-if="isCursor" class="cursor">|</span>
    <span v-else class="cursor">&nbsp;</span>
  </div>
</template>

<script lang="ts">
import { defineComponent, PropType } from "vue";

interface FontInfo {
  fontFamily: string;
  fontSize: number;
  fontUnit: string;
}

interface FontObject {
  fontFamily: string;
  fontSize: string;
}

export default defineComponent({
  name: "DynamicText",
  props: {
    msg: {
      type: String,
      default: "逆风而行，或许步履维艰，但沿途的风景会更加秀丽。",
    },
    interval: {
      type: Number,
      default: 100,
    },
    font: {
      type: Object as PropType<FontInfo>,
      default: () => ({
        fontFamily: "YEFONTYouLongXingKai",
        fontSize: 18,
        fontUnit: "px",
      }),
    },
  },
  data() {
    return {
      content: this.msg,
      showContent: "",
      inter: this.interval,
      isCursor: true,
      index: 0,
      timer: -1,
      cursorTimes: {
        index: 0,
        num: 8,
        inter: 500,
      },
    };
  },
  mounted() {
    this.timer = setInterval(this.contentIntervalTyping, this.inter);
  },
  unmounted() {
    clearInterval(this.timer);
  },
  methods: {
    contentIntervalTyping() {
      clearInterval(this.timer);
      if (this.index <= this.content.length) {
        this.inter = Math.random() * 400 + 100;
        // this.showContent += this.content[this.index];
        this.showContent = this.content.slice(0, this.index);
        this.timer = setInterval(this.contentIntervalTyping, this.inter);
        this.index++;
      } else {
        this.index = 0;
        this.timer = setInterval(
          this.cursorIntervalTyping,
          this.cursorTimes.inter
        );
      }
    },
    cursorIntervalTyping() {
      this.isCursor = !this.isCursor;
      if (this.cursorTimes.index === this.cursorTimes.num) {
        this.isCursor = true;
        this.cursorTimes.index = 0;
        clearInterval(this.timer);
        this.timer = setInterval(this.contentIntervalTyping, this.inter);
      }
      this.cursorTimes.index++;
    },
  },
  computed: {
    fontInfo(): FontObject {
      let fontInfo = {} as FontObject;
      fontInfo.fontSize = String(this.font.fontSize) + this.font.fontUnit;
      fontInfo.fontFamily = this.font.fontFamily;
      return fontInfo;
    },
  },
});
</script>

<style scoped lang="scss">
@import "@/assets/css/common.scss";
@import "@/assets/fonts/fonts.scss";
.dynamic-text {
  display: flex;
  align-items: center;
  justify-content: center;
  font-family: "YEFONTYouLongXingKai";
  .cursor {
    display: inline-block;
    vertical-align: middle;
  }
}
</style>
