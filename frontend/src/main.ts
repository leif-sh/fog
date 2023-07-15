import { createApp } from 'vue'
import App from './App.vue'
import { store, key } from './store'
import router from "./router";
import service from "./utils/https";
import urls from "./utils/urls";
import mixin from "./mixins";
import {
    ElMenu,
    ElMenuItem,
    ElMessage,
    ElLoading,
    ElButton,
    ElDialog,
    ElForm,
    ElFormItem,
    ElInput,
    ElRow,
    ElCol,
    ElDropdownMenu,
    ElTimeline,
    ElTimelineItem,
    ElDropdownItem,
    ElDropdown,
    ElCard,
    ElTag,
    ElIcon,
    ElCollapseTransition,
} from 'element-plus';

const app = createApp(App)
// app.mixin(mixin);
app.component(ElButton.name, ElButton);
app.component(ElDialog.name, ElDialog);
app.component(ElForm.name, ElForm);
app.component(ElFormItem.name, ElFormItem);
app.component(ElInput.name, ElInput);
app.use(ElMessage);
app.component(ElMenu.name, ElMenu);
app.component(ElMenuItem.name, ElMenuItem);
app.component(ElRow.name, ElRow);
app.component(ElCol.name, ElCol);
app.component(ElTimeline.name, ElTimeline);
app.component(ElTimelineItem.name, ElTimelineItem);
app.component(ElDropdownMenu.name, ElDropdownMenu);
app.component(ElDropdownItem.name, ElDropdownItem);
app.component(ElDropdown.name, ElDropdown);
app.component(ElCard.name, ElCard);
app.component(ElTag.name, ElTag);
app.component(ElIcon.name, ElIcon);
app.component(ElCollapseTransition.name, ElCollapseTransition);
app.use(ElLoading)

app.config.globalProperties.$message = ElMessage;
app.config.globalProperties.$loading = ElLoading.service;
app.config.globalProperties.$https = service;
app.config.globalProperties.$urls = urls;

app.use(store, key)
app.use(router)
app.mount('#app');
