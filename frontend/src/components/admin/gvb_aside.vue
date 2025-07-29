<template>
  <aside class="gvb_aside">
    <div class="gvb_aside_header flex">
      <div class="gvb_aside_logo flex">
        <img src="http://127.0.0.1/uploads/file/%E5%BE%AE%E4%BF%A1%E5%9B%BE%E7%89%87_20241004202055.jpg" alt="" style="width: 50%; border-radius: 50%;">
      </div>
      <div class="gvb_aside_title">悟空后台管理</div>
    </div>
    <div class="gvb_aside_body flex">
      <GVBAsideMenu></GVBAsideMenu>
    </div>
  </aside>
</template>

<script setup>
import GVBAsideMenu from "./gvb_aside_menu.vue"
</script>

<style lang="scss">
.gvb_aside {
  width: 240px;
  height: 100vh;
  background-color: var(--slide);
  overflow-y: auto;

  .gvb_aside_header {
    width: 100%;
    height: 180px;
    flex-direction: column;
  }

  .gvb_aside_logo {
    padding: 10px;

    img {
      width: 80%;
    }
  }

  .gvb_aside_title {
    font-size: 20px;
    color: #e7e7e7;
  }
}
</style>