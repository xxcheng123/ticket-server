<template>
<div class="wrapper">
  <div class="desc">
    <el-descriptions
        title="用户信息"
        direction="horizontal"
        :column="1"
        border
    >
      <el-descriptions-item label="编号">{{userInfo.id}}</el-descriptions-item>
      <el-descriptions-item label="用户名">{{userInfo.username}}</el-descriptions-item>
      <el-descriptions-item label="余额" >{{userInfo.balance}}</el-descriptions-item>
    </el-descriptions>
  </div>
  <div class="panel">
    <el-button @click="changePasswordShow=true">修改密码</el-button>
  </div>
  <el-dialog
      v-model="changePasswordShow"
      title="修改密码"
      @close="handleChangePasswordDialogClose">
    <ChangePassword ref="changePasswordRef" @toClose="closeChangePasswordDialog"/>
  </el-dialog>
</div>
</template>
<script setup>
import {reactive,ref} from "vue";
import {getInfo} from "@/api/user.js";
import {useRouter} from "vue-router";
import ChangePassword from "@/components/ChangePassword.vue";
const userInfo=reactive({
  username:'',
  id:'0',
  balance:0.0
})
const router=useRouter()
const changePasswordShow=ref(false)
getInfo().then(result=>{
  console.log(result)
  if(result.code===200){
    userInfo.username=result.data.username;
    userInfo.id=result.data.id;
    userInfo.balance=result.data.balance;
  }else{
    router.push('/login')
  }
})
const changePasswordRef=ref(null)
const handleChangePasswordDialogClose=()=>changePasswordRef.value.notifyClosed()
const closeChangePasswordDialog=()=>changePasswordShow.value=false
</script>
<style scoped>
.wrapper{
  width: 600px;
  margin: 0 auto;
  .desc{
    max-width: 400px;
    margin: 20px auto;
  }
  .panel{
    display: flex;
    justify-content: right;
    margin: 20px 0;
  }
}
</style>s