<!--  -->
<template>
    <div class=''> <el-header>
            <hr class="hr">
            <div style="display: flex;">
                <div style="display: flex;margin:10px;font-size: 20px;">出租单号：</div>
                <div style="width: 40">
                    <el-input placeholder="请输入出租单号" prefix-icon="el-icon-search" v-model="searchData.identity">
                    </el-input>
                </div>

                <div style="display: flex;margin:10px;font-size: 20px;">身份证号：</div>
                <div style="width: 40">
                    <el-input placeholder="请输入身份证号" prefix-icon="el-icon-search" v-model="searchData.customer_identity">
                    </el-input>
                </div>
                <div style="display: flex;margin:10px;font-size: 20px;">车牌号：</div>
                <div style="width: 40">
                    <el-input placeholder="请输入车牌号" prefix-icon="el-icon-search" v-model="searchData.car_number">
                    </el-input>
                </div>
                <div style="margin:10px;">
                    <el-radio label="1" v-model="searchData.status">未归还</el-radio>
                    <el-radio label="2" v-model="searchData.status">已归还</el-radio>
                </div>
                <el-button class="btns" type="primary" icon="el-icon-search" style="margin-left: 10px"
                    @click="searchDataList()">
                    查询
                </el-button>
            </div>
        </el-header>
        <el-main>
            <el-table :data="tableData" border style="width: 100%">
                <el-table-column prop="identity" label="出租单号" align="center">
                </el-table-column>
                <el-table-column prop="customer_identity" label="身份证号" align="center">
                </el-table-column>
                <el-table-column prop="car_number" label="车牌号" align="center">
                </el-table-column>
                <el-table-column prop="order_price" label="出租价格" align="center">
                </el-table-column>
                <el-table-column label="状态" width="70" align="center">
                    <template slot-scope="scope">{{
                        scope.row.is_renting === 1 ? "已归还" : "未归还"
                    }}</template>
                </el-table-column>
                <el-table-column prop="begin_date" label="起租时间" align="center">
                </el-table-column>
                <el-table-column prop="return_date" label="还车时间" align="center">
                </el-table-column>
                <el-table-column prop="customer_name" label="客户名称" align="center">
                </el-table-column>
                <el-table-column prop="created_at" label="录入时间" align="center">
                </el-table-column>
                <el-table-column fixed="right" label="操作" align="center">
                    <template slot-scope="scope">
                        <el-button @click="deleteData(scope.row.id)" type="text" size="small">删除</el-button>
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
            tableData: [],
            searchData: {
                identity: "",
                customer_identity: "",
                car_number: "",
                status: 0
            }
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
        async getDataList() {
            let params = {
                current: this.page.current,
                pageSize: this.page.size,
            };
            const { data } = await get("/rentOrder/getallrentorder", params);
            this.tableData = data.data.data;
            this.total = data.data.total;
            for (let i = 0; i < this.tableData.length; i++) {
                this.tableData[i].created_at = dataTimeFormatter(
                    this.tableData[i].created_at
                );
                this.tableData[i].updated_at = dataTimeFormatter(
                    this.tableData[i].updated_at
                );
                let customerData = (await get("/customer/getcustomerinfobyid?id=" + this.tableData[i].customer_id)).data.data;
                this.tableData[i].customer_name = customerData.customer_name;
                this.tableData[i].customer_identity = customerData.identity;
            }
            console.log(this.tableData);
        },
        async deleteData(id) {
            const confirmRes = await this.$confirm(
                "此操作将永久删除该订单, 是否继续?",
                "提示",
                {
                    confirmButtonText: "确定",
                    cancelButtonText: "取消",
                    type: "warning",
                }
            ).catch((err) => err); //用catch来捕获错误消息

            //用户单击确认，返回值为字符串confirm; 单击取消，返回值为字符串cancel
            if ("cancel" === confirmRes) {
                //用户点击了取消
                return this.$message.info("已经取消了删除");
            }
            if ("confirm" === confirmRes) {
                //用户点击了确认
                const { data } = await get("/rentOrder/deleteRentOrder", { id: id });
                if (data.code !== 200) {
                    return this.$message.error(data.message);
                } else {
                    this.$message.success("删除成功");
                    //刷新用户列表
                    this.getDataList();
                }
            }
        },
        async searchDataList() {
            let params = {
                current: this.page.current,
                pageSize: this.page.size,
                identity: this.searchData.identity,
                customer_identity: this.searchData.customer_identity,
                car_number: this.searchData.car_number,
                status: parseInt(this.searchData.status)
            };
            console.log(params);
            const { data } = await post("/rentOrder/getrentorderby", params);
            this.tableData = data.data.data;
            this.total = data.data.total;
            for (let i = 0; i < this.tableData.length; i++) {
                this.tableData[i].created_at = dataTimeFormatter(
                    this.tableData[i].created_at
                );
                this.tableData[i].updated_at = dataTimeFormatter(
                    this.tableData[i].updated_at
                );
                let customerData = (await get("/customer/getcustomerinfobyid?id=" + this.tableData[i].customer_id)).data.data;
                this.tableData[i].customer_name = customerData.customer_name;
                this.tableData[i].customer_identity = customerData.identity;
            }
            console.log(this.tableData);
        }
    },
    created() {
        this.getDataList()
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
</style>