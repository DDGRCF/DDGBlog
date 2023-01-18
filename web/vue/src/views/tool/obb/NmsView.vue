<template>
  <el-upload
    ref="uploadCmp"
    v-model:file-list="fileList"
    class="upload-demo"
    :action="getUploadUrl"
    multiple
    :with-credentials="false"
    :on-preview="handlePreview"
    :on-remove="handleRemove"
    :before-remove="beforeRemove"
    :limit="numLimit"
    :on-exceed="handleExceed"
    :on-success="handleSuccess"
    :on-error="handleError"
    :auto-upload="false"
    :drag="true"
    :data="form"
  >
    <el-icon class="el-icon--upload"><upload-filled /></el-icon>
    <div class="el-upload__text">
      将文件拖拽到这里或者<em>点击这里上传文件</em>
    </div>
    <template #tip>
      <div class="el-upload__tip">
        上传zip格式的压缩文件，且文件大小小于 {{ sizeLimitsStr }}.
      </div>
    </template>
  </el-upload>
  <div class="upload-form">
    <el-form
      ref="uploadForm"
      :model="form"
      label-width="auto"
      :scroll-to-error="true"
    >
      <el-form-item class="score-item" label="得分阈值" prop="scoreThr">
        <el-input-number
          :precision="3"
          :step="0.01"
          :min="0"
          :max="1"
          v-model="form.scoreThr"
        />
      </el-form-item>
      <el-form-item class="iou-item" label="IoU阈值" prop="iouThr">
        <el-input-number
          :min="0"
          :max="1"
          :precision="3"
          :step="0.01"
          v-model="form.iouThr"
        />
      </el-form-item>
    </el-form>
  </div>
  <el-button
    ref="uploadBtn"
    :loading="button.isLoading"
    @click="uploadFile"
    :type="button.type"
    :disabled="button.isLoading"
    >{{ button.text }}</el-button
  >
</template>

<script lang="ts">
import { defineComponent, reactive } from "vue";
import {
  ElMessage,
  ElMessageBox,
  UploadUserFile,
  FormInstance,
  UploadInstance,
  ButtonInstance,
  UploadRawFile,
} from "element-plus";
import { SizeLimitType, FormType } from "NmsCmp";
import fetch from "@/api/fetch";
import { UploadAjaxError } from "element-plus/es/components/upload/src/ajax";
import { RespType } from "Fetch";
import { RespCode } from "@/types";

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
      } as SizeLimitType,
      numLimit: 10,
      form: reactive({
        scoreThr: 0.05,
        iouThr: 0.1,
      }) as FormType,
      button: {
        text: "提交文件",
        type: "primary",
        isLoading: false,
      },
    };
  },
  mounted() {
    this.fetchSizeLimit();
  },
  methods: {
    uploadFile() {
      const uploadCmp = this.$refs.uploadCmp as UploadInstance;
      const uploadForm = this.$refs.uploadForm as FormInstance;
      const uploadBtn = this.$refs.uploadBtn as ButtonInstance;
      if (!uploadCmp) return;
      if (!uploadForm) return;
      if (!uploadBtn) return;
      uploadForm.validate((valid) => {
        if (valid) {
          console.log(this.fileList);
          if (!this.fileList.length) {
            ElMessage({
              message: "请至少提交一个文件",
              type: "error",
            });
          } else {
            ElMessage({
              message: "提交成功，请等待服务器处理",
              type: "success",
            });
            setTimeout(() => {
              uploadCmp.submit();
              this.button.isLoading = true;
              this.button.text = "loading...";
            }, 500);
          }
        } else {
          ElMessage({
            message: "提交失败，请检查输入格式是否合理",
            type: "error",
          });
          return false;
        }
      });
    },
    fetchSizeLimit() {
      fetch.get(this.url.fetchSizeLimit);
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
    handleSuccess(response: RespType, rawFile: UploadRawFile) {
      this.button.isLoading = false;
      this.button.text = "提交文件";

      if (response.code === RespCode.success) {
        this.button.type = "success";
        ElMessage({
          message: "操作已经完成",
          type: "success",
        });
        setTimeout(() => {
          this.fileList = [];
        }, 2000);
      } else {
        this.button.type = "danger";
        ElMessage({
          message: response.msg,
          type: "error",
        });
      }
    },
    handleError(err: UploadAjaxError, rawFile: UploadRawFile) {
      this.button.isLoading = false;
      this.button.text = "提交文件";
      this.button.type = "danger";
      console.log(err);
      ElMessage({
        message: "有错误发生，请检查文件",
        type: "error",
      });
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
    getUploadUrl(): string {
      return fetch.getUri() + this.url.upload;
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
  justify-content: center;
  align-items: center;
  & :first-child {
    display: flex;
    $gap: 10px;
    .score-item {
      margin-right: $gap;
    }
    .iou-item {
      margin-left: $gap;
    }
  }
}
</style>
