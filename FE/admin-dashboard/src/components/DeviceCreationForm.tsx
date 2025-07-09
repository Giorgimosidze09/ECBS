import React, { useState } from 'react';
import { createDevice } from '../api';

const inputStyle: React.CSSProperties = {
  width: '100%',
  padding: '10px 12px',
  borderRadius: 6,
  border: '1px solid #bcbcbc',
  fontSize: '1rem',
  background: '#fff',
  boxSizing: 'border-box',
  transition: 'border 0.2s ease-in-out',
  marginBottom: 12
};

const DeviceCreationForm: React.FC<{ onSuccess?: () => void }> = ({ onSuccess }) => {
  const [deviceId, setDeviceId] = useState('');
  const [location, setLocation] = useState('');
  const [loading, setLoading] = useState(false);
  const [success, setSuccess] = useState<string | null>(null);
  const [error, setError] = useState<string | null>(null);

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setLoading(true);
    setSuccess(null);
    setError(null);
    try {
      await createDevice({ device_id: deviceId, location });
      setSuccess('Device created successfully!');
      setDeviceId('');
      setLocation('');
      if (onSuccess) onSuccess();
    } catch {
      setError('Failed to create device.');
    } finally {
      setLoading(false);
    }
  };

  return (
    <form onSubmit={handleSubmit} style={{ width: '100%', display: 'flex', flexDirection: 'column', alignItems: 'center', gap: 18 }}>
      <label style={{ ...inputStyle, marginBottom: 0, background: 'none', border: 'none', color: '#4a4e69' }}>Device ID</label>
      <input
        type="text"
        value={deviceId}
        onChange={e => setDeviceId(e.target.value)}
        required
        style={inputStyle}
        placeholder="Enter Device ID"
        autoComplete="off"
      />
      <label style={{ ...inputStyle, marginBottom: 0, background: 'none', border: 'none', color: '#4a4e69' }}>Location</label>
      <input
        type="text"
        value={location}
        onChange={e => setLocation(e.target.value)}
        required
        style={inputStyle}
        placeholder="Enter Location"
        autoComplete="off"
      />
      <button
        type="submit"
        className="styled-button"
        style={{
          width: '100%',
          padding: '12px 0',
          borderRadius: 6,
          background: loading ? '#999' : '#22223b',
          color: '#fff',
          fontWeight: 600,
          fontSize: '1rem',
          marginTop: 8,
          cursor: loading ? 'not-allowed' : 'pointer',
          opacity: loading ? 0.6 : 1,
          transition: 'background 0.3s ease-in-out'
        }}
        disabled={loading}
      >
        {loading ? 'Creating...' : 'Add Device'}
      </button>
      {success && <div style={{ color: 'lightgreen', marginTop: 12 }}>{success}</div>}
      {error && <div style={{ color: 'salmon', marginTop: 12 }}>{error}</div>}
    </form>
  );
};

export default DeviceCreationForm; 