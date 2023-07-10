import axios, { AxiosInstance, AxiosRequestConfig, AxiosResponse } from "axios";
import { ElMessage } from "element-plus";

export interface ResponseData {
  code: number;
  data?: any;
  message: string;
}


// console.log('import.meta.env: ', import.meta.env);

const serverHost: string = import.meta.env.MODE === "development"?"http://localhost:3001" : "http://localhost:3001"
// 创建 axios 实例
let service: AxiosInstance | any;

service = axios.create({
  baseURL: serverHost +"/api/", // api 的 base_url
  timeout: 50000 // 请求超时时间
});


// request 拦截器 axios 的一些配置
service.interceptors.request.use(
  (config: AxiosRequestConfig) => {
    return config;
  },
  (error: any) => {
    // Do something with request error
    console.error("error:", error); // for debug
    Promise.reject(error);
  }
);

// respone 拦截器 axios 的一些配置
service.interceptors.response.use(
  (res: AxiosResponse) => {
    // Some example codes here:
    // code == 0: success
    const data: ResponseData = res.data
    if (res.status === 200) {
        return data.data;
    } else {
      ElMessage({
        message: "网络错误!",
        type: "error"
      });
      return Promise.reject(new Error(data.message || "Error"));
    }
  },
  (error: any) => Promise.reject(error)
);

export default service;