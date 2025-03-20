import axios from "../utils/request";
import type { IRuleAddDataForm, IRuleForm, IRuleList } from "../utils/types";

// 登录
export const login = (data: IRuleForm) => axios.post('/login', data);

// 数据中心列表
export const list = (data: IRuleList) => axios.post('/findAll', data);

// 添加数据
export const addDataApi = (data: IRuleAddDataForm) => axios.post("/addData", data)

// 详情
export const dataInfoApi = (data: IRuleAddDataForm) => axios.post("/dataInfo", data)

// 删除
export const deleteDataApi = (data: IRuleAddDataForm) => axios.post("/deleteData", data)

// 清空
export const clearApi = () => axios.post("/clear")
