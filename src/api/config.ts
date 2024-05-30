import  axios from 'axios';
import  type { AxiosInstance, AxiosRequestConfig, AxiosResponse } from 'axios';
import qs from 'qs';
import router from '../router/index';
import { ElMessage } from 'element-plus';

const consoleKEY = `console`;

const http: AxiosInstance = axios.create({
    baseURL: `/v1/api`,
    // baseURL: "http://localhost:5555/v1/api",
    timeout: 30000,
    method: 'post',
    headers: { 'Content-Type': 'application/json', Authorization: '' },
    params: {},
    paramsSerializer: function (params) {
        return qs.stringify(params);
    },
    withCredentials: false,
    responseType: 'json',
});

//request interceptors
http.interceptors.request.use(config => {
    let token = window.localStorage.getItem("fit-fe");
    config.headers.Authorization = token ? 'Bearer ' + token : '';
    return config;
}, error => {
    return Promise.reject(error)
});

// response interceptors
http.interceptors.response.use(
    (response: AxiosResponse) => {
        let authorization = response.headers['authorization'] || response.headers['Authorization'];
        let split = (authorization || '').split(' ');
        if (split.length === 2) {
            localStorage.setItem('fit-fe', split[1]);
        }
        return Promise.resolve(response);
    },
    (error) => {
        console.log(error.code);
        if (error.code === 'ERR_NETWORK') {
            ElMessage.error('网络错误');
            router.push('/').then(() => {});
            return Promise.reject(error);
        }
        if (error.response && error.response.status === 401) {
            ElMessage.error('未授权或登陆过期');
            router.push('/').then(() => {});
            return Promise.reject(error);
        }
        return Promise.reject(error);
    }
);

export { http, consoleKEY };
