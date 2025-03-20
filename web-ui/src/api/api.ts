import axios from "../utils/request";
import type { IRuleAddDataForm, IRuleForm, IRuleList } from "../utils/types";

// 登录
export const login = (data: IRuleForm) => {
    return axios.post('/login', data, {
        headers: {
            "Content-Type": "application/json",
        }
    });
};

// 数据中心列表
export const list = (data: IRuleList) => {
    return axios.post('/findAll', data, {
        headers: {
            "Content-Type": "application/json"
        }
    })
};

// 添加数据
export const addDataApi = (data: IRuleAddDataForm) => {
    return axios.post("/addData", data, {
        headers: {
            "Content-Type": "application/json",
        }
    });
};

// 详情
export const dataInfoApi = (data: IRuleAddDataForm) => axios.post("/dataInfo", data)

// 删除
export const deleteDataApi = (data: IRuleAddDataForm) => axios.post("/deleteData", data)

// 清空
export const clearApi = () => axios.post("/clear")
