<script setup lang="ts">
import {getCurrentInstance} from 'vue'

const {proxy}: any = getCurrentInstance();
const {api, $router, $route} = proxy;

const onBack = () => {
  localStorage.setItem('fit-fe', ``)
  $router.push('/')
}

setTimeout(onBack, 5 * 1000)
</script>

<template>
  <el-result
      icon="success"
      title="退出提示"
      sub-title="用户将在5秒内退出"
  >
    <template #extra>
      <el-button type="primary" v-on:click="onBack">马上退出</el-button>
    </template>
  </el-result>
</template>

<style scoped>

</style>