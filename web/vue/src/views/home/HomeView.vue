<template>
  <el-row class="home-affix" :gutter="colsNum.gutter">
    <el-col class="affix-left hidden-md-and-down" :lg="6">
      <el-card class="star-card" shadow="hover">
        <el-icon><Star /></el-icon>
        <p />
        <el-icon><Opportunity /></el-icon>
      </el-card>
    </el-col>
    <el-col
      class="affix-main"
      :lg="colsNum.main.lg"
      :md="colsNum.main.md"
      :sm="colsNum.main.sm"
    >
    </el-col>
    <el-col class="affix-right hidden-md-and-down" :lg="6">
      <MyCard
        :url="avatarConf.url"
        :name="avatarConf.name"
        :about="avatarConf.about"
        :items="avatarConf.items"
      />
    </el-col>
  </el-row>

  <el-row class="home" :gutter="colsNum.gutter">
    <el-col
      :lg="colsNum.left.lg"
      :md="colsNum.left.md"
      :sm="colsNum.left.sm"
    ></el-col>
    <el-col
      class="home-main"
      :lg="colsNum.main.lg"
      :md="colsNum.main.md"
      :sm="colsNum.main.sm"
    >
      <div class="home-carousel">
        <el-carousel height="250px">
          <el-carousel-item v-for="item in 4" :key="item">
            <h3 class="small justify-center" text="2xl">{{ item }}</h3>
          </el-carousel-item>
        </el-carousel>
      </div>
      <div class="home-context">
        <div
          v-for="item in 80"
          :key="item"
          style="height: 200px; border: 1px solid black"
        ></div>
      </div>
    </el-col>
    <el-col
      :lg="colsNum.main.lg"
      :md="colsNum.main.md"
      :sm="colsNum.main.sm"
    ></el-col>
  </el-row>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import { ElAvatar } from "element-plus";
import MyCard from "@/components/imgs/cards/MyCard.vue";
import { CardItemType } from "Common";

export default defineComponent({
  name: "HomeView",
  data() {
    return {
      colsNum: {
        gutter: 5,
        left: {
          lg: 6,
          md: 3,
          sm: 2,
        },
        main: {
          lg: 12,
          md: 18,
          sm: 20,
        },
        right: {
          lg: 6,
          md: 3,
          sm: 2,
        },
      },
      avatarConf: {
        url: require("../../assets/imgs/avatar/avatar-0.jpg"),
        name: "DDG RCF",
        about: "Focus on the region of Deep Learning in the computer vision. ",
        items: [
          {
            num: 80,
            word: "forks",
          },
          {
            num: 23200,
            word: "stars",
          },
          {
            num: 180000,
            word: "followers",
          },
        ] as CardItemType[],
        size: 225,
        isShow: true,
        fit: "fill",
      },
    };
  },
  components: {
    MyCard,
  },
  unmounted() {
    window.removeEventListener("resize", this.watchComponentWidth);
  },
  mounted() {
    window.addEventListener("resize", this.watchComponentWidth);
  },
  methods: {
    watchComponentWidth() {
      let infoCardClientWidth = (this.$refs.infoCard as HTMLElement)
        ?.clientWidth;
      let infoAvatarClientWidth = (this.$refs.infoAvatar as typeof ElAvatar)
        ?.$el?.clientWidth;
      if (infoCardClientWidth && infoAvatarClientWidth) {
        console.log(infoAvatarClientWidth / infoCardClientWidth);
        this.avatarConf.size = infoCardClientWidth * 0.8;
      }
      // this.avatarConf.isShow = clientWidth > this.avatarConf.size + 10;
    },
  },
});
</script>

<style scoped lang="scss">
@import "@/assets/css/common.scss";
.home-affix {
  position: fixed;
  left: 0;
  right: 0;
  z-index: 1;
  & > .affix-left {
    display: flex;
    flex-direction: column;
    align-items: flex-end;
    > .star-card {
      display: flex;
      flex-direction: column;
      justify-content: center;
      align-items: center;
      background: #ffffff;
      border: 0;
      padding: 2px;
      @include custom-shadow;
    }
  }
  & > .affix-main {
    height: 100%;
  }

  & > .affix-right {
    display: flex;
    flex-direction: column;
    align-items: flex-start;
  }
}
.home {
  height: 100%;
  & > .home-main {
    z-index: 2;
    display: flex;
    flex-direction: column;
    height: 100%;
    & > .home-context {
      margin-top: 1em;
      flex: auto;
      display: flex;
      flex-direction: column;
      @include custom-textarea;
    }
    & > .home-carousel {
      @include custom-shadow;
      .el-carousel {
        .el-carousel__item {
          h3 {
            color: #475669;
            opacity: 0.75;
            line-height: 250px;
            margin: 0;
            text-align: center;
          }
          &:nth-child(2n) {
            background-color: #99a9bf;
          }
          &:nth-child(2n + 1) {
            background-color: #d3dce6;
          }
        }
      }
    }
  }
}
</style>
