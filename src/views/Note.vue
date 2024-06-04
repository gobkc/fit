<script setup lang="ts">
import {Plus, Delete, Search, Setting} from '@element-plus/icons-vue'
import {MdEditor} from 'md-editor-v3';
import 'md-editor-v3/lib/style.css';
import {reactive, getCurrentInstance} from 'vue'
import type {
  ListCateResponse,
  ListConfigurationsResponse,
  ListNoteResponse, Note,
  Response
} from "../api/interface"

//save note event
document.onkeydown = (e: any) => {
  if (e.ctrlKey && e.key == `s`) {
    if (note.current_cate == ``) {
      return
    }
    api.newNote({
      cate: note.current_cate,
      content: note.content,
      title: note.title
    }).then((response: Response) => {
      list_note_titles(note.current_cate, note.keyword)
      console.log(response)
    })
  }
}

const aside = reactive({
  left_width: 250,
  center_width: 250,
})

const note = reactive({
  title: ``,
  update_at: `0000-00-00 00:00`,
  keyword: ``,
  content: ``,
  current_cate: ``,
  pull_tips: ``,
  push_tips: ``,
  categories: [] as string[],
  notes: [] as Note[]
})

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
  list_note_titles(note.current_cate, note.keyword)
}

const list_note_titles = (cate: string, keyword: string) => {
  if (cate == ``) {
    api.listNote(keyword).then((response: ListNoteResponse) => {
      note.notes = response.data
      setting_first_note()
    })
  } else {
    api.listCateNote(cate, keyword).then((response: ListNoteResponse) => {
      note.notes = response.data
      setting_first_note()
    })
  }
}

const setting_first_note = () => {
  if (note.notes.length > 0) {
    let current_note = note.notes[0]
    note.title = current_note.title
    note.content = current_note.content
    note.current_cate = current_note.cate
    note.update_at = current_note.updated_time
  }
}

list_note_titles(note.current_cate, note.keyword)

const select_note = (item: Note) => {
  note.current_cate = item.cate
  note.title = item.title
  note.content = item.content
  note.update_at = item.updated_time
}

const create_new_note = () => {
  note.title = ``
  note.content = ``
  note.update_at = `0000-00-00 00:00`
}

const delete_note = (title: string) => {
  if (confirm('Are you sure you want to delete this note?')) {
    api.deleteNote(note.current_cate, title).then((response: Response) => {
      list_note_titles(note.current_cate, note.keyword)
    })
  }
}

const create_new_cate = () => {
  const input = prompt('Please enter a new categories:', '');
  if (input) {
    api.newCate(input).then((response: Response) => {
      list_categories()
    })
  }
}

const delete_cate = (cate: string) => {
  if (confirm('Are you sure you want to delete this category?')) {
    api.deleteCate(cate).then((response: Response) => {
      list_categories()
      list_note_titles(note.current_cate, note.keyword)
    })
  }
}

const list_conf = () => {
  api.listConfigurations().then((response: ListConfigurationsResponse) => {
    if (response.main_conf == ``) {
      return
    }
    for (const datum of response.data) {
      if (datum.name == response.main_conf) {
        note.push_tips = datum.email.user
        note.pull_tips = datum.email.imap
      }
    }
  })
}
list_conf()

const push_notes = () => {
  api.push().then((r: Response) => {
    alert(r.msg)
  })
}

const pull_notes = () => {
  api.pull().then((r: Response) => {
    list_note_titles(``, note.keyword)
    alert(r.msg)
  })
}
</script>

