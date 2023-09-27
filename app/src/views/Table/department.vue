<template>
    <div class="tablePage-style">
        <div class="table-nav">
            <a-form layout="inline" :model="SearchFrom">
                <a-form-item :label="t('department.Name')" name="fieldA">
                    <a-input v-model:value="SearchFrom.user" placeholder="请输入用户名" />
                </a-form-item>
                <a-form-item :label="t('user.userName')" name="fieldA">
                    <a-date-picker v-model:value="SearchFrom.value1" />
                </a-form-item>

                <a-form-item>
                    <a-button type="primary" @click="on_search()">{{ t('public.Submit') }}</a-button>
                </a-form-item>
            </a-form>
        </div>
        <div>
            <a-space wrap>
                <a-button @click="addOpen = true" type="primary">新增</a-button>
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
                        <a-button type="link" @click="on_redact(record)">编辑</a-button>
                        <a-button type="link" danger @click="on_delete(record.id)">删除</a-button>
     
                    </a-space>
                </template>
            </template>
        </a-table>


        <a-modal v-model:open="addOpen" title="添加部门" :footer=null>
            <a-form :model="formState" name="basic" :label-col="{ span: 4 }" :wrapper-col="{ span: 16 }" autocomplete="off"
                @finish="onFinish">
                <a-form-item label="部门名称" name="name" :rules="[{ required: true, message: 'Please input your username!' }]">
                    <a-input v-model:value="formState.name" />
                </a-form-item>

                <a-form-item label="部门描述" name="describe"
                    :rules="[{ required: true, message: 'Please input your password!' }]">
                    <a-input v-model:value="formState.describe" />
                </a-form-item>
                <a-form-item :wrapper-col="{ offset: 4, span: 16 }">
                    <a-button type="primary" html-type="submit">确定</a-button>
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
    name: string;
    describe: string;
}

type listItemType = {
    id: string
    name: string
    describe: string
}
import { useTableHooks } from "@/Hooks/useTableHooks"
import { onMounted, ref, reactive } from 'vue';
import { addDepartment, redactDepartment, listDepartment ,deleteDepartment} from "@/api/admin"
import { useI18n } from 'vue-i18n'
import type { Dayjs } from 'dayjs';
const { t } = useI18n()
const addOpen = ref<boolean>(false);
// let listItem = reactive<listItemType>()

let { tabHeight, SearchFrom, on_search, paginationOpt, tableData } = useTableHooks<SearchType>({
    user: "",
}, listDepartment);


const formState = reactive<FormState>({
    name: '',
    describe: '',
});


const columns = [{
    title: 'user.userName',
    dataIndex: 'name',

},
{
    title: 'user.userName',
    dataIndex: 'describe',

},
{
    title: 'operation',
    dataIndex: 'operation',
},

]
const on_redact = (data: listItemType) => {
    console.log(data.id)
    addOpen.value = true
    // listItem = data
}


const onFinish = (values: any) => {
    
    redactDepartment({ ...values, id: '0fd8134d-f349-46ea-89a3-2e2a4f101a3f' }).then(() => {
        addOpen.value = false
    })

    return

    addDepartment(values).then(() => {
        addOpen.value = false
    })
};

const on_delete = (id:string) =>{
    deleteDepartment({id}).then(()=>{
        
    })
}
onMounted(() => {

})
</script>

<style lang="less" scoped></style>