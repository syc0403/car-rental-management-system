<!--  -->
<template>
    <div class=''>
        <el-header>
            <hr class="hr">
            <div style="display: flex;">
                <div style="display: flex;margin:10px;font-size: 18px;">身份证号：</div>
                <div style="width: 40">
                    <el-input placeholder="请输入身份证号" prefix-icon="el-icon-search" v-model="IdNumber">
                    </el-input>
                </div>
                <el-button class="btns" type="primary" icon="el-icon-search" style="margin-left: 10px" @click="searchCar()">
                    查询
                </el-button>
            </div>
        </el-header>
        <el-main>
            <el-table :data="tableData" border style="width: 100%">
                <el-table-column prop="car_number" label="车牌号" align="center" width="120">
                </el-table-column>
                <el-table-column prop="car_type" label="车辆类型" align="center" width="120">
                </el-table-column>
                <el-table-column prop="color" label="出租颜色" align="center" width="120">
                </el-table-column>
                <el-table-column prop="rent_price" label="出租价格" align="center" width="120">
                </el-table-column>
                <el-table-column prop="deposit" label="出租押金" align="center" width="120">
                </el-table-column>
                <el-table-column label="出租状态" width="120" align="center">
                    <template slot-scope="scope">{{
                        scope.row.is_renting === 0 ? "未租出" : "已租出"
                    }}</template>
                </el-table-column>
                <el-table-column prop="created_at" label="录入时间" align="center" width="200">
                </el-table-column>
                <el-table-column fixed="right" label="操作" align="center">
                    <template slot-scope="scope">
                        <el-button @click="rentClick(scope.row.car_number, scope.row.id)" type="text"
                            size="small">租赁汽车</el-button>
                        <el-button type="text" size="small">查看车辆大图</el-button>
                    </template>
                </el-table-column>
            </el-table>
        </el-main>
        <el-footer>
            <el-pagination @size-change="handleSizeChange" @current-change="handleCurrentChange"
                :current-page="page.current" :page-sizes="[1, 10, 20]" :page-size="page.size"
                layout="total, sizes, prev, pager, next, jumper" :total="total">
            </el-pagination>
        </el-footer>


        <el-dialog title="租赁汽车" :visible.sync="dialogRenting" width="30%">
            <el-form ref="elForm" :model="rentData" size="medium" label-width="100px">
                <el-form-item label="客户名称:" prop="">
                    <el-input v-model="rentData.customerName" placeholder="请输入客户名称:" readonly :disabled='true'
                        :style="{ width: '100%' }">
                    </el-input>
                </el-form-item>
                <el-form-item label="身份证:" prop="">
                    <el-input v-model="rentData.customerIdentity" placeholder="请输入身份证:" readonly :disabled='true'
                        :style="{ width: '100%' }">
                    </el-input>
                </el-form-item>
                <el-form-item label="出租单号:" prop="">
                    <el-input v-model="rentData.identity" placeholder="请输入出租单号:" readonly :disabled='true'
                        :style="{ width: '100%' }">
                    </el-input>
                </el-form-item>
                <el-form-item label="车牌号:" prop="">
                    <el-input v-model="rentData.car_number" placeholder="请输入车牌号:" readonly :disabled='true'
                        :style="{ width: '100%' }">
                    </el-input>
                </el-form-item>
                <el-form-item label="日期范围">
                    <el-date-picker type="datetimerange" v-model="rentData.dateRange" format="yyyy-MM-dd HH:mm:ss"
                        value-format="yyyy-MM-dd HH:mm:ss" :style="{ width: '100%' }" start-placeholder="开始日期"
                        end-placeholder="结束日期" range-separator="至" clearable></el-date-picker>
                </el-form-item>
                <el-form-item label="出租金额:" prop="">
                    <el-input v-model="rentData.order_price" placeholder="请输入出租金额:" :style="{ width: '100%' }">
                    </el-input>
                </el-form-item>
            </el-form>
            <span slot="footer" class="dialog-footer">
                <el-button @click="dialogRenting = false">取 消</el-button>
                <el-button type="primary" @click="addData()">确 定</el-button>
            </span>
        </el-dialog>
    </div>
</template>

<script>
import { dataTimeFormatter } from "@/utils/utils";
import { get, post } from "@/utils/requests";
export default {
    components: {},
    data() {
        return {
            page: {
                current: 1, // 当前分页
                size: 10, // 显示数量
            },
            total: 0,
            IdNumber: "360421199906242316",
            tableData: [],
            dialogRenting: false,
            rentData: {
                order_price: "",
                begin_date: "",
                return_date: "",
                identity: "",
                car_number: "",
                customerIdentity: "",
                customerName: "",
                customerId: "",
                dateRange: []
            },
            customerData: {}
        };
    },
    computed: {},
    watch: {},
    methods: {
        // 分页显示数量
        handleSizeChange(newSize) {
            this.page.size = newSize;
            this.getDataList();
        },
        // 分页当前分页
        handleCurrentChange(newPage) {
            this.page.current = newPage;
            this.getDataList();
        },
        async searchCar() {
            let params = this.IdNumber
            const { data } = await get("/customer/getcustomerinfobyidentity", { identity: params })
            if (data.code == 200) {
                this.getDataList();
                this.customerData = data.data
                this.rentData.customerIdentity = this.customerData.identity
                this.rentData.customerName = this.customerData.customer_name
                this.rentData.customerId = this.customerData.id

            } else {
                const h = this.$createElement;
                this.$notify.error({
                    title: '查询失败',
                    message: h('i', '客户身份证号不存在，请更正后在查询')
                });
            }
        },
        async getDataList() {
            let params = {
                current: this.page.current,
                pageSize: this.page.size,
            };
            const { data } = await post("/car/getcarinfobynumber", params);
            this.tableData = data.data.data;
            this.total = data.data.total;
            console.log(this.tableData);
            for (let i = 0; i < this.tableData.length; i++) {
                this.tableData[i].created_at = dataTimeFormatter(
                    this.tableData[i].created_at
                );
                this.tableData[i].updated_at = dataTimeFormatter(
                    this.tableData[i].updated_at
                );
            }
        },
        async rentClick(car_number, carId) {
            this.rentData.car_number = car_number
            this.rentData.identity = await get("/rentOrder/createrentidentity", { car_number: car_number, customer_identity: this.rentData.customerIdentity })
            this.rentData.identity = this.rentData.identity.data.data
            this.dialogRenting = true
        },
        async addData() {
            let params = {
                customer_id: this.rentData.customerId,
                order_price: this.rentData.order_price,
                begin_date: this.rentData.dateRange[0],
                return_date: this.rentData.dateRange[1],
                identity: this.rentData.identity,
                car_number: this.rentData.car_number,
                oper_name: this.$store.state.user.name,
                status: 1
            }
            const { data } = await post("/rentOrder/AddRentOrder", params);
            if (data.code == 200) {
                this.dialogRenting = false;
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