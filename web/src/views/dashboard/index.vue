<template>
  <div class="page">
    <div class="gva-card-box">
      <div class="gva-card gva-top-card">
        <div class="gva-top-card-left">
          <div class="gva-top-card-left-title">
            早安，{{ this.$store.state.user.name }}，请开始一天的工作吧
          </div>
          <div class="gva-top-card-left-dot"></div>
          <div class="gva-top-card-left-rows">
            <el-row>
              <el-col :span="8" :xs="24" :sm="8">
                <div class="flex-center">
                  <el-icon class="dashboard-icon">
                    <sort />
                  </el-icon>
                </div>
              </el-col>
              <el-col :span="8" :xs="24" :sm="8">
                <div class="flex-center">
                  <el-icon class="dashboard-icon">
                    <avatar />
                  </el-icon>
                </div>
              </el-col>
              <el-col :span="8" :xs="24" :sm="8">
                <div class="flex-center">
                  <el-icon class="dashboard-icon">
                    <comment />
                  </el-icon>
                </div>
              </el-col>
            </el-row>
          </div>
          <div>
            <div class="gva-top-card-left-item">
            </div>
            <div class="gva-top-card-left-item">
            </div>
          </div>
        </div>
        <img src="@/assets/dashboard.png" class="gva-top-card-right" alt />
      </div>
    </div>
    <div class="gva-card-box">
      <div class="gva-card">
        <div class="card-header">
          <span>数据统计</span>
        </div>
        <div class="echart-box">
          <el-row :gutter="20">
            <el-col :xs="24" :sm="18">
              <div class="contentr">
                <el-main>
                  <wk-charts className="echary" :optionData="option" class="echary" ref="echary"></wk-charts>
                  <div style="text-align:center;">租赁人数分布</div>
                </el-main>
              </div>

            </el-col>
            <el-col :xs="24" :sm="6">
              <el-card class="box-card">
                <div>
                  累计销售金额：{{ showData.money }}
                </div>
              </el-card>
            </el-col>
          </el-row>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import wkCharts from '@/components/Charts/wkcharts'
import { get, post } from "@/utils/requests";
export default {
  name: "Dashboard",
  data() {
    return {
      option: {},
      showData: {
        money: ""
      }
    }
  },
  components: {
    wkCharts
  },
  methods: {
    async getShowData() {
      const { data } = await get("/rentOrder/gettotalmoney")
      this.showData.money = data.data.money
    }
  },
  async mounted() {
    this.getShowData()
    const { data } = await get("/check/getadresscount")
    var addressData = []
    var countData = []
    for (let i = 0; i < data.data.data.length; i++) {
      addressData.push(data.data.data[i].address)
      countData.push(data.data.data[i].count)
    }
    console.log(countData);
    let dataName = ['人数']
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
};
</script>

<style lang="scss" scoped>
@mixin flex-center {
  display: flex;
  align-items: center;
}

.page {
  background: #f0f2f5;
  padding: 0;

  .gva-card-box {
    padding: 12px 16px;

    &+.gva-card-box {
      padding-top: 0px;
    }
  }

  .gva-card {
    box-sizing: border-box;
    background-color: #fff;
    border-radius: 2px;
    height: auto;
    padding: 26px 30px;
    overflow: hidden;
    box-shadow: 0 0 7px 1px rgba(0, 0, 0, 0.03);
  }

  .gva-top-card {
    height: 260px;
    @include flex-center;
    justify-content: space-between;
    color: #777;

    &-left {
      height: 100%;
      display: flex;
      flex-direction: column;

      &-title {
        font-size: 22px;
        color: #343844;
      }

      &-dot {
        font-size: 16px;
        color: #6b7687;
        margin-top: 24px;
      }

      &-rows {
        // margin-top: 15px;
        margin-top: 18px;
        color: #6b7687;
        width: 600px;
        align-items: center;
      }

      &-item {
        +.gva-top-card-left-item {
          margin-top: 24px;
        }

        margin-top: 14px;
      }
    }

    &-right {
      height: 600px;
      width: 600px;
      margin-top: 28px;
    }
  }

  ::v-deep(.el-card__header) {
    padding: 0;
    border-bottom: none;
  }

  .card-header {
    padding-bottom: 20px;
    border-bottom: 1px solid #e8e8e8;
  }

  .quick-entrance-title {
    height: 30px;
    font-size: 22px;
    color: #333;
    width: 100%;
    border-bottom: 1px solid #eee;
  }

  .quick-entrance-items {
    @include flex-center;
    justify-content: center;
    text-align: center;
    color: #333;

    .quick-entrance-item {
      padding: 16px 28px;
      margin-top: -16px;
      margin-bottom: -16px;
      border-radius: 4px;
      transition: all 0.2s;

      &:hover {
        box-shadow: 0px 0px 7px 0px rgba(217, 217, 217, 0.55);
      }

      cursor: pointer;
      height: auto;
      text-align: center;

      // align-items: center;
      &-icon {
        width: 50px;
        height: 50px !important;
        border-radius: 8px;
        @include flex-center;
        justify-content: center;
        margin: 0 auto;

        i {
          font-size: 24px;
        }
      }

      p {
        margin-top: 10px;
      }
    }
  }

  .echart-box {
    padding: 14px;
  }
}

.dashboard-icon {
  font-size: 20px;
  color: rgb(85, 160, 248);
  width: 30px;
  height: 30px;
  margin-right: 10px;
  @include flex-center;
}

.flex-center {
  @include flex-center;
}

//小屏幕不显示右侧，将登录框居中
@media (max-width: 750px) {
  .gva-card {
    padding: 20px 10px !important;

    .gva-top-card {
      height: auto;

      &-left {
        &-title {
          font-size: 20px !important;
        }

        &-rows {
          margin-top: 15px;
          align-items: center;
        }
      }

      &-right {
        display: none;
      }
    }

    .gva-middle-card {
      &-item {
        line-height: 20px;
      }
    }

    .dashboard-icon {
      font-size: 18px;
    }
  }
}
</style>
