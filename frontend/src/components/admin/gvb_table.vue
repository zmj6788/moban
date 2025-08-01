<template>
  <div class="gvb_container">

    <div class="gvb_search">
      <a-input-search :placeholder="props.likeTitle" v-model:value="page.key" style="width: 200px" @search="onSearch" />
      <slot name="filters"></slot>
      <div class="gvb_refresh">
        <a-button title="刷新本页" @click="refresh"><i class="fa fa-refresh"></i></a-button>
      </div>
    </div>
    <div class="gvb_actions">
      <slot name="add">

      </slot>
      <slot name="batchRemove">
        <a-button type="danger" @click="removeBatch"
          v-if="isDelete && isBatchDelete && data.selectedRowKeys.length">批量删除</a-button>
      </slot>
    </div>
    <div class="gvb_tables">
      <a-spin :spinning="data.spinning" tip="加载中..." :delay="300">
        <a-table :columns="props.columns" :row-selection="props.isSelect ? {
          selectedRowKeys: data.selectedRowKeys,
          onChange: onSelectChange
        } : undefined" :pagination="false" row-key="id" :data-source="data.list">
          <template #bodyCell="{ column, record, index }">
            <slot name="cell" v-bind="{ column, record, index }">
              <template v-if="column.key === 'created_at'">
                <span>{{ getFormatDateTime(record.created_at) }}</span>
              </template>
              <template v-if="column.key === 'action'">
                <slot name="edit" v-bind="{ column, record }">
                  <a-button type="primary" v-if="isEdit">编辑</a-button>
                </slot>
                <slot name="delete" v-bind="{ column, record }">
                  <a-popconfirm title="是否确定删除?" ok-text="删除" cancel-text="取消" @confirm="userRemove(record.id)"
                    v-if="props.isDelete">
                    <a-button class="gvb_table_action delete" type="danger">删除</a-button>
                  </a-popconfirm>
                </slot>
              </template>
            </slot>
          </template>
        </a-table>
      </a-spin>
    </div>
    <div class="gvb_pages">
      <a-pagination :showSizeChanger="false" v-model:current="page.page" v-model:page-size="page.limit"
        @change="pageChange" :total="data.count" :show-total="total => `共 ${total} 条`" />
    </div>
  </div>
</template>

<script setup>
import { reactive, ref } from "vue";
import { getFormatDateTime } from "@/utils/date";
import { message } from "ant-design-vue";
import { baseDeleteApi, baseListApi } from "@/api/base_api";

const emits = defineEmits(["delete"])
const props = defineProps({
  columns: {
    type: Array
  },
  list: {
    type: Array,
  },
  baseUrl: {
    type: String,
    default: "",
  },
  isAdd: {
    type: Boolean,
    default: true
  },
  isEdit: {
    type: Boolean,
    default: true
  },
  isDelete: {
    type: Boolean,
    default: true
  },
  isBatchDelete: {
    type: Boolean,
    default: true
  },
  isSelect: {
    type: Boolean,
    default: true
  },
  filters: {
    type: Array,
  },
  pageSize: {
    type: Number,
    default: 10,
  },
  likeTitle: {
    type: String,
    default: "模糊搜索"
  },
  defaultDelete: {
    type: Boolean,
    default: false
  }
})

// 分页
const page = reactive({
  page: 1,
  limit: props.pageSize,
  key: "",
})
// 表单ref


const data = reactive({
  list: [], // 数据
  selectedRowKeys: [], // 选择的id列表
  count: 0, // 总数
  spinning: true, // 默认是在加载中
})

// 添加用户的modal
function addModal() {
  data.modalVisible = true
}

// 选择复选框
function onSelectChange(selectedKeys) {
  data.selectedRowKeys = selectedKeys
}

// 批量删除
async function removeBatch() {
  if (props.defaultDelete) {
    let res = await baseDeleteApi(props.baseUrl, data.selectedRowKeys)
    if (res.code) {
      message.error(res.msg)
      data.selectedRowKeys = []
      return
    }
    message.success(res.msg)
    getData(page)
    data.selectedRowKeys = []
    return
  }
  emits("delete", data.selectedRowKeys)
  data.selectedRowKeys = []
}

// 获取列表页数据
async function getData(params) {
  if (props.baseUrl === "") {
    data.spinning = false
    data.list = props.list
    return
  }
  data.spinning = true
  let res = await baseListApi(props.baseUrl, params)
  data.spinning = false
  if (res.data.list === undefined) {
    data.list = res.data
    data.count = res.data.length
    return
  }
  data.list = res.data.list
  data.count = res.data.count
  data.spinning = false
}

// 分页
function pageChange(_page, limit) {
  getData(page)
}

// 删除单个用户
async function userRemove(id) {
  if (props.defaultDelete) {
    let res = await baseDeleteApi(props.baseUrl, [id])
    if (res.code) {
      message.error(res.msg)
      return
    }
    message.success(res.msg)
    getData(page)
    return
  }
  emits("delete", [id])
}

// 刷新
function refresh() {
  message.success("刷新成功")
  getData(page)
}

function onSearch() {
  page.key = page.key.trim()
  page.page = 1
  getData(page)
}

function ExportList(params) {
  if (params === undefined) {
    params = {}
  }
  page.page = 1

  // Object.assign(page, params)

  let newPage = {}
  Object.assign(newPage, page)
  Object.assign(newPage, params)
  getData(newPage)
}

getData(page)

defineExpose({
  ExportList
})


</script>

<style lang="scss">
.gvb_container {
  background-color: var(--card_bg);

  .gvb_search {
    padding: 10px;
    border-bottom: 1px solid var(--bg);
    position: relative;

    .gvb_select {
      margin-left: 10px;
    }

    .gvb_refresh {
      position: absolute;
      right: 10px;
      top: 10px;

      i {
        color: var(--text);
      }
    }
  }

  .gvb_actions {
    padding: 10px;

    .ant-btn {
      margin-right: 10px;
    }
  }

  .gvb_tables {
    padding: 0 10px 10px 10px;

    .ant-btn {
      margin-right: 10px;
    }
  }

  .gvb_pages {
    display: flex;
    justify-content: center;
    padding: 10px;
    margin-bottom: 20px;
  }

  .gvb_table_avatar {
    width: 40px;
    height: 40px;
    border-radius: 50%;
  }

  .gvb_table_action.update {
    margin-right: 10px;
  }
}
</style>