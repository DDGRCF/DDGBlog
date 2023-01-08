import { createApp } from "vue";
import App from "./App.vue";
import router from "./router";
import ElementPlus from "element-plus";
import * as ElementPlusIconsVue from "@element-plus/icons-vue";
import "element-plus/theme-chalk/display.css";
import "element-plus/dist/index.css";
import store from "./store";
// import mitt from "mitt";

const app = createApp(App);
for (const [key, component] of Object.entries(ElementPlusIconsVue)) {
  app.component(key, component);
}

app.use(router);
app.use(store);

app.use(ElementPlus);
// app.config.globalProperties.$bus = mitt;
app.mount("#app");
