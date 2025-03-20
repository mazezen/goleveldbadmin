<template>
  <div class="data-content">
    <div class="data-content-nav">
      <h3>
        Go Leveldb Admin Web UI 数据中心
        <span>
          <el-button class="lgout" type="info" @click="logout"
            >退出</el-button
          ></span
        >
      </h3>
    </div>

    <div class="data-main-contont">
      <div class="data-main-contont-left">
        <el-dropdown>
          <el-button type="primary">
            Root<el-icon class="el-icon--right"><arrow-down /></el-icon>
          </el-button>
        </el-dropdown>
      </div>

      <div class="data-main-contont-right">
        <div class="data-main-contont-right-header">
          <div class="mt-4">
            <el-input
              v-model="inputSearch"
              style="max-width: 600px"
              placeholder="请输入要搜索的key"
              class="input-with-select"
            >
              <template #append>
                <el-button :icon="Search" @click="searchAction" />
              </template>
            </el-input>
          </div>

          <el-button @click="refresh">刷新</el-button>
          <el-button @click="dialogFormVisible = true" type="success"
            >添加</el-button
          >
          <el-button type="danger" @click="clear">清空</el-button>
        </div>

        <!--数据中心-->
        <el-table :data="tableData" style="width: 100%">
          <el-table-column prop="key" label="Key" width="600" />
          <el-table-column prop="value" label="value" width="600" />
          <el-table-column fixed="right" label="操作" min-width="100">
            <template #default="scope">
              <el-button
                link
                type="danger"
                size="small"
                @click="deleteData(scope.row)"
                >删除</el-button
              >
            </template>
          </el-table-column>
        </el-table>
      </div>
      <!--分页-->
      <div class="pagination">
        <el-pagination
          background
          layout="prev, pager, next"
          :total="total"
          @current-change="handleCurrentChange"
        />
      </div>
    </div>
  </div>

  <!-- 添加弹出框 -->
  <el-dialog v-model="dialogFormVisible" title="添加" width="500">
    <el-form
      ref="ruleAddDataFormRef"
      style="max-width: 600px"
      :model="ruleAddDataForm"
      :rules="rulesAddForm"
      label-width="auto"
      class="demo-ruleForm"
      status-icon
    >
      <el-form-item label="key" prop="key">
        <el-input v-model="ruleAddDataForm.key" placeholder="请输入key" />
      </el-form-item>
      <el-form-item label="value" prop="value">
        <el-input
          v-model="ruleAddDataForm.value"
          type="textarea"
          placeholder="请输入value"
        />
      </el-form-item>
      <el-form-item>
        <el-button type="primary" @click="addData(ruleAddDataFormRef)">
          创建
        </el-button>
        <el-button @click="dialogFormVisible = false">取消</el-button>
      </el-form-item>
    </el-form>
  </el-dialog>
</template>

<script lang="ts" setup>
import router from "../../router/router";
import { list, addDataApi, deleteDataApi, clearApi } from "../../api/api";
import { reactive, ref } from "vue";
import { type FormInstance, type FormRules, ElMessage } from "element-plus";
import type { IRuleAddDataForm, IRuleList } from "../../utils/types";
import { Search } from "@element-plus/icons-vue";

const tableData = ref([]);
const total = ref(0);

// 分页
const handleCurrentChange = async (page: number) => {
  // 处理页码变化的逻辑，可以在这里发起数据请求等操作
  // 这里可以根据页码 page 发起数据请求，更新页面内容等
  // 示例：假设此处需要加载当前页的数据
  ListParam.page = page;
  ListParam.pageSize = 10;
  try {
    const res = await list(ListParam);
    if (res.data.code !== 2000) {
      ElMessage.error(res.data.message);
    } else {
      tableData.value = res.data.data.list;
      total.value = res.data.data.total;
    }
  } catch (error) {
    console.error("Failed to fetch list data:", error);
    ElMessage.error("Failed to fetch list data.");
  }
};

// 列表
const ListParam = reactive<IRuleList>({
  page: 1,
  pageSize: 10,
  key: "",
});

const getList = async () => {
  try {
    const res = await list(ListParam);
    if (res.data.code !== 2000) {
      ElMessage.error(res.data.message);
    } else {
      tableData.value = res.data.data.list;
      total.value = res.data.data.total;
    }
  } catch (error) {
    console.error("Failed to fetch list data:", error);
    ElMessage.error("Failed to fetch list data.");
  }
};
getList();

