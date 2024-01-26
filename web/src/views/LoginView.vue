<template>
  <div class="wrapper">
    <div class="left">

    </div>
    <div class="right">
      <el-tabs v-model="activeTab">
        <el-tab-pane label="登录" name="login">
          <Login/>
        </el-tab-pane>
        <el-tab-pane label="注册" name="register">
          <Register @switchTab="switchTab"/>
        </el-tab-pane>
      </el-tabs>
    </div>
  </div>
</template>
<script setup>
import { ref } from 'vue'
import Login from '@/components/Login.vue'
import Register from "@/components/Register.vue";

const activeTab = ref('login')
const switchTab=(type)=>{
  if (type==='register'){
    activeTab.value='register'
  }else{
    activeTab.value='login';
  }
}
</script>
<style scoped>
.wrapper{
  display: flex;
  justify-content: space-around;
  .right{
    min-width: 350px;
  }
}
</style>