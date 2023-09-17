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
        </div>
        <a-table class="table-style" :columns="columns" :data-source="data" :scroll="{ y: tabHeight }"
            :pagination="paginationOpt">
            <template #headerCell="{ column }">
                <span>{{ t(column.title) }}</span>
            </template>
        </a-table>
    </div>
</template>

<script setup lang="ts">
type SearchType = {
    user: string;
    value1?: Dayjs
};

// defineOptions({
//     name: 'SystemMonitorLoginLog',
//   });
import { useTableHooks } from "@/Hooks/useTableHooks"
import { onMounted } from 'vue';
import { useI18n } from 'vue-i18n'
import type { Dayjs } from 'dayjs';
const { t } = useI18n()
let { tabHeight, SearchFrom, on_search, paginationOpt } = useTableHooks<SearchType>({
    user: "",
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
onMounted(() => {

})
</script>

<style lang="less" scoped>

</style>