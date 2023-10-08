import axios, { AxiosRequestConfig } from "axios";


import { TIMEOUT, BASE_URL } from "./config";

import { message } from 'ant-design-vue';
axios.defaults.timeout = TIMEOUT;
axios.defaults.baseURL = BASE_URL;

/**
 * http request 拦截器
 */


axios.interceptors.request.use(
    (config: any) => {
        const token: string | null = localStorage.getItem("token");
        if (token) {
            config.headers = {
                "authorization": `Bearer ${token}`,
            };
        }

        return config;
    },
    (error: any) => {
        return Promise.reject(error);
    }
);

/**
 * http response 拦截器
 */
axios.interceptors.response.use(
    (response: any) => {
        return response
    },
    (error: string) => {
        console.log("请求出错：", error);
    }
);
/**
 * 封装请求
 * @param url 请求地址
 * @param param 请求参数
 * @param fecth 请求方法
 * @returns {Promise}
 */

//统一接口处理，返回数据
const requests: Function = (url: string, param: AxiosRequestConfig<string>, fecth = "post"): Promise<any> => {
    return new Promise((resolve, reject) => {
        switch (fecth) {
            case "get":
                axios.get(url, { params: param }).then((response: any) => {

                    if (response.data.code === 200) {
                        let { data } = response;
                        message.success(data.msg);
                        resolve(data);
                    } else { Landing(response?.data); }
                }).catch((error: any) => {
                    msag(error)
                    reject(error);
                });
                break;
            case "post":
                axios.post(url, param).then(
                    (response: any) => {
                        console.log(response)
                        if (response.data.code === 200) {
                            let { data } = response;
                            message.success(data.msg);
                            resolve(data);
                        } else {
                            Landing(response.data);
                        }
                    },
                    (error: any) => {
                        msag(error)
                        reject(error);
                    }
                );
                break;
            default:
                break;
        }
    });
}

//失败提示
function msag(err: { response: { status: number; data: { error: { details: string; }; }; }; }) {
    if (err && err.response) {
        switch (err.response.status) {
            case 400:
                alert(err.response.data.error.details);
                break;
            case 401:
                alert("未授权，请登录");
                break;
            case 403:
                alert("拒绝访问");
                break;
            case 404:
                alert("请求地址出错");
                break;

            case 408:
                alert("请求超时");
                break;

            case 500:
                alert("服务器内部错误");
                break;

            case 501:
                alert("服务未实现");
                break;

            case 502:
                alert("网关错误");
                break;

            case 503:
                alert("服务不可用");
                break;

            case 504:
                alert("网关超时");
                break;

            case 505:
                alert("HTTP版本不受支持");
                break;
            default:
        }
    }
}

/**
 * 查看返回的数据
 * @param data
 */
function Landing(data: { code: number, status: number | string; msg: string }) {
    console.log(data)
    message.error(data.msg);
}


export default requests

