<template>
    <div class="tablePage-style">
        <div class="table-nav">
            <a-form layout="inline" :model="SearchFrom">
                <a-form-item label="角色名称">
                    <a-input v-model:value="SearchFrom.name" placeholder="请输入角色名称" />
                </a-form-item>
                <a-form-item label="角色ID">
                    <a-input v-model:value="SearchFrom.role_id" placeholder="请输入角色ID" />
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
                <!-- <span>{{ t(column.title) }}</span> -->
                <span>{{ column.title }}</span>
            </template>
            <template #Description="{ text  }">
                <span>{{ text.String  }}</span>
            </template>
            <template #bodyCell="{ column, record }">
                <template v-if="column.dataIndex === 'operation'">
                    <a-space :size="8">
                        <a-button type="link" @click="on_redact(record)">{{ t('public.redact') }}</a-button>
                        <a-button type="link" danger @click="handleDelete(record.RoleId)">{{ t('public.delete') }}</a-button>
                    </a-space>
                </template>
            </template>
        </a-table>



        <a-modal v-model:open="addOpen" title="添加角色" :footer="null" :after-close="onCancel">
            <a-form :model="formState" name="basic" ref="formRef" :label-col="{ span: 6 }" :wrapper-col="{ span: 18 }"
                autocomplete="off" @finish="onFinish">
                <a-form-item label="角色名称" name="name"
                    :rules="[{ required: true, message: '请输入角色名称' }]">
                    <a-input v-model:value="formState.name" />
                </a-form-item>

                <a-form-item label="角色描述" name="description"
                    :rules="[{ required: true, message: '请输入角色描述' }]">
                    <a-input v-model:value="formState.description" />
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
    name: string
    role_id: number
    createTime: string
};

type formStateType = {
    role_id?: number,
    name: string,
    description: string,
    createByUid:number
}
type listItemType = {
    RoleId: number,
    Name: string,
    Description: {
        String:string
    }
}


const paramFormat = (data: any) => {
    return { ...data, role_id: Number(data.role_id) || 0 }
}

import { useTableHooks } from "@/Hooks/useTableHooks"
import { ref, reactive } from 'vue';
import { useI18n } from 'vue-i18n'
// import type { Dayjs } from 'dayjs';
import { Modal } from 'ant-design-vue';
import type { FormInstance } from 'ant-design-vue';
const { t } = useI18n()
import { listRole ,addRole,editRole,deleteRole} from "@/api/admin"
const formRef = ref<FormInstance>();
let { tabHeight, SearchFrom, on_search, paginationOpt, tableData, Fun_requestList } = useTableHooks<SearchType>({
    name: "",
    role_id: 0,
    createTime: ""
}, { listApi: listRole, paramFormat: paramFormat });
// deleteUser
const addOpen = ref<boolean>(false);
const formState = reactive<formStateType>({
    name: "",
    description: "",
    createByUid:1
});
const columns = [
    {
        title: '角色ID',
        dataIndex: 'RoleId',
    },
    {
        title: '角色名称',
        dataIndex: 'Name',
    },

    {
        title: '角色描述',
        dataIndex: 'Description',
        key: 'Description',
        slots: { customRender: 'Description' },
    },
    {
        title: '创建时间',
        dataIndex: 'CreateTime',
    },
    {
    title: 'public.operation',
    dataIndex: 'operation',
}
]

const on_redact = (data: listItemType) => { 
    addOpen.value = true
    formState.name = data.Name
    formState.description = data.Description.String
    formState.role_id = data.RoleId
} 

const onFinish = () => {
    let api
    if (formState.role_id) {
        api = editRole
    } else {
        api = addRole
    }
    api(formState).then(() => {
        addOpen.value = false
        Fun_requestList()
    })
};

    // 删除操作
    const handleDelete = (id: string) => {
        Modal.confirm({
            title: '删除操作',
            content: '确定要删除这一条记录吗？',
            async onOk() {
                try {
                    return await deleteRole({ role_id:id }).then(() => {
                        Fun_requestList()
                    })
                } catch {
                    return console.log('Oops errors!');
                }
            },
            // eslint-disable-next-line @typescript-eslint/no-empty-function
            onCancel() { },
        });
    };


const onCancel = () => {
    // console.log(formRef.value.resetFields)
    formRef.value!.resetFields();
    console.log(formState)
}
</script>

<style lang="less" scoped></style>