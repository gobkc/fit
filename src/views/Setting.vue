<script setup lang="ts">
import {reactive, getCurrentInstance} from 'vue'
import {Select} from '@element-plus/icons-vue'
import type {AppConfig, ListConfigurationsResponse, Response} from "../api/interface"
import {ElMessage} from "element-plus";

const {proxy}: any = getCurrentInstance();
const {api, $router} = proxy;

const initialAppConfig: AppConfig = {
  "name": "",
  "version": "v1/api",
  "debug": true,
  "rest_addr": ":5555",
  "dsn": "",
  "email": {
    "imap": "imap.qq.com:993",
    "smtp": "smtp.qq.com:587",
    "user": "",
    "pass": ""
  },
  "cors": {
    "enabled": true,
    "max_age": 10000000,
    "allowed_origins": [
      "*"
    ],
    "allowed_methods": [
      "GET",
      "POST",
      "PUT",
      "PATCH",
      "DELETE",
      "HEAD",
      "OPTIONS"
    ],
    "allowed_headers": [
      "*",
      "Authorization"
    ],
    "allow_credentials": true
  },
  "max_idle": 0,
  "max_conn": 0,
  "max_left_time": 0,
  "jwt_salt": ""
}

const newConfig: AppConfig = {
  "name": "",
  "version": "v1/api",
  "debug": true,
  "rest_addr": ":5555",
  "dsn": "",
  "email": {
    "imap": "imap.qq.com:993",
    "smtp": "smtp.qq.com:587",
    "user": "",
    "pass": ""
  },
  "cors": {
    "enabled": true,
    "max_age": 10000000,
    "allowed_origins": [
      "*"
    ],
    "allowed_methods": [
      "GET",
      "POST",
      "PUT",
      "PATCH",
      "DELETE",
      "HEAD",
      "OPTIONS"
    ],
    "allowed_headers": [
      "*",
      "Authorization"
    ],
    "allow_credentials": true
  },
  "max_idle": 0,
  "max_conn": 0,
  "max_left_time": 0,
  "jwt_salt": ""
}

const settings: AppConfig = reactive(initialAppConfig)

const configurations: AppConfig[] = reactive([])

const app = reactive({
  name: ``,
  create_mode: false,
})

const listConfigurations = () => {
  api.listConfigurations().then((response: ListConfigurationsResponse) => {
    configurations.splice(0, configurations.length);
    Object.assign(configurations, response.data);
    if (response.main_conf) {
      settings.name = response.main_conf
      change_conf(settings.name)
    }
    app.name = response.main_conf
    app.create_mode = settings.name == ``
  })
}

listConfigurations()

const change_conf = (config_name: string) => {
  const foundConfig = configurations.find(config => config.name === config_name);
  if (foundConfig) {
    Object.assign(settings, foundConfig);
    app.create_mode = false
  } else {
    newConfig.name = ''
    newConfig.email.user = ''
    newConfig.email.pass = ''
    newConfig.email.imap = 'imap.qq.com:993'
    newConfig.email.smtp = 'smtp.qq.com:587'
    newConfig.version = `v1/api`
    newConfig.rest_addr = `:5555`
    Object.assign(settings, newConfig);
    app.create_mode = true
  }
  app.create_mode = settings.name == ``
}

const deleteConfiguration = (config_name: string) => {
  api.deleteConfigurations(config_name).then((r: Response) => {
    newConfig.name = ''
    newConfig.email.user = ''
    newConfig.email.pass = ''
    Object.assign(settings, newConfig);
    listConfigurations()
    ElMessage.warning(r.msg)
  })
}

const upsert = () => {
  if (app.create_mode) {
    //create a new configuration
    api.createConfiguration(settings).then((r: Response) => {
      listConfigurations()
    })
  } else {
    //enable configuration
    api.enableConfiguration(settings).then((r: Response) => {
      listConfigurations()
    })
  }
}

</script>

<template>
  <el-container>
    <el-header>
      <el-page-header @back="$router.push('/')" style="display: flex; align-items: center; vertical-align: middle">
        <template #content>
          <span class="text-large font-600 mr-3" style="color: white;"> Setting </span>
        </template>
      </el-page-header>
    </el-header>
    <el-main>
      <div class="center-container">
        <el-form :model="settings" label-width="140px">
          <el-form-item label="configuration">
            <el-select v-model="settings.name" v-on:change="change_conf" placeholder="Create a new configuration">
              <el-option
                  key="Create a new configuration"
                  label="Create a new configuration"
                  value=""
              />
              <el-option
                  v-for="item in configurations"
                  :key="item.name"
                  :label="item.name"
                  :value="item.name"
              />
            </el-select>
          </el-form-item>
          <el-form-item label="Configuration Name">
            <el-input v-model="settings.name" :suffix-icon="app.name==settings.name?Select:''" class="green_icon"/>
          </el-form-item>
          <el-form-item label="Api Version">
            <el-input v-model="settings.version"/>
          </el-form-item>
          <el-form-item label="Rest Address">
            <el-input v-model="settings.rest_addr"/>
          </el-form-item>
          <el-form-item label="Imap">
            <el-input v-model="settings.email.imap"/>
          </el-form-item>
          <el-form-item label="Smtp">
            <el-input v-model="settings.email.smtp"/>
          </el-form-item>
          <el-form-item label="User">
            <el-input v-model="settings.email.user"/>
          </el-form-item>
          <el-form-item label="Password">
            <el-input v-model="settings.email.pass" type="password" show-password/>
          </el-form-item>
          <el-form-item>
            <el-button :type="app.create_mode?`primary`:`success`" v-on:click="upsert">
              {{ app.create_mode ? `Create Configuration` : `Enable Configuration` }}
            </el-button>
            <el-button type="danger" v-on:click="deleteConfiguration(settings.name)">Delete</el-button>
          </el-form-item>
        </el-form>
      </div>
    </el-main>
  </el-container>
</template>

<style scoped>

</style>