<template>
  <div>
    <el-row :gutter="colNums.gutter" class="tools-affix">
      <el-col
        class="tools-left hidden-sm-and-down"
        ref="toolsLeftSide"
        :lg="colNums.left.lg"
        :md="colNums.left.md"
      >
        <div class="menu">
          <div class="menu-title" v-if="!menuStatus.isCollapse">
            <el-icon><Tools /></el-icon><span>&nbsp;&nbsp;</span
            ><span>工具箱</span>
          </div>
          <el-menu
            :collapse="menuStatus.isCollapse"
            class="menu-content"
            :default-active="activeIndex"
            @select="handleSelect"
          >
            <el-sub-menu index="1">
              <template #title
                ><el-icon><Position /></el-icon
                ><span>旋转框处理工具</span></template
              >
              <el-menu-item index="1-1">DOTA提交格式转表格</el-menu-item>
              <el-menu-item index="1-2">非极大值抑制合并</el-menu-item>
              <el-menu-item index="1-3">十字交叉法转换</el-menu-item>
            </el-sub-menu>
          </el-menu>
        </div>
      </el-col>
      <el-col
        class="tools-right"
        :lg="colNums.main.lg"
        :md="colNums.main.md"
        :sm="colNums.main.sm"
      ></el-col>
    </el-row>
  </div>
  <el-row :gutter="colNums.gutter" class="tools">
    <el-col
      class="hidden-sm-and-down"
      :lg="colNums.left.lg"
      :md="colNums.left.md"
    ></el-col>
    <el-col
      class="tools-main"
      :lg="colNums.main.lg"
      :md="colNums.main.md"
      :sm="colNums.main.sm"
    >
      <el-row>
        <el-col
          class="tools-context"
          :lg="colNums.main.left.lg"
          :md="colNums.main.left.md"
        >
          <router-view></router-view>
        </el-col>
        <el-col
          class="tools-right hidden-sm-and-down"
          :lg="colNums.main.right.lg"
          :md="colNums.main.right.md"
        ></el-col>
      </el-row>
    </el-col>
  </el-row>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import { ElAside } from "element-plus";

export default defineComponent({
  name: "ToolsView",
  data() {
    return {
      colNums: {
        gutter: 5,
        left: {
          lg: 4,
          md: 2,
        },
        main: {
          lg: 20,
          md: 22,
          sm: 24,
          left: {
            lg: 19,
            md: 21,
          },
          right: {
            lg: 5,
            md: 3,
          },
        },
      },
      menuStatus: {
        isCollapse: false,
        collapseSize: 200,
        menuUrl: {
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
  unmounted() {
    window.removeEventListener("resize", this.watchComponentWidth);
  },
  mounted() {
    window.addEventListener("resize", this.watchComponentWidth);
  },
  methods: {
    watchComponentWidth() {
      let clientWidth = (this.$refs.toolsLeftSide as typeof ElAside).$el
        .clientWidth;
      this.menuStatus.isCollapse = clientWidth < this.menuStatus.collapseSize;
    },
    handleSelect(key: string, keyPath: string[]) {
      this.$store.commit("changeUpInd", "3");
      if (key === "1-1") {
        this.$router.push(this.menuStatus.menuUrl.tools.obb.submit);
        this.$store.commit("changeActInd", "1-1");
      } else if (key === "1-2") {
        this.$router.push(this.menuStatus.menuUrl.tools.obb.nms);
        this.$store.commit("changeActInd", "1-2");
      }
      console.log(key, keyPath);
    },
  },
  computed: {
    activeIndex() {
      return this.$store.getters.getActInd();
    },
  },
});
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped lang="scss">
@import "@/assets/css/common.scss";
$padding-size: 5px;

.tools-affix {
  position: fixed;
  left: 0;
  right: 0;
  z-index: 1;
  height: calc(100% - 150px);

  & > .tools-left {
    height: 100%;
    flex: 1 1 auto;
    & > .menu {
      display: flex;
      flex-direction: column;
      height: 100%;
      $shadow-color: rgb(199, 193, 193);
      margin-left: 20px;
      & > .menu-title {
        justify-content: center;
        align-items: center;
        padding: $padding-size;
        font-weight: bold;
        box-shadow: 0 0 0.5em $shadow-color;
      }
      & > .menu-content {
        @include custom-scrollbar;
        flex: auto;
        margin-top: 10px;
        box-shadow: 0 0 0.5em $shadow-color;
      }
    }
  }
}

$context-color: #ffffff;
.tools {
  height: 100%;
  left: 0;
  right: 0;
  // flex 的直系子元素是flex元素，而非直系不是
  // 层级迭代只用使用el元素失效
  .tools-main {
    z-index: 2;
    height: 100%;
    & > :first-child {
      height: 100%;
      .tools-context {
        z-index: inherit;
        height: 100%;
        @include custom-textarea;
      }
    }
  }
}
</style>
