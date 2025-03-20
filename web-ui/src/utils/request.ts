import axios, { type AxiosRequestConfig, type AxiosResponse } from "axios";


// 路径
axios.defaults.baseURL = `http://localhost:9091/lb`;

// 请求拦截器
axios.interceptors.request.use((config: AxiosRequestConfig | any) => {
    // console.log("adad" + localStorage.getItem('token'));
    let token = localStorage.getItem('token');
    (config: AxiosRequestConfig) => {
        if (localStorage.getItem('token')) {
            if (!config.headers) {
                config.headers = {}; // 如果 headers 未定义，则初始化为一个空对象
            }
            config.headers.Authorization = `Bearer ${token}`
        }
    }
    return config
},
    (error) => {
        return Promise.reject(error);
    }
)

// 响应拦截器
axios.interceptors.response.use((res: AxiosResponse) => {
    return res
}, (err: any) => {
    return Promise.reject(err)
});

export default axios;

