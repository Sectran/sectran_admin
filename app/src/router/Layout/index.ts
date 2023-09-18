//配置管理
let Layout = [
    {
        path: '/admin',
        component: import('views/Layout/index.vue'),
        name:'admin',
        meta: {
            // title: '配置管理',
            title: 'configuration',
        },
        children: [
            {
                path: 'user',
                name: 'user',
                component: () => import('views/Table/user.vue'),
                meta: {
                    // title: '人员管理',
                    title: 'userAdmin',
                },
            },
            {
                path: 'department',
                name: 'department',
                component: () => import('views/Table/department.vue'),
                meta: {
                    // title: '部门管理',
                    title: 'departmentAdmin',
                },
            },

        ]
    }]

export default Layout;

