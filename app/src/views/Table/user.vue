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
        <a-table class="table-style" :columns="columns" :data-source="tableData" :scroll="{ y: tabHeight }"
            :pagination="paginationOpt">
            <template #headerCell="{ column }">
                <span>{{ t(column.title) }}</span>
            </template>

            <template #bodyCell="{ column, record }">
                <template v-if="column.dataIndex === 'operation'">
                    <a-space :size="8">
                        <a-button type="link" @click="on_redact(record)">{{ t('public.redact') }}</a-button>
                        <a-button type="link" danger @click="handleDelete(record.id)">{{ t('public.delete') }}</a-button>
                    </a-space>
                </template>
            </template>
        </a-table>



        <a-modal v-model:open="addOpen" title="添加用户" :footer="null" :after-close="onCancel">
            <a-form :model="formState" name="basic" ref="formRef" :label-col="{ span: 6 }" :wrapper-col="{ span: 18 }"
                autocomplete="off" @finish="onFinish">
                <a-form-item :label="t('user.userName')" name="userName"
                    :rules="[{ required: true, message: t('user.usernameVerification') }]">
                    <a-input v-model:value="formState.userName" />
                </a-form-item>

                <a-form-item :label="t('user.password')" name="password"
                    :rules="[{ required: true, message: t('user.passwordVerification') }]">
                    <a-input v-model:value="formState.password" />
                </a-form-item>
                <a-form-item :wrapper-col="{ offset: 4, span: 16 }">
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
    id?: string;
    userName: string;
    password: string;
}
type listItemType = {
    id: string,
    userName: string,
    password: string
}
import { useTableHooks } from "@/Hooks/useTableHooks"
import { ref, reactive } from 'vue';
import { useI18n } from 'vue-i18n'
import type { Dayjs } from 'dayjs';
import type { FormInstance } from 'ant-design-vue';
const { t } = useI18n()
import { adduser, listUser, deleteUser ,edituser} from "@/api/admin"
const formRef = ref<FormInstance>();
let { tabHeight, SearchFrom, on_search, handleDelete, paginationOpt, tableData } = useTableHooks<SearchType>({
    user: "",
}, listUser, deleteUser);
const addOpen = ref<boolean>(false);
const formState = reactive<formStateType>({
    userName: '',
    password: '',
});
const columns = [{
    title: '用户名',
    dataIndex: 'userName',
},
{
    title: '修改时间',
    dataIndex: 'revampTime',
},

{
    title: 'public.operation',
    dataIndex: 'operation',
},]

const on_redact = (data: listItemType) => {
    console.log(data.userName)
    addOpen.value = true
    formState.userName = data.userName
    // formState.password = data.password
    formState.id = data.id
}


const onFinish = () => {
    let api
    if(formState.id){
        api = edituser
    }else {
        api = adduser
    }
    api(formState).then((res: any) => {
        console.log(res)
    })
};

const onCancel = () => {
    // console.log(formRef.value.resetFields)
    formRef.value!.resetFields();
    console.log(formState)
}
</script>

<style lang="less" scoped></style>