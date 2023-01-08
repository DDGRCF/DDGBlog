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
        path: "tools",
        name: "tools",
        component: () => import("@/views/tools/ToolsView.vue"),
        children: [
          {
            path: "obb/nms",
            name: "nms",
            component: () => import("@/views/tools/obb/NmsView.vue"),
          },
          {
            path: "obb/submit",
            name: "submit",
            component: () => import("@/views/tools/obb/SubmitView.vue"),
          },
        ],
      },
    ],
  },
  // {
  //   path: "/about",
  //   name: "about",
  //   // route level code-splitting
  //   // this generates a separate chunk (about.[hash].js) for this route
  //   // which is lazy-loaded when the route is visited.
  //   component: () =>
  //     import(/* webpackChunkName: "about" */ "../views/AboutView.vue"),
  // },
];

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes,
});

export default router;
