<template>
  <div>
    <div class="search-term">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
        <el-form-item label="路径">
          <el-input placeholder="路径" v-model="searchInfo.path"></el-input>
        </el-form-item>
        <el-form-item label="描述">
          <el-input placeholder="描述" v-model="searchInfo.description"></el-input>
        </el-form-item>
        <el-form-item label="api组">
          <el-input placeholder="api组" v-model="searchInfo.apiGroup"></el-input>
        </el-form-item>
        <el-form-item label="请求">
          <el-select clearable placeholder="请选择" v-model="searchInfo.method">
            <el-option
              :key="item.value"
              :label="`${item.label}(${item.value})`"
              :value="item.value"
              v-for="item in methodOptions"
            ></el-option>
          </el-select>
        </el-form-item>
        <el-form-item>
          <el-button @click="onSubmit" type="primary">查询</el-button>
        </el-form-item>
        <el-form-item>
          <el-button @click="openDialog('addApi')" type="primary">新增api</el-button>
        </el-form-item>
      </el-form>
    </div>
    <el-table :data="tableData" @sort-change="sortChange" border stripe>
      <el-table-column label="id" min-width="60" prop="ID" sortable="custom"></el-table-column>
      <el-table-column label="论文名称" min-width="150" prop="title" sortable="custom"></el-table-column>
       

      <el-table-column fixed="right" label="操作" width="200">
        <template slot-scope="scope">
          <el-button @click="editApi(scope.row)" size="small" type="primary" icon="el-icon-edit">编辑</el-button>
          <el-button
            @click="deleteApi(scope.row)"
            size="small"
            type="danger"
            icon="el-icon-delete"
          >删除</el-button>
        </template>
      </el-table-column>
    </el-table>
    <el-pagination
      :current-page="page"
      :page-size="pageSize"
      :page-sizes="[10, 30, 50, 100]"
      :style="{float:'right',padding:'20px'}"
      :total="total"
      @current-change="handleCurrentChange"
      @size-change="handleSizeChange"
      layout="total, sizes, prev, pager, next, jumper"
    ></el-pagination>

    <el-dialog :before-close="closeDialog" :title="dialogTitle" :visible.sync="dialogFormVisible">
      <el-form :inline="true" :model="form" :rules="rules" label-width="80px" ref="apiForm">
        <el-form-item label="论文名称" prop="path">
          <el-input autocomplete="off" v-model="form.title"></el-input>
        </el-form-item>
         
      </el-form>
      <div class="warning">新增Api需要在角色管理内配置权限才可使用</div>
      <div class="dialog-footer" slot="footer">
        <el-button @click="closeDialog">取 消</el-button>
        <el-button @click="enterDialog" type="primary">确 定</el-button>
      </div>
    </el-dialog>
  </div>
</template>


<script>
// 获取列表内容封装在mixins内部  getTableData方法 初始化已封装完成 条件搜索时候 请把条件安好后台定制的结构体字段 放到 this.searchInfo 中即可实现条件搜索

import {
  getPaperList,
  create,
  update,
  deleteItem,
  getInfo
} from "@/api/modules/research/paper";
import infoList from "@/mixins/infoList";
import { toSQLLine } from "@/utils/stringFun";
const methodOptions = [
  {
    value: "POST",
    label: "创建",
    type: "success"
  },
  {
    value: "GET",
    label: "查看",
    type: ""
  },
  {
    value: "PUT",
    label: "更新",
    type: "warning"
  },
  {
    value: "DELETE",
    label: "删除",
    type: "danger"
  }
];

export default {
  name: "Api",
  mixins: [infoList],
  data() {
    return {
      listApi: getPaperList,
      dialogFormVisible: false,
      dialogTitle: "新增Api",
      form: {
        title: "",
      },
      methodOptions: methodOptions,
      type: "",
      rules: {
        title: [{ required: true, message: "请输入论文名称", trigger: "blur" }],
     
      }
    };
  },
  methods: {
    // 排序
    sortChange({ prop, order }) {
      if (prop) {
        this.searchInfo.orderKey = toSQLLine(prop);
        this.searchInfo.desc = order == "descending";
      }
      this.getTableData();
    },
    //条件搜索前端看此方法
    onSubmit() {
      this.page = 1;
      this.pageSize = 10;
      this.getTableData();
    },
    initForm() {
      this.$refs.apiForm.resetFields();
      this.form = {
        path: "",
        apiGroup: "",
        method: "",
        description: ""
      };
    },
    closeDialog() {
      this.initForm();
      this.dialogFormVisible = false;
    },
    openDialog(type) {
      switch (type) {
        case "addApi":
          this.dialogTitlethis = "新增";
          break;
        case "edit":
          this.dialogTitlethis = "编辑";
          break;
        default:
          break;
      }
      this.type = type;
      this.dialogFormVisible = true;
    },
    async editApi(row) {
      const res = await getInfo({ id: row.ID });
      this.form = res.data.api;
      this.openDialog("edit");
    },
    async deleteApi(row) {
      this.$confirm("此操作将永久删除所有角色下该api, 是否继续?", "提示", {
        confirmButtonText: "确定",
        cancelButtonText: "取消",
        type: "warning"
      })
        .then(async () => {
          const res = await deleteInfo(row);
          if (res.code == 0) {
            this.$message({
              type: "success",
              message: "删除成功!"
            });
            if (this.tableData.length == 1) {
              this.page--;
            }
            this.getTableData();
          }
        })
        .catch(() => {
          this.$message({
            type: "info",
            message: "已取消删除"
          });
        });
    },
    async enterDialog() {
      this.$refs.apiForm.validate(async valid => {
        if (valid) {
          switch (this.type) {
            case "addApi":
              {
                const res = await create(this.form);
                if (res.code == 0) {
                  this.$message({
                    type: "success",
                    message: "添加成功",
                    showClose: true
                  });
                }
                this.getTableData();
                this.closeDialog();
              }

              break;
            case "edit":
              {
                const res = await update(this.form);
                if (res.code == 0) {
                  this.$message({
                    type: "success",
                    message: "编辑成功",
                    showClose: true
                  });
                }
                this.getTableData();
                this.closeDialog();
              }
              break;
            default:
              {
                this.$message({
                  type: "error",
                  message: "未知操作",
                  showClose: true
                });
              }
              break;
          }
        }
      });
    }
  },
  filters: {
    methodFiletr(value) {
      const target = methodOptions.filter(item => item.value === value)[0];
      // return target && `${target.label}(${target.value})`
      return target && `${target.label}`;
    },
    tagTypeFiletr(value) {
      const target = methodOptions.filter(item => item.value === value)[0];
      return target && `${target.type}`;
    }
  },
  created() {
    this.getTableData();
  }
};
</script>
<style scoped lang="scss">
.button-box {
  padding: 10px 20px;
  .el-button {
    float: right;
  }
}
.el-tag--mini {
  margin-left: 5px;
}
.warning {
  color: #dc143c;
}
</style>