<script setup lang="ts">
import {reactive, getCurrentInstance} from 'vue'
import type {AppConfig} from "../api/interface"

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

const settings: AppConfig = reactive(initialAppConfig)

const configurations: AppConfig[] = reactive([])

const conf = api.listConfigurations().then((appConfigs: AppConfig[]) => {
  Object.assign(configurations, appConfigs);
})

const change_conf = (config_name: string) => {
  const foundConfig = configurations.find(config => config.name === config_name);
  if (foundConfig) {
    Object.assign(settings, foundConfig);
  }
}

</script>

<template>
  <el-container>
    <el-header>
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
            <el-input v-model="settings.name"/>
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
            <el-input v-model="settings.email.pass"/>
          </el-form-item>
          <el-form-item>
            <el-button :type="settings.name==``?`primary`:`success`">
              {{ settings.name == `` ? `Create Configuration` : `Enable Configuration` }}
            </el-button>
            <el-button type="danger">Delete</el-button>
          </el-form-item>
        </el-form>
      </div>
    </el-main>
  </el-container>
</template>

<style scoped>

</style>