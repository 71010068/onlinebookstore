<template>
  <div>
    <a-card>
         <a-divider>更新用户个人户信息</a-divider>
      <h2>{{id}}-----------------</h2>
      <a-form-model>
        <a-form-model-item
          :label-col="labelCol"
          :wrapper-col="wrapperCol"
          label="用户ID ：">
          <a-input placeholder="请输入用户的昵称" v-model="personalInfo.ID"></a-input>
        </a-form-model-item>

        <a-form-model-item
          :label-col="labelCol"
          :wrapper-col="wrapperCol"
          label="用户昵称 ：">
          <a-input placeholder="请输入用户的昵称" v-model="personalInfo.name"></a-input>
        </a-form-model-item>

        <a-form-model-item
          :label-col="labelCol"
          :wrapper-col="wrapperCol"
          style="width: 760px; margin: 0 auto"
          label="上传头像"
        >
          <a-upload
            name="file"
            listType="picture"
            :action="upUrl"
            :headers="headers"
            @change="avatarChange"
          >
            <a-button> <a-icon type="upload" /> 点击上传 </a-button>
            <br />
            <template v-if="personalInfo.avatar">
              <img :src="personalInfo.avatar" style="width: 120px; height: 120px" />
            </template>
          </a-upload>
        </a-form-model-item>


        <a-form-model-item
          :label-col="labelCol"
          :wrapper-col="wrapperCol"
          label="QQ号 ：">
          <a-input placeholder="请输入QQ号" v-model="personalInfo.qq_chat"></a-input>
        </a-form-model-item>
        
        <a-form-model-item
          :label-col="labelCol"
          :wrapper-col="wrapperCol"
          label="微信号 ：">
          <a-input placeholder="请输入微信号" v-model="personalInfo.wechat"></a-input>
        </a-form-model-item>

        <a-form-model-item
          :label-col="labelCol"
          :wrapper-col="wrapperCol"
          label="个人简介 ：">
          <a-input type="textarea" placeholder="请输入个人介绍" v-model="personalInfo.desc"></a-input>
        </a-form-model-item>


        <a-divider></a-divider>

        <a-form-model-item>
          <a-button
            type="danger"
            style="margin-right: 15px"
            @click="updatePersonal"
            >更新</a-button
          >
        </a-form-model-item>
      </a-form-model>
    </a-card>
  </div>
</template>

<script>
import { Url } from "../../plugin/http";
export default {
  props:['id'],
  data() {
    return {
      
      labelCol: { span: 9 },
      wrapperCol: { span: 6 },

      personalInfo:{
          id:this.id,
          name:'',
          desc:'',
          qq_chat:'',
          wechat:'',
          avatar:'',
      },
      upUrl: Url + "/upload",
      headers: {}, // 请求头
    };
  },
  created(){    // 生命周期函数进入页面渲染时的操作
    this.getPersonalInfo()
    this.headers = {
      Authorization: `Bearer ${window.sessionStorage.getItem("token")}`,
    };
  },
  methods:{
      // 获取个人信息设置
      async getPersonalInfo() {
      const { data: res } = await this.$http.get(`personal/${this.personalInfo.id}`);
      if (res.status !== 200) return this.$message.error(res.message);
      this.personalInfo = res.data;
    //   this.bookInfo.id = res.data.ID;
      //console.log(this.personalInfo);
      console.log(res);
    },

      // 更新
      async updatePersonal(id) {
        //   console.log(res+"++++++++++++++++++++++++++++++++++++++++++++++")
        const { data: res } = await this.$http.put(`personal/${this.personalInfo.id}`,this.personalInfo);
        if (res.status !== 200) return this.$message.error(res.message);
        //   console.log(res+"----------------------------------------------")
        this.$message.success(`个人信息更新成功。`);
        this.$router.push('/index')
      },
      
    // 上传头像接口的逻辑
    avatarChange(info) {
      if (info.file.status !== "uploading") {
        console.log(info.file, info.fileList);
      }
      if (info.file.status === "done") {
        this.$message.success(`图片上传成功。`);
        const imgUrl = info.file.response.url;
        this.personalInfo.avatar = imgUrl;
      } else if (info.file.status === "error") {
        this.$message.error(`图片上传失败。`);
      }
    },

    // 上传头像背景接口的逻辑
    imgChange(info) {
      if (info.file.status !== "uploading") {
        console.log(info.file, info.fileList);
      }
      if (info.file.status === "done") {
        this.$message.success(`图片上传成功。`);
        const imgUrl = info.file.response.url;
        this.personalInfo.img = imgUrl;
      } else if (info.file.status === "error") {
        this.$message.error(`图片上传失败。`);
      }
    },

  }
};
</script>

<style>
</style>