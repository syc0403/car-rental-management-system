import Layout from '@/layout'

const basemanagementRouter = {
    path: '/basemanagement',
    component: Layout,
    redirect: '/basemanagement',
    name: 'basemanagement',
    meta: {
        title: '基础管理',
        icon: 'nested'
    },
    children: [
        {
          path: 'customermanagement',
          component: () => import('@/views/basemanagement/customermanagement/index'), // Parent router-view
          name: '客户管理',
          meta: { title: '客户管理' },
        },
        {
          path: 'carmanagement',
          name: '车辆管理',
          component: () => import('@/views/basemanagement/carmanagement/index'),
          meta: { title: '车辆管理' }
        }
      ]
}

export default basemanagementRouter