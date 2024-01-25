<script setup>
import { RouterLink, RouterView } from 'vue-router'
import HelloWorld from './components/HelloWorld.vue'
import Header from "@/components/Header.vue";
import Footer from "@/components/Footer.vue";
</script>

<template>
  <div class="all">
    <el-container >
      <el-header >
        <Header/>
      </el-header>
      <el-main><RouterView /></el-main>
      <el-footer>
        <Footer/>
      </el-footer>
    </el-container>
  </div>

</template>

<style scoped>
.el-header{
  --el-header-padding:0;
  --el-header-height:100px;
}
.el-main{
  min-height: 100px;
  --el-main-padding:0;
}
.el-footer{
  --el-footer-padding:0;
  --el-footer-height:100px;
}

</style>
