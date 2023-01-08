<template>
  <div class="logo" ref="logo">
    <!-- <img class="img" :src="url" :width="width" :height="height" /> -->
    LOGO
  </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import logoUrl from "@/assets/imgs/logo/blob-logo.svg";

export default defineComponent({
  name: "LogoView",
  props: {
    width: {
      type: Number,
      default: 55,
    },
    height: {
      type: Number,
      default: 40,
    },
  },
  data() {
    return {
      url: logoUrl,
    };
  },
});
</script>

<style scoped lang="scss">
.logo {
  vertical-align: middle;
  text-align: center;
  .img {
    display: inline-flex;
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
