<!--  -->
<template>
    <div class=''>
        <el-header>
            <el-button class="btns" type="primary" icon="el-icon-circle-plus-outline" @click="dialogAdd = true">
                添加
            </el-button>
        </el-header>
        <el-main>
            <el-table :data="tableData" border style="width: 100%">
                <el-table-column type="index" width="70" label="序号" align="center">
                </el-table-column>
                <el-table-column prop="user_type" label="标签" width="120" align="center">
                </el-table-column>
                <el-table-column prop="roles" label="角色名称" align="center">
                </el-table-column>
                <el-table-column prop="description" label="描述" align="center">
                </el-table-column>
                <el-table-column label="操作" width="150" align="center">
                    <template slot-scope="scope">
                        <el-button size="mini" type="primary" @click="editData(scope.row)">编辑</el-button>
                        <el-button size="mini" type="danger" @click="deleteData(scope.row.id)">删除</el-button>
                    </template>
                </el-table-column>
            </el-table>
        </el-main>
        <!-- 分页 -->
        <el-footer>
            <el-pagination @size-change="handleSizeChange" @current-change="handleCurrentChange"
                :current-page="page.current" :page-sizes="[1, 10, 20]" :page-size="page.size"
                layout="total, sizes, prev, pager, next, jumper" :total="total">
            </el-pagination>
        </el-footer>

        <el-dialog title="编辑用户信息" :visible.sync="dialogEdit" width="30%">
            <el-form ref="editDataList" :model="editDataList" label-width="80px">
                <el-form-item label="角色标签">
                    <el-input placeholder="请输入角色标签" v-model="editDataList.user_type">
                    </el-input>
                </el-form-item>
                <el-form-item label="角色名称">
                    <el-input v-model="editDataList.roles" placeholder="请输入角色名称"></el-input>
                </el-form-item>
                <el-form-item label="描述">
                    <el-input v-model="editDataList.description" placeholder="请输入描述"></el-input>
                </el-form-item>
            </el-form>
            <span slot="footer" class="dialog-footer">
                <el-button @click="dialogEdit = false">取 消</el-button>
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
            editDataList: {
                id: "",
                user_type: "",
                roles: "",
                description: "",
            },
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
            const { data } = await get("/role/getallrole", params);
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
        async editData(data) {
            this.editDataList.id = data.id;
            this.editDataList.user_type = data.user_type;
            this.editDataList.roles = data.roles;
            this.editDataList.description = data.description;
            this.dialogEdit = true;
        },
        async updateData() {
            const { data } = await post(
                "/role/updaterole",
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
        async deleteData(id) {
            const confirmRes = await this.$confirm(
                "此操作将永久删除该用户, 是否继续?",
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
                const { data } = await get("/role/deleterole", { id: id });
                if (data.code !== 200) {
                    return this.$message.error(data.message);
                } else {
                    this.$message.success("删除成功");
                    //刷新用户列表
                    this.getDataList();
                }
            }
        },
    },
    created() {
        this.getDataList();
    },
    mounted() {

    },
}
</script>
<style>
.lookup {
    display: flex;
    margin: 20px;
}

.aInput {
    width: 70%;
}

.btns {
    margin-top: 15px;
}
</style>