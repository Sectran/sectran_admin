// import { LOGIN_NAME } from '@/router/constant';
/**
 * layout布局之外的路由
 */
export const LoginRoute = {
  path: '/login',
  name: 'Login',
  component: () => import(/* webpackChunkName: "login" */ 'views/Login/index.vue'),
  meta: {
    title: '登录',
  },
};

/**
 * layout布局之外的路由
 */
export const LinkConfiguration = {
  path: '/linkConfiguration',
  name: 'LinkConfiguration',
  component: () => import(/* webpackChunkName: "login" */ 'views/linkConfiguration/index.vue'),
  meta: {
    title: '连接配置',
  },
};

export default [LoginRoute,LinkConfiguration];
