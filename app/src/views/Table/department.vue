<template>
    <div class="tablePage-style">
        <div class="table-nav">
            <a-form layout="inline" :model="SearchFrom">

                <a-form-item label="部门ID">
                    <a-input v-model:value="SearchFrom.dept_id" placeholder="请输入部门ID" />
                </a-form-item>
                <a-form-item :label="t('department.name')">
                    <a-input v-model:value="SearchFrom.name" placeholder="请输入部门名称" />
                </a-form-item>

                <a-form-item label="部门位置">
                    <a-input v-model:value="SearchFrom.region" placeholder="请输入部门位置" />
                </a-form-item>



                <a-form-item label="上级部门">
                    <a-input v-model:value="SearchFrom.parentId" placeholder="请输入上级部门" />
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
                <span>{{ column.title }}</span>
            </template>
            <template #Description="{ text }">
                <span>{{ text.String }}</span>
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


        <a-modal v-model:open="addOpen" title="添加部门" :footer=null>
            <a-form :model="formState" name="basic" :label-col="{ span: 6 }" :wrapper-col="{ span: 16 }" autocomplete="off"
                @finish="onFinish">
                <a-form-item label="部门名称" name="name" :rules="[{ required: true, message: 'Please input your username!' }]">
                    <a-input v-model:value="formState.name" />
                </a-form-item>

                <a-form-item label="部门描述" name="description"
                    :rules="[{ required: true, message: 'Please input your password!' }]">
                    <a-input v-model:value="formState.description" />
                </a-form-item>

                <a-form-item label="部门位置" name="region"
                    :rules="[{ required: true, message: 'Please input your password!' }]">
                    <a-input v-model:value="formState.region" />
                </a-form-item>

                <a-form-item label="上级部门ID" name="parentId"
                    :rules="[{ required: true, message: 'Please input your password!' }]">
                    <a-input v-model:value="formState.parentId" />
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

interface FormState {
    dept_id?: number
    name: string;
    parentId: number
    description: string,
    childIds: string,
    region: string,
    createByUid: number
}

type listItemType = {
    id: string
    name: string
    describe: string
}
const paramFormat = (data: any) => {
    return { ...data, dept_id: Number(data.dept_id) || 0, parentId: Number(data.parentId) || 0 }
}

import { useTableHooks } from "@/Hooks/useTableHooks"
import { onMounted, ref, reactive } from 'vue';
import { addDepartment, editDepartment, listDepartment, deleteDepartment } from "@/api/admin"
import { useI18n } from 'vue-i18n'
import type { Dayjs } from 'dayjs';
const { t } = useI18n()
const addOpen = ref<boolean>(false);
// let listItem = reactive<listItemType>()

let { tabHeight, SearchFrom, on_search, paginationOpt, tableData, Fun_requestList } = useTableHooks<SearchType>({
    dept_id: 0,
    name: "",
    parentId: 0,
    region: ""
}, { listApi: listDepartment, paramFormat: paramFormat });


const formState = reactive<FormState>({
    name: "",
    description: "",
    childIds: "",
    parentId: 0,
    region: "",
    createByUid: 1
});


const columns = [
    {
        title: '部门ID',
        dataIndex: 'DeptId',

    },

    {
        title: '部门名称',
        dataIndex: 'Name',

    },

    {
        title: '部门描述',
        dataIndex: 'Description',
        key: 'Description',
        slots: { customRender: 'Description' },
    },
    {
        title: '部门地址',
        dataIndex: 'Region',
    },

    {
        title: '上级部门ID',
        dataIndex: 'ParentId',
    },

    {
        title: '下级部门ID',
        dataIndex: 'ChildIds',
    },


    {
        title: '创建时间',
        dataIndex: 'CreateTime',
    },
    {
        title: '操作',
        dataIndex: 'operation',
    }


]
const on_redact = (data: any) => {
    console.log(data.id)
    addOpen.value = true
    formState.name = data.Name
    formState.description = data.Description.String
    formState.dept_id = data.DeptId
    formState.parentId = data.ParentId
    formState.childIds = data.ChildIds
    formState.region = data.Region
    formState.createByUid = data.CreateByUid.int64

}


const onFinish = () => {

    let api
    if (formState.dept_id) {
        api = editDepartment
    } else {
        api = addDepartment
    }
    formState.parentId = formState.parentId ? Number(formState.parentId) : 0
    api(formState).then(() => {
        addOpen.value = false
        Fun_requestList()
    })
};

onMounted(() => {

})
</script>

<style lang="less" scoped></style>