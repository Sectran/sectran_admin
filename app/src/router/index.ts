import { createRouter, createWebHashHistory } from 'vue-router';
import outsideLayout from './outsideLayout';
import Layout from "./Layout/index"
const routes = [
  ...Layout,
  ...outsideLayout
];

const router = createRouter({
  history: createWebHashHistory(),
  routes,
});

export default router;
