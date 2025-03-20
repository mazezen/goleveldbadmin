<template>
  <div>
    <div class="login">
      <div class="login-container">
        <div class="login-main">
          <div class="login-container">
            <el-form
              ref="ruleFormRef"
              :model="ruleForm"
              :rules="rules"
              label-width="120px"
              class="demo-ruleForm"
            >
              <el-form-item label="用户名 " prop="account">
                <el-input
                  v-model="ruleForm.account"
                  placeholder="请输入用户名"
                />
              </el-form-item>

              <el-form-item label="密码 " prop="password">
                <el-input
                  type="password"
                  v-model="ruleForm.password"
                  placeholder="请输入密码"
                />
              </el-form-item>

              <el-form-item>
                <el-button type="primary" @click="submitForm(ruleFormRef)"
                  >登录</el-button
                >
                <el-button @click="resetForm(ruleFormRef)">取消登录</el-button>
              </el-form-item>
            </el-form>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
  
<script lang="ts" setup>
import router from "../../router/router";
import { ref, reactive } from "vue";
import { type FormInstance, type FormRules, ElMessage } from "element-plus";
import type { IRuleForm } from "../../utils/types";
import { login } from "../../api/api";

// 表单信息
const ruleFormRef = ref<FormInstance>();
const ruleForm = reactive<IRuleForm>({
  account: "",
  password: "",
});

// 校验
const rules = reactive<FormRules<IRuleForm>>({
  account: [
    { required: true, message: "请输入账号", trigger: "blur" },
    { min: 1, max: 30, message: "账号太长了", trigger: "blur" },
  ],
  password: [
    { required: true, message: "请输入密码", trigger: "change" },
    { min: 6, max: 18, message: "密码不正确", trigger: "blur" },
  ],
});

// 按钮
const submitForm = (formEl: FormInstance | undefined) => {
  if (!formEl) return;

  formEl.validate(async (valid: boolean) => {
    if (valid) {
      try {
        let res = await login(ruleForm); // 假设 ruleForm 已经定义
        console.log(res);
        if (res.data.code === 5000) {
          ElMessage.error(res.data.message);
        }
        if (res.data.code === 2000) {
          ElMessage.success(res.data.message);
          localStorage.setItem("token", res.data.data);
          router.push("/home");
        }
      } catch (error) {
        console.error("登录失败:", error);
      }
    } else {
      console.log("账号和密码错误!");
    }
  });
};

const resetForm = (formEl: FormInstance | undefined) => {
  if (!formEl) return;
  formEl.resetFields();
};
</script>



<style scoped>
.login-main {
  text-align: center; /*让div内部文字居中*/
  background-color: #fff;
  border-radius: 20px;
  width: 300px;
  height: 350px;
  margin: auto;
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
}
</style>