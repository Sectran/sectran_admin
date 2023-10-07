<template>
    <div class="Login">
        <div class="Login-from">
            <div class="Login-title">{{ uselocals('login.login') }}</div>
            <a-form :model="formState" layout="vertical" name="basic" :label-col="{ span: 8 }" autocomplete="off"
                @finish="onFinish" @finishFailed="onFinishFailed">
                <!-- :rules="[{ required: true, message: 'Please input your username!' }]" -->
                <a-form-item :label="uselocals('login.userName')" name="username">
                    <a-input v-model:value="formState.username" />
                </a-form-item>
                <!-- :rules="[{ required: true, message: 'Please input your password!' }]" -->
                <a-form-item :label="uselocals('login.password')" name="password">
                    <a-input-password v-model:value="formState.password" />
                </a-form-item>

                <a-form-item name="remember" :wrapper-col="{ offset: 8, span: 16 }">
                    <a-checkbox v-model:checked="formState.remember">{{ uselocals('login.record') }}</a-checkbox>
                </a-form-item>
                <a-form-item :wrapper-col="{ offset: 8, span: 16 }">
                    <a-button type="primary" html-type="submit">{{ uselocals('public.Submit') }}</a-button>
                </a-form-item>
            </a-form>
        </div>
    </div>
</template>

<script setup lang="ts">
import { reactive } from 'vue';
import { uselocals } from "@/Hooks/localsHooks"
// import { useRouter } from 'vue-router';
import { login } from "@/api/Login"
// const router = useRouter();
interface FormState {
    username: string;
    password: string;
    remember: boolean;
}

const formState = reactive<FormState>({
    username: '',
    password: '',
    remember: true,
});

const onFinish = (values: { username: string, password: string }) => {
    console.log('Success:', values);

    login<{ username: string, password: string }>({ password: values.password, username: values.username }).then((res: { data: { token: string } }) => {
        console.log(res)
        let { token } = res.data
        localStorage.setItem('token', token)
    })
    // router.replace('/')                                                                                                                 
};

const onFinishFailed = (errorInfo: any) => {
    console.log('Failed:', errorInfo);
};
</script>

<style lang="less" scoped>
.Login {
    position: relative;
    width: 100vw;
    height: 100vh;

    .Login-title {
        font-size: 26px;
        text-align: center;
        margin-bottom: 20px;
    }

    .Login-from {
        width: 300px;
        padding: 20px;
        border-radius: 20px;
        position: absolute;
        top: 50%;
        left: 50%;
        transform: translate(-50%, -50%);
        box-shadow: 0px 6px 24px 1px rgba(110, 110, 110, 0.18);
    }
}
</style>