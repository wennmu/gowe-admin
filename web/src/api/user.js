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
    headers: {
      'Content-Type': 'application/x-www-form-urlencoded'
      },
    method: 'get',
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
