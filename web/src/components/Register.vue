<template>
  <div class="form-wrapper">
    <el-form
        label-position="top"
        :model="formInfo"
    >
      <el-form-item label="用户名" >
        <el-input v-model="formInfo.username" style="width: 320px;"/>
      </el-form-item>
      <el-form-item label="密码">
        <el-input v-model="formInfo.password"  style="width: 320px;"/>
      </el-form-item>
      <el-button type="primary" @click="handleLoginBtn" :loading="loginBtnInfo.loading" :disabled="loginBtnInfo.disabled">注册</el-button>
    </el-form>
  </div>
</template>
<script setup>
import { reactive } from 'vue'
import axios from 'axios';
import {ElMessage} from "element-plus";
const formInfo = reactive({
  username: '',
  password: '',
})
const loginBtnInfo=reactive({
  loading:false,
  disabled:false,
})
const handleLoginBtn=()=>{
  loginBtnInfo.disabled=true;
  loginBtnInfo.loading=true;
  axios.post("/api/user/register",{
    username:formInfo.username,
    password:formInfo.password
  }).then(response=>{
    console.log(response.data)
    if(response.data.code!==200){
      ElMessage.error(response.data.msg);
    }else{
      ElMessage.success("注册成功！")
    }
  }).finally(e=>{
    loginBtnInfo.disabled=false;
    loginBtnInfo.loading=false;
  })
}
</script>
<style scoped>
.form-wrapper{
  display: flex;
  justify-content: left;
  padding: 20px 0;
}
</style>