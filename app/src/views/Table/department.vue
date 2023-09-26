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
                <a-button @click="open = true" type="primary">新增</a-button>
            </a-space>
        </div>
        <a-table class="table-style" :columns="columns" :data-source="data" :scroll="{ y: tabHeight }"
            :pagination="paginationOpt">
            <template #headerCell="{ column }">
                <span>{{ t(column.title) }}</span>
            </template>
        </a-table>


        <a-modal v-model:open="open" title="添加部门" :footer=null>
            <a-form :model="formState" name="basic" :label-col="{ span: 4 }" :wrapper-col="{ span: 16 }" autocomplete="off"
                @finish="onFinish">
                <a-form-item label="部门名称" name="name"
                    :rules="[{ required: true, message: 'Please input your username!' }]">
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


// defineOptions({
//     name: 'SystemMonitorLoginLog',
//   });
import { useTableHooks } from "@/Hooks/useTableHooks"
import { onMounted, ref, reactive } from 'vue';
import { addDepartment } from "@/api/admin"
import { useI18n } from 'vue-i18n'
import type { Dayjs } from 'dayjs';
const { t } = useI18n()
const open = ref<boolean>(false);
let { tabHeight, SearchFrom, on_search, paginationOpt } = useTableHooks<SearchType>({
    user: "",
});


const formState = reactive<FormState>({
    name: '',
    describe: '',
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

const onFinish = (values: any) => {
    addDepartment(values).then((res:any)=>{
        console.log(res)
    })
  console.log('Success:', values);
};
onMounted(() => {

})
</script>

<style lang="less" scoped></style>