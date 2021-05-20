import request from '@/utils/request'

export function login(data) {
  return request({
    url: '/login',
    method: 'post',
    data
  })
}

export function getInfo(token) {
  return request({
    url: '/admin/info',
    method: 'get',
    params: {token}
    
  }).catch((error) => {
    console.log(error)
  })
}

export function logout() {
  return request({
    url: '/logout',
    method: 'post'
  })
}
