export interface IRuleForm {
    account: string;
    password: string;
}

export interface IRuleList {
    page: number,
    pageSize: number,
    key: string
}

export interface IRuleAddDataForm {
    key: string
    value: string
}