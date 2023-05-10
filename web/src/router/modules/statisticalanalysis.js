import Layout from '@/layout'

const statisticalanalysis = {
    path: '/statisticalanalysis',
    component: Layout,
    redirect: '/statisticalanalysis',
    name: 'statisticalanalysis',
    meta: {
        title: '统计分析',
        icon: 'nested'
    },
    children: [
        {
            path: 'customerregionstatistics',
            component: () => import('@/views/statisticalanalysis/customerregionstatistics/index'), // Parent router-view
            name: '客户地区统计',
            meta: { title: '客户地区统计'},
            
          },
          {
            path: 'customeroccupation',
            component: () => import('@/views/statisticalanalysis/customeroccupation/index'), // Parent router-view
            name: '客户职业统计',
            meta: { title: '客户职业统计'},
            
          },
          {
            path: 'usersalesvolume',
            component: () => import('@/views/statisticalanalysis/usersalesvolume/index'), // Parent router-view
            name: '员工销售额统计',
            meta: { title: '员工销售额统计'},
            
          },
    ]
}

export default statisticalanalysis