// 搜索
const inputSearch = ref("");
const searchAction = async () => {
  console.log(inputSearch.value);
  ListParam.key = inputSearch.value;
  try {
    const res = await list(ListParam);
    if (res.data.code !== 2000) {
      ElMessage.error(res.data.message);
    } else {
      tableData.value = res.data.data.list;
      total.value = res.data.data.total;
    }
  } catch (error) {
    console.error("Failed to fetch list data:", error);
    ElMessage.error("Failed to fetch list data.");
  }
};

// 刷新
const refresh = () => {
  ListParam.key = "";
  inputSearch.value = "";
  getList();
};

// 清空
const clear = async () => {
  let res = await clearApi();
  if (res.data.code === 2000) {
    ElMessage.success(res.data.message);
    getList();
  } else {
    ElMessage.error(res.data.message);
    getList();
  }
};

// 添加
const dialogFormVisible = ref(false);
const ruleAddDataFormRef = ref<FormInstance>();
const ruleAddDataForm = reactive<IRuleAddDataForm>({
  key: "",
  value: "",
});
const rulesAddForm = reactive<FormRules<IRuleAddDataForm>>({
  key: [
    { required: true, message: "请输入key", trigger: "blur" },
    { min: 1, max: 255, message: "Key太长了", trigger: "blur" },
  ],
  value: [
    { required: true, message: "请输入value", trigger: "change" },
    { min: 1, max: 10000, message: "Value超出了长度限制", trigger: "blur" },
  ],
});

// 添加数据的异步函数
const addData = async (formEl: FormInstance | undefined) => {
  console.log(formEl);
  if (!formEl) return;
  formEl.validate(async (valid) => {
    if (valid) {
      let res = await addDataApi(ruleAddDataForm);
      console.log(res);
      if (res.data.code === 2000) {
        ElMessage.success(res.data.message);
        getList();
      }
      if (res.data.code === 5000) {
        ElMessage.error(res.data.message);
        getList();
      }
    }
  });
};

// 删除
const deleteData = async (row: IRuleAddDataForm) => {
  // console.log(row);
  let res = await deleteDataApi(row);
  // console.log(res);
  if (res.data.code === 2000) {
    ElMessage.success(res.data.message);
    getList();
  } else {
    ElMessage.error(res.data.message);
    getList();
  }
};

// 退出
const logout = () => {
  localStorage.clear();
  router.push("/login");
};
</script>

<style c scope>
.lgout {
  margin-left: 600px;
}
.data-content {
  width: 100%;
  height: 100%;
  color: black;
}
.data-content .data-content-nav {
  display: grid;
  place-items: center; /* 同时实现水平和垂直居中 */
  height: 50px; /* 容器的高度 */
}

.data-main-contont {
  width: 80%;
  height: 800px;
  margin-left: 10%;
  margin-right: 10%;
  margin-top: 50px;
  border: 2px solid;
  border-color: rgb(233, 232, 232);
  filter: drop-shadow(5px 5px 5px rgb(233, 232, 232));
}
.data-main-contont-left {
  width: 10%;
  height: 600px;
  margin-left: 5%;
  margin-top: 20px;
  margin-right: 5%;
  border: 2px solid;
  border-color: rgb(233, 232, 232);
  filter: drop-shadow(5px 5px 5px rgb(233, 232, 232));
  display: inline-block;
}

.data-main-contont-right {
  width: 75%;
  height: auto;
  margin-top: 20px;
  border: 2px solid;
  border-color: rgb(233, 232, 232);
  filter: drop-shadow(5px 5px 5px rgb(233, 232, 232));
  display: inline-grid;
}
.el-dropdown {
  display: grid;
  place-items: center; /* 同时实现水平和垂直居中 */
  height: 50px; /* 容器的高度 */
}
.pagination {
  margin-top: 40px;
  width: 10%;
  margin-left: auto;
  margin-right: auto;
}
.data-main-contont-right-header {
  width: 100%;
  height: 60px;
  display: flex;
  justify-content: right;
  align-items: center;
}
.mt-4 {
  margin-right: 20px;
}
</style>