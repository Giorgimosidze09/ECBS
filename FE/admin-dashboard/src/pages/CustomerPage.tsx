import React, { useEffect, useState } from 'react';
import { fetchCustomerSumBalance } from '../api';
import { useAuth } from '../auth';
import { useHistory } from 'react-router-dom';

const CustomerPage: React.FC = () => {
  const { deviceId, logout } = useAuth();
  const history = useHistory();
  const [totalBalance, setTotalBalance] = useState<number | null>(null);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState('');

  useEffect(() => {
    if (!deviceId) return;
    setLoading(true);
    setError('');
    setTotalBalance(null);
    fetchCustomerSumBalance(deviceId)
      .then(res => {
        setTotalBalance(res.total_balance);
      })
      .catch(err => {
        setError(err?.response?.data || 'Failed to fetch balance');
      })
      .finally(() => setLoading(false));
  }, [deviceId]);

  const handleLogout = () => {
    logout();
    history.push('/login');
  };

  return (
    <div style={{ minHeight: '100vh', background: 'linear-gradient(135deg, #4a4e69 0%, #9a8c98 100%)', display: 'flex', alignItems: 'center', justifyContent: 'center' }}>
      <div style={{ maxWidth: 400, width: '100%', background: '#fff', borderRadius: 16, boxShadow: '0 8px 32px rgba(34,34,59,0.18)', padding: 36, display: 'flex', flexDirection: 'column', alignItems: 'center' }}>
        <h1 style={{ color: '#22223b', fontWeight: 800, fontSize: 32, marginBottom: 24 }}>Your Balance</h1>
        {!deviceId ? (
          <div style={{ color: '#c72c41', fontWeight: 600, margin: '24px 0' }}>No device assigned to your account. Please contact support.</div>
        ) : loading ? (
          <div style={{ color: '#4a4e69', fontWeight: 600 }}>Loading...</div>
        ) : error ? (
          <div style={{ color: '#c72c41', fontWeight: 600 }}>{error}</div>
        ) : totalBalance !== null ? (
          <div style={{ fontSize: 48, color: '#4a4e69', fontWeight: 800, margin: '24px 0' }}>{totalBalance} ₾</div>
        ) : (
          <div style={{ color: '#4a4e69', fontWeight: 600, margin: '24px 0' }}>No balance found for your device.</div>
        )}
        <button
          onClick={handleLogout}
          style={{
            marginTop: '2rem',
            background: 'linear-gradient(90deg, #c72c41 60%, #9a8c98 100%)',
            color: '#fff',
            border: 'none',
            borderRadius: 8,
            padding: '12px 0',
            fontWeight: 700,
            fontSize: 18,
            cursor: 'pointer',
            boxShadow: '0 2px 8px #c9ada7',
            width: '100%'
          }}
        >
          Logout
        </button>
      </div>
    </div>
  );
};

export default CustomerPage; 