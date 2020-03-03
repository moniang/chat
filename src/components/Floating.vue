<template>
    <div>
        <div class="floatingBox" ref="floatingBox">
            <div class="gif" v-bind:class="gifClassName" />
            <div class="floating" v-bind:class="bgClassName" >
                <span class="text">{{ text }}</span>
            </div>
        </div>
<!--        <transition name="slide-floating" v-on:enter="floatingBoxEnter" v-on:after-enter="floatingBoxAfterEnter">-->
<!--            -->
<!--        </transition>-->
    </div>
</template>

<script>
    export default {
        name: "Floating",
        data(){
            return{
                logoAnimateTimer:0,
                left:0,
                step:5,
                upTime:0,
                max:0,
                text:"",
                bgClassName:"",
                gifClassName:"",
                inHand:false,
                task:[]
            }
        },
        methods:{
            init(text,vip){
                if(this.inHand){ // 如果有飘屏正在进行，则压入队列
                    this.task.push({
                        text,vip
                    })
                    return
                }
                this.inHand = true
                if(vip > 3 && vip < 8){
                    this.bgClassName="vipBg-"+vip
                    this.gifClassName="gif-"+vip
                }else{
                    this.inHand = false
                    return
                }

                this.text = text
                let domLeft = window.document.body.clientWidth - 640
                let el = this.$refs.floatingBox
                el.style.display = "block"
                el.style.left = domLeft.toString() +"px"
                this.left = domLeft
                setTimeout(()=>{
                    this.move();
                },2000)
            },
            move(timestamp){
                this.logoAnimateTimer = requestAnimationFrame(() => {
                    if(timestamp - this.upTime > 30){
                        this.left = this.left - this.step;
                        if(this.left < 0){
                            this.left = 0
                        }

                        this.$refs.floatingBox.style.left = this.left + "px"
                        this.upTime = timestamp
                    }

                    if(this.left > this.max){
                        this.logoAnimateTimer = requestAnimationFrame(this.move)
                    }else{
                        cancelAnimationFrame(this.logoAnimateTimer)
                        setTimeout(()=>{
                            this.$refs.floatingBox.style.display = "none"
                            this.inHand = false // 修改任务进行状态
                            let task = this.task.pop();
                            if(task){
                                this.init(task.text,task.vip)
                            }
                        },3000)
                    }
                })
            },
            floatingBoxEnter(el){
                el.style.display = "block"
                el.style.left = (window.document.body.clientWidth - 640).toString() +"px"
            },
            floatingBoxAfterEnter(el){
                console.log(el)
            }
        }
    }
</script>

<style scoped>
    .floating{
        height:68px;
        width:520px;
        position:absolute;
        text-align:center;
        margin-left:70px;
    }
    .vipBg-7{
        background-repeat: no-repeat;
        background-image: url('../assets/vipenter-bg-7.webp');
    }
    .vipBg-6{
        background-repeat: no-repeat;
        background-image: url('../assets/vipenter-bg-6.webp');
    }
    .vipBg-5{
        background-repeat: no-repeat;
        background-image: url('../assets/vipenter-bg-5.webp');
    }
    .vipBg-4{
        background-repeat: no-repeat;
        background-image: url('../assets/vipenter-bg-4.webp');
    }
    .gif-7{
        background-repeat:no-repeat;
        background-image:url('../assets/vipenter-icon-7.webp');
    }
    .gif-6{
        background-repeat:no-repeat;
        background-image:url('../assets/vipenter-icon-6.webp');
    }
    .gif-5{
        background-repeat:no-repeat;
        background-image:url('../assets/vipenter-icon-5.webp');
    }
    .gif-4{
        background-repeat:no-repeat;
        background-image:url('../assets/vipenter-icon-4.webp');
    }
    .gif {
        height:68px;
        width:120px;
        float:left;
    }
    .text{
        margin-top:10px;
        font-size:14px;
        width:100%;
        position:relative;
        top:23px;
        color:#fff;
        font-weight:600;
        font-family: -apple-system,Microsoft Yahei,sans-serif;
    }
    .floatingBox{
        display:none;
        width:640px;
        position:absolute;
    }
</style>
