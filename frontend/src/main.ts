import { createApp } from 'vue'
import App from './App.vue'
import { store, key } from './store'
import router from "./router";
import service from "./utils/https";
import urls from "./utils/urls";
import mixin from "./mixins";
import {
    ElMessage,
    ElLoading,
} from 'element-plus';

const app = createApp(App)
// app.mixin(mixin);


app.config.globalProperties.$message = ElMessage;
app.config.globalProperties.$loading = ElLoading.service;
// app.config.globalProperties.productionTip = false;
app.config.globalProperties.$https = service;
app.config.globalProperties.$urls = urls;

app.use(store, key)
app.use(router)
app.mount('#app');
