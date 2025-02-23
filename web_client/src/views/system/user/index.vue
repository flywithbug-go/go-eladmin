<template>
  <div class="app-container">
    <eHeader :roles="roles" :query="query"/>
    <!--表格渲染-->
    <el-table
      v-loading="loading"
      :data="data"
      size="small"
      border
      style="width: 100%;">
      <el-table-column prop="username" label="用户名"/>
      <el-table-column label="头像">
        <template slot-scope="scope">
          <img :src="scope.row.avatar+ '?size=120'" class="avatar" width="100px" align="center">
        </template>
      </el-table-column>
      <el-table-column prop="email" label="邮箱"/>
      <el-table-column label="状态">
        <template slot-scope="scope">
          <span>{{ scope.row.enabled ? '激活':'锁定' }}</span>
        </template>
      </el-table-column>
      <el-table-column prop="createTime" label="注册日期">
        <template slot-scope="scope">
          <span>{{ parseTime(scope.row.createTime*1000) }}</span>
        </template>
      </el-table-column>
      <el-table-column label="操作" width="150px" align="center">
        <template slot-scope="scope">
          <edit v-if="checkPermission(['ADMIN','USER_ALL','USER_EDIT'])" :data="scope.row" :roles="roles" :sup_this="sup_this"/>
          <el-popover
            v-if="checkPermission(['ADMIN','USER_ALL','USER_DELETE'])"
            v-model="scope.row.delPopover"
            placement="top"
            width="180">
            <p>确定删除本条数据吗？</p>
            <div style="text-align: right; margin: 0">
              <el-button size="mini" type="text" @click="scope.row.delPopover = false">取消</el-button>
              <el-button :loading="delLoading" type="primary" size="mini" @click="subDelete(scope.$index, scope.row)">确定</el-button>
            </div>
            <el-button v-if="scope.row.id != 10000" slot="reference" type="danger" size="mini" @click="scope.row.delPopover = true">删除</el-button>
          </el-popover>
        </template>
      </el-table-column>
    </el-table>
    <!--分页组件-->
    <el-pagination
      :total="total"
      style="margin-top: 8px;"
      layout="total, prev, pager, next, sizes"
      @size-change="sizeChange"
      @current-change="pageChange"/>
  </div>
</template>

<script>
import checkPermission from '@/utils/permission'
import initData from '@/mixins/initData'
import { del } from '@/api/user'
import { getRoleTree } from '@/api/role'
import { parseTime } from '@/utils/index'
import eHeader from './module/header'
import edit from './module/edit'
export default {
  components: { eHeader, edit },
  mixins: [initData],
  data() {
    return {
      roles: [], delLoading: false, sup_this: this
    }
  },
  created() {
    this.query.type = 'username'
    this.query.enabled = '激活'
    this.getRoles()
    this.$nextTick(() => {
      this.init()
    })
  },
  methods: {
    parseTime,
    checkPermission,
    beforeInit() {
      this.url = '/user/tree'
      const sort = '+id'
      const query = this.query
      const type = query.type
      const value = query.value
      const enabled = query.enabled
      this.params = { page: this.page, size: this.size, sort: sort }
      if (type && value) { this.params[type] = value }
      if (enabled !== '' && enabled !== null) { this.params['enabled'] = enabled }
      return true
    },
    subDelete(index, row) {
      this.delLoading = true
      del(row.id).then(() => {
        this.delLoading = false
        row.delPopover = false
        this.init()
        this.$notify({
          title: '删除成功',
          type: 'success',
          duration: 1500
        })
      }).catch(err => {
        this.delLoading = false
        row.delPopover = false
        console.log(err.msg)
      })
    },
    getRoles() {
      getRoleTree().then(res => {
        this.roles = res.list
      })
    }
  }
}
</script>

<style scoped>

</style>
