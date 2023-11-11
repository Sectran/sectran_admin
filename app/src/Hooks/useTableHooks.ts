import { reactive, ref, onMounted, } from "vue";
import { Modal } from 'ant-design-vue';
import type { Ref } from "vue"
type pageData = {
    pageNum: number,
    pageSize: number
}

type resTable = {
    list: any
    PageInfo: {
        total:Number
    }
}



/**
 * 
 * @param SearchObject 搜索表单数据
 * @param Listapi 请求列表方法
 * @param deleteApi  删除方法
 * @returns sizeChange 数据
 */
export const useTableHooks = <K extends object>(SearchObject: K, paramList: {listApi:Function,paramFormat:Function}, deleteApi: Function) => {
    //表格头部颜色
    const headerStyle = { background: '#F8F8F9' }
    //分页可以选择的条数
    const pageSizes = [10, 20, 50, 100, 200]
    const FromSearchRef: Ref = ref<any>()
    //表格高度
    let tabHeight = ref<number>(100)
    //总条数
    let pageTotal = ref(200)
    //分页
    let pageData = reactive<pageData>({
        pageNum: 1,
        pageSize: 10
    })
    //表格是否正在加载
    const Table_loading = ref(false)
    //输入的搜索条件
    let SearchFrom = reactive<K>(SearchObject);

    //确认后的搜索条件
    let notarizeFrom = reactive({})
    //当前表格数据
    const tableData = ref([]);
    //分页参数
    const paginationOpt = reactive({
        current: 1,
        pageSize: 10,
        pageSizeOptions: ["10", "30", "50"],
        total: 100,
        onChange: (current: number, size: number) => {
            console.log(current)
            console.log(size)
            paginationOpt.current = current
            paginationOpt.pageSize = size
            pageData.pageNum = current
            pageData.pageSize = size
            Fun_requestList()
        },
    })



    /**
     * 点击搜索，确认搜索条件
     * @param data 搜索条件，如果没有传入就用初始化传入的搜素条件(主要功能是为了搜索条件可能需要二次处理)  非必传
     */
    const on_search = (data?: any) => {
        if (data) {
            notarizeFrom = { ...data }
        } else {
            notarizeFrom = { ...SearchFrom }
        }
        pageData.pageNum = 1
        Fun_requestList()
    }




    //表单重置
    const Fromreset = (FromRef: any) => {
        if (FromRef) {
            FromRef.resetFields()
        }
    }

    //点击分页
    const pagingChange = (val: number) => {
        pageData.pageNum = val;
        Fun_requestList()
    };
    //修改每页条数
    const sizeChange = (pageSize: number) => {
        pageData.pageSize = pageSize
        Fun_requestList()
    }

    onMounted(() => {
        let tableDom = document.querySelector('.table-style')
        if (tableDom) {
            let Height = tableDom.getBoundingClientRect().height
            tabHeight.value = Height - 120
        }
        Fun_requestList()

    })

    //请求接口
    const Fun_requestList = () => {
        console.log(pageData)
        let fromData = paramList.paramFormat({...SearchObject,...pageData})
        paramList.listApi(fromData).then((res: { data: resTable }) => {
            let { list, PageInfo } = res.data
            tableData.value = list
            paginationOpt.total = PageInfo.total
        })
    }


    return {
        sizeChange,
        Fromreset,
        pagingChange,
        on_search,
        SearchFrom,
        pageData,
        headerStyle,
        pageTotal,
        tableData,
        FromSearchRef,
        pageSizes,
        Table_loading,
        tabHeight,
        paginationOpt,
        Fun_requestList
    };
}