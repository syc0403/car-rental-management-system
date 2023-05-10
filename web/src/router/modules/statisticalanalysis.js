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
    ]
}

export default statisticalanalysis