<template>
  <div class="logo">
    <embed class="embed" :src="urlD" v-if="url.length > 0" /><Picture v-else />
  </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";

export default defineComponent({
  name: "LogoView",
  props: {
    url: {
      type: String,
      default: "",
    },
  },
  data() {
    return {
      urlD: this.url,
    };
  },
});
</script>

<style scoped lang="scss">
.logo {
  display: inline-flex;
  height: 100%;
  flex-direction: column;
  justify-content: center;
}
</style>
