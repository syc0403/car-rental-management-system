/**
 * axios 工具
 */
import axios from 'axios'
let instance=axios.create({
    baseURL:process.env.VUE_APP_BASE_API,
    timeout:5000,
})
//-------------拦截器---------------------
//----------------request----------------------
instance.interceptors.request.use(
    function(config){
        return config;
    },
    function(err){
        return new Promise.reject(err);
    }
);
//-------------respones---------------
instance.interceptors.response.use(
    function(response){
        return response;
    },
    function(err){
        return new Promise.reject(err);
    }
);
//--------------------------------
/**
 * 远程通信get请求
 * 返回json对象
 * @param {*} url 请求地址 baseUrl+url
 * @param {*} param {}
 */
async function get(url,params){

    return instance.get(url,{params});
}
async function getNoParam(url){
    return instance.get(url)
}
/**
 * 远程通信post请求
 * 返回json对象
 * @param {*} url 请求地址 baseUrl+url
 * @param {*} data  {}
 */
async function post(url,data){
    return instance.post(url,data)
}
export {get,getNoParam,post}