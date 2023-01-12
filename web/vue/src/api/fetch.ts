import axios from "axios";
import storage from "@/utils/storage";
import router from "@/router";
// import { ElLoading } from "element-plus";
import { AxiosRequestConfig, AxiosResponse } from "axios";

const fetch = axios.create({
  baseURL: process.env.VUE_APP_SRV,
  timeout: 30000,
});

fetch.interceptors.request.use(
  function (config: AxiosRequestConfig) {
    // const loadingInstance = ElLoading.service({ fullscreen: false });
    return config;
  },
  function (err) {
    return Promise.reject(err);
  }
);

fetch.interceptors.response.use(
  function (resp: AxiosResponse) {
    console.log("recv");
    if (resp.data.code === "400") {
      console.log("error");
    }
    return resp;
  },
  function (err) {
    return Promise.reject(err);
  }
);

export default fetch;
