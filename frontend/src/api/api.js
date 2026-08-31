import axios from 'axios'

const api = axios.create()

export default {
  async getPosts() {
    return await api.post('/posts')
  }
}
