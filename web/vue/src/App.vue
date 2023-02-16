<template>
  <router-view />
</template>

<script lang="ts">
import particles from "@/conf/particles";
import type { Engine } from "tsparticles-engine";
import type { Container } from "tsparticles-engine";
import { loadFull } from "tsparticles";

export default {
  data() {
    return {
      particlesconf: particles.conf,
    };
  },
  methods: {
    async particlesInit(engine: Engine) {
      await loadFull(engine);
    },
    async particlesLoaded(container: Container) {
      console.log("Particles container loaded", container);
    },
  },
};
</script>

<style lang="scss">
html,
body {
  width: 100%;
  height: 100%;
  background: #ffff;
  overflow: hidden;
  font-family: "Helvetica Neue", Helvetica, "PingFang SC", "Hiragino Sans GB",
    "Microsoft YaHei", "微软雅黑", Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin: 0;
  padding: 0;
  --el-color-primary: #2c3e50;
}
#tsparticles {
  display: flex;
  flex: auto;
  height: 100vh;
  width: 100%;
  background-size: cover;
  background-repeat: no-repeat;
}
</style>
