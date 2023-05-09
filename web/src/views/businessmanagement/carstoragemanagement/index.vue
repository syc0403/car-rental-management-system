<!--  -->
<template>
    <div class=''>
        <el-header>
            <hr class="hr">
            <div style="display: flex;">
                <div style="display: flex;margin:10px;font-size: 18px;">订单号：</div>
                <div style="width: 40">
                    <el-input placeholder="请输入订单号" prefix-icon="el-icon-search" v-model="searchIdentity">
                    </el-input>
                </div>
                <el-button class="btns" type="primary" icon="el-icon-search" style="margin-left: 10px"
                    @click="searchData()">
                    查询
                </el-button>
            </div>
        </el-header>
        <hr class="hr1">
        <div style="margin-left: 2%;">
            <div>
                <el-row :gutter="24">
                    <el-col :span="8">检查单号：<el-input v-model="checkList.checkIdentity" placeholder="请输入内容"
                            style="width: 70%;"></el-input></el-col>
                    <el-col :span="8">出租单号：<el-input v-model="checkList.rentOrderIdentity" placeholder="请输入内容"
                            style="width: 70%;"></el-input></el-col>
                    <el-col :span="8">客户名称：<el-input v-model="checkList.customerName" placeholder="请输入内容"
                            style="width: 70%;"></el-input></el-col>
                </el-row>
            </div>
            <div style="margin-top: 20px;">
                <el-row :gutter="24">
                    <el-col :span="8">存在问题：<el-input v-model="checkList.problem" placeholder="请输入内容"
                            style="width: 70%;"></el-input></el-col>
                    <el-col :span="8">赔付金额：<el-input v-model="checkList.payPrice" placeholder="请输入内容"
                            style="width: 70%;"></el-input></el-col>
                    <el-col :span="8">检查时间：<el-input v-model="checkList.checkTime" placeholder="请输入内容"
                            style="width: 70%;"></el-input></el-col>
                </el-row>
            </div>
            <div style="margin-top: 20px;">
                <div style="float: left;">问题描述：</div>
                <el-input v-model="checkList.problemDesc" type="textarea" :rows="3" placeholder="请输入内容" style="width: 90%;">
                </el-input>
            </div>
            <div style="width: 100%;">
                <el-button type="primary" icon="el-icon-s-promotion" @click="dialogAdd = true" round
                    style="margin: 15px 90%;">保存</el-button>
            </div>
        </div>
        <el-main>
            <el-row :gutter="24">
                <el-col :span="8"><el-card class="box-card">
                        <div slot="header" class="clearfix">
                            <span>车辆信息</span>
                        </div>
                        <div style="margin-bottom: 5px;">
                            车牌号码：{{ carData.car_number }}
                        </div>
                        <div style="margin-bottom: 5px;">
                            车辆类型：{{ carData.car_type }}
                        </div>
                        <div style="margin-bottom: 5px;">
                            车辆颜色：{{ carData.color }}
                        </div>
                        <div style="margin-bottom: 5px;">
                            购买价格：{{ carData.price }}
                        </div>
                        <div style="margin-bottom: 5px;">
                            出租价格：{{ carData.rent_price }}
                        </div>
                        <div style="margin-bottom: 5px;">
                            出租押金：{{ carData.deposit }}
                        </div>
                    </el-card></el-col>
                <el-col :span="8"><el-card class="box-card">
                        <div slot="header" class="clearfix">
                            <span>出租单信息</span>
                        </div>
                        <div style="margin-bottom: 5px;">
                            出租单号：{{ rentOrderData.identity }}
                        </div>
                        <div style="margin-bottom: 5px;">
                            出租价格：{{ rentOrderData.order_price }}
                        </div>
                        <div style="margin-bottom: 5px;">
                            起租时间：{{ rentOrderData.begin_date }}
                        </div>
                        <div style="margin-bottom: 5px;">
                            还车时间：{{ rentOrderData.return_date }}
                        </div>
                        <div style="margin-bottom: 5px;">
                            客户姓名：{{ customerData.customer_name }}
                        </div>
                    </el-card></el-col> <el-col :span="8"><el-card class="box-card">
                        <div slot="header" class="clearfix">
                            <span>客户信息</span>
                        </div>
                        <div style="margin-bottom: 5px;">
                            身份证号：{{ customerData.identity }}
                        </div>
                        <div style="margin-bottom: 5px;">
                            客户姓名：{{ customerData.customer_name }}
                        </div>
                        <div style="margin-bottom: 5px;">
                            客户地址：{{ customerData.address }}
                        </div>
                        <div style="margin-bottom: 5px;">
                            客户电话：{{ customerData.phone }}
                        </div>
                        <div style="margin-bottom: 5px;">
                            客户职位：{{ customerData.occupation }}
                        </div>
                    </el-card></el-col>
            </el-row>
        </el-main>

        <el-footer>
        </el-footer>


        <el-dialog title="提示" :visible.sync="dialogVisible" width="30%">
            <span> 您输入的出租单号相关车辆已经归还，无需再入库
            </span>
            <span slot="footer" class="dialog-footer">
                <el-button @click="dialogVisible = false">取 消</el-button>
                <el-button type="primary" @click="dialogVisible = false">确 定</el-button>
            </span>
        </el-dialog>
        <el-dialog title="提示" :visible.sync="dialogAdd" width="30%">
            <span> 请确定是否保存
            </span>
            <span slot="footer" class="dialog-footer">
                <el-button @click="dialogAdd = false">取 消</el-button>
                <el-button type="primary" @click="addData()">确 定</el-button>
            </span>
        </el-dialog>
    </div>
