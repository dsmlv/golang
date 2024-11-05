// src/Login.js

import React, { useState, useContext } from 'react';
import axios from 'axios';
import { AuthContext } from './context/AuthContext'; // Import AuthContext

function Login() {
    const [username, setUsername] = useState('');
    const [password, setPassword] = useState('');
    const { login } = useContext(AuthContext); // Use the AuthContext

    const handleLogin = async (e) => {
        e.preventDefault();
        try {
            const response = await axios.post('http://localhost:8080/login', {
                username,
                password
            });
            const { token } = response.data;
            login(token); // Call the login function from the context
            alert('Login successful!');
        } catch (error) {
            alert('Login failed!');
        }
    };

    return (
        <div>
            <h2>Login</h2>
            <form onSubmit={handleLogin}>
                <input
                    type="text"
                    placeholder="Username"
                    value={username}
                    onChange={(e) => setUsername(e.target.value)}
                    required
                />
                <input
                    type="password"
                    placeholder="Password"
                    value={password}
                    onChange={(e) => setPassword(e.target.value)}
                    required
                />
                <button type="submit">Login</button>
            </form>
        </div>
    );
}

export default Login;
