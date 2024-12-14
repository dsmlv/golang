import React, { useState, useEffect } from "react";
import API from "../api/axiosInstance";

const AdminPanel = () => {
  const [products, setProducts] = useState([]);
  const [formData, setFormData] = useState({
    name: "",
    description: "",
    price: "",
    stock: "",
    category_id: "",
  });
  const [editingProduct, setEditingProduct] = useState(null);
  const [error, setError] = useState(null);

  // Fetch products
  useEffect(() => {
    const fetchProducts = async () => {
      try {
        const response = await API.get("/products/");
        setProducts(response.data);
      } catch (err) {
        setError("Failed to fetch products");
      }
    };
    fetchProducts();
  }, []);

  // Handle form input changes
  const handleChange = (e) => {
    setFormData({ ...formData, [e.target.name]: e.target.value });
  };

  // Create a new product
  const handleCreate = async (e) => {
    e.preventDefault();
    try {
      const response = await API.post("/products/", formData);
      setProducts([...products, response.data]);
      setFormData({ name: "", description: "", price: "", stock: "", category_id: "" });
      setError(null);
    } catch (err) {
      setError("Failed to create product");
    }
  };

  // Start editing a product
  const handleEditClick = (product) => {
    setEditingProduct(product);
    setFormData({
      name: product.name,
      description: product.description,
      price: product.price,
      stock: product.stock,
      category_id: product.category_id,
    });
  };

  // Update an existing product
  const handleUpdate = async (e) => {
    e.preventDefault();
    try {
      const response = await API.put(`/products/${editingProduct.product_id}`, formData);
      setProducts(
        products.map((product) =>
          product.product_id === editingProduct.product_id ? response.data : product
        )
      );
      setEditingProduct(null);
      setFormData({ name: "", description: "", price: "", stock: "", category_id: "" });
      setError(null);
    } catch (err) {
      setError("Failed to update product");
    }
  };

  // Delete a product
  const handleDelete = async (id) => {
    try {
      await API.delete(`/products/${id}`);
      setProducts(products.filter((product) => product.product_id !== id));
      setError(null);
    } catch (err) {
      setError("Failed to delete product");
    }
  };

  return (
    <div>
      <h1>Admin Panel</h1>

      {/* Product List */}
      <h2>Products</h2>
      {error && <p style={{ color: "red" }}>{error}</p>}
      <ul>
        {products.map((product) => (
          <li key={product.product_id}>
            {product.name} - ${product.price}{" "}
            <button onClick={() => handleEditClick(product)}>Edit</button>
            <button onClick={() => handleDelete(product.product_id)}>Delete</button>
          </li>
        ))}
      </ul>

      {/* Product Form */}
      <h2>{editingProduct ? "Edit Product" : "Create Product"}</h2>
      <form onSubmit={editingProduct ? handleUpdate : handleCreate}>
        <input
          type="text"
          name="name"
          placeholder="Product Name"
          value={formData.name}
          onChange={handleChange}
          required
        />
        <textarea
          name="description"
          placeholder="Product Description"
          value={formData.description}
          onChange={handleChange}
          required
        ></textarea>
        <input
          type="number"
          name="price"
          placeholder="Price"
          value={formData.price}
          onChange={handleChange}
          required
        />
        <input
          type="number"
          name="stock"
          placeholder="Stock Quantity"
          value={formData.stock}
          onChange={handleChange}
          required
        />
        <input
          type="text"
          name="category_id"
          placeholder="Category ID"
          value={formData.category_id}
          onChange={handleChange}
          required
        />
        <button type="submit">{editingProduct ? "Update Product" : "Create Product"}</button>
      </form>
    </div>
  );
};

export default AdminPanel;
