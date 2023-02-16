<template>
  <div class="submit">{{ msg }}</div>
  test
</template>

<script lang="ts">
import { defineComponent } from "vue";

export default defineComponent({
  name: "SubmitView",
  props: {
    msg: {
      type: String,
      default: "this is submit view",
    },
  },
});
</script>

<style scoped lang="scss"></style>
