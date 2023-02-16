<template>
  <div class="sign-up">
    <h1>Create Account</h1>
    <div class="social-login">
      <el-icon><ChromeFilled /></el-icon>
      <el-icon><Eleme /></el-icon>
      <el-icon><ElemeFilled /></el-icon>
    </div>
    <span>或者使用你的邮箱进行注册</span>
    <el-form :mode="signInInfo">
      <el-form-item>
        <el-input placeholder="用户名" v-model="signInInfo.username" />
      </el-form-item>
      <el-form-item>
        <el-input placeholder="邮箱" v-model="signInInfo.email" />
      </el-form-item>
      <el-form-item>
        <el-input
          placeholder="密码"
          v-model="signInInfo.password"
          type="password"
        />
      </el-form-item>
      <el-form-item class="checkpassword-item">
        <el-input
          placeholder="再次输入密码"
          v-model="signInInfo.checkPassword"
          type="password"
          autocomplete="off"
        />
      </el-form-item>
      <div class="button-group">
        <div class="signup-tip">Have no count? SignUp</div>
        <el-button class="signin-button" type="primary">登录</el-button>
      </div>
    </el-form>
  </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";
interface SignInInfo {
  username: string;
  email: string;
  password: string;
  checkPassword: string;
}

export default defineComponent({
  name: "SignIn",
  data() {
    return {
      signInInfo: {
        username: "",
        email: "",
        password: "",
        checkPassword: "",
      },
    };
  },
});
</script>

<style scoped lang="scss">
.sign-up {
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  .checkpassword-item {
    margin: 0 0 5px;
  }
  .button-group {
    display: flex;
    flex-direction: column;
    // justify-content: space-between;
    padding: 0 5px;
    .signup-tip {
      text-align: left;
      font-size: 10px;
    }
    .signin-button {
      flex: auto;
      margin-top: 15px;
    }
  }
}
</style>
