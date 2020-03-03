<template>
  <el-row type="flex" class="login-wrapper">
    <div class="left-wrapper">
      <el-row type="flex" align="middle" class="header-wrapper">
        <el-row type="flex" align="middle">
          <i class="logo" />
          <span class="text">墨娘的聊天室</span>
        </el-row>
      </el-row>

      <div v-if="loginStatus === 0" class="login-form-wrapper">
        <div class="title-wrapper pr-large">
          登录
        </div>
        <el-form ref="formHook">
          <el-form-item prop="phone">
            <el-input  v-model="login.user" placeholder="账号" >
            </el-input>
          </el-form-item>
          <el-form-item prop="password">
            <el-input  v-model="login.pass" type="password" placeholder="密码" v-on:keyup.enter="loginHandle()">
            </el-input>
          </el-form-item>
        </el-form>

        <el-row type="flex" align="middle" justify="end" class="footer-wrapper">
          <a  class="link base-font" style="margin-top: -10px;" @click="switchLoginStatus(1)">注册账号</a>
        </el-row>

        <el-button type="primary" style="width: 100%; margin-top: 60px;" @click="loginHandle()">
          登录
        </el-button>
      </div>

      <div v-if="loginStatus === 1" class="login-form-wrapper">
        <div class="title-wrapper pr-large">
          注册
        </div>
        <el-form ref="formHook">
          <el-form-item>
            <el-input  v-model="register.user" placeholder="账号" >
            </el-input>
          </el-form-item>
          <el-form-item>
            <el-input  v-model="register.pass" type="password" placeholder="密码">
            </el-input>
          </el-form-item>
          <el-form-item>
            <el-input  v-model="register.nick" placeholder="昵称" >
            </el-input>
          </el-form-item>
        </el-form>

        <el-row type="flex" align="middle" justify="end" class="footer-wrapper">
          <a  class="link base-font" style="margin-top: -10px;" @click="switchLoginStatus(0)">已有账号</a>
        </el-row>

        <el-button type="primary" style="width: 100%; margin-top: 60px;" @click="registerHandle()">
          确认注册
        </el-button>
      </div>

    </div>
    <el-row type="flex" align="middle" justify="center" class="right-wrapper">
      <div class="title">
        墨娘的聊天室
      </div>
      <div class="desc">
        一个初学者写的练手小项目
      </div>
      <a href="https://chat.imoniang.com/static/imoniangClient.zip" target="_blank" class="link-button pr-large">客户端下载</a>
    </el-row>
  </el-row>
</template>

<script>
  import {login, register} from "../api/user";

  export default {
    name: 'login',
    components: {
    },
    data(){
      return {
        loginStatus:0, // 0 为登录 1为注册
        login:{
          user: "",
          pass: "",
        },
        register:{
          user: "",
          pass: "",
          nick: ""
        }
      }
    },
    methods:{
      switchLoginStatus(status){
        this.loginStatus = status
        return false
      },
      loginHandle(){
        login(this.login).then(res=>{
          if(res.code !== 200){
            this.$message.warning(res.msg)
          }else{
            localStorage.setItem("token",res.data.token)
            localStorage.setItem("nick",res.data.nick)
            localStorage.setItem("fontColor",res.data.font_color)
            this.$message.success(res.msg)
            this.$router.push({name: 'chat'})
          }
        }).catch(()=>{})
      },
      registerHandle(){
        register(this.register).then(res=>{
          if(res.code !== 200){
            this.$message.warning(res.msg)
          }else{
            this.login.user = this.register.user
            this.$message.success(res.msg)
            this.loginStatus = 0
            this.register = {
              user: "",
              pass: "",
              nick: ""
            }
          }
        })
      }
    }
  }
</script>

<style lang="scss" scoped>
.link {
  color: #4381c6;
}

.login-wrapper {
  min-width: 1300px;
  height: 100%;
  min-height: 100%;
  background-color: #f4f7fa;

  .left-wrapper,
  .right-wrapper {
    flex: 1;
    width: 100%;
    height: 100%;
  }

  // 左容器内容
  .left-wrapper {
    .header-wrapper {
      height: 76px;
      padding: 0 50px;
      background-color: #ffffff;
      .logo {
        margin-right: 15px;
        background:url("../assets/images/logo.gif") no-repeat center center;
        background-size: cover;
        height: 50px;
        width: 50px;
      }

      .text {
        color: #7b7b7b;
        .link {
          margin-left: 10px;
        }
      }
    }

    .login-form-wrapper {
      width: 368px;
      margin: 120px auto 0;
      text-align: center;

      .title-wrapper {
        margin-bottom: 38px;
        color: #2b2b2b;
      }

      .icon {
        margin-left: 5px;
      }
    }
  }

  // 右容器内容
  .right-wrapper {
    flex-direction: column;
    background-color:#FFCCCC;
    // background: url('../assets/images/bg.jpg') no-repeat center center;
    background-size: cover;

    .title {
      font-size: 40px;
      color: white;
    }

    .desc {
      margin-top: 10px;
      font-size: 26px;
      color: white;
    }

    .link-button {
      margin-top: 30px;
      padding: 15px 45px;
      color: #2b2b2b;
      background-color: #ffffff;
      cursor: pointer;
      text-decoration: none;
    }
  }
}
</style>
