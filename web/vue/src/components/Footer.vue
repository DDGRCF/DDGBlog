<template>
  <el-divider class="divider" content-position="center"></el-divider>
  <div class="desc">
    <span>CopyRight DDGRCF</span>
    <el-icon :size="icon.size"><Github /></el-icon>
    <el-icon :size="icon.size"><Gmail /></el-icon>
    <el-icon :size="icon.size"><QQ /></el-icon>
  </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import Github from "./imgs/Github.vue";
import QQ from "./imgs/QQ.vue";
import Gmail from "./imgs/Gmail.vue";

export default defineComponent({
  name: "WebFooter",
  components: {
    Github,
    QQ,
    Gmail,
  },
  data() {
    return {
      icon: {
        size: 25,
      },
    };
  },
});
</script>
<style scoped lang="scss">
.divider {
  width: 60%;
  margin: 1em auto 1em;
}

.desc {
  display: flex;
  justify-content: center;
  align-items: center;
  & > * {
    margin: 0 5px 0;
  }
  & > span {
    font-size: 0.8em;
    font-weight: bold;
    vertical-align: middle;
    margin-bottom: 0.1em;
  }
}
</style>
