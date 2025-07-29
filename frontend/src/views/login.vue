<template>
  <div class="gvb_login_bg">
    <div class="gvb_login_container">
      <div class="gvb_login_main">
        <div class="gvb_login_head">用户登录</div>
        <div class="gvb_login_form">
          <div class="gvb_login_form_item">
            <a-input placeholder="用户名" v-model:value="data.user_name">
              <template #prefix><i class="fa fa-user-o"></i></template>
            </a-input>
          </div>
          <div class="gvb_login_form_item">
            <a-input type="password" placeholder="密码" v-model:value="data.password">
              <template #prefix><i class="fa fa-key"></i></template>
            </a-input>
          </div>
          <div class="gvb_login_form_item">
            <a-button type="primary" @click="emailLogin">登录</a-button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import {reactive} from "vue";
import {message} from 'ant-design-vue';
import {emailLoginApi, getQQLoginLinkApi, qqLoginApi} from "@/api/user_api";
import {parseToken} from "@/utils/jwt";
import {useStore} from "@/stores/store";
import {useRouter, useRoute} from "vue-router";

const router = useRouter()
const store = useStore()
const data = reactive({
  user_name: "",
  password: "",
})


async function emailLogin() {
  if (data.user_name.trim() === "") {
    message.error("请输入用户名")
    return
  }
  if (data.password.trim() === "") {
    message.error("请输入密码")
    return
  }

  let res = await emailLoginApi(data)
  if (res.code) {
    message.error(res.msg)
    return
  }
  // res.data就是jwt
  message.success(res.msg)
  let userInfo = parseToken(res.data)
  userInfo.token = res.data
  store.setUserInfo(userInfo)
  router.push("/admin/home")

  return

}



</script>

<style lang="scss">
.gvb_login_bg {
  background: url("/bg.png") 50%/cover no-repeat;;
  width: 100%;
  height: 100vh;
}

.gvb_login_container {
  position: fixed;
  right: 0;
  top: 0;
  width: 400px;
  height: 100vh;
  background-color: rgba(white, 0.8);
  display: flex;
  justify-content: center;
  align-items: center;

  .gvb_login_main {
    width: 75%;

    .gvb_login_head {
      font-size: 22px;
      font-weight: 600;
      color: var(--active);
      text-align: center;
      margin-bottom: 10px;
    }

    .gvb_login_form {
      .gvb_login_form_item {
        margin-bottom: 10px;

        .ant-btn {
          width: 100%;
        }
      }
    }

    .gvb_login_other {
      font-size: 14px;
      color: #666;
      display: flex;
      justify-content: center;
      align-items: center;

      &::before {
        width: 34%;
        height: 1px;
        background-color: #666;
        display: inline-block;
        content: "";
        margin-right: 5px;
      }

      &::after {
        width: 34%;
        height: 1px;
        background-color: #666;
        display: inline-block;
        content: "";
        margin-left: 5px;
      }
    }

    .gvb_login_other_icons {
      margin-top: 10px;
      display: flex;
      justify-content: center;
      align-items: center;

      .hvb_login_ohter_icon {
        width: 40px;
        height: 40px;
        cursor: pointer;
        margin-right: 10px;

        &:last-child {
          margin-right: 0;
        }
      }
    }
  }
}
</style>