<template>
  <el-header>
    <div class="header-left-box">
      <div class="logo">
        <el-image src="/assets/logo.svg" fit="fill"/>
      </div>
      <div class="search-box">
        <el-input
            v-model="note.keyword"
            style="max-width: 100%"
            placeholder="Please enter cate | title | content for fuzzy search"
            class="search-input"
            v-on:change="list_note_titles(``, note.keyword)"
        >
          <template #append>
            <el-button :icon="Search" v-on:click="list_note_titles(``, note.keyword)"/>
          </template>
        </el-input>
      </div>
    </div>
    <div class="header-right-box">
      <div class="fetch-box">
        <div class="pull high-light-hover" v-on:click="pull_notes">
          <div class="fetch-tips">{{ note.pull_tips }}</div>
          <div>Pull</div>
        </div>
        <div class="push high-light-hover" v-on:click="push_notes">
          <div class="fetch-tips">{{ note.push_tips }}</div>
          <div>Push</div>
        </div>
      </div>
      <div class="setting-box">
        <el-link :icon="Setting" :underline="false" v-on:click="$router.push('/setting')">Setting</el-link>
      </div>
    </div>
  </el-header>
  <el-container>
    <!--categories-->
    <el-aside :width="aside.left_width+'px'">
      <div class="item-box">
        <div class="item-title">
          Categories
          <div class="btn-plus" v-on:click="create_new_cate">
            <el-icon>
              <Plus/>
            </el-icon>
          </div>
        </div>
        <div class="item-content scrollable">
          <div :class="{item:true, 'high-light':item==note.current_cate}" v-for="(item,index) of note.categories"
               :key="index" v-on:click="select_cate(item)"
               :title="item">
            {{ item }}
            <div class="remove-item" v-on:click="delete_cate(item)">
              <el-icon>
                <Delete/>
              </el-icon>
            </div>
          </div>
        </div>
      </div>
    </el-aside>
    <!--note titles-->
    <el-aside :width="aside.center_width+'px'">
      <div class="item-box">
        <div class="item-title">
          {{ note.current_cate ? note.current_cate.toUpperCase() + ' Titles' : 'Titles' }}
          <div class="btn-plus" v-on:click="create_new_note">
            <el-icon>
              <Plus/>
            </el-icon>
          </div>
        </div>
        <div class="item-content scrollable">
          <div :class="{item:true, 'high-light':item.title==note.title}" v-for="(item,index) of note.notes" :key="index"
               :title="item.title"
               v-on:click="select_note(item)">
            {{ item.title }}
            <div class="remove-item" v-on:click="delete_note(item.title)">
              <el-icon>
                <Delete/>
              </el-icon>
            </div>
          </div>
        </div>
      </div>
    </el-aside>
    <!--note contents-->
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
                :disabled="note.current_cate==``"
                type="text"
            />
          </div>
          <div class="note-uptime" style="min-width: 120px">{{ note.update_at }}</div>
        </div>
        <div class="note-content-box">
          <MdEditor v-model="note.content" theme="dark" :codeFoldable="false" style="height: 100%"/>
        </div>
      </div>
    </el-main>
  </el-container>
</template>

<style scoped>
.header-left-box {
  width: 480px;
  display: flex;
}

.header-right-box {
  flex: 1;
  display: flex;
}

.fetch-box {
  width: 400px;
  display: flex;
}

.pull {
  flex: 1;
  height: 60px;
  line-height: 20px;
  display: flex;
  justify-content: center;
  align-items: center;
  border-left: 1px solid black;
  border-right: 1px solid black;
  cursor: pointer;
  user-select: none;
  font-size: 13px;
  line-height: 20px;
  flex-direction: column;
  color: rgba(255, 255, 255, 0.7);
}

.push {
  flex: 1;
  height: 60px;
  line-height: 20px;
  border-right: 1px solid black;
  cursor: pointer;
  user-select: none;
  font-size: 13px;
  display: flex;
  justify-content: center;
  align-items: center;
  flex-direction: column;
  color: rgba(255, 255, 255, 0.7);
}

.fetch-tips {
  font-size: 12px;
  color: rgba(255, 255, 255, 0.5);
}

.setting-box {
  flex: 1;
  line-height: 60px;
  text-align: right;
}

.logo {
  width: 60px;
  height: 60px;
  display: inline-block;
}

.search-box {
  padding-left: 20px;
  padding-right: 20px;
  flex: 1;
}

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
  height: 45px;
  line-height: 45px;
  padding: 0px 15px 0px 15px;
  position: relative;
  color: rgba(255, 255, 255, 0.7);
  cursor: pointer;
  user-select: none;
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
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

.item:hover, .high-light-hover:hover, .high-light {
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