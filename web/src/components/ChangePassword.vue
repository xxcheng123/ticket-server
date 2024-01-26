<template>
<div class="wrapper-cg">
  <el-form
  v-model="info"
  label-position="top"
  >
    <el-form-item label="新密码" prop="password">
      <el-input v-model="info.password" type="password"/>
    </el-form-item>
    <el-form-item label="确认密码" prop="againPassword">
      <el-input v-model="info.againPassword" type="password"/>
    </el-form-item>
    <el-form-item>
      <el-button type="primary" @click="handleConfirm">
        确认
      </el-button>
      <el-button @click="handleCancel">取消</el-button>
    </el-form-item>
  </el-form>
</div>
</template>
<script setup>
import {reactive,defineExpose,defineEmits} from "vue";
import {ElMessage} from "element-plus";
import {changePassword} from "@/api/user.js";
import {useRouter} from "vue-router";
const router=useRouter()
const info=reactive({
  password:'',
  againPassword:''
})

const emits=defineEmits(['toClose'])

const handleCancel=()=>{emits('toClose')}
const handleConfirm=()=>{
  if (info.password===""||info.againPassword==="")return
  if(info.password!==info.againPassword){
    ElMessage.error("两次密码不一致！")
  }
  changePassword(info).then(response=>{
    if(response.code===200){
      router.push('/login')
      sessionStorage.removeItem('token')
    }else{
      ElMessage.error(response.msg)
    }
  })
}
const notifyClosed=()=>{
  info.password=''
  info.againPassword=''
}
defineExpose({notifyClosed})
</script>
<style scoped>
.wrapper-cg{
  width: 100%;
}
</style>