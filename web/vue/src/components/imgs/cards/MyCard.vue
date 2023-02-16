<template>
  <el-card
    :style="{ width: String(width) + 'px' }"
    :body-style="{ padding: '0px' }"
    class="card"
    ref="infoCard"
  >
    <div class="card-header">
      <el-image class="avatar" ref="infoAvatar" :src="url" :fit="fit">
        <template #error>
          <div class="image-slot">
            <el-icon><icon-picture /></el-icon>
          </div>
        </template>
      </el-image>
      <div class="name">
        <span class="last">{{ nameData.first }}</span>
        <span class="first">{{ nameData.last }}</span>
      </div>
      <div class="menu">
        <div class="menu-icon">
          <el-icon><MoreFilled /></el-icon>
        </div>
        <ul>
          <li>
            <el-icon><Github fill="#ffffff" /></el-icon>
          </li>
          <li>
            <el-icon><Gmail fill="#ffffff" /></el-icon>
          </li>
          <li>
            <el-icon><QQ fill="#ffffff" /></el-icon>
          </li>
        </ul>
      </div>
    </div>
    <div class="container">
      <div class="left-section">
        <h3>About</h3>
        <p>{{ about }}</p>
        <WaveBtn class="follow-btn" msg="Follow Me" />
      </div>
      <div class="right-section">
        <template v-for="item in strItems" :key="item.num">
          <div class="item" :style="{ order: item.order }">
            <span class="num">{{ item.num }}</span>
            <span class="word">{{ item.word }}</span>
          </div>
        </template>
      </div>
    </div>
    <div class="downdraw">
      <Transition name="slide-fade">
        <div class="downdraw-content" v-if="downDraw.isExpand">
          <div class="main">
            <h3>Detail</h3>
            <p>
              Focus on deep learning applications and learn various programming
              knowledge
            </p>
          </div>
        </div>
      </Transition>
      <div class="downdraw-icon" @click="clickExpand">
        <el-icon v-if="!downDraw.isExpand" class="downarrow"
          ><ArrowDown
        /></el-icon>
        <el-icon v-else class="uparrow"><ArrowUp /></el-icon>
      </div>
    </div>
  </el-card>
</template>

<script lang="ts">
import { CardItemType } from "Common";
import { defineComponent, PropType } from "vue";
import Gmail from "../icons/Gmail.vue";
import Github from "../icons/Github.vue";
import QQ from "../icons/QQ.vue";
import WaveBtn from "../../btns/WaveBtn.vue";

interface StrItemType {
  num: string;
  word: string;
  order: number;
}

