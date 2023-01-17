<template>
  <el-row class="home-affix" :gutter="colsNum.gutter">
    <el-col
      :lg="colsNum.left.lg - 1"
      :md="colsNum.left.md"
      :sm="colsNum.left.sm"
    ></el-col>
    <el-col class="affix-left hidden-md-and-down" :lg="1">
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
    ></el-col>
    <el-col class="affix-right hidden-md-and-down" :lg="4">
      <div v-if="avatarConf.isShow" ref="infoCard">
        <el-card :body-style="{ padding: '5px' }">
          <el-avatar
            class="avatar"
            ref="infoAvatar"
            :size="avatarConf.size"
            :src="avatarConf.avatarUrl"
          />
          <div style="padding: 10px; margin-top: 10px">
            <span> {{ avatarConf.name }}</span>
            <div class="bottom">
              <!-- <time class="time">{{ currentDate }}</time> -->
              <el-button text class="button">Follow Me</el-button>
            </div>
          </div>
        </el-card>
      </div>
      <p />
    </el-col>
    <el-col
      :lg="colsNum.right.lg - 4"
      :md="colsNum.right.md"
      :sm="colsNum.right.sm"
    ></el-col>
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
        <el-button type="primary" @click="printEnv">This is test</el-button>
        <div v-for="item in 20" :key="item">{{ item }}</div>
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
import menuObj from "@/conf/menu";
import { ElAvatar } from "element-plus";
import { Background } from "tsparticles-engine";

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
        avatarUrl: require("../../assets/imgs/avatar/avatar-0.jpg"),
        size: 225,
        isShow: true,
        name: "DDGRCF",
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
    printEnv() {
      console.log(menuObj.getItemName("0-1"));
    },
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
      background-color: #eeeeee;
    }
  }
  & > .affix-main {
    height: 100%;
  }

  & > .affix-right {
    & > :hover {
      box-shadow: 0 0 5px #dcdde1;
    }
    .avatar {
      margin-top: 20px;
    }
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
