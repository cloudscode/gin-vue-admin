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
      <el-form :inline="true" :model="formData" :rules="rules"  size="medium" label-width="110px" ref="formData">
         <el-row >
            <el-form-item label="论文名称" prop="path">
          <el-input autocomplete="off" v-model="formData.title"></el-input>
        </el-form-item>
         
         
        </el-row>      
  <el-row>
             <el-form-item label="刊物类别" prop="path">
            <el-radio-group v-model="formData.publicationType" size="medium">
              <el-radio v-for="(item, index) in publicationType" :key="index" :label="item.value"
                :disabled="item.disabled">{{item.label}}</el-radio>
            </el-radio-group>
          </el-form-item>          

  </el-row>
             <el-row>
                <el-col >
                     <el-form-item label="论文类型" prop="kind">
            <el-radio-group v-model="formData.kind" size="medium">
              <el-radio v-for="(item, index) in kind" :key="index" :label="item.value"
                :disabled="item.disabled">{{item.label}}</el-radio>
            </el-radio-group>
          </el-form-item>      
                </el-col>
                <el-col >
                       <el-form-item label="刊物名称" prop="publicationName">
                <el-input v-model="formData.publicationName" placeholder="请输入刊物名称" clearable
                  :style="{width: '100%'}"></el-input>
              </el-form-item>
                </el-col>
                 
            </el-row> 
       <el-row>
                <el-col >
                       <el-form-item label="期号" prop="period">
                <el-input v-model="formData.period" placeholder="期号" clearable
                 ></el-input>
              </el-form-item>
                </el-col>
                <el-col >
                      <el-form-item label="卷号" prop="volumeNo">
                <el-input v-model="formData.volumeNo" placeholder="卷号" clearable
                 ></el-input>
              </el-form-item>
                </el-col>
                 
            </el-row> 
            <el-row>
                <el-col >
                  <el-form-item label="学科门类" prop="knowledgeClass">
                    <el-select v-model="formData.knowledgeClass" placeholder="请输入学科门类" clearable
                      :style="{width: '100%'}">
                      <el-option v-for="(item, index) in knowledgeClass" :key="index" :label="item.label"
                        :value="item.value" :disabled="item.disabled"></el-option>
                    </el-select>
                  </el-form-item>
                </el-col>
                <el-col >                     
             <el-form-item label="一级学科:">
                <el-input v-model="formData.firstKnowledge" clearable placeholder="请输入一级学科" ></el-input>
          </el-form-item>
                </el-col>
                 
            </el-row> 
           
          <el-row>
                <el-col >
                          <el-form-item label="成果来源:">
                <el-input v-model="formData.source" clearable placeholder="请输入成果来源" ></el-input>
          </el-form-item>
                </el-col>
                <el-col >
                    
                </el-col>
                     <el-form-item label="发表范围:">
                <el-input v-model="formData.publishingRange" clearable placeholder="请输入发表范围" ></el-input>
          </el-form-item>
            </el-row> 
            <el-row>
                <el-col >
                       <el-form-item label="ISSN号:">
                <el-input v-model="formData.issn" clearable placeholder="请输入ISSN号" ></el-input>
          </el-form-item>
                </el-col>
                <el-col >
                       <el-form-item label="ＣＮ号:">
                <el-input v-model="formData.cn" clearable placeholder="请输入ＣＮ号" ></el-input>
          </el-form-item>
                </el-col>
                 
            </el-row> 

              <el-row>
                <el-col >
                  <el-form-item label="发表/出版时间" prop="publishingDate">
                <el-date-picker v-model="formData.publishingDate" format="yyyy-MM-dd"
              value-format="yyyy-MM-dd"
                 placeholder="请选择发表/出版时间" clearable   :style="{width: '95%'}"></el-date-picker>
              </el-form-item>       
                </el-col>
                <el-col >
                       <el-form-item label="字数" prop="wordLength">
                  <el-input-number v-model="formData.wordLength" placeholder="字数" :precision="2" ></el-input-number>
           </el-form-item>
                </el-col>
                 
            </el-row> 
      </el-form>
       
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
      dialogTitle: "新增",
      formData: {
        title: "",
        publicationType: 1,
        kind:1,
        publicationName:"",
        period:"",
        volumeNo:"",
        publishingDate: new Date(),
        wordLength:0,
        knowledgeClass:"",
        firstKnowledge:"",
        source:"",
        publishingRange:"",
        issn:"",
        cn:"",
      },
      methodOptions: methodOptions,
      type: "",
      rules: {
        title: [{ required: true, message: "请输入论文名称", trigger: "blur" }],
     
      },
      publicationType: [{
        "label": "A类",
        "value": 1
      }, {
        "label": "B类",
        "value": 2
      }, {
        "label": "C类",
        "value": 3
      }, {
        "label": "D类",
        "value": 4
      }, {
        "label": "E类",
        "value": 5
      }, {
        "label": "F类·",
        "value": 6
      }],
      kind: [{
        "label": "期刊论文",
        "value": 1
      }, {
        "label": "论文集",
        "value": 2
      }, {
        "label": "文章",
        "value": 3
      }],
      knowledgeClass: [{
        "label": "科技类",
        "value": '1'
      }, {
        "label": "社科类",
        "value": '2'
      }],
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
      //this.$refs.form.resetFields();
      this.formData = {
        title: "",
        kind: 1,
        publicationType: 1,
        publicationName:"",
        period:"",
        publishingDate: new Date(),
        wordLength:0,
        knowledgeClass:"",
        firstKnowledge:"",
        source:"",
        publishingRange:"",
        issn:"",
        cn:"",
      };
    },
    closeDialog() {
      this.initForm();
      this.dialogFormVisible = false;
    },
    openDialog(type) {
      switch (type) {
        case "addApi":this.initForm();
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
      const res = await getInfo({ ID: row.ID });
      this.formData = res.data.detail;
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
      // console.log('****' + JSON.stringify(this.formData) + '***')
      this.formData.publishingDate = new Date(this.formData.publishingDate).toISOString()
      this.$refs.formData.validate(async valid => {
        if (valid) {         
          switch (this.type) {
            case "addApi":
              {               
                const res = await create(this.formData);
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
                const res = await update(this.formData);              
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
