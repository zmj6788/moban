<template>
  <a-menu v-model:selectedKeys="selectedKeys" :open-keys="data.openKeys" mode="inline" @openChange="onOpenChange"
    :inline-collapsed="false">
    <template v-for="menu in data.menuList" :key="menu.name">
      <a-menu-item :key="menu.name" v-if="menu.children.length === 0" @click="goto(menu)">
        <template #icon>
          <i :class="'fa ' + menu.icon"></i>
        </template>
        <span>{{ menu.title }}</span>
      </a-menu-item>
      <a-sub-menu :key="menu.name" v-else>
        <template #icon>
          <i :class="'fa ' + menu.icon"></i>
        </template>
        <template #title>{{ menu.title }}</template>
        <a-menu-item v-for="sub_menu in menu.children" @click="goto(sub_menu, menu)" :key="sub_menu.name">
          <template #icon>
            <i :class="'fa ' + sub_menu.icon"></i>
          </template>
          <span>{{ sub_menu.title }}</span>
        </a-menu-item>
      </a-sub-menu>
    </template>


  </a-menu>
</template>

<script setup>
import { reactive, ref } from "vue";
import { useRouter, useRoute, onBeforeRouteUpdate } from "vue-router";
import { useStore } from "@/stores/store";

const store = useStore()
const data = reactive({
  menuList: [
    {
      id: 1,
      icon: "fa-home",
      title: "首页",
      name: "home",
      children: []
    },

    {
      id: 4,
      icon: "fa-male",
      title: "用户管理",
      name: "users",
      children: [
        {
          id: 41,
          icon: "fa-users",
          title: "用户列表",
          name: "user_list",
          children: []
        },
      ]
    },
    {
      id: 5,
      icon: "fa-book",
      title: "货物管理",
      name: "goods",
      children: [

        {
          id: 51,
          icon: "fa-gift",
          title: "货物管理",
          name: "goods_list",
          children: []
        },
        {
          id: 52,
          icon: "fa-gift",
          title: "货物类型管理",
          name: "goods_type",
          children: []
        },
        {
          id: 53,
          icon: "fa-gift",
          title: "货物地区管理",
          name: "goods_region",
          children: []
        }
        ,
        {
          id: 54,
          icon: "fa-gift",
          title: "关税管理",
          name: "tariff_list",
          children: []
        }
      ]
    },
    {
      id: 6,
      icon: "fa-cogs",
      title: "系统管理",
      name: "sys",
      children: [
        {
          id: 61,
          icon: "fa-tree",
          title: "菜单列表",
          name: "menu_list",
          children: []
        },
        // {
        //   id: 52,
        //   icon: "fa-sun-o",
        //   title: "广告列表",
        //   name: "advert_list",
        //   children: []
        // },

        {
          id: 63,
          icon: "fa-server",
          title: "系统日志",
          name: "log_list",
          children: []
        },
        // {
        //   id: 64,
        //   icon: "fa-globe",
        //   title: "项目配置",
        //   name: "system",
        // },
      ]
    },
  ],
  openKeys: []
})




const selectedKeys = ref([])
const router = useRouter()
const route = useRoute()

function goto(item, parentItem) {
  if (parentItem !== undefined) {
    store.setCrumb([parentItem.title, item.title])
  } else {
    store.setCrumb([item.title])
  }
  store.addTab({
    name: item.name,
    title: item.title,
    parentTitle: parentItem?.title // 一级菜单的名称，如果为undefined，那么这个tab就是一级菜单
  })

  router.push({
    name: item.name
  })
}

function onOpenChange(openKeys) {
  const latestOpenKey = openKeys.find(key => data.openKeys.indexOf(key) === -1);
  data.openKeys = latestOpenKey ? [latestOpenKey] : [];
}

function loadRoute(name) {
  for (const menu of data.menuList) {
    if (menu.name === name) {
      selectedKeys.value = [menu.name]
      return
    }
    for (const subMenu of menu.children) {
      if (subMenu.name === name) {
        selectedKeys.value = [subMenu.name]

        data.openKeys = [menu.name]
        return
      }
    }
  }
}

onBeforeRouteUpdate((to, from, next) => {
  loadRoute(to.name)
  next()
})

loadRoute(route.name)

</script>

<style>
.ant-menu {
  background-color: var(--slide_bg);
  color: white;
}

.ant-menu-submenu-arrow {
  color: white;
}

.ant-menu-sub.ant-menu-inline {
  background-color: var(--slide_sub_bg);
}

.ant-menu-submenu-selected {
  color: inherit;
}

/*.ant-menu-inline .ant-menu-item {*/
/*  margin-top: 0;*/
/*  margin-bottom: 0;*/
/*}*/
</style>