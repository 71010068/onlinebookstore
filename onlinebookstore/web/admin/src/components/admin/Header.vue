<template>
<div>
    <a-layout-header class="col headerBtn">
    <a-button type="danger" @click="loginOut">退出</a-button>
</a-layout-header>
</div>
</template>

<script>
export default {
    data() {
    return {
      collapsed: false,
    };
  },
  methods: {
    loginOut() {
      window.sessionStorage.clear("token"); //清除token
      this.$router.push("/");
    },
    
  },
};
</script>

<style scoped>
.col{
     background-color: rgb(50, 208, 236);
}
.headerBtn{
    display: flex;
    justify-content: flex-end;
    align-items: center;
}
</style>