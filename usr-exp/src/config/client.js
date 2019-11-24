import axios from 'axios'

export const clientConfig = axios.create({
   baseURL: 'http://localhost:9000/',
   timeout: 2000,
   headers: {"Access-Control-Allow-Origin": "*"}
});

export default clientConfig;