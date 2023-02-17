<template>
  <div class="login-view">
    <div class="container" ref="LoginViewContainer">
      <div class="form-container sign-up-container">
        <el-form
          :rules="signUpRelus"
          :model="signUpInfo"
          ref="signUpForm"
          status-icon
        >
          <h1>创建账号</h1>
          <div class="social-container">
            <el-icon class="icon"><ChromeFilled /></el-icon>
            <el-icon class="icon"><Eleme /></el-icon>
            <el-icon class="icon"><ElemeFilled /></el-icon>
          </div>
          <span class="tip">或者使用你自己的邮箱进行登录</span>
          <el-form-item class="form-item" prop="username">
            <el-input
              placeholder="用户名"
              v-model="signUpInfo.username"
              :prefix-icon="icon.user"
            />
          </el-form-item>
          <el-form-item class="form-item" prop="email">
            <el-input
              placeholder="邮箱"
              v-model="signUpInfo.email"
              type="email"
              :prefix-icon="icon.email"
            />
          </el-form-item>
          <el-form-item class="form-item" prop="password">
            <el-input
              placeholder="密码"
              v-model="signUpInfo.password"
              type="password"
              :prefix-icon="icon.password"
              autocomplete="off"
              @paste.capture.prevent="handlePaste"
              @copy.capture.prevent="handlePaste"
              @cut.capture.prevent="handlePaste"
            />
          </el-form-item>
          <el-form-item class="form-item" prop="checkPassword">
            <el-input
              placeholder="再次输入密码"
              v-model="signUpInfo.checkPassword"
              type="password"
              autocomplete="off"
              :prefix-icon="icon.password"
              @paste.capture.prevent="handlePaste"
              @copy.capture.prevent="handlePaste"
              @cut.capture.prevent="handlePaste"
            />
          </el-form-item>
          <div class="button-group">
            <div class="signup-tip">以及拥有密码?</div>
            <el-button class="login-button signup-button">注 册</el-button>
          </div>
        </el-form>
      </div>
      <div class="form-container sign-in-container">
        <el-form
          :rules="signInRelus"
          :model="signInInfo"
          ref="signInForm"
          status-icon
        >
          <h1>登 录</h1>
          <div class="social-container">
            <el-icon class="icon"><ChromeFilled /></el-icon>
            <el-icon class="icon"><Eleme /></el-icon>
            <el-icon class="icon"><ElemeFilled /></el-icon>
          </div>
          <span class="tip">或者使用你已经注册的账号</span>
          <el-form-item class="form-item" prop="email">
            <el-input
              placeholder="邮箱"
              v-model="signInInfo.email"
              type="email"
              :prefix-icon="icon.email"
            />
          </el-form-item>
          <el-form-item class="form-item" prop="password">
            <el-input
              placeholder="密码"
              v-model="signInInfo.password"
              type="password"
              :prefix-icon="icon.password"
              autocomplete="off"
              @paste.capture.prevent="handlePaste"
              @copy.capture.prevent="handlePaste"
              @cut.capture.prevent="handlePaste"
            />
          </el-form-item>
          <div class="button-group">
            <div class="signin-tip">忘记密码?</div>
            <el-button class="login-button signin-button" type="primary"
              >登 录</el-button
            >
          </div>
        </el-form>
      </div>
      <div class="overlay-container">
        <div class="overlay">
          <div class="overlay-panel overlay-left">
            <router-link to="/api/home">
              <el-icon class="me-logo" :size="80"><MyLogo></MyLogo></el-icon>
            </router-link>
            <h1>嗨, 我的朋友</h1>
            <p>请输入你的个人信息进行注册，然后开始本博客的体验吧！</p>
            <el-button
              class="login-button ghost-button"
              type="primary"
              @click="signInMoveClick"
              >登 录</el-button
            >
          </div>
          <div class="overlay-panel overlay-right">
            <router-link to="/api/home">
              <el-icon class="me-logo" :size="80"><MyLogo></MyLogo></el-icon>
            </router-link>
            <h1>欢迎回来，我的朋友</h1>
            <p>为了进入本博客系统，请使用你已经注册的个人信息进行登录吧！</p>
            <el-button
              class="login-button ghost-button"
              type="primary"
              @click="signUpMoveClick"
              >注 册</el-button
            >
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import { User, Message, Lock } from "@element-plus/icons-vue";
import { ElMessage } from "element-plus";
import { FormInstance } from "element-plus";
import MyLogo from "@/components/imgs/icons/MyLogo.vue";

