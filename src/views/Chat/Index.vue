<template >
    <div id="chat">
        <floating ref="floating" />
        <div ref="mainContent" class="main-content">
            <div class="content" v-html="allMessage">

            </div>
        </div>
        <div class="chat">
            <div class="chat-tools">
                <el-button icon="el-icon-setting" size="mini" style="padding: 7px 7px" @click="openUserSet()" title="用户设置"/>
                <el-button icon="el-icon-edit" size="mini" style="padding: 7px 7px" @click="openFontSet()" title="字体设置"/>
                <el-button icon="el-icon-delete" size="mini" style="padding: 7px 7px;float: right" @click="()=>{this.allMessage=''}" title="清空页面"/>
            </div>
            <textarea v-model="message" rows="6" cols="20" id = "message" v-on:keyup.enter="sendMsg"></textarea>
            <button class="send" @click="sendMsg()">
                <span id="send">发送(enter)</span>
            </button>
        </div>
        <el-dialog title="用户设置" :visible.sync="userSetVisible" width="500px">
            <el-form ref="form" label-width="80px">
                <el-form-item label="用户昵称">
                    <el-input v-model="revise.nick"></el-input>
                </el-form-item>
                <el-form-item>
                    <el-button type="primary" @click="onSetSubmit">保存</el-button>
                    <el-button @click="cancel()">取消</el-button>
                </el-form-item>
            </el-form>
        </el-dialog>
        <el-dialog title="字体设置" :visible.sync="fontSetVisible" width="500px">
            <el-form>
                <el-form-item label="字体颜色">
                    <el-color-picker v-model="revise.fontColor" id="selectColor" size="mini"></el-color-picker>
                </el-form-item>
                <el-form-item>
                    <el-button type="primary" @click="onFontSubmit">保存</el-button>
                    <el-button @click="cancel()">取消</el-button>
                </el-form-item>
            </el-form>
        </el-dialog>
    </div>
</template>

