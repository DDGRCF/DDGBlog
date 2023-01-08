<template>
  <div class="common-layout">
    <el-container class="container">
      <el-header class="header"><HeaderVue /></el-header>
      <el-main class="main">
        <router-view></router-view>
      </el-main>
      <el-footer class="footer"><FooterVue /></el-footer>
    </el-container>
  </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import HeaderVue from "@/components/Header.vue";
import FooterVue from "@/components/Footer.vue";

export default defineComponent({
  name: "HomeView",
  components: {
    HeaderVue,
    // MainVue,
    FooterVue,
  },
});
</script>

<style scoped lang="scss">
.common-layout {
  display: flex;
  flex-direction: column;
  position: absolute;
  width: 100%;
  height: 100%;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  padding: 0;
  margin: 0;

  // 需要加overflow: auto
  overflow: auto;
  &::-webkit-scrollbar {
    width: 8px;
    height: 8px;
  }
  &::-webkit-scrollbar-track {
    background: rgb(239, 239, 239);
    border-radius: 2px;
  }
  &::-webkit-scrollbar-thumb {
    background: #bfbfbf;
    border-radius: 10px;
  }
  &::-webkit-scrollbar-thumb:hover {
    background: #aaaaaa;
  }
  &::-webkit-scrollbar-corner {
    background: #179a16;
  }

  .container {
    flex-direction: column;
    .header {
      flex: 0 1 0;
    }
    .main {
      flex: 1 0 auto;
    }
    .footer {
      flex: 0 1 0;
      margin-bottom: 0.8em;
    }
  }
}
</style>
