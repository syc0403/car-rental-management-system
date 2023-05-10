<template>
  <div class="contentr">
    <el-main>
    <wk-charts className="echary" :optionData="option" class="echary" ref="echary"></wk-charts>
    </el-main>
  </div>
</template>
   
<script>
import wkCharts from '@/components/Charts/wkcharts'
import { get, post } from "@/utils/requests";
export default {
  name: "dataEcharts",
  data() {
    return {
      option: {}
    }
  },
  components: {
    wkCharts
  },
  created() {
  },
  async mounted() {
    const {data}  = await get("/rentOrder/getmoneybyuser")
    var addressData = []
    var countData = []
    for (let i = 0; i < data.data.data.length; i++) {
      addressData.push(data.data.data[i].oper_name)
      countData.push(data.data.data[i].money)
    }
    let dataName = ['金额']
    let option = JSON.parse(JSON.stringify(this.$refs.echary.getData))
    let series = []
    option.legend.data = dataName
    option.grid.top = '15%'
    option.xAxis[0].data = addressData
    dataName.map(res => {
        series.push({ name: res, type: 'bar', data: countData })
    })
    option.series = series
    this.option = JSON.parse(JSON.stringify(option))
    // 示例数据
    let optionDatas = {
      tooltip: {
        trigger: 'axis',
        axisPointer: {
          type: 'cross',
          crossStyle: {
            color: '#999'
          }
        }
      },
      // 工具栏
      toolbox: {
        feature: {
          // 显示表格数据
          dataView: { show: true, readOnly: false },
          // 切换为 k线图 或柱形图
          magicType: { show: true, type: ['line', 'bar'] },
          // 刷新
          restore: { show: true },
          // 保存为本地图片
          saveAsImage: { show: true }
        }
      },
      // 说明 数据与 series 的name值对应 才会显示
      legend: {
        data: ['蒸发量', '降水量', '温度', '光照', '平均温度']
      },
      xAxis: [
        // 设置x轴底部显示的标题数据 及样式
        {
          type: 'category',
          data: ['1月', '2月', '3月', '4月', '5月', '6月', '7月', '8月', '9月', '10月'],
          axisPointer: {
            type: 'shadow'
          },
          axisLine: {
            lineStyle: {
              color: 'red'
            }
          }
        }
      ],
      yAxis: [
        // 显示样式设置  此处显示两个 y轴  看个人需求 是一个就写一个
        {
          type: 'value',
          name: '水量',
          min: 0,
          max: 250,
          interval: 50,
          axisLabel: {
            formatter: '{value} ml'
          },
          axisLine: {
            lineStyle: {
              color: 'red'
            }
          }
        },
        {
          type: 'value',
          name: '温度',
          min: 0,
          max: 25,
          interval: 5,
          axisLabel: {
            formatter: '{value} °C'
          },
          axisLine: {
            lineStyle: {
              color: 'blue'
            }
          }
        }
      ],
      series: [
        // bar 为柱形图 一个隔断 分为四个柱形  外加一个line K线图  数据全放在series
        {
          name: '蒸发量',
          type: 'bar',
          data: [2.0, 4.9, 7.0, 23.2, 25.6, 76.7, 135.6, 162.2, 32.6, 20.0],
          stack: '同', // 名称相同 数据呈现在同一柱形图上 上下呈现
          color: 'red',
          itemStyle: {
            normal: {
              label: {
                show: false, //开启显示数据图上的数值
                position: 'top', //在上方显示
                textStyle: { //数值样式
                  color: 'red',
                  fontSize: 12
                }
              }
            }
          }
        },
        {
          name: '降雨量',
          type: 'bar',
          data: [2.6, 5.9, 9.0, 26.4, 28.7, 70.7, 175.6, 182.2, 48.7, 18.8],
          color: 'pink',
          itemStyle: {
            normal: {
              label: {
                show: false, //开启显示数据图上的数值
                position: 'top', //在上方显示
                textStyle: { //数值样式
                  color: 'pink',
                  fontSize: 12
                }
              }
            }
          }
        },
        {
          name: '温度',
          type: 'bar',
          data: [2.6, 2.9, 1.0, 26.4, 28.7, 70.7, 15.6, 12.2, 2.7, 38.8],
          color: 'blue',
          itemStyle: {
            normal: {
              label: {
                show: false, //开启显示数据图上的数值
                position: 'top', //在上方显示
                textStyle: { //数值样式
                  color: 'blue',
                  fontSize: 12
                }
              }
            }
          }
        },
        {
          name: '光照',
          type: 'bar',
          data: [2.6, 5.9, 9.0, 26.4, 28.7, 20.7, 11.6, 112.2, 48.7, 18.8],
          itemStyle: {
            normal: {
              label: {
                show: false, //开启显示数据图上的数值
                position: 'top', //在上方显示
                textStyle: { //数值样式
                  color: 'black',
                  fontSize: 12
                }
              }
            }
          }
        },
        {
          name: '平均温度',
          type: 'line',
          yAxisIndex: 0,
          data: [2.0, 2.2, 3.3, 4.5, 6.3, 10.2, 20.3, 23.4, 23.0, 16.5, 12.0, 10],
          color: 'rgb(64, 158, 255)',
          itemStyle: {
            normal: {
              label: {
                show: false, //开启显示数据图上文字
                position: 'top', //在上方显示
                textStyle: { //数值样式
                  color: 'rgb(64, 158, 255)',
                  fontSize: 12
                }
              }
            }
          }
        }
      ]
    }
  },
  methods: {
  }
}
</script>
   
<style lang="scss" scoped></style>