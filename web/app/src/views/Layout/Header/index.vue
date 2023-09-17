<template>
    <a-layout-header class="flex-space-between-center Header-style">
        <div class="Header-letf">
            <menu-unfold-outlined v-if="collapsed" class="nav-icon" @click="on_icon(false)" />
            <menu-fold-outlined v-else class="nav-icon" @click="on_icon(true)" />
            <!-- <Breadcrumb>
                <Breadcrumb.Item>Home</Breadcrumb.Item>
            </Breadcrumb> -->
        </div>
        <div class="nav-right">
            <a-dropdown>
                <a class="ant-dropdown-link" @click.prevent>
                    <TranslationOutlined class="nav-icon" />
                </a>
                <template #overlay>
                    <a-menu>
                        <a-menu-item :class="{ optCalss: langSrt === item.lang }" v-for="item in menuItem" :key="item.lang"
                            @click='on_language(item.lang)'>{{
                                item.name }}</a-menu-item>
                    </a-menu>
                </template>
            </a-dropdown>
            <!-- 
            <a-avatar shape="square">
                <template #icon>
                    <UserOutlined class="nav-icon" />
                </template>
            </a-avatar> -->
        </div>
    </a-layout-header>
</template>

<script setup lang="ts">
import {
    MenuUnfoldOutlined,
    MenuFoldOutlined,
    TranslationOutlined,
    // UserOutlined
} from "@ant-design/icons-vue";
// Breadcrumb
import { useI18n } from 'vue-i18n'
const { locale } = useI18n()
import { ref, onMounted } from 'vue';
let langSrt = ref(localStorage.getItem('lang') || 'zh')
type menuItemType = {
    name: string,
    lang: string
}
let menuItem: menuItemType[] = [
    {
        name: '中文',
        lang: 'zh'
    },
    {
        name: 'English',
        lang: 'en'
    }]
onMounted(() => {
    console.log('3.-组件挂载到页面之后执行-------onMounted,Header')
})
/**
 * 点击切换全局化语言
 * @param value 
 */
const on_language = (lang: string) => {
    langSrt.value = lang
    locale.value = lang
    localStorage.setItem('lang', lang)
}
const { collapsed } = defineProps(['collapsed'])
const emit = defineEmits(['update:collapsed'])
const on_icon = (type: boolean) => {
    emit("update:collapsed", type)
}
</script>

<style scoped lang="less">
.Header-style {
    background: #ffffff;
    padding: 0 20px;
    border-bottom: 2px solid #F1F3F4;

    .Header-letf {
        display: flex;
    }
}

.nav-right {
    display: flex;
    align-items: center;

    span {
        margin-left: 20px;
    }
}

::v-deep(.optCalss) {
    color: #2D93FC !important;
    background: #E7F7FF !important;
}
</style>