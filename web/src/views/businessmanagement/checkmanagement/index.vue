<!--  -->
<template>
    <div class=''> <el-header>
            <hr class="hr">
            <div style="display: flex;">
                <div style="display: flex;margin:10px;font-size: 20px;">出租单号：</div>
                <div style="width: 40">
                    <el-input placeholder="请输入出租单号" prefix-icon="el-icon-search" v-model="searchData.rentorder_identity">
                    </el-input>
                </div>

                <div style="display: flex;margin:10px;font-size: 20px;">检查单号：</div>
                <div style="width: 40">
                    <el-input placeholder="请输入检查单号" prefix-icon="el-icon-search" v-model="searchData.check_identity">
                    </el-input>
                </div>
                <el-button class="btns" type="primary" icon="el-icon-search" style="margin-left: 10px"
                    @click="searchDataList()">
                    查询
                </el-button>
            </div>
        </el-header>
        <el-main>
            <el-table :data="tableData" border style="width: 100%">
                <el-table-column prop="identity" label="检查单单号" align="center">
                </el-table-column>
                <el-table-column prop="rentorder_identity" label="出租单单号" align="center">
                </el-table-column>
                <el-table-column prop="problem" label="存在问题" align="center">
                </el-table-column>
                <el-table-column prop="problem_desc" label="问题描述" align="center">
                </el-table-column>
                <el-table-column prop="pay_price" label="赔付金额" align="center" width="100">
                </el-table-column>
                <el-table-column prop="customer_name" label="客户名称" align="center">
                </el-table-column>
                <el-table-column prop="check_time" label="检查时间" align="center">
                </el-table-column>
                <el-table-column prop="created_at" label="录入时间" align="center">
                </el-table-column>
                <el-table-column fixed="right" label="操作" align="center">
                    <template slot-scope="scope">
                        <el-button size="mini" type="primary" @click="editData(scope.row)">编辑</el-button>
                        <el-button size="mini" type="primary" @click="deleteData(scope.row.id)">删除</el-button>
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

        <el-dialog title="编辑车辆信息" :visible.sync="dialogEdit" width="30%">
            <el-form ref="editDataList" :model="editDataList" label-width="80px">
                <el-form-item label="存在问题">
                    <el-input placeholder="存在问题" v-model="editDataList.problem">
                    </el-input>
                </el-form-item>
                <el-form-item label="问题描述">
                    <el-input v-model="editDataList.problem_desc" placeholder="请输入问题描述"></el-input>
                </el-form-item>
                <el-form-item label="赔付金额">
                    <el-input v-model="editDataList.pay_price" placeholder="请输入赔付金额"></el-input>
                </el-form-item>
            </el-form>
            <span slot="footer" class="dialog-footer">
                <el-button @click="dialogAdd = false">取 消</el-button>
                <el-button type="primary" @click="updateData()">确 定</el-button>
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
            tableData: [],
            searchData: {
                check_identity: "",
                rentorder_identity: "",
            },
            editDataList: {
                id: "",
                problem: "",
                problem_desc: "",
                pay_price: 0,
            },
            dialogEdit: false,
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
            const { data } = await get("/check/getallcheck", params);
            this.tableData = data.data.data;
            this.total = data.data.total;
            for (let i = 0; i < this.tableData.length; i++) {
                this.tableData[i].created_at = dataTimeFormatter(
                    this.tableData[i].created_at
                );
                this.tableData[i].updated_at = dataTimeFormatter(
                    this.tableData[i].updated_at
                );
                this.tableData[i].customer_name = (await get("/customer/getcustomerbyrentorder", { rentorder_identity: "CZ_2316U960_1683472026" })).data.data.customer_name
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
                const { data } = await get("/check/deletecheck", { id: id });
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
            if (this.searchData.check_identity == "" && this.searchData.rentorder_identity == "") { } else {
                let params = {
                    current: this.page.current,
                    pageSize: this.page.size,
                    identity: this.searchData.check_identity,
                    rentorder_identity: this.searchData.rentorder_identity,
                };
                const { data } = await post("/rentOrder/getcheckby", params);
                this.tableData = data.data.data;
                this.total = data.data.total;
                for (let i = 0; i < this.tableData.length; i++) {
                    this.tableData[i].created_at = dataTimeFormatter(
                        this.tableData[i].created_at
                    );
                    this.tableData[i].updated_at = dataTimeFormatter(
                        this.tableData[i].updated_at
                    );
                    this.tableData[i].customer_name = (await get("/customer/getcustomerbyrentorder", { rentorder_identity: "CZ_2316U960_1683472026" })).data.data.customer_name
                }
            }
        },
        async editData(data) {
            this.editDataList.id = data.id;
            this.editDataList.problem = data.problem;
            this.editDataList.problem_desc = data.problem_desc;
            this.editDataList.pay_price = data.pay_price;
            this.dialogEdit = true;
        },
        async updateData() {
            const { data } = await post(
                "/rentOrder/updatecheck",
                this.editDataList
            );
            console.log(data);
            if (data.code == 200) {
                this.dialogEdit = false;
                this.$notify({
                    message: "修改成功",
                    type: "success",
                });
                this.getDataList();
            } else {
                this.$notify({
                    message: "修改失败",
                    type: "warning",
                });
            }
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