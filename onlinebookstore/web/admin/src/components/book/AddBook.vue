<template>
  <div id="addbook">
    <a-card>
      <a-divider
        ><h3>{{ id ? "编辑图书" : "添加图书" }}</h3></a-divider
      >
      <a-form-model
        :model="bookInfo"
        ref="bookInfoRef"
        :rules="bookInfoRules"
        :hideRequiredMark="true"
      >
        <a-form-model-item
          :label-col="labelCol"
          :wrapper-col="wrapperCol"
          label="图书书名 ："
          prop="title"
        >
          <a-input placeholder="请输入该图书的名称" v-model="bookInfo.title">
            <a-icon slot="prefix" type="robot" />
            <a-tooltip slot="suffix" title="该图书的书名">
              <a-icon type="info-circle" style="color: rgba(0, 0, 0, 0.45)" />
            </a-tooltip>
          </a-input>
        </a-form-model-item>

        <a-form-model-item
          :label-col="labelCol"
          :wrapper-col="wrapperCol"
          label="图书作者 ："
          prop="author"
        >
          <a-input placeholder="请输入该图书的作者" v-model="bookInfo.author">
            <a-icon slot="prefix" type="user" />
            <a-tooltip slot="suffix" title="撰写该图书的作者">
              <a-icon type="info-circle" style="color: rgba(0, 0, 0, 0.45)" />
            </a-tooltip>
          </a-input>
        </a-form-model-item>

        <a-form-model-item
          :label-col="labelCol"
          :wrapper-col="wrapperCol"
          label="图书价格 ："
          prop="price"
        >
          <a-input
            placeholder="请输入该图书的价格"
            prefix="￥"
            suffix="RMB"
            v-model="bookInfo.price"
          ></a-input>
        </a-form-model-item>

        <a-form-model-item
          :label-col="labelCol"
          :wrapper-col="wrapperCol"
          label="图书分类 ："
          prop="cid"
        >
          <a-select
            placeholder="请选择该图书的分类"
            v-model="bookInfo.cid"
            @change="cateChange"
            :defaultFileList="fileList"
          >
            <a-select-option
              v-for="item in Catelist"
              :key="item.id"
              :value="item.id"
              >{{ item.name }}</a-select-option
            >
          </a-select>
        </a-form-model-item>

        <a-form-model-item
          :label-col="labelCol"
          :wrapper-col="wrapperCol"
          label="图书描述 ："
          prop="Desc"
        >
          <a-input
            placeholder="请输入该图书的简介"
            type="textarea"
            v-model="bookInfo.Desc"
          >
          </a-input>
        </a-form-model-item>

        <a-form-model-item
          :label-col="labelCol"
          :wrapper-col="wrapperCol"
          style="width: 760px; margin: 0 auto"
          label="图书缩略图"
          prop="img"
        >
          <a-upload
            name="file"
            listType="picture"
            :action="upUrl"
            :headers="headers"
            @change="upChange"
          >
            <a-button> <a-icon type="upload" /> 点击上传 </a-button>
            <br />
            <template v-if="id">
              <img :src="bookInfo.img" style="width: 120px; height: 120px" />
            </template>
          </a-upload>
        </a-form-model-item>

        <a-divider></a-divider>

        <a-form-model-item >
          <a-button
            type="danger"
            style="margin-right: 15px"
            @click="bookOk(bookInfo.id)"
            >{{ bookInfo.id ? "更新" : "提交" }}</a-button
          >
          <a-button type="primary" @click="addCancel">取消</a-button>
        </a-form-model-item>
      </a-form-model>
    </a-card>
  </div>
</template>

<script>
import { Url } from "../../plugin/http";
export default {
  props: ["id"],
  data() {
    return {
      labelCol: { span: 9 },
      wrapperCol: { span: 6 },
      bookInfo: {
        id: 0,
        title: "",
        author: "",
        cid: undefined,
        desc: "",
        price: "",
        img: "",
      },
      Catelist: [],
      upUrl: Url + "/upload",
      headers: {}, // 请求头
      fileList: [], // 当前操作的文件列表
      bookInfoRules: {
        // 表单验证
        title: [{ required: true, message: "请输入图书标题", trigger: "blur" }],
        author: [
          { required: true, message: "请输入作者名称", trigger: "blur" },
        ],
        Desc: [
          { required: true, message: "请输入图书描述", trigger: "blur" },
          { max: 120, message: "描述最多可写120个字符", trigger: "change" },
        ],
        cid: [{ required: true, message: "请输入图书分类", trigger: "blur" }],
        img: [{ required: true, message: "请选择图书缩略图", trigger: "blur" }],
        price: [
          { required: true, message: "请输入图书价格", trigger: "blur" },
        ],
      },
    };
  },
  // 进入页面渲染做什么
  created() {
    this.getCateList();
    this.headers = {
      Authorization: `Bearer ${window.sessionStorage.getItem("token")}`,
    };
    if (this.id) {
      this.getBookInfo(this.id);
    }
  },
  methods: {
    // 查询图书信息
    async getBookInfo(id) {
      const { data: res } = await this.$http.get(`/book/info/${id}`);
      if (res.status !== 200) return this.$message.error(res.message);
      this.bookInfo = res.data;
      this.bookInfo.id = res.data.ID;
      // console.log(res);
    },
    // 获取分类列表
    async getCateList() {
      const { data: res } = await this.$http.get("/category");
      if (res.status !== 200) return this.$message.error(res.message);
      this.Catelist = res.data;
    },
    // 选择分类
    cateChange(value) {
      this.bookInfo.cid = value;
    },
    // 上传图片接口的逻辑
    upChange(info) {
      if (info.file.status !== "uploading") {
        console.log(info.file, info.fileList);
      }
      if (info.file.status === "done") {
        this.$message.success(`图片上传成功。`);
        const imgUrl = info.file.response.url;
        this.bookInfo.img = imgUrl;
      } else if (info.file.status === "error") {
        this.$message.error(`图片上传失败。`);
      }
    },

    // 提交&&更新图书
    bookOk(id) {
      this.$refs.bookInfoRef.validate(async (valid) => {
        if (id === 0) {
          const { data: res } = await this.$http.post(
            "/book/add",
            this.bookInfo
          );
          if (res.status !== 200) return this.$message.error(res.message);
          this.$router.push("/booklist");
          this.$message.success("添加图书成功。");
        } else {
          const { data: res } = await this.$http.put(
            `/book/${id}`,
            this.bookInfo
          );
          if (res.status !== 200) return this.$message.error(res.message);
          // this.$router.push("/admin/booklist");
          this.$router.push("/booklist");
          this.$message.success("更新图书成功。");
        }
      });
    },
    // 取消提交&&更新图书
    addCancel() {
      this.$refs.bookInfoRef.resetFields();
    },
  },
};
</script>