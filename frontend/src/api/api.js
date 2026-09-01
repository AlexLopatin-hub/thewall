import axios from 'axios'

const api = axios.create()

export default {
  async getPosts() {
    return await api.get('http://localhost:8080/posts')
  }
}
