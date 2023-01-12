<template>
  <el-affix>
    <el-menu
      :default-active="activeIndex"
      class="el-menu-demo"
      mode="horizontal"
      :ellipsis="false"
      @select="handleSelect"
    >
      <el-menu-item index="0"><LogoView /></el-menu-item>
      <div style="flex-grow: 1" />
      <el-menu-item index="1">登录</el-menu-item>
      <el-menu-item index="2">关于</el-menu-item>
      <el-sub-menu index="3">
        <template #title>小工具</template>
        <el-sub-menu index="3-1">
          <template #title>旋转框处理工具</template>
          <el-menu-item index="3-1-1">DOTA提交格式转表格</el-menu-item>
          <el-menu-item index="3-1-2">非极大值抑制合并</el-menu-item>
          <el-menu-item index="3-1-3">十字交叉法转换</el-menu-item>
        </el-sub-menu>
        <el-menu-item index="3-2">item two</el-menu-item>
        <el-menu-item index="3-3">item three</el-menu-item>
      </el-sub-menu>
    </el-menu>
  </el-affix>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import { ElMessage } from "element-plus";
import LogoView from "./imgs/Logo.vue";
import fetch from "@/api/fetch";

export default defineComponent({
  name: "WebHeader",
  components: {
    LogoView,
  },
  data() {
    return {
      menuStatus: {
        activeIndex: "0",
        menuUrl: {
          hello: "/hello",
          home: "/api/home",
          tools: {
            obb: {
              nms: "/api/tools/obb/nms",
              submit: "/api/tools/obb/submit",
            },
          },
        },
      },
    };
  },
  methods: {
    handleSelect(key: string, keyPath: string[]) {
      if (key.startsWith("0")) {
        this.$router.push(this.menuStatus.menuUrl.home);
        this.menuStatus.activeIndex = key;
      } else if (key.startsWith("3")) {
        this.$store.commit("changeUpInd", "3");
        if (key.startsWith("3-1")) {
          if (key === "3-1-1") {
            this.$router.push(this.menuStatus.menuUrl.tools.obb.submit);
            this.menuStatus.activeIndex = key;
            this.$store.commit("changeActInd", "1-1");
          } else if (key === "3-1-2") {
            this.$router.push(this.menuStatus.menuUrl.tools.obb.nms);
            this.menuStatus.activeIndex = key;
            this.$store.commit("changeActInd", "1-2");
          }
        }
      } else {
        fetch
          .get(this.menuStatus.menuUrl.hello)
          .then(function (response) {
            let responseData = response.data;
            console.log(responseData.msg);
            ElMessage({
              message: responseData.code + ": " + responseData.msg,
              type: "success",
            });
          })
          .catch(function (reason) {
            ElMessage({
              message: reason,
              type: "error",
            });
          });
      }
      console.log(key, keyPath);
    },
  },
  computed: {
    activeIndex() {
      return this.$store.getters.getActInd(true);
    },
  },
});
</script>

<style scoped lang="scss"></style>
