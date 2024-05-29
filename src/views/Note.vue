<script setup lang="ts">
import {reactive} from "vue";
import {Plus, Delete} from '@element-plus/icons-vue'
import {MdEditor} from 'md-editor-v3';
import 'md-editor-v3/lib/style.css';

document.onmouseup = (e: any) => {
  console.log(e)
}

const aside = reactive({
  left_width: 250,
  center_width: 250,
})

const note = reactive({
  title: ``,
  update_at: `2024-01-02 03:04`,
  content: ``,
  current_cate: ``,
  categories: []
})

import {reactive, getCurrentInstance} from 'vue'
import type {AppConfig, ListCateResponse, ListConfigurationsResponse, Response} from "../api/interface"

const {proxy}: any = getCurrentInstance();
const {api, $router} = proxy;

const list_categories = () => {
  api.listCate().then((response: ListCateResponse) => {
    note.categories = response.data
  })
}

list_categories()

const select_cate = (cate: string) => {
  note.current_cate = cate
}

</script>

<template>
  <el-header>
    <el-image style="width: 60px; height: 60px" src="/logo.svg" fit="fill"/>
  </el-header>
  <el-container>
    <el-aside :width="aside.left_width+'px'">
      <div class="item-box">
        <div class="item-title">
          Categories
          <div class="btn-plus">
            <el-icon>
              <Plus/>
            </el-icon>
          </div>
        </div>
        <div class="item-content scrollable">
          <div class="item" v-for="(item,index) of note.categories" :key="index" v-on:click="select_cate(item)">
            {{ item }}
            <div class="remove-item">
              <el-icon>
                <Delete/>
              </el-icon>
            </div>
          </div>
        </div>
      </div>
    </el-aside>
    <el-aside :width="aside.center_width+'px'">
      <div class="item-box">
        <div class="item-title">
          {{note.current_cate?note.current_cate+' Titles':'Titles'}}
          <div class="btn-plus">
            <el-icon>
              <Plus/>
            </el-icon>
          </div>
        </div>
        <div class="item-content scrollable">
          <div class="item">
            item 1
            <div class="remove-item">
              <el-icon>
                <Delete/>
              </el-icon>
            </div>
          </div>
        </div>
      </div>
    </el-aside>
    <el-main class="scrollable">
      <div class="note-box">
        <div class="note-title-box">
          <div class="note-title">
            <el-input
                v-model="note.title"
                style="min-width: 400px"
                maxlength="100"
                placeholder="Please enter a note title"
                show-word-limit
                type="text"
            />
          </div>
          <div class="note-uptime" style="min-width: 120px">{{ note.update_at }}</div>
        </div>
        <div class="note-content-box">
          <MdEditor v-model="note.content" theme="dark" style="height: 100%"/>
        </div>
      </div>
    </el-main>
  </el-container>
</template>

<style scoped>
.item-box {
  display: flex;
  flex-direction: column;
  height: 100%;
}

.item-title {
  height: 35px;
  line-height: 35px;
  background-color: rgba(0, 0, 0, 0.25);
  border-bottom: 1px solid black;
  text-align: center;
  position: relative;
  font-weight: bold;
  user-select: none;
}

.item-content {
  flex: 1;
  background-color: rgba(0, 0, 0, 0.15);
}

.btn-plus {
  display: inline-block;
  width: 35px;
  position: absolute;
  right: 5px;
  user-select: none;
  cursor: pointer;
}

.item {
  height: 35px;
  line-height: 35px;
  padding: 0px 15px 0px 15px;
  position: relative;
  color: rgba(255, 255, 255, 0.7);
  overflow-x: hidden;
  cursor: pointer;
  user-select: none;
}

.remove-item {
  position: absolute;
  right: 5px;
  width: 35px;
  display: none;
  text-align: center;
  cursor: pointer;
}

.item:hover .remove-item {
  display: inline-block;
}

.item:hover {
  background-color: rgba(255, 255, 255, 0.05);
  color: rgba(255, 255, 255, 1);
}

.note-box {
  display: flex;
  flex-direction: column;
  height: 100%;
  font-size: 13px;
}

.note-title-box {
  height: 35px;
  line-height: 35px;
  border-bottom: 1px solid black;
  text-align: center;
  position: relative;
  font-weight: bold;
  user-select: none;
  display: flex;
}

.note-content-box {
  flex: 1;
}

.note-title {
  flex: 1;
  text-align: left;
}

.note-uptime {
  width: 300px;
  text-align: right;
  color: rgba(255, 255, 255, 0.55);
}
</style>