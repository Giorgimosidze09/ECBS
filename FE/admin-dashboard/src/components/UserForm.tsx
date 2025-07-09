import React, { useState } from 'react';
import { createUser } from '../api'; // same pattern as rideCost

const inputStyle: React.CSSProperties = {
  width: '100%',
  padding: '10px 12px',
  borderRadius: 6,
  border: '1px solid #bcbcbc',
  fontSize: '1rem',
  background: '#fff',
  boxSizing: 'border-box',
  marginBottom: 12,
};

const UserForm: React.FC = () => {
  const [name, setName] = useState('');
  const [email, setEmail] = useState('');
  const [phone, setPhone] = useState('');
  const [result, setResult] = useState<boolean | null>(null);
  const [error, setError] = useState('');

  const isFormReady = name.trim() !== '' && email.trim() !== '' && phone.trim() !== '';

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setError('');
    setResult(null);

    try {
      const response = await createUser({ name, email, phone });
      setResult(response);
      if (response) {
        setName('');
        setEmail('');
        setPhone('');
      }
    } catch {
      setError('Failed to create user. Please try again.');
    }
  };

  return (
    <form
      onSubmit={handleSubmit}
      style={{
        width: '100%',
        display: 'flex',
        flexDirection: 'column',
        alignItems: 'center',
      }}
    >
      <h2 style={{ textAlign: 'center', color: '#4a4e69', marginBottom: 20 }}>Create User</h2>

      {error && <div style={{ color: 'red', marginBottom: 12 }}>{error}</div>}

      <div style={{ width: '100%' }}>
        <label
          htmlFor="user-name"
          style={{ display: 'block', marginBottom: 6, color: '#4a4e69' }}
        >
          Name:
        </label>
        <input
          id="user-name"
          type="text"
          value={name}
          onChange={e => setName(e.target.value)}
          required
          style={inputStyle}
        />
      </div>

      {name.trim() !== '' && (
        <div style={{ width: '100%' }}>
          <label
            htmlFor="user-email"
            style={{ display: 'block', marginBottom: 6, color: '#4a4e69' }}
          >
            Email:
          </label>
          <input
            id="user-email"
            type="email"
            value={email}
            onChange={e => setEmail(e.target.value)}
            required
            style={inputStyle}
          />
        </div>
      )}

      {email.trim() !== '' && (
        <div style={{ width: '100%' }}>
          <label
            htmlFor="user-phone"
            style={{ display: 'block', marginBottom: 6, color: '#4a4e69' }}
          >
            Phone:
          </label>
          <input
            id="user-phone"
            type="text"
            value={phone}
            onChange={e => setPhone(e.target.value)}
            required
            style={inputStyle}
          />
        </div>
      )}

      <button
        type="submit"
        disabled={!isFormReady}
        style={{
          width: '100%',
          padding: '12px 0',
          borderRadius: 6,
          background: isFormReady ? '#22223b' : '#999',
          color: '#fff',
          fontWeight: 600,
          fontSize: '1rem',
          marginTop: 8,
          cursor: isFormReady ? 'pointer' : 'not-allowed',
        }}
      >
        Create User
      </button>

      {result !== null && (
        <div
          style={{
            marginTop: 16,
            background: '#fff',
            padding: 16,
            borderRadius: 8,
            color: result ? 'green' : 'red',
            textAlign: 'center',
          }}
        >
          {result ? 'User created successfully!' : 'Failed to create user.'}
        </div>
      )}
    </form>
  );
};

export default UserForm;
