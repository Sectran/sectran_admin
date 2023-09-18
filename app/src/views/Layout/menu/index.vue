<template>
  <div class="menu-style">
    <div class="logo"></div>
    <a-menu v-model:selectedKeys="menuState.selectedKeys" v-model:openKeys="menuState.openKeys" mode="inline">
      <a-sub-menu v-for="item in administration" :key="item.name">
        <template #title>
          <span>
            <pie-chart-outlined />
            <span>{{ t(`menu.${item.meta.title}`) }}</span>
          </span>
        </template>
        <a-menu-item v-for="el in item.children" :key="el.name" @click="onPath(el)">{{ t(`menu.${el.meta.title}`)
        }}</a-menu-item>
      </a-sub-menu>
    </a-menu>
  </div>
</template>

<script lang="ts" setup>
import { reactive, watch } from 'vue';
import { useRouter, useRoute } from 'vue-router';
import { PieChartOutlined, } from '@ant-design/icons-vue';
import { useStore } from 'vuex'
import administration from "@/router/Layout/index"
import { routerType } from "@/utils/type"
import { useI18n } from 'vue-i18n'
const { t } = useI18n()
const router = useRouter();
const route = useRoute();
const store = useStore()
let tabs = {
  name: route.name,
  title: route.meta.title
}
store.commit('TabsStore/addtabsArr', tabs)
const menuState = reactive<{ openKeys: string[], selectedKeys: string[] }>({
  openKeys: ['admin'],
  selectedKeys: [route.name],
})
watch(() => store.state.TabsStore.pitchTabs, (oldValue: string) => {
  menuState.selectedKeys = [oldValue]
})
const onPath = (el: routerType) => {
  let data = {
    name: el.name,
    title: el.meta.title
  }
  store.commit('TabsStore/addtabsArr', data)
  router.push(data.name)
}
</script>

<style lang="less" scoped>
.logo {
  height: 32px;
  background: rgba(255, 255, 255, 0.3);
  margin: 16px;
}
</style>