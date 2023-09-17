//tabs的处理
type tabsArrType = {
    tabsArr: tabsType[],
    pitchTabs: string
}
type tabsType = {
    title: string,
    name: string,
}
import router from "@/router"
export default {
    namespaced: true,
    state: {
        //tabs数组
        tabsArr: [],
        pitchTabs: ''
    },
    mutations: {
        /**
         * 添加tabs
         * @param state 数据源
         * @param data 添加的tabs信息
         */
        addtabsArr(state: tabsArrType, data: tabsType) {
            if (!state.tabsArr.some(item => item.name === data.name)) {
                state.tabsArr.push(data)
            }
            state.pitchTabs = data.name
        },
        /**
         * 跳转tabs
         * @param state 数据源 
         * @param name 路由信息
         */
        pitchTabsChange(state: tabsArrType, name: string) {
            state.pitchTabs = name
        },
        /**
         * 删除tabs
         * @param state 
         * @param name 当前删除的name
         * @param routeName 当前路由的name
         */
        deleteTabsArr(state: tabsArrType, NameObj: { name: string, routeName: string }) {
            state.tabsArr.splice(state.tabsArr.findIndex(item => item.name === NameObj.name), 1)
            //如果删除的是当前页面，就要把路由跳转到tab的最后一页
            if (NameObj.name === NameObj.routeName) {
                let routerPushName = state.tabsArr[state.tabsArr.length - 1].name
                router.push(routerPushName)
                state.pitchTabs = routerPushName
            }
        }
    }
}