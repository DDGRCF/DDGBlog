<template>
  <el-button class="btn">
    <span class="span-group1" v-for="(item, index) in msg" :key="item">
      <span :style="{ transition: String((index + 1) * 0.1) + 's' }">
        <template v-if="item !== ' '">{{ item }}</template>
        <template v-else>&nbsp;</template>
      </span>
    </span>
    <span class="span-group2" v-for="(item, index) in msg" :key="item">
      <span :style="{ transition: String((index + 2) * 0.1) + 's' }">
        <template v-if="item !== ' '">{{ item }}</template>
        <template v-else>&nbsp;</template>
      </span>
    </span>
  </el-button>
</template>

<script lang="ts">
import { defineComponent } from "vue";

export default defineComponent({
  name: "HelloWorld",
  props: {
    msg: String,
  },
  computed: {
    msgSplit(): string[] | undefined {
      return this.msg?.split("");
    },
  },
});
</script>

<style scoped lang="scss">
.btn {
  position: relative;
  display: flex;
  justify-content: center;
  align-items: center;
  border: none;
  background-color: #419cd8;
  font-weight: bold;
  &:hover {
    animation: color-move 1s forwards;
    background: #0675ca;
  }
  &:active {
    transform: scale(0.95);
  }

  &:hover {
    .span-group1 {
      position: absolute;
      opacity: 0;
      & > span {
        transform: translateY(100%);
      }
    }
  }
  .span-group1 {
    display: flex;
    overflow: hidden;
  }

  &:hover {
    .span-group2 {
      position: relative;
      left: 0;
      opacity: 1;
      & > span {
        transform: translateY(0);
      }
    }
  }
  .span-group2 {
    display: flex;
    overflow: hidden;
    position: absolute;
    opacity: 0;
    & > span {
      transition: 0.1s;
      transform: translateY(-100%);
    }
  }
}
</style>
