import React, { useState, useEffect } from "react";
import API from "../api/axiosInstance";

const Dashboard = () => {
  const [user, setUser] = useState(null); // User details
  const [orders, setOrders] = useState([]); // List of user orders
  const [selectedOrder, setSelectedOrder] = useState(null); // Details of a selected order
  const [error, setError] = useState(null);

  // Profile Update State
  const [updateProfile, setUpdateProfile] = useState({
    username: "",
    email: "",
  });

  // Fetch user details
  useEffect(() => {
    const fetchUser = async () => {
      try {
        const response = await API.get("/users/me"); // Endpoint to fetch current user
        setUser(response.data);
        setUpdateProfile({
          username: response.data.username,
          email: response.data.email,
        });
        setError(null);
      } catch (err) {
        setError("Failed to fetch user details");
      }
    };
    fetchUser();
  }, []);

  // Fetch user orders
  useEffect(() => {
    const fetchOrders = async () => {
      try {
        const response = await API.get("/orders/"); // Endpoint to fetch user orders
        setOrders(response.data);
        setError(null);
      } catch (err) {
        setError("Failed to fetch orders");
      }
    };
    fetchOrders();
  }, []);

  // Fetch order details
  const fetchOrderDetails = async (orderId) => {
    try {
      const response = await API.get(`/orders/${orderId}`); // Endpoint to fetch order details
      setSelectedOrder(response.data);
      setError(null);
    } catch (err) {
      setError("Failed to fetch order details");
    }
  };

  // Handle profile update
  const handleProfileUpdate = async (e) => {
    e.preventDefault();
    try {
      const response = await API.put("/users/me", updateProfile); // Update profile endpoint
      setUser(response.data); // Update user details
      setError(null);
      alert("Profile updated successfully");
    } catch (err) {
      setError("Failed to update profile");
    }
  };

  // Cancel an order
const handleCancelOrder = async (orderId) => {
    try {
      await API.put(`/orders/${orderId}/cancel`); // Endpoint to cancel an order
      setOrders(
        orders.map((order) =>
          order.order_id === orderId ? { ...order, status: "Cancelled" } : order
        )
      );
      setError(null);
      alert("Order cancelled successfully");
    } catch (err) {
      setError("Failed to cancel order");
    }
  };

  return (
    <div>
      <h1>Dashboard</h1>
      {error && <p style={{ color: "red" }}>{error}</p>}

      {/* User Details */}
      {user && (
        <div>
          <h2>Welcome, {user.username}!</h2>
          <p>Email: {user.email}</p>
          <p>Role: {user.role}</p>
        </div>
      )}

      {/* Profile Update Section */}
      <h2>Update Profile</h2>
      <form onSubmit={handleProfileUpdate}>
        <input
          type="text"
          name="username"
          placeholder="Username"
          value={updateProfile.username}
          onChange={(e) => setUpdateProfile({ ...updateProfile, username: e.target.value })}
          required
        />
        <input
          type="email"
          name="email"
          placeholder="Email"
          value={updateProfile.email}
          onChange={(e) => setUpdateProfile({ ...updateProfile, email: e.target.value })}
          required
        />
        <button type="submit">Update Profile</button>
      </form>

      {/* Orders List */}
      <h2>Your Orders</h2>
      {orders.length > 0 ? (
        <ul>
          {orders.map((order) => (
            <li key={order.order_id}>
              Order #{order.order_id} - {order.status} - ${order.total_amount}
              <button onClick={() => fetchOrderDetails(order.order_id)}>View Details</button>
              {order.status === "Pending" && (
                <button onClick={() => handleCancelOrder(order.order_id)}>Cancel Order</button>
              )}
            </li>
          ))}
        </ul>
      ) : (
        <p>You have no orders yet.</p>
      )}

      {/* Order Details */}
      {selectedOrder && (
        <div>
          <h3>Order Details</h3>
          <p>Order ID: {selectedOrder.order_id}</p>
          <p>Status: {selectedOrder.status}</p>
          <p>Total Amount: ${selectedOrder.total_amount}</p>
          <h4>Items:</h4>
          <ul>
            {selectedOrder.items.map((item) => (
              <li key={item.order_item_id}>
                {item.product_name} - ${item.price} x {item.quantity}
              </li>
            ))}
          </ul>
        </div>
      )}
    </div>
  );
};

export default Dashboard;
