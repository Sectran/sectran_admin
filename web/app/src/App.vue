<template>
  <!-- :locale="getAntdLocale" -->
  <!-- :locale="enUS.locale" -->
  <ConfigProvider :locale="localeValue == 'en' ? en : zh">
    <router-view #="{ Component }">
      <component :is="Component" />
    </router-view>
    <!-- <LockScreen /> -->
  </ConfigProvider>
</template>

<script setup lang="ts">
import { watchEffect, watch, ref } from 'vue';
import { useRoute } from 'vue-router';
import { ConfigProvider } from 'ant-design-vue';
import en from 'ant-design-vue/es/locale/en_US';
import zh from 'ant-design-vue/es/locale/zh_CN';
import 'dayjs/locale/zh-cn';
import { useI18n } from 'vue-i18n'
import dayjs from 'dayjs';
const { locale } = useI18n()
const route = useRoute();
let localeValue = ref<string>(locale.value)
dayjs.locale(locale.value);
watch(locale, (val: string) => {
  localeValue.value = val
  dayjs.locale(val);
});

watchEffect(() => {
  if (route.meta?.title) {
    // 翻译网页标题
    document.title = '管理系统';
  }
});
</script>

<style lang="less">
* {
  padding: 0;
  margin: 0;
}
</style>
