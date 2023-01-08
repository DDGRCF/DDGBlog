<template>
  <el-row class="tools" :gutter="15">
    <el-col class="tools-left" :lg="4" :md="4" :sm="4">
      <div class="menu">
        <div class="menu-title">
          <el-icon><Tools /></el-icon><span>&nbsp;&nbsp;</span
          ><span>工具箱</span>
        </div>
        <el-menu class="menu-content" default-active="activeIndex">
          <el-sub-menu index="1">
            <template #title>
              <span>旋转框处理工具</span>
            </template>
            <el-menu-item index="1-1">DOTA提交格式转表格</el-menu-item>
            <el-menu-item index="1-2">非极大值抑制合并</el-menu-item>
          </el-sub-menu>
        </el-menu>
      </div>
    </el-col>
    <el-col class="tools-main" :lg="16" :md="20" :sm="24"
      ><router-view></router-view
    ></el-col>
    <el-col class="tools-right" :lg="4" :md="4" :sm="4"></el-col>
  </el-row>
</template>

<script lang="ts">
import { defineComponent } from "vue";

export default defineComponent({
  name: "ToolsView",
  props: {
    msg: String,
  },
  data() {
    return {
      activeIndex: "0",
    };
  },
});
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped lang="scss">
$padding-size: 5px;
.tools {
  height: 100%;
  // flex 的直系子元素是flex元素，而非直系不是
  // 层级迭代只用使用el元素失效
  .tools-left {
    height: 100%;
    .menu {
      $shadow-color: rgb(199, 193, 193);
      display: flex;
      height: 100%;
      flex-direction: column;
      .menu-title {
        justify-content: center;
        align-items: center;
        padding: $padding-size;
        font-weight: bold;
        box-shadow: 0 0 1em $shadow-color;
      }
      .menu-content {
        flex: auto;
        margin-top: 10px;
        box-shadow: 0 0 1em $shadow-color;
      }
    }
  }
  .tools-main {
    padding: $padding-size;
  }
  .tools-right {
    background-color: yellow;
  }
}
// h3 {
//   margin: 40px 0 0;
// }
// ul {
//   list-style-type: none;
//   padding: 0;
// }
// li {
//   display: inline-block;
//   margin: 0 10px;
// }
// a {
//   color: #42b983;
// }
</style>
