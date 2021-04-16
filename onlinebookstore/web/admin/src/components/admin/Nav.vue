<template>
  <div>
    <a-layout-sider
      v-model="collapsed"
      breakpoint="lg"
      style="width: 256px; max-width: 256px; min-width: 256px"
      theme="light"
      class="col"
      ><div class="log">
        <span>{{ collapsed ? "商城" : "图书商城" }}</span>
      </div>

      <a-menu
        style="width: 256px"
        :open-keys.sync="openKeys"
        mode="inline"
        @click="goToPage"
      >
        <a-menu-item key="index">
          <a-icon type="heart" theme="twoTone" two-tone-color="#eb2f96" spin />
          <span id="tit">欢迎进入后台首页</span></a-menu-item
        >
        <a-sub-menu key="sub1">
          <span slot="title"><a-icon type="user" /><span>用户管理</span></span>
          <a-menu-item-group key="g1"
            ><template slot="title"
              ><a-icon type="frown" theme="filled" /><span> 操作 </span>
            </template>
            <a-menu-item key="userlist">用户列表</a-menu-item>
            <!-- <a-menu-item key="personal">用户个人设置</a-menu-item> -->
          </a-menu-item-group>
        </a-sub-menu>
        <a-sub-menu key="sub2">
          <span slot="title"><a-icon type="book" /><span>图书管理</span></span>
          <a-menu-item-group key="g2"
            ><template slot="title"
              ><a-icon type="smile" theme="filled" /><span> 操作 </span>
            </template>
            <a-menu-item key="booklist"> 图书列表 </a-menu-item>
            
            <a-menu-item key="addbook"> 添加图书 </a-menu-item>
          </a-menu-item-group>
        </a-sub-menu>
        <a-sub-menu key="sub3">
          <span slot="title"
            ><a-icon type="container" /><span>分类管理</span></span
          >
          <a-menu-item-group key="g3"
            ><template slot="title"
              ><a-icon type="meh" theme="filled" /><span> 操作 </span>
            </template>
            <a-menu-item key="catelist"> 图书分类 </a-menu-item>  
            <!-- <a-menu-item key="cart"> 购物车 </a-menu-item> -->
          </a-menu-item-group>
        </a-sub-menu>
      </a-menu>
    </a-layout-sider>
  </div>
</template>


<script>
export default {
  data() {
    return {
      collapsed: false,
      openKeys: ["sub1", "sub2", "sub3"],
    };
  },
  watch: {
    openKeys(val) {
      console.log("openKeys", val);
    },
  },
  methods: {
    goToPage(item) {
      // this.$router.push("/admin/" + item.key).catch((err) => err);
      this.$router.push("/" + item.key).catch((err) => err);
    },
  },
};
</script>


<style scoped>
#tit {
  font-size: 1.3em;
}
.col {
  background-color: rgb(50, 208, 236);
}
.log {
  height: 32px;
  margin: 16px;
  background: #fafafa;
  display: flex;
  justify-content: center;
  align-items: center;
  font: 18px bold 幼圆;
}
</style>
