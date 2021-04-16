<template>
  <div>
    <a-card>
      <a-row :gutter="20">
        <a-col :span="4">
          <a-button icon="plus" type="primary" @click="addCateVisible = true"
            >新增分类</a-button
          >
        </a-col>
      </a-row>
      <a-table
        rowKey="id"
        :columns="columns"
        :pagination="pagination"
        :dataSource="Catelist"
        @change="handleTableChange"
        bordered
      >
        <template slot="action" slot-scope="data">
          <div class="actionSlot">
            <a-button type="primary" icon="edit" @click="editCate(data.id)"
              >编辑</a-button
            >
            <a-popconfirm
              title=" 确定要删除该分类吗，想好了吗 ？"
              placement="left"
              ok-text="还要继续删除。"
              cancel-text="不删除了。"
              @confirm="deleteCate(data.id)"
              @keyup.enter.native="deleteCate(data.id)"
              @cancel="cancel"
            >
              <a-button icon="delete" style="margin: 0 20px" type="danger"
                >删除</a-button
              >
              <!-- @click="deleteCate(data.ID)" -->
            </a-popconfirm>
          </div>
        </template>
      </a-table>
    </a-card>

    <!-- ------------------------------------新增分类区域---------------------------------------- -->
    <a-modal
      closable
      width="30%"
      destroyOnClose
      title="新增分类"
      :visible="addCateVisible"
      @ok="addCateOk"
      @keyup.enter.native="addCateOk"
      @cancel="addCateCancel"
    >
      <a-form-model :model="newCate" :rules="addCateRules" ref="addCateRef">
        <a-form-model-item has-feedback label="分类名称" prop="name">
          <a-input v-model="newCate.name"></a-input>
        </a-form-model-item>
      </a-form-model>
    </a-modal>

    <!-- ------------------------------------编辑分类区域---------------------------------------- -->
    <a-modal
      closable
      destroyOnClose
      width="30%"
      title="编辑分类"
      :visible="editCateVisible"
      @ok="editCateOk"
      @keyup.enter.native="editCateOk"
      @cancel="editCateCancel"
    >
      <a-form-model :model="CateInfo" :rules="CateRules" ref="addCateRef">
        <a-form-model-item has-feedback label="分类名称" prop="name">
          <a-input v-model="CateInfo.name"></a-input>
        </a-form-model-item>
      </a-form-model>
    </a-modal>
  </div>
</template>

<script>
const columns = [
  {
    title: "ID",
    dataIndex: "id",
    width: "30%",
    key: "id",
    align: "center",
  },
  {
    title: "分类名",
    dataIndex: "name",
    width: "30%",
    key: "name",
    align: "center",
  },
  {
    title: "操作",
    width: "40%",
    key: "action",
    align: "center",
    scopedSlots: { customRender: "action" },
  },
];
export default {
  data() {
    return {
      Catelist: [],
      columns,
      pagination: {
        pageSizeOptions: ["5", "10", "20"],
        pageSize: 5,
        total: 0,
        showSizeChanger: true,
        showTotal: (total) => `共${total}条`,
      },
      queryParam: {
        pagesize: 5,
        pagenum: 1,
      },
      CateInfo: {
        // 绑定编辑分类信息
        id: 0,
        name: "",
      },
      newCate: {
        // 绑定信息
        id: 0,
        name: "",
      },
      visible: false, // 删除的
      addCateVisible: false, // 新增用户的
      CateRules: {
        // 编辑用户的
        name: [
          {
            validator: (rule, value, callback) => {
              if (this.CateInfo.name === "") {
                callback(new Error("请输入分类名。"));
              } else {
                callback();
              }
            },
            trigger: "blur",
          },
        ],
      },
      addCateRules: {
        // 新增用户的
        name: [
          {
            validator: (rule, value, callback) => {
              if (this.newCate.name === "") {
                callback(new Error("请输入分类名。"));
              } else {
                callback();
              }
            },
            trigger: "blur",
          },
        ],
      },
      editCateVisible: false, // 编辑用户的
    };
  },
  created() {
    this.getCateList();
  },
  methods: {
    // 获取分类列表
    async getCateList() {
      const { data: res } = await this.$http.get("category", {
        params: {
          pagesize: this.queryParam.pagesize,
          pagenum: this.queryParam.pagenum,
        },
      });
      if (res.status !== 200) return this.$message.error(res.message);
      this.Catelist = res.data;
      this.pagination.total = res.total;
      
    },
    // 更改分页
    handleTableChange(pagination, filters, sorter) {
      var pager = { ...this.pagination };
      pager.current = pagination.current;
      pager.pageSize = pagination.pageSize;
      this.queryParam.pagesize = pagination.pageSize;
      this.queryParam.pagenum = pagination.current;

      if (pagination.pageSize !== this.pagination.pageSize) {
        this.queryParam.pagenum = 1;
        pager.current = 1;
      }
      this.pagination = pager;
      this.getCateList();
    },

    // 删除分类
    // async deleteCate(id) {
    //   const res = await this.$http.delete(`Cate/${id}`);
    //       if (res.status != 200) return this.$message.error(res.message);
    //       this.$message.success("删除成功。");
    //       this.getCateList();
    //   // console.log(res)
    // },
    async deleteCate(id) {
      const res = await this.$http.delete(`category/${id}`);
      if (res.status != 200) return this.$message.error(res.message);
      this.getCateList();
      this.$message.success("删除成功。");
      console.log(res);
    },
    cancel() {
      this.$message.error("已取消删除。");
    },

    // 新增分类
    addCateOk() {
      this.$refs.addCateRef.validate(async (valid) => {
        if (!valid) return this.$message.error("参数不符合要求，请重新输入。");
        const { data: res } = await this.$http.post("category/add", {
          name: this.newCate.name,
        });
        if (res.status !== 200) return this.$message.error(res.message);

        this.addCateVisible = false;
        this.$message.success("添加分类成功。");
        await this.getCateList();
      });
    },
    addCateCancel() {
      this.$refs.addCateRef.resetFields();
      this.addCateVisible = false;
      this.$message.info("新增分类已取消。");
    },

    // 编辑分类
    async editCate(id) {
      this.editCateVisible = true;
      //   console.log(res+"++++++++++++++++++++++++++++++++++++++++++++++")
      const { data: res } = await this.$http.get(`category/${id}`);
      //   console.log(res+"----------------------------------------------")
      this.CateInfo = res.data;
      this.CateInfo.id = id;
    },
    editCateOk() {
      this.$refs.addCateRef.validate(async (valid) => {
        if (!valid) return this.$message.error("参数不符合要求，请重新输入。");
        const { data: res } = await this.$http.put(
          `category/${this.CateInfo.id}`,
          {
            name: this.CateInfo.name,
          }
        );
        if (res.status != 200) return this.$message.error(res.message);

        this.editCateVisible = false;
        this.$message.success("更新分类信息成功。");
        this.getCateList();
      });
    },
    editCateCancel() {
      this.$refs.addCateRef.resetFields();
      this.editCateVisible = false;
      this.$message.info("编辑已取消。");
    },
  },
};
</script>

<style scoped>
.actionSlot {
  display: flex;
  justify-content: center;
}
</style>