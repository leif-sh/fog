import axios, { AxiosInstance, AxiosRequestConfig, AxiosResponse } from "axios";
import { ElMessage } from "element-plus";
import {serverHost} from "./settings";

export interface ResponseData {
  code: number;
  data?: any;
  message: string;
}


// 创建 axios 实例
let service: AxiosInstance | any;

service = axios.create({
  baseURL: serverHost +"/api/", // api 的 base_url
  timeout: 30 * 1000 // 请求超时时间 ms
});


// request 拦截器 axios 的一些配置
service.interceptors.request.use(
  (config: AxiosRequestConfig) => {
    return config;
  },
  (error: any) => {
    // Do something with request error
    console.error("error:", error); // for debug
  }
);

// respone 拦截器 axios 的一些配置
service.interceptors.response.use(
  (res: AxiosResponse) => {
    return res.data.data;
  },
  (error: any) => {
    ElMessage({
      message: error.response.data.message,
      type: "error"
    });
    return error.response.data
    // Promise.reject(error)
  }
);

export default service;