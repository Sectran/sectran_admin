/// <reference types="vite/client" />
declare module "*.vue" {
    import type { DefineComponent } from "vue";
    const vueComponent: DefineComponent<{}, {}, any>;
    export default vueComponent;
}
declare module 'vue-i18n';
declare module 'vuex';
declare module 'vue-router';
declare module 'Vue';
declare module 'xterm';
declare module 'xterm-addon-fit';
