import axios from "axios";
import AsyncStorage from "@react-native-async-storage/async-storage";


const DEVICE_IP = "192.168.100.89";
export const API_BASE_URL = `http://${DEVICE_IP}:8080/api`;


const api = axios.create({
  baseURL: API_BASE_URL,
  timeout: 5000,
});

api.interceptors.request.use(
  async (config) => {
    if (config.url?.startsWith('/auth')) {
      return config;
    }

    const token = await AsyncStorage.getItem('token');
    if (token && config.headers) {
      config.headers.Authorization = `Bearer ${token}`;
    }
    return config;
  },
  (err) => Promise.reject(err)
);

export default api;