</template>

<script>
import { getNowTime, getTimeStamp } from "@/utils/utils";
import { get, post } from "@/utils/requests";
export default {
    components: {},
    data() {
        return {
            dialogVisible: false,
            dialogAdd: false,
            searchIdentity: "CZ_1131U960_1683472049",
            checkList: {
                checkIdentity: "",
                rentOrderIdentity: "",
                customerName: "",
                problem: "",
                problemDesc: "",
                payPrice: 0,
                checkTime: ""
            },
            customerData: {},
            rentOrderData: {},
            carData: {},
        };
    },
    computed: {},
    watch: {},
    methods: {
        async searchData() {
            const { data } = await post("/rentOrder/getrentorderby", { identity: this.searchIdentity })
            if (data.data.data[0].status == 2) {
                this.dialogVisible = true
            } else {
                this.rentOrderData = data.data.data[0]
                console.log(this.rentOrderData);
                this.getCustomerData(data.data.data[0].customer_id)
                this.getCarData(data.data.data[0].car_number)
                this.checkList.checkIdentity = "JC" + this.searchIdentity.slice(2, 12) + getTimeStamp();
                this.checkList.rentOrderIdentity = this.searchIdentity
                this.checkList.checkTime = getNowTime()
            }

        },
        async getCustomerData(id) {
            const { data } = await get("/customer/getcustomerinfobyid", { id: id })
            this.customerData = data.data
            this.checkList.customerName = data.data.customer_name
        },
        async getCarData(car_number) {
            const { data } = await post("/car/getcarinfobynumber", { car_number: car_number })
            this.carData = data.data.data[0]
        },
        async addData() {
            let params = {
                identity: this.checkList.checkIdentity,
                rentorder_identity: this.checkList.rentOrderIdentity,
                problem: this.checkList.problem,
                problem_desc: this.checkList.problemDesc,
                pay_price: this.checkList.payPrice,
                check_time: this.checkList.checkTime,
                oper_name: this.$store.state.user.name
            }
            console.log(params);
            const { data } = await post("/check/addcheck", params)
            if (data.code == 200) {
                this.updateRentOrderStatus()
                this.$notify({
                    message: "添加成功",
                    type: "success",
                });
            } else {
                this.$notify({
                    message: "添加失败",
                    type: "warning",
                });
            }
            console.log(this.checkList);
            this.dialogAdd = false
            location.reload()
        },
        async updateRentOrderStatus() {
            let params = {
                identity: this.checkList.rentOrderIdentity,
                status: 2
            }
            console.log(params);
            const { data } = await post("/check/updaterentorderstatus", params)
        }
    },
    created() {
    },
    mounted() {

    },
}
</script>
<style>
.hr {
    color: #000;
    border: 0;
    font-size: 20px;
    padding: 10px 0;
    position: relative;
}

.hr::before {
    content: "查询条件";
    position: absolute;
    padding: 0 10px;
    line-height: 1px;
    border: solid rgba(0, 51, 255, 0.491);
    border-width: 0 100vw;
    white-space: nowrap;
    left: 50%;
    transform: translateX(-50%);
}

.hr1 {
    color: #000;
    border: 0;
    font-size: 20px;
    padding: 10px 0;
    position: relative;
}

.hr1::before {
    content: "检查单表单";
    position: absolute;
    padding: 0 10px;
    line-height: 1px;
    border: solid rgba(0, 51, 255, 0.491);
    border-width: 0 100vw;
    white-space: nowrap;
    left: 50%;
    transform: translateX(-50%);
}

.lookup {
    display: flex;
    margin: 20px 20px;
}

.aInput {
    width: 70%;
}

.btns {
    margin-left: 10px;
}
</style>