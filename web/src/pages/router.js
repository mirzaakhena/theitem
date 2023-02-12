import {createRouter, createWebHashHistory} from 'vue-router'

const routes = [

    {
        path: '/',
        component: () => import('./PageWithSidebar.vue'),
        children: [
            {
                path: '/payment',
                redirect: '/payment/waiting',
                component: () => import('./yourpage/ViewTab.vue'),
                children: [
                    {
                        path: '/payment/waiting',
                        component: () => import('../usecase/sample/Page1.vue'),
                    },
                    {
                        path: '/payment/processing',
                        component: () => import('../usecase/sample/Page2.vue'),
                    },
                ],
            },
            {
                path: '/order',
                component: () => import('../usecase/sample/Page3.vue'),
            },
            {
                path: '/item',
                component: () => import('./item/PageTable.vue'),
            },
        ],
    },

]

const router = createRouter({
    history: createWebHashHistory(),
    routes
})

export default router