<!--  -->
<template>
  <div>
    <el-header>
      <div style="display: flex; margin: 20px">
        <div style="width: 40">
          <el-input placeholder="请输入车牌号" prefix-icon="el-icon-search" v-model="carNumber">
          </el-input>
        </div>
        <el-button class="btns" type="primary" icon="el-icon-search" style="margin-left: 10px" @click="getDataList()">
          查找
        </el-button>
        <el-button class="btns" type="primary" icon="el-icon-circle-plus-outline" style="margin-left: 10px"
          @click="dialogAdd = true">
          添加
        </el-button>
      </div>
    </el-header>
    <el-main>
      <el-table :data="tableData" border style="width: 100%">
        <el-table-column type="index" width="50" label="序号" align="center">
        </el-table-column>
        <el-table-column prop="car_number" label="车牌号" width="100" align="center">
        </el-table-column>
        <el-table-column prop="car_type" label="品牌" align="center">
        </el-table-column>
        <el-table-column prop="color" label="颜色" width="70" align="center">
        </el-table-column>
        <el-table-column prop="price" label="车辆价格" align="center">
        </el-table-column>
        <el-table-column prop="rent_price" label="租赁价格/天" align="center">
        </el-table-column>
        <el-table-column prop="deposit" label="押金" align="center">
        </el-table-column>
        <el-table-column label="缩略图" width="200">
          <template slot-scope="scope">
            <!-- trigger(触发方式)、placement(出现位置) -->
            <el-popover trigger="hover" placement="right">
              <!-- table中原本显示的图片 -->
              <img slot="reference"
                src="https://images.weserv.nl/?url=http://43.136.71.96:443/images/1.png"
                alt="https://images.weserv.nl/?url=http://43.136.71.96:443/images/1.png"
                style="width: 100px;height: 100px;">
              <!-- 鼠标移入时弹出的图片 -->
              <img src="https://images.weserv.nl/?url=http://43.136.71.96:443/images/1.png"
                alt="" style="width: 500px;height: 500px;">
            </el-popover>
          </template>
        </el-table-column>

        <el-table-column label="状态" width="70" align="center">
          <template slot-scope="scope">{{
            scope.row.is_renting === 0 ? "未租出" : "已租出"
          }}</template>
        </el-table-column>
        <el-table-column prop="created_at" label="创建时间" align="center">
        </el-table-column>
        <el-table-column prop="updated_at" label="更新时间" align="center">
        </el-table-column>
        <el-table-column label="操作" width="150" align="center">
          <template slot-scope="scope">
            <el-button size="mini" type="primary" @click="editData(scope.row)">编辑</el-button>
            <el-button size="mini" type="danger" @click="deleteData(scope.row.id)">删除</el-button>
          </template>
        </el-table-column>
      </el-table>
    </el-main>
    <el-dialog title="添加车辆" :visible.sync="dialogAdd" width="30%">
      <el-form ref="newData" :model="newData" label-width="80px">
        <el-form-item label="车牌号">
          <el-input placeholder="车牌号" v-model="newData.car_number">
          </el-input>
        </el-form-item>
        <el-form-item label="品牌">
          <el-input v-model="newData.car_type" placeholder="请输入品牌"></el-input>
        </el-form-item>
        <el-form-item label="颜色">
          <el-input v-model="newData.color" placeholder="请输入颜色"></el-input>
        </el-form-item>
        <el-form-item label="车辆价格">
          <el-input v-model="newData.price" placeholder="请输入车辆价格"></el-input>
        </el-form-item>
        <el-form-item label="租金/天">
          <el-input v-model="newData.rent_price" placeholder="请输入租金价格"></el-input>
        </el-form-item>
        <el-form-item label="押金">
          <el-input v-model="newData.deposit" placeholder="请输入押金"></el-input>
        </el-form-item>

        <el-upload class="avatar-uploader" action="http://43.136.71.96:21/images/" :show-file-list="false"
          :on-success="handleAvatarSuccess" :before-upload="beforeAvatarUpload">
          <img v-if="imageUrl" :src="imageUrl" class="avatar">
          <i v-else class="el-icon-plus avatar-uploader-icon"></i>
        </el-upload>

      </el-form>
      <span slot="footer" class="dialog-footer">
        <el-button @click="dialogAdd = false">取 消</el-button>
        <el-button type="primary" @click="addData()">确 定</el-button>
      </span>
    </el-dialog>
    <el-dialog title="编辑车辆信息" :visible.sync="dialogEdit" width="30%">
      <el-form ref="editDataList" :model="editDataList" label-width="80px">
        <el-form-item label="车牌号">
          <el-input placeholder="车牌号" v-model="editDataList.car_number">
          </el-input>
        </el-form-item>
        <el-form-item label="品牌">
          <el-input v-model="editDataList.car_type" placeholder="请输入品牌"></el-input>
        </el-form-item>
        <el-form-item label="颜色">
          <el-input v-model="editDataList.color" placeholder="请输入颜色"></el-input>
        </el-form-item>
        <el-form-item label="车辆价格">
          <el-input v-model="editDataList.price" placeholder="请输入车辆价格"></el-input>
        </el-form-item>
        <el-form-item label="租金/天">
          <el-input v-model="editDataList.rent_price" placeholder="请输入租金价格"></el-input>
        </el-form-item>
        <el-form-item label="押金">
          <el-input v-model="editDataList.deposit" placeholder="请输入押金"></el-input>
        </el-form-item>
      </el-form>
      <span slot="footer" class="dialog-footer">
        <el-button @click="dialogAdd = false">取 消</el-button>
        <el-button type="primary" @click="updateData()">确 定</el-button>
      </span>
    </el-dialog>
    <el-footer>
      <el-pagination @size-change="handleSizeChange" @current-change="handleCurrentChange" :current-page="page.current"
        :page-sizes="[1, 10, 20]" :page-size="page.size" layout="total, sizes, prev, pager, next, jumper" :total="total">
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
      carNumber: "",
      total: 0,
      tableData: [],
      newData: {
        car_number: "",
        car_type: "",
        color: "",
        deposit: "",
        description: "",
        price: "",
        rent_price: "",
        img_url: "",
      },
      editDataList: {
        car_number: "",
        car_type: "",
        color: "",
        deposit: "",
        description: "",
        price: "",
        rent_price: "",
        img_url: "",
      },
      dialogAdd: false,
      dialogEdit: false,
      imageUrl: '',
    };

  },
  computed: {},
  watch: {},
  methods: {
    handleAvatarSuccess(res, file) {
      this.imageUrl = URL.createObjectURL(file.raw);
    },
    beforeAvatarUpload(file) {
      const isJPG = file.type === 'image/jpeg';
      const isLt2M = file.size / 1024 / 1024 < 2;

      if (!isJPG) {
        this.$message.error('上传头像图片只能是 JPG 格式!');
      }
      if (!isLt2M) {
        this.$message.error('上传头像图片大小不能超过 2MB!');
      }
      return isJPG && isLt2M;
    },
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
        car_number: this.carNumber,
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
    async editData(data) {
      this.editDataList.id = data.id;
      this.editDataList.car_number = data.car_number;
      this.editDataList.car_type = data.car_type;
      this.editDataList.color = data.color;
      this.editDataList.deposit = data.deposit;
      this.editDataList.description = data.description;
      this.editDataList.price = data.price;
      this.editDataList.rent_price = data.rent_price;
      this.editDataList.img_url = data.img_url;
      this.dialogEdit = true;
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
        const { data } = await get("/car/deletecar", { id: id });
        if (data.code !== 200) {
          return this.$message.error(data.message);
        } else {
          this.$message.success("删除成功");
          //刷新用户列表
          this.getDataList();
        }
      }
    },
    async addData() {
      const { data } = await post("/car/addcar", this.newData);
      console.log(data);
      if (data.code == 200) {
        this.dialogAdd = false;
        this.$notify({
          message: "添加成功",
          type: "success",
        });
        this.getDataList();
      } else {
        this.$notify({
          message: "添加失败",
          type: "warning",
        });
      }
    },
    async updateData() {
      const { data } = await post(
        "/car/updatecarinfo",
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
    this.getDataList();
  },
  mounted() { },
};
</script>
<style>
.avatar-uploader .el-upload {
  margin-left: 30%;
  border: 1px dashed #d9d9d9;
  border-radius: 6px;
  cursor: pointer;
  position: relative;
  overflow: hidden;
}

.avatar-uploader .el-upload:hover {
  border-color: #409EFF;
}

.avatar-uploader-icon {
  font-size: 28px;
  color: #8c939d;
  width: 178px;
  height: 178px;
  line-height: 178px;
  text-align: center;
}

.avatar {
  width: 178px;
  height: 178px;
  display: block;
}</style>