import React from 'react';
import { Link } from 'react-router-dom';
import logoImage from '../assets/images/joyklub-logo.jpg';

function Header() {
  return (
    <header className="bg-white shadow-md">
      <div className="container mx-auto px-4 py-4 flex justify-between items-center">
        <Link to="/" className="flex items-center">
          <img src={logoImage} alt="JoyKlub Logo" className="h-12 w-auto" /> {/* Adjust height as needed */}
        </Link>
        <nav>
          <ul className="flex space-x-4">
            <li><Link to="/" className="text-[#4A90E2] hover:text-[#E24A4A]">Home</Link></li>
            <li><Link to="/products" className="text-[#4A90E2] hover:text-[#E24A4A]">Products</Link></li>
            <li><Link to="/cart" className="text-[#4A90E2] hover:text-[#E24A4A]">Cart</Link></li>
            <li><Link to="/login" className="text-[#4A90E2] hover:text-[#E24A4A]">Login</Link></li>
            <li><Link to="/register" className="text-[#4A90E2] hover:text-[#E24A4A]">Register</Link></li>
            <li><Link to="/users" className="text-[#4A90E2] hover:text-[#E24A4A]">User List</Link></li>
          </ul>
        </nav>
      </div>
    </header>
  );
}

export default Header;