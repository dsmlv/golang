import React, { createContext, useState, useEffect } from 'react';

// Create the AuthContext
export const AuthContext = createContext();

// AuthProvider component to wrap the app and provide the context
export const AuthProvider = ({ children }) => {
    const [token, setToken] = useState(null);

    // Load token from localStorage when the app starts
    useEffect(() => {
        const storedToken = localStorage.getItem('token');
        if (storedToken) {
            setToken(storedToken);
        }
    }, []);

    // Function to handle login and store the token
    const login = (newToken) => {
        setToken(newToken);
        localStorage.setItem('token', newToken);
    };

    // Function to handle logout and clear the token
    const logout = () => {
        setToken(null);
        localStorage.removeItem('token');
    };

    return (
        <AuthContext.Provider value={{ token, login, logout }}>
            {children}
        </AuthContext.Provider>
    );
};