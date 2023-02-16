<template>
  <div class="hello">{{ msg }}</div>
  <router-view></router-view>
</template>

<script lang="ts">
import { defineComponent } from "vue";

export default defineComponent({
  name: "AdmView",
  props: {
    msg: {
      type: String,
      default: "this is administrator",
    },
  },
});
</script>

<style scoped lang="scss"></style>