<script>
    import {reviseName,reviseFontColor} from "../../api/user";
    import Floating from "../../components/Floating";
    export default {
        name: "chat",
        components:{
            Floating
        },
        data(){
            return {
                show:false,
                ws: undefined,
                token:'',
                nick:'',
                message : '',
                msgId:-1, // 最近消息ID
                allMessage:'', // 全部消息
                timeHwnd:false, // 心跳包定时器
                tryTime: 0, // 掉线重登次数
                socketAddr: '',
                userSetVisible:false,
                fontSetVisible:false,
                revise:{
                    nick:"",
                    fontColor:"#000000",
                },
                fontColor:"#000000",
            }
        },
        watch:{
            allMessage:{
                handler(){
                    this.$nextTick(()=>{
                        this.$refs.mainContent.scrollTop = this.$refs.mainContent.scrollHeight
                    })
                }
            }
        },
        mounted() {
            this.token = localStorage.getItem('token')
            this.nick = localStorage.getItem('nick')
            this.fontColor = localStorage.getItem("fontColor")
            if(process.env.NODE_ENV === 'production'){
                this.socketAddr = "wss://"+ window.location.host +"/ws?token=" +this.token
            }else{
                this.socketAddr = "ws://127.0.0.1:8080/ws?token=" +this.token
            }

            this.ws = new WebSocket(this.socketAddr)
            this.ws.onopen = () => {
                console.log("连接成功")
                this.tryTime =  0
                if(this.timeHwnd){
                    clearInterval(this.timeHwnd)
                    this.timeHwnd = false
                }

                this.timeHwnd = setInterval(()=>{
                    if(this.ws.readyState === WebSocket.OPEN){
                        this.ws.send('ping')
                    }else{
                        clearInterval(this.timeHwnd)
                        this.timeHwnd = false
                    }
                },15000)

                this.ws.onmessage = e => {
                    let msgInfo = JSON.parse(e.data)
                    if(msgInfo.id === -1){
                        this.$notify({
                            message: '您的账号在别处登录，请重新登录',
                            duration: 0,
                            type:"waring"
                        });
                        localStorage.clear() // 清空本地储存的token
                        setTimeout(function () {
                            location.href = "/"
                        },1000)
                        return
                    }else if(msgInfo.id === -2){
                        this.makeAddMessage(msgInfo.nick)
                        let vip = parseInt(msgInfo.message)
                        if(vip > 3 && vip < 8){
                            this.$refs.floating.init("欢迎尊贵的"+msgInfo.nick+"进入本聊天室",msgInfo.message)
                        }
                        return
                    }
                    if(this.msgId !== msgInfo.id){ // 如同一人连续发信息，则无需再加载名字
                        this.makeOtherName(msgInfo.nick,msgInfo.id)
                    }
                    this.makeContent(msgInfo.message,msgInfo.font_color)
                }

                this.ws.onerror = e =>{
                    this.$notify({
                        message: "您已经与服务器断开连接，错误代码：" + e.code,
                        duration: 0,
                        type:"waring"
                    });
                    setTimeout(function () {
                        location.href = "/"
                    },1000)
                }

                this.ws.onclose= () => {
                    this.tryTime ++;
                    if(this.tryTime > 3){
                        alert("与服务器断开连接")
                        return
                    }
                    this.ws = new WebSocket(this.socketAddr)
                }
            }
        },
        beforeDestroy() {
          this.ws.close()
        },
        methods:{
            openFontSet(){
              this.fontSetVisible = true
            },
            onFontSubmit(){
                reviseFontColor(this.revise).then(res=>{
                    if(res.code === 200){
                        this.$message.success("修改成功")
                        this.fontColor = this.revise.fontColor
                        localStorage.setItem("fontColor",this.revise.fontColor)
                        this.fontSetVisible = false
                    }else{
                        this.$message.warning(res.msg)
                    }
                })
            },
            onSetSubmit(){
                if(this.revise.nick.length === 0){
                    this.$message.warning("请输入用户昵称")
                    return
                }
                if(this.revise.nick === this.nick){
                    this.$message.success("修改成功")
                    this.userSetVisible = false
                    return
                }
                reviseName(this.revise).then(res=>{
                    if(res.code !== 200){
                        this.$message.warning(res.msg)
                    }else{
                        this.$message.success(res.msg)
                        localStorage.setItem("nick",this.revise.nick)
                        this.nick = this.revise.nick
                        this.revise.nick = ""
                        this.userSetVisible = false
                    }
                }).catch(()=>{})
            },
            cancel(){
                this.revise.nick = ""
                this.userSetVisible = false
                this.fontSetVisible = false
            },
            openUserSet(){
              this.userSetVisible = true
            },
            sendMsg(){
                if(this.message.length === 0){
                    this.$message.warning("发送的消息不可为空")
                    return
                }
                if(this.message.length > 2000){
                    this.$message.warning("消息长度过长!")
                    return
                }

                if(this.msgId !== 0){
                    this.makeMyName(this.nick)
                }

                if(this.ws.readyState !== WebSocket.OPEN){
                    this.$message.warning("已和服务器断开，请重新登录")
                    return
                }
                let msg = {message:this.message}
                this.ws.send(JSON.stringify(msg))


                this.makeContent(this.makeMessage(this.message),this.fontColor)
                this.message = ""
            },
            makeMyName(name){
                this.msgId = 0;
                let now = new Date();
                let hour = now.getHours();
                if(hour.toString().length === 1) hour = "0" + hour
                let minu = now.getMinutes();
                if(minu.toString().length === 1) minu = "0" + minu
                let sec = now.getSeconds();
                if(sec.toString().length === 1) sec = "0" + sec
                const htmlstr = "<span class='content-name name-me'>" + name + " "+hour+":"+minu+":"+sec+"</span>";
                this.allMessage += htmlstr
            },
            makeOtherName(name,user_id){
                this.msgId = user_id;
                let now = new Date();
                let hour = now.getHours();
                if(hour.toString().length === 1) hour = "0" + hour
                let minu = now.getMinutes();
                if(minu.toString().length === 1) minu = "0" + minu
                let sec = now.getSeconds();
                if(sec.toString().length === 1) sec = "0" + sec
                const htmlstr = "<span class='content-name name-other'>" + name + " "+hour+":"+minu+":"+sec+"</span>";
                this.allMessage += htmlstr
            },
            makeAddMessage(name){
                let str="<span class='content-msg content-notice'><span>" +name+ "加入了聊天室</span></span>"
                this.allMessage += str
            },
            makeContent(message,color){
                let str = "<span class='content-msg'><li class='li-show'><span style='color:"+color+"'>"+message+"</span></li></span>"
                this.allMessage += str
            },
            makeMessage(message){
                message = message.replace(/&/g,"&amp;")
                message = message.replace(/</g,"&lt;")
                message = message.replace(/>/g,"&gt;")
                return message
            }
        }
    }
</script>

<style>
    html,
    body,
    #app {
        background-color: #4a4a4a !important;
    }
</style>

<style scoped>
@import "index.css";
</style>
