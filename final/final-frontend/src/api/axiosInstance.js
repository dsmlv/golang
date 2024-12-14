import axios from "axios";

// Create Axios instance
const API = axios.create({
  baseURL: process.env.REACT_APP_API_URL, // Use environment variable for the API URL
  withCredentials: true,
});

// Add JWT token to headers if available
API.interceptors.request.use((config) => {
  const token = localStorage.getItem("token"); // Retrieve token from localStorage
  if (token) {
    config.headers.Authorization = `Bearer ${token}`; // Attach token to Authorization header
  }
  return config;
});

export default API;
