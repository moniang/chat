import Vue from 'vue'
import App from './App.vue'
import Element from 'element-ui'
//import { Row, Form, Input, FormItem, Button,Backtop } from 'element-ui'
import 'element-ui/lib/theme-chalk/index.css'
import '@/styles/element-ui.scss'
import {Message} from "element-ui";

Vue.config.productionTip = false
Vue.prototype.$message = Message

//Vue.use(Form).use(Input).use(FormItem).use(Button).use(Row).use(Backtop)
Vue.use(Element)
import './styles/index.scss'
import router from './router'

router.beforeEach((to,from,next)=>{
  if(to.name === "chat"){
    if(!localStorage.getItem("token")){
      Message.warning("请先登录！")
      next({path:"/"})
      return
    }
  }
  next()
})
new Vue({
  router,
  render: h => h(App),
}).$mount('#app')
