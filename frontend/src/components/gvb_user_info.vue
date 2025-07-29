<template>
  <div class="gvb_user_info">
    <div class="avatar" v-if="props.isAvatar">
      <img :src="store.userInfo.avatar" alt="">
    </div>
    <div class="drop_menu">
      <a-dropdown placement="bottomRight">
        <a class="ant-dropdown-link" @click.prevent>
          {{ store.userInfo.nick_name }}
          <i class="fa fa-angle-down"></i>
        </a>
        <template #overlay>
          <a-menu>
            <template v-for="(item,index) in data.menu_list" :key="index">
              <a-menu-item v-if="item.title !== 'line'" @click="goto(item)">
                <a href="javascript:void (0)">{{ item.title }}</a>
              </a-menu-item>
              <a-menu-divider v-else/>
            </template>
            <a-menu-divider/>
            <a-menu-item @click="logoutUtil">
              <a href="javascript:;">注销退出</a>
            </a-menu-item>
          </a-menu>
        </template>
      </a-dropdown>
    </div>
  </div>
</template>

<script setup>
import {useRoute, useRouter} from "vue-router"
import {logoutApi} from "@/api/user_api";
import {message} from "ant-design-vue";
import {useStore} from "@/stores/store";
import {reactive} from "vue";
import {logoutUtil} from "@/utils/logout";

const router = useRouter()
const route = useRoute()
const store = useStore()
const props = defineProps({
  // 是否显示头像部分
  isAvatar: {
    type: Boolean,
    default: false,
  }
})

const data = reactive({
  menu_list: [

  ]
})

function goto(item) {
  store.setCrumb([item.parentTitle, item.title])
  store.addTab({
    name: item.name,
    title: item.title,
    parentTitle: item.parentTitle // 一级菜单的名称，如果为undefined，那么这个tab就是一级菜单
  })
  router.push({
    name: item.name
  })
}

</script>

<style>
.gvb_user_info {
  display: flex;
  align-items: center;
}
</style>