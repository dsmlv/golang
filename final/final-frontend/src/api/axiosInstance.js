import axios from "axios";

const API = axios.create({
  baseURL: "http://localhost:8080", // Backend API URL
});

// Add JWT token to requests if available
API.interceptors.request.use((config) => {
  const token = localStorage.getItem("token"); // Store token securely
  if (token) {
    config.headers.Authorization = `Bearer ${token}`;
  }
  return config;
});

export default API;
