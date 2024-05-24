import './assets/main.css'

import { createApp } from 'vue'
import { createPinia } from 'pinia'
import {Api} from './api/api'

import App from './App.vue'
import router from './router'

const app = createApp(App)

const api = new Api()

app.use(createPinia())
app.use(router)
app.config.globalProperties.api = api

app.mount('#app')