export default defineComponent({
  name: "MyCard",
  props: {
    url: String,
    fit: {
      type: String,
      default: "contain",
    },
    name: {
      type: String,
      default: "",
    },
    width: {
      type: Number,
      default: 300,
    },
    about: {
      type: String,
      default: "",
    },
    items: {
      type: Object as PropType<CardItemType[]>,
    },
  },
  components: {
    QQ,
    Github,
    Gmail,
    WaveBtn,
  },
  data() {
    return {
      nameData: {
        first: this.name.split(" ")[1],
        last: this.name.split(" ")[0],
      },
      downDraw: {
        isExpand: false,
      },
    };
  },
  methods: {
    clickExpand() {
      this.downDraw.isExpand = !this.downDraw.isExpand;
    },
  },
  computed: {
    strItems(): StrItemType[] {
      let strItems = [] as StrItemType[];
      this.items?.forEach((value: CardItemType, index: number) => {
        let item = {} as StrItemType;
        item.word = value.word;
        if (Math.floor(value.num / 1000) > 0) {
          item.num = String(Math.floor(value.num / 1000)) + "K";
        } else if (Math.floor(value.num / 100000) > 0) {
          item.num = String(Math.floor(value.num / 1000)) + "M";
        } else {
          item.num = String(value.num);
        }
        item.order = item.num.length;
        strItems[index] = item;
      });
      return strItems;
    },
  },
});
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped lang="scss">
@import "@/assets/css/common.scss";
$deep-blue: #2980b9;
$light-blue: #3498db;
.card {
  position: relative;
  padding: 0;
  margin: 0;
  border: 0px;
  background: $textarea-color;
  * {
    margin: 0;
    padding: 0;
  }
  .card-header {
    height: 100%;
    position: relative;
    .name {
      position: absolute;
      bottom: 20px;
      left: 20px;
      color: white;
      font-family: "Century Gothic";
      text-align: left;
      .first {
        display: block;
        font-size: 32px;
      }
      .last {
        display: block;
        font-size: 32px;
      }
    }
    .avatar {
      padding: 0;
      margin: 0;
      position: relative;
      width: 100%;
      font-size: 0;
      overflow: hidden;
      &::after {
        content: "";
        display: block;
        position: absolute;
        top: 0;
        left: 0;
        width: 100%;
        height: 100%;
        background-color: $deep-blue;
        opacity: 55%;
      }
    }

    .menu {
      position: absolute;
      top: 10px;
      right: 5px;
      display: flex;
      flex-direction: row-reverse;
      align-items: center;
      color: white;
      opacity: 0.6;
      cursor: pointer;
      & > .menu-icon {
        transform: rotate(90deg);
        flex: 0 0 auto;
        font-size: 22px;
        padding: 5px 0;
      }
      & > ul {
        display: none;
        background-color: #33333355;
        & > li {
          display: flex;
          height: 28px;
          align-items: center;
          list-style-type: none;
          padding: 0 5px;
          &:hover {
            background-color: #333333b3;
            transition: background-color 0.2s ease-in-out;
          }
        }
      }
      &:hover > ul {
        display: flex;
        flex-direction: row;
        flex: auto;
        align-items: center;
      }
    }
  }
  .container {
    padding: 20px 20px 0px 20px;
    overflow: hidden;
    display: flex;
    position: relative;
    background: $textarea-color;
    z-index: 1;

    .left-section {
      display: flex;
      flex: 1 1 auto;
      flex-direction: column;
      overflow: hidden;
      align-items: flex-start;
      text-align: start;
      justify-content: space-between;
      h3 {
        margin-top: 5px;
        color: $light-blue;
        margin-bottom: 0px;
      }
      p {
        font-family: "Century Gothic";
        font-size: 20px;
        margin: 10px 0;
      }
      .follow-btn {
        width: 100px;
        border-radius: 12px;
        color: white;
        text-align: center;
        display: block;
        padding: 5px;
        margin-bottom: 10px;
      }
    }

    .right-section {
      display: flex;
      flex-direction: column;
      align-items: flex-end;
      flex: 0 1 auto;
      margin-left: 10px;
      .item {
        margin: 6px 0;
        text-align: right;
        .num {
          display: block;
          font-size: 28px;
        }
        .word {
          font-size: 14px;
        }
      }
    }
  }

  .downdraw {
    position: relative;
    display: flex;
    flex-direction: column;
    align-items: center;
    width: 100%;

    @keyframes bound {
      from {
        transform: translateY(-5%);
      }
      50% {
        transform: translateY(25%);
      }
      to {
        transform: translateY(-5%);
      }
    }
    .downdraw-icon {
      display: flex;
      position: relative;
      justify-content: center;
      align-items: center;
      height: 20px;
      width: 100%;
      flex: 0 1 auto;
      padding: 0 0 2px 0;
      margin: 0;
      &:hover {
        background-color: #353b4833;
      }
      .downarrow {
        animation: bound 1s ease-in-out infinite;
      }
      .uparrow {
      }
    }
    .slide-fade-enter-active {
      transition: all 0.3s ease-out;
    }

    .slide-fade-leave-active {
      transition: all 0.3s ease-in-out;
    }

    .slide-fade-enter-from,
    .slide-fade-leave-to {
      visibility: hidden;
      transform: translateY(-100%);
    }
    .downdraw-content {
      display: flex;
      flex-direction: column;
      align-items: flex-start;
      width: 100%;
      padding: 20px;
      box-sizing: border-box;
      box-shadow: #bbb7aa18 0px 30px 50px -12px inset,
        #948e7a52 0px 18px 26px -18px inset;
      background-color: $textarea-color;
      z-index: 0;
      .main {
        h3 {
          color: $light-blue;
          margin: 5px 0;
          text-align: left;
        }
        p {
          text-align: left;
          font-size: 20px;
          font-family: "Century Gothic";
        }
      }
    }
  }
}
</style>
