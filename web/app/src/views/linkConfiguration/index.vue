<template>
    <div class="configuration-style">
        <div class="configuration-nav"></div>
        <a-row class="Content-style">
            <a-col :span="4" class="Content-left">
                <a-directory-tree multiple default-expand-all>
                    <a-tree-node key="0-0" title="parent 0">
                        <a-tree-node key="0-0-0" title="leaf 0-0" is-leaf />
                        <a-tree-node key="0-0-1" title="leaf 0-1" is-leaf />
                    </a-tree-node>
                    <a-tree-node key="0-1" title="parent 1">
                        <a-tree-node key="0-1-0" title="leaf 1-0" is-leaf />
                        <a-tree-node key="0-1-1" title="leaf 1-1" is-leaf />
                    </a-tree-node>
                </a-directory-tree>
            </a-col>
            <a-col :span="20">
                <div id="terminal" ref="terminal"></div>
            </a-col>
        </a-row>
    </div>
</template>

<script setup lang='ts'>
import { onBeforeMount, onMounted, watchEffect, ref, reactive } from 'vue';
import { Terminal } from "xterm";
import { FitAddon } from "xterm-addon-fit";
import "xterm/css/xterm.css";
let terminal = ref(null)

let path = ref<string>('ws://127.0.0.1:19527')
let term = reactive<any>({})
let socket = reactive<any>({})
// ref, reactive, 
// import { useStore } from 'vuex';
// import { useRoute, useRouter } from 'vue-router';
// const store = useStore();
// const route = useRoute();
// const router = useRouter();
// const data = reactive({})
onBeforeMount(() => {
    //console.log('2.组件挂载页面之前执行----onBeforeMount')
})
onMounted(() => {
    initXterm()
    initSocket()
    //console.log('3.-组件挂载到页面之后执行-------onMounted')
})

const initXterm = () => {
    let terms: any = new Terminal({
        rendererType: "canvas", //渲染类型
        // rows: _this.rows, //行数
        // cols: _this.cols, // 不指定行数，自动回车后光标从下一行开始
        convertEol: true, //启用时，光标将设置为下一行的开头
        disableStdin: false, //是否应禁用输入
        cursorBlink: true, //光标闪烁
        theme: {
            foreground: "#ECECEC", //字体
            background: "#000000", //背景色
            cursor: "help", //设置光标
            lineHeight: 20,
        },
    });
    // 创建terminal实例
    terms.open(terminal.value);
    // 换行并输入起始符 $
    terms.prompt = (_: any) => {
        terms.write("\r\n\x1b[33m$\x1b[0m ");
    };
    // canvas背景全屏
    const fitAddon = new FitAddon();
    terms.loadAddon(fitAddon);
    fitAddon.fit();

    window.addEventListener("resize", resizeScreen);
    function resizeScreen() {
        try {
            fitAddon.fit();
        } catch (e: any) {
            console.log("e", e.message);
        }
    }
    term = terms
    runFakeTerminal();
}
const runFakeTerminal = () => {
    if (term._initialized) return;
    // 初始化
    term._initialized = true;
    term.onData((raw: any) => {
        write(raw);
    });
}
const write = (data: any) => {
    socket.send(stringToUint8Array(data));
}
const stringToUint8Array = (str: string) => {
    console.log(str)
    var arr = [];
    for (var i = 0, j = str.length; i < j; ++i) {
        arr.push(str.charCodeAt(i));
    }
    var tmpUint8Array = new Uint8Array(arr);
    return tmpUint8Array;
}

const initSocket = () => {
    if (typeof WebSocket === "undefined") {
        alert("您的浏览器不支持socket");
    } else {
        var websocket = new WebSocket(path.value);
        websocket.onopen = () => {
            onOpen();
        };
        websocket.onmessage = (evt) => {
            onData(evt.data);
        };
        websocket.onerror = (evt) => {
            onError(evt);
        };
        websocket.onclose = (evt) => {
            console.log(evt)
            onClose(evt);
        };
        websocket.onclose = () => {
            onClose();
        };
        socket = websocket;
    }
}
const onData = (msg: any) => {
    msg.text().then((data: any) => {
        term.write(data);
    })
        .catch((err: string) => {
            console.log(err);
        });
}
const onOpen = () => {
    console.log("socket连接成功");
}
const onError = (evt?: Event) => {
    console.log(evt)
    console.log("socket连接错误");
}

// const onSend = () => {
//     //   this.socket.send();
// }
const onClose = (evt?: CloseEvent) => {
    console.log(evt)
    console.log("socket已经关闭");
}
watchEffect(() => {
})


</script>
<style scoped lang='less'>
.configuration-style {
    width: 100%;
    height: 100vh;

    .configuration-nav {
        height: 30px;
        background: #463E3E;
    }

    .Content-style {
        height: calc(100vh - 30px);

        .Content-left {
            background: #2F2A2A;
            padding: 20px;


            ::v-deep(.ant-tree-list) {
                background: #2F2A2A;
                color: #ffffff;
            }
        }
    }
}
</style>