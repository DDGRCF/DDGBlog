<template>
  <el-upload
    v-model:file-list="fileList"
    class="upload-demo"
    :action="url.upload"
    multiple
    :on-preview="handlePreview"
    :on-remove="handleRemove"
    :before-remove="beforeRemove"
    :limit="numLimit"
    :on-exceed="handleExceed"
    :auto-upload="false"
    :drag="true"
  >
    <el-icon class="el-icon--upload"><upload-filled /></el-icon>
    <div class="el-upload__text">
      Drop file here or <em>click to upload</em>
    </div>
    <template #tip>
      <div class="el-upload__tip">
        上传zip格式的压缩文件，大小小于 {{ sizeLimitsStr }}.
      </div>
    </template>
  </el-upload>
  <div class="upload-form">
    <el-form :model="form" label-width="auto" :scroll-to-error="true">
      <el-form-item label="得分阈值">
        <el-input @change="printChange" v-model="form.scoreThr" />
      </el-form-item>
      <el-form-item label="IoU阈值">
        <el-input v-model="form.iouThr" />
      </el-form-item>
    </el-form>
  </div>
  <el-button :type="submitType">提交文件</el-button>
</template>

<script lang="ts">
import { defineComponent, reactive } from "vue";
import { ElMessage, ElMessageBox, UploadUserFile } from "element-plus";
import fetch from "@/api/fetch";

interface sizeLimitType {
  size: number;
  unit: string;
}

export default defineComponent({
  name: "ToolObbNms",
  data() {
    return {
      url: {
        upload: "/api/tool/obb/nms/upload",
        fetchSizeLimit: "/api/tool/obb/nms/maxsize",
      },
      fileList: [] as UploadUserFile[],
      sizeLimit: {
        size: 100,
        unit: "MB",
      } as sizeLimitType,
      numLimit: 10,
      form: reactive({
        scoreThr: 0.05,
        iouThr: 0.1,
      }),
      submitType: "primary",
    };
  },
  mounted() {
    this.fetchSizeLimit();
  },
  methods: {
    printChange() {
      console.log(this.form.scoreThr);
    },
    fetchSizeLimit() {
      fetch
        .get(this.url.fetchSizeLimit)
        .then((response) => {
          let responseData = response.data;
          if (responseData.code === 2) {
            this.sizeLimit = responseData.data as sizeLimitType;
          } else {
            ElMessage({
              message: "初始化网页失败! 请联系管理员",
              type: "error",
            });
          }
        })
        .catch(function (reason) {
          ElMessage({
            message: reason,
            type: "error",
          });
        });
    },
    handleRemove(file: UploadUserFile, uploadFiles: UploadUserFile[]) {
      console.log(file, uploadFiles);
    },
    handlePreview(uploadFile: UploadUserFile) {
      console.log(uploadFile);
    },
    handleExceed(files: UploadUserFile[], uploadFiles: UploadUserFile[]) {
      ElMessage.warning(
        `最多上传 ${this.numLimit}个文件，你选择了 ${files.length} 个文件，
        总共有${files.length + uploadFiles.length}个文件`
      );
    },
    beforeRemove(uploadFile: UploadUserFile, uploadFiles: UploadUserFile[]) {
      return ElMessageBox.confirm(
        `真的取消上传 ${uploadFile.name} 文件吗? `
      ).then(
        () => true,
        () => false
      );
    },
  },
  computed: {
    // 这里需要加一个输出参数说明
    sizeLimitsStr(): string {
      return String(this.sizeLimit.size) + " " + this.sizeLimit.unit;
    },
  },
});
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped lang="scss">
.upload-form {
  display: flex;
  flex-direction: column;
  margin: 1em 4em 0 4em;
}
</style>
