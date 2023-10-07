<template>
    <div class="tablePage-style">
        <div class="table-nav">
            <a-form layout="inline" :model="SearchFrom">
                <a-form-item :label="t('user.userName')" name="fieldA">
                    <a-input v-model:value="SearchFrom.user" placeholder="请输入用户名" />
                </a-form-item>
                <a-form-item :label="t('user.userName')" name="fieldA">
                    <a-date-picker v-model:value="SearchFrom.value1" />
                </a-form-item>

                <a-form-item>
                    <a-button type="primary" @click="on_search()">{{ t('public.Submit') }}</a-button>
                </a-form-item>
            </a-form>

            <a-space wrap>
                <a-button @click="addOpen = true" type="primary">{{ t('public.add') }}</a-button>
            </a-space>
        </div>
        <a-table class="table-style" :columns="columns" :data-source="data" :scroll="{ y: tabHeight }"
            :pagination="paginationOpt">
            <template #headerCell="{ column }">
                <span>{{ t(column.title) }}</span>
            </template>
        </a-table>



        <a-modal v-model:open="addOpen" title="添加用户" :footer=null>
            <a-form :model="formState" name="basic" :label-col="{ span: 6 }" :wrapper-col="{ span: 18 }" autocomplete="off"
                @finish="onFinish">
                <a-form-item :label="t('user.userName')" name="name"
                    :rules="[{ required: true, message: t('user.usernameVerification') }]">
                    <a-input v-model:value="formState.name" />
                </a-form-item>

                <a-form-item :label="t('user.password')" name="password"
                    :rules="[{ required: true, message: t('user.passwordVerification') }]">
                    <a-input v-model:value="formState.password" />
                </a-form-item>
                <a-form-item :wrapper-col="{ offset: 4, span: 16 }" >
                    <a-button type="primary" html-type="submit">{{ t('public.Submit') }}</a-button>
                </a-form-item>
            </a-form>
        </a-modal>
    </div>
</template>

<script setup lang="ts">
type SearchType = {
    user: string;
    value1?: Dayjs
};

type formStateType = {
    name: string;
    password: string;
}

// defineOptions({
//     name: 'SystemMonitorLoginLog',
//   });
import { useTableHooks } from "@/Hooks/useTableHooks"
import { ref, reactive } from 'vue';
import { useI18n } from 'vue-i18n'
import type { Dayjs } from 'dayjs';
const { t } = useI18n()
import {adduser} from "@/api/admin"

let { tabHeight, SearchFrom, on_search, paginationOpt } = useTableHooks<SearchType>({
    user: "",
}, '');
const addOpen = ref<boolean>(false);

const formState = reactive<formStateType>({
    name: '',
    password: '',
});
const columns = [{
    title: 'user.userName',
    dataIndex: 'name',
    key: 1
},
{
    title: 'user.userName',
    dataIndex: 'age',
    key: 2
}]
const data = [...Array(100)].map((_, i) => ({
    key: i,
    name: `Edward King ${i}`,
    age: 32,
    address: `London, Park Lane no. ${i}`,
}));
const onFinish = (values: formStateType) => {
    console.log(values)
    adduser(values).then((res:any)=>{
        console.log(res)
    })
};
</script>

<style lang="less" scoped></style>