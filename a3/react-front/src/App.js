import React, { useContext, useEffect, useState } from 'react';
import axios from 'axios';
import Login from './Login';
import { AuthContext } from './context/AuthContext'; // Import AuthContext

function App() {
    const { token, logout } = useContext(AuthContext); // Use the AuthContext
    const [data, setData] = useState(null);
    const [loading, setLoading] = useState(true);
    const [error, setError] = useState(null);

    useEffect(() => {
        if (token) {
            axios.get('http://localhost:8080/read/item-id', {
                headers: {
                    Authorization: `Bearer ${token}`
                }
            })
            .then((response) => {
                setData(response.data);
                setLoading(false);
            })
            .catch((err) => {
                setError(err);
                setLoading(false);
            });
        } else {
            setLoading(false);
        }
    }, [token]);

    if (loading) return <p>Loading...</p>;
    if (error) return <p>Error: {error.message}</p>;

    return (
        <div className="App">
            <h1>Simple Task Manager</h1>
            {!data ? (
                <Login />
            ) : (
                <div>
                    <h2>Task Details</h2>
                    <p>ID: {data.id}</p>
                    <p>Name: {data.name}</p>
                    <button onClick={logout}>Logout</button> {/* Logout button */}
                </div>
            )}
        </div>
    );
}

export default App;
