import React, { useEffect, useState } from 'react';
import { fetchBalanceList } from '../api';
import { Link } from 'react-router-dom';

const balanceCardStyle: React.CSSProperties = {
  background: 'linear-gradient(135deg, var(--color-card) 60%, var(--color-accent) 100%)',
  borderRadius: 'var(--radius)',
  boxShadow: '0 4px 24px var(--color-shadow)',
  padding: 24,
  margin: 16,
  minWidth: 260,
  maxWidth: 320,
  display: 'flex',
  flexDirection: 'column',
  alignItems: 'flex-start',
  position: 'relative',
  transition: 'box-shadow var(--transition)',
};

const BalancesListPage: React.FC = () => {
  const [balances, setBalances] = useState<any[]>([]);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    setLoading(true);
    fetchBalanceList({ limit: 1000, offset: 0 }).then(data => {
      const rows = Array.isArray(data) ? data : data.rows || data.balances || [];
      setBalances(rows);
      setLoading(false);
    });
  }, []);

  return (
    <div className="fade-in" style={{ minHeight: '100vh', background: 'var(--color-bg)', padding: '40px 0', display: 'flex', flexDirection: 'column', alignItems: 'center', justifyContent: 'center' }}>
      <div style={{ maxWidth: 1200, margin: '0 auto', display: 'flex', justifyContent: 'space-between', alignItems: 'center', marginBottom: 32, width: '100%' }}>
        <h1 style={{ color: 'var(--color-primary)', letterSpacing: 1, fontWeight: 800, fontSize: 32 }}>All Balances</h1>
        <Link
          to="/balances"
          className="scale-hover"
          style={{
            display: 'inline-block',
            background: 'linear-gradient(90deg, var(--color-primary) 60%, var(--color-accent) 100%)',
            color: '#fff',
            border: 'none',
            borderRadius: 8,
            padding: '10px 28px',
            fontWeight: 700,
            fontSize: 16,
            boxShadow: '0 2px 8px var(--color-shadow)',
            textDecoration: 'none',
            transition: 'transform 0.15s cubic-bezier(.4,2,.6,1), box-shadow 0.15s',
            cursor: 'pointer',
          }}
        >
          Back to Balance Top-Up
        </Link>
      </div>
      {loading ? (
        <div style={{ textAlign: 'center', color: 'var(--color-primary)', fontWeight: 500 }}>Loading...</div>
      ) : (
        <div style={{ display: 'flex', flexWrap: 'wrap', justifyContent: 'center', width: '100%', maxWidth: 1200 }}>
          {balances.map(balance => (
            <div key={balance.id} className="scale-hover fade-in" style={balanceCardStyle}>
              <div style={{ fontWeight: 700, fontSize: 18, color: 'var(--color-bg-dark)', marginBottom: 8 }}>Balance #{balance.id}</div>
              <div style={{ marginBottom: 6 }}><span style={{ color: 'var(--color-primary)', fontWeight: 600 }}>User ID:</span> {balance.user_id}</div>
              <div style={{ marginBottom: 6 }}><span style={{ color: 'var(--color-primary)', fontWeight: 600 }}>Card ID:</span> {balance.card_id}</div>
              <div style={{ marginBottom: 6 }}><span style={{ color: 'var(--color-primary)', fontWeight: 600 }}>Balance:</span> {balance.balance}</div>
              <div style={{ marginBottom: 6 }}><span style={{ color: 'var(--color-primary)', fontWeight: 600 }}>Ride Cost:</span> {balance.ride_cost}</div>
              <div style={{ marginBottom: 6 }}><span style={{ color: 'var(--color-primary)', fontWeight: 600 }}>Updated At:</span> {balance.updated_at ? new Date(balance.updated_at).toLocaleString() : 'N/A'}</div>
            </div>
          ))}
        </div>
      )}
    </div>
  );
};

export default BalancesListPage; 