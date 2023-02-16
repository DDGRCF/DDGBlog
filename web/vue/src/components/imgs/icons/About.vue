<template>
  <i>
    <svg
      t="1673855744504"
      class="icon"
      viewBox="0 0 1024 1024"
      version="1.1"
      xmlns="http://www.w3.org/2000/svg"
      p-id="1287"
      width="240"
      height="240"
    >
      <path
        d="M512 0C229.376 0 0 229.376 0 512s229.376 512 512 512 512-229.376 512-512S794.624 0 512 0z m62.464 774.144c0 27.648-22.528 50.176-50.176 50.176h-25.088c-27.648 0-50.176-22.528-50.176-50.176V474.624c0-27.648 22.528-50.176 50.176-50.176h25.088c27.648 0 50.176 22.528 50.176 50.176v299.52zM512 349.696c-34.304 0-62.464-28.16-62.464-62.464 0-34.304 28.16-62.464 62.464-62.464s62.464 28.16 62.464 62.464c0 34.304-28.16 62.464-62.464 62.464z"
        :fill="fill"
        p-id="1288"
      ></path>
    </svg>
  </i>
</template>

<script lang="ts">
import { defineComponent } from "vue";

export default defineComponent({
  name: "AboutView",
  props: {
    fill: {
      type: String,
      default: "#000000",
    },
  },
});
</script>
