import axios from 'axios'
import qs from 'qs'

// 创建axios实例
const service = axios.create({
  baseURL: "/",
  timeout: 30000 // 请求超时时间
})

// request拦截器
service.interceptors.request.use(
  config => {
    config.transformRequest = [function(data) {
      if (data instanceof FormData) {
        return data
      } else {
        return qs.stringify(data)
      }
    }]
    config.headers['Content-Type'] = 'application/x-www-form-urlencoded';
    config.headers['token'] = localStorage.getItem('token');
    return config
  },
  error => {
    Promise.reject(error)
  }
)

service.interceptors.response.use(
    response=>{
        return response.data
    }
)
export default service
