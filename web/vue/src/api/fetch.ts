import axios from "axios";
import storage from "@/utils/storage";
import router from "@/router";
import { ElLoading } from "element-plus";
import { AxiosRequestConfig, AxiosResponse } from "axios";

let loadingInstance = ElLoading.service();
loadingInstance.close();

const fetch = axios.create({
  baseURL: process.env.VUE_APP_SRV,
  timeout: 30000,
});

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
    if (resp.data.code === "400") {
      console.log("error");
    }
    return resp;
  },
  function (err) {
    loadingInstance.close();
    return Promise.reject(err);
  }
);

export default fetch;
