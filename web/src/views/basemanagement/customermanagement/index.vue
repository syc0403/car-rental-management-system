<template>
  <div>
    <!-- 混合查询 -->
    <div class="lookup">
      <div style="width: 40">
        <el-input
          placeholder="请输入客户姓名"
          prefix-icon="el-icon-search"
          v-model="customerName"
        >
        </el-input>
      </div>
      <el-button
        class="btns"
        size="small"
        type="primary"
        icon="el-icon-search"
        @click="getCustomerList()"
      >
        查找
      </el-button>
      <el-button
        class="btns"
        size="small"
        type="primary"
        icon="el-icon-circle-plus-outline"
        @click="dialogAdd = true"
      >
        添加
      </el-button>
    </div>
    <el-table :data="tableData" border style="width: 100%">
      <el-table-column prop="id" label="id" width="70" align="center">
      </el-table-column>
      <el-table-column prop="customer_name" label="客户姓名" align="center">
      </el-table-column>
      <el-table-column prop="phone" label="电话" align="center">
      </el-table-column>
      <el-table-column prop="identity" label="身份证" align="center">
      </el-table-column>
      <el-table-column prop="sex" label="性别" width="70" align="center">
      </el-table-column>
      <el-table-column prop="address" label="地址" align="center">
      </el-table-column>
      <el-table-column prop="occupation" label="职业" align="center">
      </el-table-column>
      <el-table-column prop="created_at" label="创建时间" align="center">
      </el-table-column>
      <el-table-column prop="updated_at" label="更新时间" align="center">
      </el-table-column>
      <el-table-column label="操作" align="center">
        <template slot-scope="scope">
          <el-button size="mini" type="primary">编辑</el-button>
          <el-button size="mini" type="danger" @click="deleteData(scope.row.id)"
            >删除</el-button
          >
        </template>
      </el-table-column>
    </el-table>
    <el-dialog title="添加客户" :visible.sync="dialogAdd" width="30%">
      <el-form ref="newData" :model="newData" label-width="80px">
        <el-form-item label="姓名">
          <el-input
            placeholder="请输入客户姓名"
            v-model="newData.customer_name"
          >
          </el-input>
        </el-form-item>
        <el-form-item label="电话">
          <el-input
            v-model="newData.phone"
            placeholder="请输入客户电话"
          ></el-input>
        </el-form-item>
        <el-form-item label="身份证">
          <el-input
            v-model="newData.identity"
            placeholder="请输入客户身份证"
          ></el-input>
        </el-form-item>
        <el-form-item label="地址">
          <el-input
            v-model="newData.address"
            placeholder="请输入客户地址"
          ></el-input>
        </el-form-item>
        <el-form-item label="职业">
          <el-input
            v-model="newData.occupation"
            placeholder="请输入客户职业"
          ></el-input>
        </el-form-item>
        <el-form-item label="性别">
          <el-radio v-model="newData.sex" label=0>男</el-radio>
          <el-radio v-model="newData.sex" label=1>女</el-radio>
        </el-form-item>
      </el-form>
      <span slot="footer" class="dialog-footer">
        <el-button @click="dialogAdd = false">取 消</el-button>
        <el-button type="primary" @click="addCustomer()">确 定</el-button>
      </span>
    </el-dialog>
    <!-- 分页 -->
    <el-pagination
      @size-change="handleSizeChange"
      @current-change="handleCurrentChange"
      :current-page="page.current"
      :page-sizes="[1, 10, 20]"
      :page-size="page.size"
      layout="total, sizes, prev, pager, next, jumper"
      :total="total"
    >
    </el-pagination>
  </div>
</template>

<script>
import {} from "@/utils/utils";
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
      customerName: "",
      localcommunity: "",
      tableData: [],
      newData: {
        customer_name: "",
        phone: "",
        identity: "",
        sex: 0,
        address: "",
        occupation: "",
      },
      dialogAdd: false,
    };
  },
  computed: {},
  watch: {},
  methods: {
    // 分页显示数量
    handleSizeChange(newSize) {
      this.page.size = newSize;
      this.getCustomerList();
    },
    // 分页当前分页
    handleCurrentChange(newPage) {
      this.page.current = newPage;
      this.getCustomerList();
    },

    async getCustomerList() {
      let params = {
        customer_name: this.customerName,
        current: this.page.current,
        pageSize: this.page.size,
      };
      const { data } = await post("/customer/getcustomerinfobyname", params);
      this.tableData = data.data.data;
      this.total = data.data.total;
    },
    async addCustomer() {
      console.log(this.newData);
      const { data } = await post("/customer/addcustomer", this.newData);
      console.log(data);
      if (data.code == 200) {
        this.dialogAdd = false;
        this.$notify({
          message: "添加成功",
          type: "success",
        });
        this.getCustomerList();
      } else {
        this.$notify({
          message: "添加失败",
          type: "warning",
        });
      }
    },
    async deleteData(id) {
      const { data } = await get("/customer/deletecustomer", { id: id });
      if (data.code == 200) {
        this.$notify({
          message: "删除成功",
          type: "success",
        });
        this.getCustomerList();
      } else {
        this.$notify({
          message: "删除失败",
          type: "warning",
        });
      }
    },
  },
  created() {
    this.getCustomerList();
  },
  mounted() {},
};
</script>
<style scoped>
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