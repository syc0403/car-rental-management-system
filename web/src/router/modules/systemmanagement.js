import Layout from '@/layout'

const systemmanagement = {
    path: '/systemmanagement',
    component: Layout,
    redirect: '/systemmanagement',
    name: 'systemmanagement',
    meta: {
        title: '系统管理',
        icon: 'nested'
    },
    children: [
        {
            path: 'usermanagement',
            component: () => import('@/views/systemmanagement/usermanagement/index'), // Parent router-view
            name: '用户管理',
            meta: { title: '用户管理'},
            
          },
          {
            path: 'rolemanagement',
            component: () => import('@/views/systemmanagement/rolemanagement/index'), // Parent router-view
            name: '角色管理',
            meta: { title: '角色管理'},
            
          },
    ]
}

export default systemmanagement