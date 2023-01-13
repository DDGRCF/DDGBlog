import { createRouter, createWebHistory, RouteRecordRaw } from "vue-router";
import ApiView from "../views/ApiView.vue";
import AdmView from "../views/AdmView.vue";

const routes: Array<RouteRecordRaw> = [
  {
    path: "/",
    name: "root",
    component: AdmView,
  },
  {
    path: "/api",
    name: "api",
    component: ApiView,
    children: [
      {
        path: "home",
        name: "home",
        component: () => import("@/views/home/HomeView.vue"),
      },
      {
        path: "tool",
        name: "tool",
        component: () => import("@/views/tool/ToolView.vue"),
        children: [
          {
            path: "obb/nms",
            name: "nms",
            component: () => import("@/views/tool/obb/NmsView.vue"),
          },
          {
            path: "obb/submit",
            name: "submit",
            component: () => import("@/views/tool/obb/SubmitView.vue"),
          },
        ],
      },
    ],
  },
];

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes,
});

export default router;
