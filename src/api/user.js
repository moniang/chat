import request from '@/utils/request'

/**
 * 登录
 * @param data
 */
export function login(data) {
    return request({
        url: '/login',
        method: 'post',
        data
    })
}

/**
 * 注册
 * @param data
 */
export function register(data) {
    return request({
        url: '/register',
        method: 'post',
        data
    })
}

/**
 * 修改昵称
 * @param data
 */
export function reviseName(data) {
    return request({
        url: '/revise/name',
        method: 'post',
        data
    })
}

/**
 * 修改字体颜色
 * @param data
 */
export function reviseFontColor(data) {
    return request({
        url: '/revise/fontColor',
        method: 'post',
        data
    })
}
