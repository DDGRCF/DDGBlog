<template>
  <el-affix class="affix">
    <el-menu
      :default-active="activeIndex"
      class="top-menu"
      mode="horizontal"
      :ellipsis="false"
      active-text-color="black"
      @select="handleSelect"
    >
      <el-menu-item index="0">
        <div>
          <el-icon :size="icon.size"><MyLogo /></el-icon>
        </div>
      </el-menu-item>
      <div style="flex-grow: 1" />
      <el-menu-item index="1">
        <div>
          <el-icon :size="icon.size"><Avatar /></el-icon>
        </div>
      </el-menu-item>
      <el-menu-item index="2">
        <div>
          <el-icon :size="icon.size"><About /></el-icon>
        </div>
      </el-menu-item>
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
import MyLogo from "./imgs/icons/MyLogo.vue";
import Avatar from "./imgs/icons/Avatar.vue";
import About from "./imgs/icons/About.vue";
import fetch from "@/api/fetch";

export default defineComponent({
  name: "WebHeader",
  components: {
    MyLogo,
    Avatar,
    About,
  },
  data() {
    return {
      menuStatus: {
        activeIndex: "0",
        menuUrl: {
          hello: "/api/hello",
          home: "/api/home",
          tool: {
            obb: {
              nms: "/api/tool/obb/nms",
              submit: "/api/tool/obb/submit",
            },
          },
        },
      },
      icon: {
        size: 28,
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
            this.$router.push(this.menuStatus.menuUrl.tool.obb.submit);
            this.menuStatus.activeIndex = key;
            this.$store.commit("changeActInd", "1-1");
          } else if (key === "3-1-2") {
            this.$router.push(this.menuStatus.menuUrl.tool.obb.nms);
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

<style scoped lang="scss">
$menu-hover-color: #7f8fa6;
.top-menu {
  width: 100%;
  padding: 0;
  margin: 0;
  background-color: #ffffffbb;
  border-radius: 0 0 5px 5px;
}

.el-menu--horizontal .el-sub-menu .el-menu-item:not(.is-disabled):hover,
.el-menu--horizontal .el-sub-menu .el-menu-item:not(.is-disabled):focus {
  color: #2980b9 !important;
  animation: action 1s forwards !important;
}

@keyframes action {
  from {
    width: 0;
  }
  to {
    width: 100%;
  }
}
</style>
