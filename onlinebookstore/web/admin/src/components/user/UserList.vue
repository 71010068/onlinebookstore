<template>
  <div>
    <a-card>
      <a-row :gutter="20">
        <a-col :span="6">
          <a-input-search
            v-model="queryParam.username"
            placeholder="输入用户名查找"
            enter-button
            allowClear
            @search="getUserList"
          />
        </a-col>
        <a-col :span="2">
          <a-button icon="plus" type="primary" @click="addUserVisible = true"
            >新增</a-button
          >
        </a-col>
      </a-row>
      <a-table
        rowKey="username"
        :columns="columns"
        :pagination="pagination"
        :dataSource="userlist"
        @change="handleTableChange"
        bordered
      >
        <span slot="role" slot-scope="role">{{
          role == 1 ? "管理员" : "用户"
        }}</span>
        <template slot="action" slot-scope="data">
          <div class="actionSlot">
            <a-button
              type="dashed"
              style="margin-right: 8px"
              icon="edit"
              @click="editUser(data.ID)"
              >编辑</a-button
            >
            <a-button
              icon="minus-circle"
              type="dashed"
              @click="ChangePassword(data.ID)"
              >修改密码</a-button
            >
            <a-popconfirm
              title=" 确定要删除吗，想好了吗 ？"
              placement="left"
              ok-text="还要继续删除。"
              cancel-text="不删除了。"
              @confirm="deleteUser(data.ID)"
              @keyup.enter.native="deleteUser(data.ID)"
              @cancel="cancel"
            >
              <a-button icon="delete" style="margin: 0 8px" type="danger"
                >删除</a-button
              >
              <!-- @click="deleteUser(data.ID)" -->
            </a-popconfirm>
            <a-button
              icon="info-circle"
              type="link"
              @click="getPersonalInfo(data.ID)"
              >个人信息</a-button
            >
          </div>
        </template>
      </a-table>
    </a-card>

    <!-- 修改密码 ------------------------------------------------------------------------ -->
    <a-modal
      closable
      title="修改密码"
      :visible="changePasswordVisible"
      width="20%"
      @ok="changePasswordOk"
      @cancel="changePasswordCancel"
      destroyOnClose
    >
      <a-form-model
        :model="changePassword"
        :rules="changePasswordRules"
        ref="changePasswordRef"
      >
        <a-form-model-item has-feedback label="密码" prop="password">
          <a-input-password
            v-model="changePassword.password"
          ></a-input-password>
        </a-form-model-item>
        <a-form-model-item has-feedback label="确认密码" prop="checkpass">
          <a-input-password
            v-model="changePassword.checkpass"
          ></a-input-password>
        </a-form-model-item>
      </a-form-model>
    </a-modal>

    <!-- ------------------------------------新增用户区域---------------------------------------- -->
    <a-modal
      closable
      width="25%"
      destroyOnClose
      title="新增用户"
      :visible="addUserVisible"
      @ok="addUserOk"
      @keyup.enter.native="addUserOk"
      @cancel="addUserCancel"
    >
      <a-form-model :model="newUser" :rules="addUserRules" ref="addUserRef">
        <a-form-model-item
          :label-col="labelCol"
          :wrapper-col="wrapperCol"
          has-feedback
          label="用户名"
          prop="username"
        >
          <a-input v-model="newUser.username"></a-input>
        </a-form-model-item>

        <a-form-model-item
          :label-col="labelCol"
          :wrapper-col="wrapperCol"
          has-feedback
          label="密码"
          prop="password"
        >
          <a-input-password v-model="newUser.password"></a-input-password>
        </a-form-model-item>

        <a-form-model-item
          :label-col="labelCol"
          :wrapper-col="wrapperCol"
          has-feedback
          label="确认密码"
          prop="checkpass"
        >
          <a-input-password v-model="newUser.checkpass"></a-input-password>
        </a-form-model-item>

        <a-form-model-item
          :label-col="labelCol"
          :wrapper-col="wrapperCol"
          has-feedback
          label="手机号码"
          prop="phone"
        >
          <a-input v-model="newUser.phone"></a-input>
        </a-form-model-item>

        <a-form-model-item
          :label-col="labelCol"
          :wrapper-col="wrapperCol"
          has-feedback
          label="电子邮箱"
          prop="email"
        >
          <a-input v-model="newUser.email"></a-input>
        </a-form-model-item>

        <a-form-model-item label="是否为管理员" prop="role">
          <!-- <a-select defauleValue="2" style="120px" @change="adminChange">
            <a-select-option key="1" value="1">是</a-select-option>
            <a-select-option key="2" value="2">否</a-select-option>
          </a-select> -->
          <a-switch
            checked-children="是"
            un-checked-children="否"
            @change="adminChange"
          />
        </a-form-model-item>
      </a-form-model>
    </a-modal>

    <!-- ------------------------------------编辑用户区域---------------------------------------- -->
    <a-modal
      closable
      width="25%"
      title="编辑用户"
      :visible="editUserVisible"
      @ok="editUserOk"
      @keyup.enter.native="editUserOk"
      @cancel="editUserCancel"
    >
      <a-form-model :model="userInfo" :rules="userRules" ref="addUserRef">
        <a-form-model-item has-feedback label="用户名" prop="username">
          <a-input v-model="userInfo.username"></a-input>
        </a-form-model-item>
        <a-form-model-item has-feedback label="手机号码" prop="phone">
          <a-input v-model="userInfo.phone"></a-input>
        </a-form-model-item>
        <a-form-model-item has-feedback label="电子邮箱" prop="email">
          <a-input v-model="userInfo.email"></a-input>
        </a-form-model-item>
        <a-form-model-item label="是否为管理员">
          <!-- <a-select defauleValue="2" style="120px" @change="adminChange">
            <a-select-option key="1" value="1">是</a-select-option>
            <a-select-option key="2" value="2">否</a-select-option>
          </a-select> -->
          <a-switch
            checked-children="是"
            un-checked-children="否"
            @change="adminChange"
          />
        </a-form-model-item>
      </a-form-model>
    </a-modal>

    <!-- 修改个人信息设置 -------------------------------------------------------------------- -->
    <a-modal
      closable
      width="30%"
      title="个人信息设置"
      :visible="personalVisible"
      @ok="personalOk"
      @keyup.enter.native="personalOk"
      @cancel="personalCancel"
    >
      <a-form-model
        :model="personalInfo"
        :rules="personalRules"
        ref="personalRef"
      >
        <a-form-model-item
          :label-col="labelCol"
          :wrapper-col="wrapperCol"
          has-feedback
          label="昵称："
          prop="name"
        >
          <a-input v-model="personalInfo.name"></a-input>
        </a-form-model-item>

        <a-form-model-item
          :label-col="labelCol"
          :wrapper-col="wrapperCol"
          has-feedback
          label="关于我："
          prop="desc"
        >
          <a-input v-model="personalInfo.desc"></a-input>
        </a-form-model-item>

        <a-form-model-item
          :label-col="labelCol"
          :wrapper-col="wrapperCol"
          has-feedback
          label="QQ："
          prop="qq_chat"
        >
          <a-input v-model="personalInfo.qq_chat"></a-input>
        </a-form-model-item>

        <a-form-model-item
          :label-col="labelCol"
          :wrapper-col="wrapperCol"
          has-feedback
          label="微信："
          prop="wechat"
        >
          <a-input v-model="personalInfo.wechat"></a-input>
        </a-form-model-item>

        <a-form-model-item
          :label-col="labelCol"
          :wrapper-col="wrapperCol"
          style="margin: 0 auto"
          label="用户头像"
          prop="avatar"
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
              <img
                :src="personalInfo.avatar"
                style="width: 120px; height: 120px; margin-top: 8px"
              />
            </template>
          </a-upload>
        </a-form-model-item>
      </a-form-model>
    </a-modal>
  </div>
