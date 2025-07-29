import { createRouter, createWebHistory } from 'vue-router'
import { useStore } from "@/stores/store";
import { message } from "ant-design-vue";

const router = createRouter({
    history: createWebHistory(import.meta.env.BASE_URL),
    routes: [
        {
            path: "/login",
            name: "login",
            component: () => import("../views/login.vue")
        },
        {
            path: "/admin",
            name: "index",
            component: () => import("../views/admin/admin.vue"),
            children: [
                {
                    path: "home",
                    name: "home",
                    component: () => import("../views/admin/home/home.vue")
                },
                {
                    path: "user_list",
                    name: "user_list",
                    component: () => import("../views/admin/user/user_list.vue")
                },
            ]
        }
    ]
})

export default router


