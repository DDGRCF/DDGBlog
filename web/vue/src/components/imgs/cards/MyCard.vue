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
      <div class="cover"></div>
      <div class="name">
        <span class="last">{{ nameData.first }}</span>
        <span class="first">{{ nameData.last }}</span>
      </div>
    </div>
    <div class="container">
      <div class="left-section">
        <h3>About</h3>
        <p>{{ about }}</p>
        <WaveBtn class="follow-btn" msg="Follow-Me" />
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
  </el-card>
</template>

<script lang="ts">
import { CardItemType } from "Common";
import { defineComponent, PropType } from "vue";
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
    WaveBtn,
  },
  data() {
    return {
      nameData: {
        first: this.name.split(" ")[1],
        last: this.name.split(" ")[0],
      },
    };
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
$deep-blue: #2980b9;
$light-blue: #3498db;
.card {
  position: relative;
  padding: 0;
  margin: 0;
  background: white;
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
    }
    .cover {
      position: absolute;
      top: 0;
      left: 0;
      right: 0;
      height: calc(100% - 3px);
      background: $deep-blue;
      opacity: 55%;
    }
  }
  .container {
    padding: 20px;
    overflow: hidden;
    display: flex;
    position: relative;

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
        // background: $light-blue;
        border-radius: 12px;
        color: white;
        text-align: center;
        display: block;
        padding: 5px;
        // & :hover {
        //   background: $deep-blue;
        // }
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
}
</style>
