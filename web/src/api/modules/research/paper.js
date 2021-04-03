import service from '@/utils/request'

export const getPaperList = (data) => {
    return service({
        url: "/paper/getUserList",
        method: 'post',
        data: data
    })
}


export const create = (data) => {
    return service({
        url: "/paper/createInfo",
        method: 'post',
        data: data
    })
}

export const update = (data) => {
    return service({
        url: "/paper/updateInfo",
        method: 'put',
        data: data
    })
}

export const getInfo = (params) => {
    return service({
        url: "/paper/getInfo",
        method: 'get',
        params
    })
}

export const deleteItem = (data) => {
    return service({
        url: "/paper/deleteInfo",
        method: 'delete',
        data
    })
}