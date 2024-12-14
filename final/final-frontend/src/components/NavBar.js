import React from "react";
import { Link } from "react-router-dom";

const Navbar = () => {
  const role = localStorage.getItem("role"); // Store user role on login

  return (
    <nav>
      <Link to="/dashboard">Dashboard</Link>
      {role === "admin" && <Link to="/admin">Admin Panel</Link>}
      <Link to="/logout">Logout</Link>
    </nav>
  );
};

export default Navbar;
