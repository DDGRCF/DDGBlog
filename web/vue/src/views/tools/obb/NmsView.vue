<template>
  <el-upload
    v-model:file-list="fileList"
    class="upload-demo"
    :action="uploadUrl"
    multiple
    :on-preview="handlePreview"
    :on-remove="handleRemove"
    :before-remove="beforeRemove"
    :limit="10"
    :on-exceed="handleExceed"
  >
    <el-button type="primary">Click to upload</el-button>
    <template #tip>
      <div class="el-upload__tip">
        jpg/png files with a size less than {{ sizeLimits }}.
      </div>
    </template>
  </el-upload>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import { UploadUserFile } from "element-plus";

export default defineComponent({
  name: "ToolsOBBNMS",
  props: {
    sizeLimits: {
      type: String,
      default: "500 MB",
    },
  },
  data() {
    return {
      fileList: [] as UploadUserFile[],
      uploadUrl: "http://localhost:12000/upload",
    };
  },
  methods: {
    handleRemove(file: UploadUserFile, uploadFiles: UploadUserFile[]) {
      console.log(file, uploadFiles);
    },
    handlePreview(uploadFile: UploadUserFile) {
      console.log(uploadFile);
    },
    handleExceed(files: UploadUserFile, uploadFiles: UploadUserFile[]) {
      console.log(files, uploadFiles);
    },
    beforeRemove(uploadFile: UploadUserFile, uploadFiles: UploadUserFile[]) {
      console.log(uploadFile, uploadFiles);
    },
  },
});
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped lang="scss"></style>
