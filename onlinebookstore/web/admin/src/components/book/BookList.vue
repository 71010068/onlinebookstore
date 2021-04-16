<template>
  <div>
    <a-card>
      <a-row :gutter="15">
        <a-col :span="6">
          <a-input-search
            v-model="queryParam.title"
            placeholder="输入图书名查找"
            enter-button
            allowClear
            @search="getBookList"
          />
        </a-col>
        <a-col :span="2">
          <a-button icon="plus" type="primary" @click="$router.push(`/addbook`)"
            >新增</a-button
          >
        </a-col>

        <a-col :span="5" :push="3">
          <a-select
            allowClear
            placeholder="请选择分类"
            style="width: 200px"
            @change="CateChange"
            
          >
            <a-select-option
              v-for="item in Catelist"
              :key="item.id"
              :value="item.id"
              >{{ item.name }}</a-select-option
            >
          </a-select>
        </a-col>
        <a-col :span="2" :push="3">
          <a-button type="info" @click="getBookList()">显示全部</a-button>
        </a-col>
      </a-row>
      <a-table
        rowKey="ID"
        :columns="columns"
        :pagination="pagination"
        :dataSource="Booklist"
        @change="handleTableChange"
        bordered
      >
        <span class="BookImg" slot="img" slot-scope="img"
          ><img :src="img"
        /></span>
        <template slot="action" slot-scope="data">
          <div class="actionSlot">
            <a-button
              size="small"
              type="primary"
              icon="edit"
              
              @click="$router.push(`/addbook/${data.ID}`)"
              >编辑</a-button
            >
            <!-- @click="$router.push(`/admin/addbook/${data.ID}`)" -->
            <a-popconfirm
              title=" 确定要删除该图书吗，想好了吗 ？"
              placement="left"
              ok-text="还要继续删除。"
              cancel-text="不删除了。"
              @confirm="deleteBook(data.ID)"
              @keyup.enter.native="deleteBook(data.ID)"
              @cancel="cancel"
            >
              <a-button
                size="small"
                icon="delete"
                style="margin: 0 12px"
                type="danger"
                >删除</a-button
              >
              <!-- @click="deleteBook(data.ID)" -->
            </a-popconfirm>
          </div>
        </template>
      </a-table>
    </a-card>
  </div>
</template>

<script>
const columns = [
  {
    title: "ID",
    dataIndex: "ID",
    width: "5%",
    key: "id",
    align: "center",
  },
  {
    title: "分类",
    dataIndex: "category.name",
    width: "10%",
    key: "name",
    align: "center",
  },
  {
    title: "图书书名",
    dataIndex: "title",
    width: "20%",
    key: "title",
    align: "center",
  },
  {
    title: "图书作者",
    dataIndex: "author",
    width: "20%",
    key: "author",
    align: "center",
  },
  {
    title: "图书价格",
    dataIndex: "price",
    width: "10%",
    key: "price",
    align: "center",
  },
  {
    title: "图书描述",
    dataIndex: "Desc",
    width: "20%",
    key: "Desc",
    align: "center",
  },
  {
    title: "缩略图",
    dataIndex: "img",
    width: "15%",
    key: "img",
    align: "center",
    scopedSlots: { customRender: "img" },
  },
  {
    title: "操作",
    width: "20%",
    key: "action",
    align: "center",
    scopedSlots: { customRender: "action" },
  },
];
export default {
  data() {
    return {
      columns,
      Booklist: [],
      Catelist: [],
      pagination: {
        pageSizeOptions: ["5", "10", "20"],
        pageSize: 5,
        total: 0,
        showSizeChanger: true,
        showTotal: (total) => `共${total}条`,
      },
      queryParam: {
        title: "",
        pagesize: 5,
        pagenum: 1,
      },
    };
  },
  created() {
    this.getCateList();
    // this.getBookList();
  },
  methods: {
    // 获取图书列表
    async getBookList() {
      const { data: res } = await this.$http.get("book", {
        params: {
          title: this.queryParam.title,
          pagesize: this.queryParam.pagesize,
          pagenum: this.queryParam.pagenum,
        },
      });
      // console.log(res);
      if (res.status != 200) return this.$message.error(res.message);
      this.Booklist = res.data;
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
      this.getBookList();
    },

  
    // 获取分类
    async getCateList() {
      const { data: res } = await this.$http.get("/category");
      if (res.status !== 200) return this.$message.error(res.message);
      this.Catelist = res.data;
      this.pagination.total = res.total;
      this.getBookList();
      // console.log(res);
    },

    // 查询分类下的图书
    CateChange(value) {
      this.getCateBook(value);
    },
    async getCateBook(id) {
      const { data: res } = await this.$http.get(`/book/list/${id}`, {
        params: {
          pagesize: this.queryParam.pagesize,
          pagenum: this.queryParam.pagenum,
        },
      });
      if (res.status !== 200) return this.$message.error(res.message);
      this.Booklist = res.data;
      this.pagination.total = res.total;
    },


      // 删除用户
    // async deleteBook(id) {
    //   const res = await this.$http.delete(`Book/${id}`);
    //       if (res.status != 200) return this.$message.error(res.message);
    //       this.$message.success("删除成功。");
    //       this.getBookList();
    //   // console.log(res)
    // },
    async deleteBook(id) {
      const res = await this.$http.delete(`/book/${id}`);
      if (res.status != 200) return this.$message.error(res.message);
      this.getBookList();
      this.$message.success("删除成功。");
    },
    cancel() {
      this.$message.error("已取消删除。");
    },
  },
};
</script>

<style scoped>
.BookImg {
  height: 100%;
  width: 100%;
}
.BookImg img {
  width: 80px;
  height: 80px;
}
.actionSlot {
  display: flex;
  justify-content: center;
}
</style>