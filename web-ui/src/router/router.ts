import { createRouter, createWebHashHistory, type RouteRecordRaw } from "vue-router";

const routes: Array<RouteRecordRaw> = [
    {
        path: "/",
        redirect: '/login',
    },
    {
        path: "/login",
        name: "login",
        component: () => import('../views/login/Login.vue'),
        meta: {
            title: '登录页',
            requiresAuth: false,
        }
    },
    {
        path: "/home",
        name: "home",
        component: () => import('../views/home/Home.vue'),
        meta: {
            title: '数据中心',
            requiresAuth: true,
        }
    }
];

const router = createRouter({
    history: createWebHashHistory(),
    routes
});

router.beforeEach((to, from, next) => {
    console.log(from);
    const isAuthenticated = localStorage.getItem('token') !== null;
    const requiresAuth = to.matched.some(record => record.meta.requiresAuth);

    if (requiresAuth && !isAuthenticated) {
        next('/login'); // 需要认证但未认证，重定向到登录页面
    } else {
        next(); // 其他情况下，继续访问请求的路由
    }
});

export default router;
