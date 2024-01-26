import axios from 'axios';

const instance=axios.create({
    baseURL:"/api/user"
});
instance.interceptors.request.use(config=>{
    config.headers.set("Authorization",sessionStorage.getItem("token"))
    return config
})
instance.interceptors.response.use(response=>response.data)


export const getInfo=()=>instance({
    method:'GET',
    url:'/info'
})

export const changePassword=(data)=>instance({
    method:'POST',
    url:'/change_password',
    data:{
        password:data.password,
        againPassword:data.againPassword
    }
})