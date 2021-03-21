import service from '@/utils/request'

export const getPaperList = (data) => {
    return service({
        url: "/paper/getUserList",
        method: 'post',
        data: data
    })
}
