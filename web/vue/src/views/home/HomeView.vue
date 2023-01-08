<template>
  <div class="home">
    <el-row :gutter="15">
      <el-col class="left-span-0" :lg="5" :md="3" :sm="2"></el-col>
      <el-col class="left-span-1 hidden-md-and-down" :lg="1">
        <el-card class="card" shadow="hover">
          <el-icon><Star /></el-icon>
          <p />
          <el-icon><Opportunity /></el-icon>
        </el-card>
      </el-col>
      <el-col class="main-span" :lg="12" :md="18" :sm="20">
        <div class="block text-center">
          <el-carousel height="250px">
            <el-carousel-item v-for="item in 4" :key="item">
              <h3 class="small justify-center" text="2xl">{{ item }}</h3>
            </el-carousel-item>
          </el-carousel>
        </div>
        <button @click="printEnv">This is test</button>
      </el-col>
      <el-col class="right-span-1 hidden-md-and-down" :lg="4">
        <div>
          <el-card :body-style="{ padding: '10px' }">
            <el-avatar
              class="avatar"
              :size="avatarConf.size"
              :src="avatarConf.avatarUrl"
            />
            <div style="padding: 14px">
              <span>DDGRCF</span>
              <div class="bottom">
                <!-- <time class="time">{{ currentDate }}</time> -->
                <el-button text class="button">Follow Me</el-button>
              </div>
            </div>
          </el-card>
        </div>
        <p />
      </el-col>
      <el-col class="right-span-0" :lg="2" :md="3" :sm="2"></el-col>
    </el-row>
  </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";

export default defineComponent({
  name: "HomeView",
  data() {
    return {
      avatarConf: {
        avatarUrl: require("../../assets/imgs/avatar/avatar-0.jpg"),
        size: 225,
      },
    };
  },
  methods: {
    printEnv() {
      console.log(process.env.VUE_APP_SRV);
    },
  },
});
</script>

<style scoped lang="scss">
.home {
  display: flex;
  flex-direction: column;
  .avatar {
    margin-top: 20px;
    background-color: white;
  }
  .left-span-1 {
    display: flex;
    flex-direction: column;
    margin: 0;
    padding: 0;
    .affix {
      .card {
        display: flex;
        flex-direction: column;
        align-items: center;
        background-color: #eeeeee;
      }
    }
  }
  .main-span {
    .block.text-center {
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