export default defineComponent({
  name: "LoginView",
  props: {
    mode: {
      type: String,
      default: "signIn",
    },
  },
  components: {
    MyLogo,
  },
  mounted() {
    if (this.mode == "signin") {
      this.signInClassChange();
    } else if (this.mode == "signup") {
      this.signUpClassChange();
    } else {
      ElMessage({
        message: "路由错误",
        type: "error",
      });
    }
  },
  unmounted() {
    let classList = (this.$refs.LoginViewContainer as HTMLElement)?.classList;
    if (classList) {
      classList.remove("right-panel-active");
    }
  },
  data() {
    return {
      loginMode: this.mode,
      signUpInfo: {
        username: "",
        email: "",
        password: "",
        checkPassword: "",
      },
      signInInfo: {
        email: "",
        password: "",
      },
      icon: {
        user: User,
        email: Message,
        password: Lock,
      },
    };
  },
  computed: {
    signUpRelus() {
      return {
        username: [
          { required: true, message: "请输入姓名", trigger: "blur" },
          {
            min: 2,
            max: 10,
            message: "长度在 2 到 10 个字符之间",
            trigger: "blur",
          },
          {
            required: true,
            pattern: /^[\u4e00-\u9fa5_a-zA-Z0-9.·-]+$/,
            message: "用户名不支持特殊字符",
          },
        ],
        email: [
          { required: true, message: "请输入邮箱", trigger: "blur" },
          {
            required: true,
            pattern: /^\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*$/,
            message: "请输入正确格式的邮箱",
          },
        ],
        password: [
          { required: true, message: "请输入密码", trigger: "blur" },
          {
            required: true,
            pattern:
              /^(?=.*[A-Za-z])(?=.*\d)(?=.*[$@$!%*#?&])[A-Za-z\d$@$!%*#?&]{8,}$/,
            message: "请输入正确的密码",
            trigger: "blur",
          },
          {
            min: 8,
            max: 16,
            message: "密码长度必须在8-16之间",
            trigger: "blur",
          },
        ],
        checkPassword: [
          { required: true, message: "请输入密码", trigger: "blur" },
          {
            validator: (rule: any, value: any, callback: any) => {
              console.log(rule);
              if (value !== this.signUpInfo.password) {
                callback(new Error("请保证两次输入的密码一致"));
              } else {
                callback();
              }
            },
            message: "请保证两次密码输入相同",
          },
        ],
      };
    },
    signInRelus() {
      return {
        email: this.signUpRelus.email,
        password: this.signUpRelus.password,
      };
    },
  },
  methods: {
    signInClassChange() {
      let classList = (this.$refs.LoginViewContainer as HTMLElement)?.classList;
      if (classList) {
        classList.remove("right-panel-active");
      }
    },
    signInMoveClick() {
      this.resetSignUpForm();
      this.$router.replace("/login/signin");
      this.signInClassChange();
    },
    signUpClassChange() {
      let classList = (this.$refs.LoginViewContainer as HTMLElement)?.classList;
      if (classList) {
        classList.add("right-panel-active");
      }
    },
    signUpMoveClick() {
      this.resetSignInForm();
      this.$router.replace("/login/signup");
      this.signUpClassChange();
    },
    resetSignInForm() {
      let form = this.$refs.signInForm as FormInstance;
      form?.resetFields(Object.keys(this.signInInfo));
    },
    resetSignUpForm() {
      let form = this.$refs.signUpForm as FormInstance;
      form?.resetFields(Object.keys(this.signUpInfo));
    },
    handlePaste(event: Event) {
      console.log((event.target as HTMLElement)?.tagName); // event 是dom原生事件的
      return false;
    },
  },
});
</script>

<style scoped lang="scss">
@import "@/assets/css/common.scss";
.login-view {
  display: flex;
  flex-direction: column;
  flex: auto;
  position: absolute;
  top: 0;
  left: 0;
  height: 100%;
  width: 100%;
  align-items: center;
  justify-content: center;
  background-color: $main-bg-color;
  h1 {
    font-weight: bold;
    margin: 0;
  }
  p {
    font-size: 14px;
    font-weight: 100;
    line-height: 20px;
    letter-spacing: 0.5px;
    margin: 20px 0 30px;
  }

  @keyframes signin-disapper {
    0%,
    59.99% {
      opacity: 1;
      z-index: 2;
    }

    60% {
      opacity: 0.1;
      z-index: 2;
    }

    100% {
      opacity: 0;
      z-index: -1;
    }
  }

  @keyframes signup-disapper {
    0%,
    59.99% {
      opacity: 1;
      z-index: 5;
    }
    60% {
      opacity: 0.1;
      z-index: 5;
    }

    100% {
      opacity: 0;
      z-index: 1;
    }
  }

  @keyframes show {
    0%,
    49.99% {
      opacity: 0;
      z-index: 1;
    }

    50% {
      opacity: 1;
      z-index: 5;
    }
    100% {
      opacity: 1;
      z-index: 5;
    }
  }

  .ghost-button {
    background-color: transparent;
    border-color: #ffffff;
  }

  .login-button {
    border-radius: 25px;
    font-size: 15px;
    font-weight: bold;
    padding: 20px 55px;
    width: 130px;
    letter-spacing: 1px;
    transition: transform 80ms ease-in;
    text-align: center;
    border: 2px solid #ffffff;
    color: #ffffff;
    background-color: #ff4b2b;
    &:active {
      transform: scale(0.95);
    }
    &:focus {
      outline: none;
    }
  }
  .container {
    box-sizing: border-box;
    background-color: #fff;
    border-radius: 10px;
    box-shadow: 0 15px 20px rgba(80, 79, 79, 0.25),
      0 10px 10px rgba(68, 66, 66, 0.034);
    position: relative;
    overflow: hidden;
    width: 768px;
    max-width: 100%;
    min-height: 520px;
    .form-container {
      position: absolute;
      display: flex;
      flex-direction: column;
      align-items: center;
      justify-content: center;
      top: 0;
      height: 100%;
      transition: all 0.6s ease-in-out;
      .tip {
        font-size: 10px;
        display: block;
        margin-bottom: 5px;
      }
      .form-item {
        width: 250px;
        ::v-deep .el-input__inner {
          height: 35px;
        }
      }
    }
    .sign-in-container {
      left: 0;
      width: 50%;
      z-index: 2;
      .button-group {
        margin-top: 20px;
        .signin-tip {
          font-size: 15px;
        }
        .signin-button {
          margin-top: 10px;
          border: 2px solid #ff416c;
        }
      }
    }
    .sign-up-container {
      left: 0;
      width: 50%;
      opacity: 0;
      z-index: 1;
      animation: signup-disapper 0.6s;
      .button-group {
        margin-top: 10px;
        .signin-tip {
          font-size: 15px;
        }
        .signup-button {
          margin-top: 10px;
          border: 2px solid #ff416c;
        }
      }
    }

    .overlay-container {
      position: absolute;
      top: 0;
      left: 50%;
      width: 50%;
      height: 100%;
      overflow: hidden;
      transition: transform 0.6s ease-in-out;
      z-index: 100;
      .overlay {
        background: #ff416c;
        background: -webkit-linear-gradient(to right, #ff4b2b, #ff416c);
        background: linear-gradient(to right, #ff4b2b, #ff416c);
        background-repeat: no-repeat;
        background-size: cover;
        background-position: 0 0;
        color: #ffffff;
        position: relative;
        left: -100%;
        height: 100%;
        width: 200%;
        transform: translateX(0);
        transition: transform 0.6s ease-in-out;
        .overlay-panel {
          .me-logo {
            margin-bottom: 40px;
            &:active {
              transform: scale(0.95);
            }
            cursor: pointer;
            // box-shadow: -5px 0 2px 2px rgba(0, 0, 0, 0.25);
            border-radius: 50%;
          }
          box-sizing: border-box;
          position: absolute;
          display: flex;
          align-items: center;
          justify-content: center;
          flex-direction: column;
          padding: 0 40px;
          text-align: center;
          top: 0;
          height: 100%;
          width: 50%;
          transform: translateX(0);
          transition: transform 0.6s ease-in-out;
          &.overlay-left {
            transform: translateX(-20%);
          }
          &.overlay-right {
            right: 0;
            transform: translateX(0);
          }
        }
      }
    }

    .social-container {
      & > .icon {
        margin: 0 5px;
        border: 1px solid #dddddd;
        border-radius: 50%;
        padding: 8px;
        display: inline-flex;
        justify-content: center;
        align-items: center;
      }
      display: flex;
      justify-content: center;
      align-items: center;
      font-size: 22px;
      margin: 20px 0 25px;
    }

    .social-container a {
      border: 1px solid #dddddd;
      border-radius: 50%;
      display: inline-flex;
      justify-content: center;
      align-items: center;
      margin: 0 5px;
      height: 40px;
      width: 40px;
    }

    &.right-panel-active {
      .sign-in-container {
        transform: translateX(100%);
        opacity: 0;
        z-index: -1;
        animation: signin-disapper 0.6s;
      }
      .sign-up-container {
        transform: translateX(100%);
        opacity: 1;
        z-index: 5;
        animation: show 0.6s;
      }
      .overlay-container {
        transform: translateX(-100%);
        .overlay {
          transform: translateX(50%);
          .overlay-panel {
            &.overlay-left {
              transform: translateX(0);
            }
            &.overlay-right {
              transform: translateX(20%);
            }
          }
        }
      }
    }
  }
}
</style>
