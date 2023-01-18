import axios from "axios";
import { ElLoading, ElMessage } from "element-plus";
import { AxiosRequestConfig, AxiosResponse } from "axios";
import { RespType } from "Fetch";
import { RespCode } from "@/types";

let loadingInstance = ElLoading.service();
loadingInstance.close();
const baseConfig = {
  baseURL: process.env.VUE_APP_SRV,
  timeout: 30000,
};

const fetch = axios.create(baseConfig);

// const fetch_clear = axios.create(baseConfig);

// fetch_clear.interceptors.request.use(
//   function (config: AxiosRequestConfig) {
//     return config;
//   },
//   function (err) {
//     return Promise.reject(err);
//   }
// );

// fetch_clear.interceptors.response.use(
//   function (resp: AxiosResponse) {
//     console.log("recv: ", resp.data);
//     return resp;
//   },
//   function (err) {
//     return Promise.reject(err);
//   }
// );

// Promise.reject 和 Promise.resolve 是传输失败和成功的信息给catch或者then
// catch(function (reason) {}) 或者 then(function (result) {})
fetch.interceptors.request.use(
  function (config: AxiosRequestConfig) {
    loadingInstance = ElLoading.service();
    return config;
  },
  function (err) {
    loadingInstance.close();
    return Promise.reject(err);
  }
);

fetch.interceptors.response.use(
  function (resp: AxiosResponse) {
    loadingInstance.close();
    console.log("recv: ", resp.data);
    const data = resp.data as RespType;
    if (data.code === RespCode.failure) {
      ElMessage({
        message: "error: " + (resp.data as RespType).msg,
        type: "error",
      });
    }
    return resp;
  },
  function (err) {
    loadingInstance.close();
    ElMessage({
      message: "error: " + err,
      type: "error",
    });
    return Promise.reject(err);
  }
);

// export { fetch_clear };
export default fetch;
