//对路由表的处理
interface stateinterface {
    if_routeMenu: Boolean,
}
export default {
    namespaced: true,
    state: {
        //是否重新渲染路由表
        if_routeMenu: true, 
    },
    mutations: {
        routeUpdate(state: stateinterface) {
            state.if_routeMenu = false
        }
    }

}