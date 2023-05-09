import Layout from '@/layout'

const businessmanagement = {
    path: '/businessmanagementt',
    component: Layout,
    redirect: '/businessmanagement',
    name: 'businessmanagement',
    meta: {
        title: '业务管理',
        icon: 'nested'
    },
    children: [
        {
          path: 'carrentalmanagement',
          component: () => import('@/views/businessmanagement/carrentalmanagement/index'), // Parent router-view
          name: '>>汽车出租',
          meta: { title: '>>汽车出租'},
          
        },
        {
          path: 'rentordermanagement',
          name: '出租单管理',
          component: () => import('@/views/businessmanagement/rentordermanagement/index'),
          meta: { title: '出租单管理' }
        },
        {
          path: 'carstoragemanagement',
          component: () => import('@/views/businessmanagement/carstoragemanagement/index'), // Parent router-view
          name: '>>汽车入库',
          meta: { title: '>>汽车入库'},
          
        },
      ]
}

export default businessmanagement