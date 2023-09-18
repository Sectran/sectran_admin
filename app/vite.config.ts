import { defineConfig } from 'vite';
import vue from '@vitejs/plugin-vue';
// import VueSetupExtend from 'vite-plugin-vue-setup-extend';
import { join } from "path";
import { defineConfig } from 'vite';
import Components from 'unplugin-vue-components/vite';
import { AntDesignVueResolver } from 'unplugin-vue-components/resolvers';
export default defineConfig({
	base: './',
	plugins: [
		vue(),
		// VueSetupExtend(),
		Components({
			resolvers: [
				AntDesignVueResolver({
					importStyle: false, // css in js
				}),
			],
		}),
	],

	server: {
		proxy: {
			'/api': {
				target: 'http://127.0.0.1:8082',
				changeOrigin: true,
				rewrite: (path) => path.replace(/^\/api/, ''),
			}
		}
	},

	resolve: {
		alias: {
			'@': join(__dirname, "src"),
			'views': join(__dirname, "src/views"),
		}
	},
	css: {
		// css预处理器
		preprocessorOptions: {
			less: {
				charset: false,
				// additionalData: '@import "./src/assets/style/global.less";',
			},
		},
	},

	optimizeDeps: {
		include: ['schart.js']
	},
	define: { __VUE_I18N_FULL_INSTALL__: true, __VUE_I18N_LEGACY_API__: true, __INTLIFY_PROD_DEVTOOLS__: false },
});
