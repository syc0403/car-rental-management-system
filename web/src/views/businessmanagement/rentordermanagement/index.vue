<!--  -->
<template>
    <div class=''> <el-header>
            <hr class="hr">
            <div style="display: flex;">
                <div style="display: flex;margin:10px;font-size: 20px;">出租单号：</div>
                <div style="width: 40">
                    <el-input placeholder="请输入出租单号" prefix-icon="el-icon-search">
                    </el-input>
                </div>

                <div style="display: flex;margin:10px;font-size: 20px;">身份证号：</div>
                <div style="width: 40">
                    <el-input placeholder="请输入身份证号" prefix-icon="el-icon-search">
                    </el-input>
                </div>
                <div style="display: flex;margin:10px;font-size: 20px;">车牌号：</div>
                <div style="width: 40">
                    <el-input placeholder="请输入车牌号" prefix-icon="el-icon-search">
                    </el-input>
                </div>
                <div style="margin:10px;">
                    <el-radio label="1">已归还</el-radio>
                    <el-radio label="0">未归还</el-radio>
                </div>
                <el-button class="btns" type="primary" icon="el-icon-search" style="margin-left: 10px"
                    @click="getDataList()">
                    查询
                </el-button>
            </div>
        </el-header>
        <el-main>
            <el-table :data="tableData" border style="width: 100%">
                <el-table-column prop="car_number" label="车牌号" width="120">
                </el-table-column>
                <el-table-column prop="car_type" label="车辆类型" width="120">
                </el-table-column>
                <el-table-column prop="color" label="出租颜色" width="120">
                </el-table-column>
                <el-table-column prop="price" label="汽车价格" width="120">
                </el-table-column>
                <el-table-column prop="rent_price" label="出租价格" width="120">
                </el-table-column>
                <el-table-column prop="deposit" label="出租押金" width="120">
                </el-table-column>
                <el-table-column label="状态" width="70" align="center">
                    <template slot-scope="scope">{{
                        scope.row.is_renting === 0 ? "未租出" : "已租出"
                    }}</template>
                </el-table-column>
                <el-table-column prop="img_url" label="缩略图" width="120">
                </el-table-column>
                <el-table-column prop="created_at" label="录入时间" width="120">
                </el-table-column>
                <el-table-column fixed="right" label="操作" width="100">
                    <template slot-scope="scope">
                        <el-button @click="handleClick(scope.row)" type="text" size="small">查看</el-button>
                        <el-button type="text" size="small">编辑</el-button>
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
                this.tableData[i].customer_name = (await get("/customer/getcustomerinfobyid?id=" + this.tableData[i].customer_id)).data.data.customer_name;
            }
            console.log(this.tableData);

        },
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