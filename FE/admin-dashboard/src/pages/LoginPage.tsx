import React, { useState } from 'react';
import { login } from '../api';
import { useAuth } from '../auth';
import { useHistory, Link } from 'react-router-dom';

const LoginPage: React.FC = () => {
  const [username, setUsername] = useState('');
  const [password, setPassword] = useState('');
  const [error, setError] = useState('');
  const { login: doLogin } = useAuth();
  const history = useHistory();

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setError('');
    try {
      const res = await login({ username, password });
      doLogin(res.token, res.role);
      if (res.role === 'admin') history.push('/admin');
      else if (res.role === 'customer') history.push('/customer');
      else history.push('/');
    } catch (err: any) {
      setError(err?.response?.data || 'Login failed');
    }
  };

  return (
    <div style={{ minHeight: '100vh', background: 'linear-gradient(135deg, #4a4e69 0%, #9a8c98 100%)', display: 'flex', alignItems: 'center', justifyContent: 'center' }}>
      <div style={{ maxWidth: 400, width: '100%', background: '#fff', borderRadius: 16, boxShadow: '0 8px 32px rgba(34,34,59,0.18)', padding: 36, display: 'flex', flexDirection: 'column', alignItems: 'center' }}>
        <h1 style={{ color: '#22223b', fontWeight: 800, fontSize: 32, marginBottom: 24 }}>Sign In</h1>
        <form onSubmit={handleSubmit} style={{ width: '100%' }}>
          <div style={{ marginBottom: 18 }}>
            <label style={{ display: 'block', color: '#4a4e69', fontWeight: 600, marginBottom: 6 }}>Username</label>
            <input value={username} onChange={e => setUsername(e.target.value)} required style={{ width: '100%', padding: 10, borderRadius: 8, border: '1px solid #c9ada7', fontSize: 16 }} />
          </div>
          <div style={{ marginBottom: 18 }}>
            <label style={{ display: 'block', color: '#4a4e69', fontWeight: 600, marginBottom: 6 }}>Password</label>
            <input type="password" value={password} onChange={e => setPassword(e.target.value)} required style={{ width: '100%', padding: 10, borderRadius: 8, border: '1px solid #c9ada7', fontSize: 16 }} />
          </div>
          {error && <div style={{ color: '#c72c41', marginBottom: 12, fontWeight: 600 }}>{error}</div>}
          <button type="submit" style={{ width: '100%', background: 'linear-gradient(90deg, #4a4e69 60%, #9a8c98 100%)', color: '#fff', border: 'none', borderRadius: 8, padding: '12px 0', fontWeight: 700, fontSize: 18, marginTop: 8, cursor: 'pointer', boxShadow: '0 2px 8px #c9ada7' }}>Login</button>
        </form>
        <div style={{ marginTop: 24, color: '#4a4e69', fontWeight: 500 }}>
          Don't have an account?{' '}
          <Link to="/register" style={{ color: '#9a8c98', fontWeight: 700, textDecoration: 'underline' }}>Register</Link>
        </div>
      </div>
    </div>
  );
};

export default LoginPage; 