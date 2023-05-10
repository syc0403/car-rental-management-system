 
<template>
    <div :class="className" :style="{ height: height, width: width }" />
</template>
   
<script>
import * as echarts from 'echarts';
export default {
    props: {
        // class 为 当前图表的唯一标识
        className: {
            type: String,
            default: "chart"
        },
        width: {
            type: String,
            default: "100%"
        },
        height: {
            type: String,
            default: "350px"
        },
        // option 为图表数据 包括呈现的方式 数据
        optionData: {}
    },
    data() {
        return {
            chart: null,
            getData: {
                tooltip: {
                    trigger: 'axis',
                    axisPointer: {
                        type: 'cross',
                        crossStyle: {
                            color: '#000'
                        }
                    }
                },
                grid: {
                    left: '2%',
                    right: '2%',
                    bottom: '3%',
                    top: '10%',
                    containLabel: true
                },
                legend: {},
                xAxis: [
                    {
                        type: 'category',
                        data: ['1月', '2月', '3月', '4月', '5月', '6月', '7月', '8月', '9月', '10月', '11月', '12月'],
                        axisPointer: {
                            type: 'shadow'
                        }
                    }
                ],
                yAxis: [
                    {
                        type: 'value',
                        name: ' ',
                        min: 0,
                        axisLabel: {
                            formatter: '{value} '
                        }
                    },
                ],
                series: []
            }
        }
    },
    // 监听数据变化 进行试试刷新
    watch: {
        optionData(n) {
            if (n) {
                this.$nextTick(() => {
                    this.chart.setOption(this.optionData, true) //设置为true时不会合并数据，而是重新刷新数据
                })
            }
        }
    },
    mounted() {
        // 防止未加载完毕 报错
        this.$nextTick(() => {
            this.chart = echarts.init(this.$el, "macarons")
            this.chart.setOption(this.getData)
        });
    },
    // 关闭 及 删除图表
    beforeDestroy() {
        if (!this.chart) {
            return;
        }
        this.chart.dispose();
        this.chart = null;
    },
    methods: {
    }
};
</script>