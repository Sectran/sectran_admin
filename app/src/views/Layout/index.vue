<template>
  <a-layout class="layout-style" style="background: #fff">
    <a-layout-sider v-model:collapsed="collapsed" :trigger="null" collapsible style="background: #fff">
      <AsideMenu></AsideMenu>
    </a-layout-sider>
    <a-layout>
      <Headers v-model:collapsed="collapsed"></Headers>
      <Tabs></Tabs>
      <a-layout-content class="Content-style">
        <router-view v-slot="{ Component }">
          <template v-if="Component">
            <transition :name="Object.is(route.meta?.transitionName, false) ? '' : 'fade-transform'" mode="out-in" appear>
              <keep-alive :include="[store.state.TabsStore.tabsArr.map((item: any) => item.name)]">
                <component :is="Component" :key="route.fullPath" />
              </keep-alive>
            </transition>
          </template>
        </router-view>
      </a-layout-content>
    </a-layout>
  </a-layout>
</template>
<script lang="ts" setup>
import { ref, onMounted } from "vue";
import AsideMenu from './menu/index.vue';
import Headers from './Header/index.vue';
import Tabs from './Tabs/index.vue';
const collapsed = ref<boolean>(false);
import { useRoute } from 'vue-router';
const route = useRoute();
import { useStore } from 'vuex'
const store = useStore()
onMounted(() => {
  console.log('3.-组件挂载到页面之后执行-------onMounted,index')
})
</script>

<style lang="less" scoped>
@import "@/assets/css/transition.less";

.layout-style {
  width: 100vw;
  height: 100vh;
}


.Content-style {
  height: calc(100vh - 64px - 50px);
  padding: 20px;
  background: #F1F3F4;
}
</style>
