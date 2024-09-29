import React, { useState } from 'react';
import axios from 'axios';

function Register() {
  const [username, setUsername] = useState('');
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [error, setError] = useState('');
  const [userExists, setUserExists] = useState(false);

  const checkUserExists = async (email) => {
    try {
      const response = await axios.get(`http://localhost:8080/api/check-user?email=${email}`);
      setUserExists(response.data.exists);
    } catch (error) {
      console.error('Error checking user existence:', error);
    }
  };

  const handleEmailChange = (e) => {
    const newEmail = e.target.value;
    setEmail(newEmail);
    if (newEmail) {
      checkUserExists(newEmail);
    } else {
      setUserExists(false);
    }
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    setError('');
    try {
      const response = await axios.post('http://localhost:8080/api/register', { username, email, password });
      console.log('Registration successful:', response.data);
      // TODO: Handle successful registration (e.g., redirect to login)
    } catch (error) {
      console.error('Registration error:', error);
      if (error.response) {
        console.error('Error data:', error.response.data);
        console.error('Error status:', error.response.status);
        setError(error.response.data.message || 'Registration failed');
      } else if (error.request) {
        console.error('No response received:', error.request);
        setError('No response from server. Please try again.');
      } else {
        console.error('Error message:', error.message);
        setError('An error occurred. Please try again.');
      }
    }
  };

  return (
    <div>
      <h2>Register</h2>
      {error && <p style={{color: 'red'}}>{error}</p>}
      <form onSubmit={handleSubmit}>
        <input
          type="text"
          placeholder="Username"
          value={username}
          onChange={(e) => setUsername(e.target.value)}
          required
        />
        <input
          type="email"
          placeholder="Email"
          value={email}
          onChange={handleEmailChange}
          required
        />
        {userExists && <p style={{color: 'red'}}>This email is already registered</p>}
        <input
          type="password"
          placeholder="Password"
          value={password}
          onChange={(e) => setPassword(e.target.value)}
          required
        />
        <button type="submit" disabled={userExists}>Register</button>
      </form>
    </div>
  );
}

export default Register;