</template>

<script>
import { Url } from "../../plugin/http";
const columns = [
  {
    title: "ID",
    dataIndex: "ID",
    width: "8%",
    key: "id",
    align: "center",
  },
  {
    title: "用户名",
    dataIndex: "username",
    width: "16%",
    key: "username",
    align: "center",
  },
  {
    title: "电话号码",
    dataIndex: "phone",
    width: "16%",
    key: "phone",
    align: "center",
  },
  {
    title: "邮箱",
    dataIndex: "email",
    width: "20%",
    key: "email",
    align: "center",
  },
  {
    title: "角色",
    dataIndex: "role",
    width: "10%",
    key: "role",
    align: "center",
    scopedSlots: { customRender: "role" },
  },
  {
    title: "操作",
    width: "30%",
    key: "action",
    align: "center",
    scopedSlots: { customRender: "action" },
  },
];
export default {
  data() {
    return {
      upUrl: Url + "/upload",
      headers: {}, // 请求头
      userlist: [],
      columns,
      labelCol: { span: 6 },
      wrapperCol: { span: 14 },
      pagination: {
        pageSizeOptions: ["5", "10", "20"],
        pageSize: 5,
        total: 0,
        showSizeChanger: true,
        showTotal: (total) => `共${total}条`,
      },
      queryParam: {
        username: "",
        pagesize: 5,
        pagenum: 1,
      },
      userInfo: {
        // 绑定用户信息
        id: 0,
        username: "",
        password: "",
        checkpass: "",
        phone: "",
        email: "",
        role: 2,
      },
      personalInfo: {
        id: 0,
        name: "",
        desc: "",
        qq_chat: "",
        wechat: "",
        avatar: "",
      },
      newUser: {
        // 绑定用户信息
        id: 0,
        username: "",
        password: "",
        checkpass: "",
        phone: "",
        email: "",
        role: 2,
      },
      changePassword: {
        id: 0,
        password: "",
        checkPass: "",
      },
      visible: false, // 删除的
      addUserVisible: false, // 新增用户的
      personalRules: {
        // 个人信息的
        name: [
          {
            validator: (rule, value, callback) => {
              if (this.personalInfo.name == "") {
                callback(new Error("请输入昵称。"));
              }
              if (
                [...this.personalInfo.name].length < 1 ||
                [...this.personalInfo.name].length > 12
              ) {
                callback(new Error("昵称应该在 1 ~ 12 位之间"));
              } else {
                callback();
              }
            },
            trigger: "change",
          },
        ],
        desc: [
          {
            validator: (rule, value, callback) => {
              if (this.personalInfo.desc == "") {
                callback(new Error("请输入关于该用户的介绍。"));
              }
              if (
                [...this.personalInfo.desc].length < 1 ||
                [...this.personalInfo.desc].length > 50
              ) {
                callback(new Error("字数在在 1 ~ 50 字之间"));
              } else {
                callback();
              }
            },
            trigger: "change",
          },
        ],
        qq_chat: [
          {
            validator: (rule, value, callback) => {
              if (this.personalInfo.qq_chat == "") {
                callback(new Error("请输入QQ号。"));
              }
              if (
                [...this.personalInfo.qq_chat].length < 5 ||
                [...this.personalInfo.qq_chat].length > 10
              ) {
                callback(new Error("QQ号在5 ~ 10 位之间哦。"));
              } else {
                callback();
              }
            },
            trigger: "change",
          },
        ],
        wechat: [
          {
            validator: (rule, value, callback) => {
              if (this.personalInfo.wechat == "") {
                callback(new Error("请输入微信号。"));
              } else {
                callback();
              }
            },
            trigger: "change",
          },
        ],
      },
      userRules: {
        // 编辑用户的
        username: [
          {
            validator: (rule, value, callback) => {
              if (this.userInfo.username == "") {
                callback(new Error("请输入用户名。"));
              }
              if (
                [...this.userInfo.username].length < 4 ||
                [...this.userInfo.username].length > 12
              ) {
                callback(new Error("用户名应该在 4 ~ 12 位之间"));
              } else {
                callback();
              }
            },
            trigger: "change",
          },
        ],
        password: [
          {
            validator: (rule, value, callback) => {
              if (this.userInfo.password == "") {
                callback(new Error("请输入密码。"));
              }
              if (
                [...this.userInfo.password].length < 6 ||
                [...this.userInfo.password].length > 20
              ) {
                callback(new Error("密码应该在 6 ~ 20 位之间"));
              } else {
                callback();
              }
            },
            trigger: "change",
          },
        ],
        checkpass: [
          {
            validator: (rule, value, callback) => {
              if (this.userInfo.checkpass == "") {
                callback(new Error("请输入密码。"));
              }
              if (this.userInfo.password !== this.userInfo.checkpass) {
                callback(new Error("密码不一致，请重新输入。"));
              } else {
                callback();
              }
            },
            trigger: "change",
          },
        ],
        phone: [
          {
            validator: (rule, value, callback) => {
              var reg = /^1(3|4|5|6|7|8)\d{9}$/;
              if (!reg.test(value)) {
                callback(new Error("手机号输入错误，请重新输入。"));
              } else {
                callback();
              }
            },
            trigger: "change",
          },
        ],
        email: [
          {
            validator: (rule, value, callback) => {
              var reg = /^([a-zA-Z]|[0-9])(\w|-)+@[a-zA-Z0-9]+\.([a-zA-Z]{2,4})$/;
              if (!reg.test(value)) {
                callback(new Error("邮箱地址不正确，请重新输入。"));
              } else {
                callback();
              }
            },
            trigger: ["blur", "change"],
          },
        ],
      },

      addUserRules: {
        // 新增用户的
        username: [
          {
            validator: (rule, value, callback) => {
              if (this.newUser.username == "") {
                callback(new Error("请输入用户名。"));
              }
              if (
                [...this.newUser.username].length < 4 ||
                [...this.newUser.username].length > 12
              ) {
                callback(new Error("用户名应该在 4 ~ 12 位之间"));
              } else {
                callback();
              }
            },
            trigger: ["blur", "change"],
          },
        ],
        password: [
          {
            validator: (rule, value, callback) => {
              if (this.newUser.password == "") {
                callback(new Error("请输入密码。"));
              }
              if (
                [...this.newUser.password].length < 6 ||
                [...this.newUser.password].length > 20
              ) {
                callback(new Error("密码应该在 6 ~ 20 位之间"));
              } else {
                callback();
              }
            },
            trigger: ["blur", "change"],
          },
        ],
        checkpass: [
          {
            validator: (rule, value, callback) => {
              if (this.newUser.checkpass == "") {
                callback(new Error("请输入密码。"));
              }
              if (this.newUser.password !== this.newUser.checkpass) {
                callback(new Error("密码不一致，请重新输入。"));
              } else {
                callback();
              }
            },
            trigger: ["blur", "change"],
          },
        ],
        phone: [
          {
            validator: (rule, value, callback) => {
              var reg = /^1(3|4|5|6|7|8)\d{9}$/;
              if (!reg.test(value)) {
                callback(new Error("手机号输入错误，请重新输入。"));
              } else {
                callback();
              }
            },
            trigger: ["blur", "change"],
          },
        ],
        email: [
          {
            validator: (rule, value, callback) => {
              var reg = /^([a-zA-Z]|[0-9])(\w|-)+@[a-zA-Z0-9]+\.([a-zA-Z]{2,4})$/;
              if (!reg.test(value)) {
                callback(new Error("邮箱地址不正确，请重新输入。"));
              } else {
                callback();
              }
            },
            trigger: ["blur", "change"],
          },
        ],
      },
      changePasswordRules: {
        password: [
          {
            validator: (rule, value, callback) => {
              if (this.changePassword.password == "") {
                callback(new Error("请输入密码"));
              }
              if (
                [...this.changePassword.password].length < 6 ||
                [...this.changePassword.password].length > 20
              ) {
                callback(new Error("密码应当在6到20位之间"));
              } else {
                callback();
              }
            },
            trigger: ["blur", "change"],
          },
        ],
        checkpass: [
          {
            validator: (rule, value, callback) => {
              if (this.changePassword.checkpass == "") {
                callback(new Error("请输入密码"));
              }
              if (
                this.changePassword.password !== this.changePassword.checkpass
              ) {
                callback(new Error("密码不一致，请重新输入"));
              } else {
                callback();
              }
            },
            trigger: ["blur", "change"],
          },
        ],
      },
      changePasswordVisible: false, // 改密码的
      editUserVisible: false, // 编辑用户的
      personalVisible: false, // 个人信息设置的
    };
  },
  created() {
    this.getUserList();
    this.headers = {
      Authorization: `Bearer ${window.sessionStorage.getItem("token")}`,
    };
  },
  methods: {
    // 获取用户列表
    async getUserList() {
      const { data: res } = await this.$http.get("users", {
        params: {
          username: this.queryParam.username,
          pagesize: this.queryParam.pagesize,
          pagenum: this.queryParam.pagenum,
        },
      });
      if (res.status != 200) return this.$message.error(res.message);
      this.userlist = res.data;
      this.pagination.total = res.total;
      console.log(res);
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
      this.getUserList();
    },

    // 删除用户
    // async deleteUser(id) {
    //   const res = await this.$http.delete(`user/${id}`);
    //       if (res.status != 200) return this.$message.error(res.message);
    //       this.$message.success("删除成功。");
    //       this.getUserList();
    //   // console.log(res)
    // },
    async deleteUser(id) {
      const res = await this.$http.delete(`user/${id}`);
      if (res.status != 200) return this.$message.error(res.message);
      this.getUserList();
      this.$message.success("删除成功。");
    },
    cancel() {
      this.$message.error("已取消删除。");
    },

    // 新增用户---------------------------------------------------------------------------
    addUserOk() {
      this.$refs.addUserRef.validate(async (valid) => {
        if (!valid) return this.$message.error("参数不符合要求，请重新输入。");
        const { data: res } = await this.$http.post("user/add", {
          username: this.newUser.username,
          password: this.newUser.password,
          phone: parseInt(this.newUser.phone),
          email: this.newUser.email,
          role: this.newUser.role,
        });
        if (res.status !== 200) return this.$message.error(res.message);

        this.addUserVisible = false;
        this.$message.success("添加用户成功。");
        this.getUserList();
      });
    },
    addUserCancel() {
      this.$refs.addUserRef.resetFields();
      this.addUserVisible = false;
      this.$message.info("新增用户已取消。");
    },
    adminChange(checked) {
      //this.userInfo.role = Number(value);
      // console.log(this.userInfo.role)
      //console.log(checked);
      if (checked) {
        this.userInfo.role = Number(1);
      } else {
        this.userInfo.role = Number(2);
      }
    },

    // 编辑用户------------------------------------------------------------------------------------
    async editUser(id) {
      this.editUserVisible = true;
      const { data: res } = await this.$http.get(`user/${id}`);
      this.userInfo = res.data;
      this.userInfo.id = id;
    },
    editUserOk() {
      this.$refs.addUserRef.validate(async (valid) => {
        if (!valid) return this.$message.error("参数不符合要求，请重新输入。");
        const { data: res } = await this.$http.put(`user/${this.userInfo.id}`, {
          username: this.userInfo.username,
          phone: parseInt(this.userInfo.phone),
          email: this.userInfo.email,
          role: this.userInfo.role,
        });
        if (res.status !== 200) return this.$message.error(res.message);

        this.editUserVisible = false;
        this.$message.success("更新用户信息成功。");
        this.getUserList();
      });
    },
    editUserCancel() {
      this.$refs.addUserRef.resetFields();
      this.editUserVisible = false;
      this.$message.info("编辑已取消。");
    },

    //  用户信息------------------------------------------------------------------------
    // 获取个人信息设置
    async getPersonalInfo(id) {
      this.personalVisible = true;
      const { data: res } = await this.$http.get(`personal/${id}`);
      if (res.status !== 200) return this.$message.error(res.message);
      this.personalInfo = res.data;
      this.personalInfo.id = id;
      //   this.bookInfo.id = res.data.ID;
      //console.log(this.personalInfo);
      //console.log(res);
    },
    // 更新
    async personalOk(id) {
      this.$refs.personalRef.validate(async (valid) => {
        if (!valid) return this.$message.error("参数不符合要求，请重新输入。");
        const { data: res } = await this.$http.put(
          `personal/${this.personalInfo.id}`,
          {
            name: this.personalInfo.name,
            desc: this.personalInfo.desc,
            qq_chat: this.personalInfo.qq_chat,
            wechat: this.personalInfo.wechat,
            avatar: this.personalInfo.avatar,
          }
        );
        if (res.status !== 200) return this.$message.error(res.message);
        this.personalVisible = false;
        this.$message.success(`个人信息更新成功。`);
        this.getUserList();
      });
    },
    // 取消
    personalCancel() {
      this.$refs.personalRef.resetFields();
      this.personalVisible = false;
      this.$message.info("修改用户信息已取消。");
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

    //  -----修改密码------------------------------------------------------------------------
    async ChangePassword(id) {
      this.changePasswordVisible = true;
      const { data: res } = await this.$http.get(`user/${id}`);
      this.changePassword.id = id;
    },
    changePasswordOk() {
      this.$refs.changePasswordRef.validate(async (valid) => {
        if (!valid) return this.$message.error("参数不符合要求，请重新输入");
        const { data: res } = await this.$http.put(
          `changepw/${this.changePassword.id}`,
          {
            password: this.changePassword.password,
          }
        );
        if (res.status != 200) return this.$message.error(res.message);
        this.changePasswordVisible = false;
        this.$message.success("修改密码成功");
        this.getUserList();
      });
    },
    changePasswordCancel() {
      this.$refs.changePasswordRef.resetFields();
      this.changePasswordVisible = false;
      this.$message.info("已取消");
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