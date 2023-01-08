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
import axios from "axios";
import LogoView from "./imgs/Logo.vue";

export default defineComponent({
  name: "WebHeader",
  components: {
    LogoView,
  },
  data() {
    return {
      menuUrl: {
        hello: "http://localhost:12000/hello",
        home: "/api/home",
        tools: {
          obb: {
            nms: "/api/tools/obb/nms",
            submit: "/api/tools/obb/submit",
          },
        },
      },
      activeIndex: "0",
    };
  },
  methods: {
    handleSelect(key: string, keyPath: string[]) {
      if (key === "0") {
        this.$router.push(this.menuUrl.home);
      } else if (key === "3-1-1") {
        this.$router.push(this.menuUrl.tools.obb.submit);
      } else if (key === "3-1-2") {
        this.$router.push(this.menuUrl.tools.obb.nms);
      } else {
        axios
          .get(this.menuUrl.hello)
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
});
</script>

<style scoped lang="scss"></style>
