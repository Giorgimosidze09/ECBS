import React, { useEffect, useState } from 'react';
import { fetchBalanceList } from '../api';
import { Link } from 'react-router-dom';

const balanceCardStyle: React.CSSProperties = {
  background: 'linear-gradient(135deg, #fff 60%, #b5ead7 100%)',
  borderRadius: 16,
  boxShadow: '0 4px 24px #b5ead722',
  padding: 24,
  margin: 16,
  minWidth: 260,
  maxWidth: 320,
  display: 'flex',
  flexDirection: 'column',
  alignItems: 'flex-start',
  position: 'relative',
  transition: 'box-shadow 0.2s',
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
    <div style={{ minHeight: '100vh', background: '#f2e9e4', padding: '40px 0' }}>
      <div style={{ maxWidth: 1200, margin: '0 auto', display: 'flex', justifyContent: 'space-between', alignItems: 'center', marginBottom: 32 }}>
        <h1 style={{ color: '#4a4e69', letterSpacing: 1 }}>All Balances</h1>
        <Link
          to="/balances"
          style={{
            display: 'inline-block',
            background: 'linear-gradient(90deg, #4a4e69 60%, #b5ead7 100%)',
            color: '#fff',
            border: 'none',
            borderRadius: 8,
            padding: '10px 28px',
            fontWeight: 700,
            fontSize: 16,
            boxShadow: '0 2px 8px #b5ead722',
            textDecoration: 'none',
            transition: 'transform 0.15s cubic-bezier(.4,2,.6,1), box-shadow 0.15s',
            cursor: 'pointer',
          }}
          onMouseOver={e => { (e.currentTarget as HTMLAnchorElement).style.transform = 'scale(1.06)'; (e.currentTarget as HTMLAnchorElement).style.boxShadow = '0 4px 16px #b5ead744'; }}
          onMouseOut={e => { (e.currentTarget as HTMLAnchorElement).style.transform = 'scale(1)'; (e.currentTarget as HTMLAnchorElement).style.boxShadow = '0 2px 8px #b5ead722'; }}
        >
          Back to Balance Top-Up
        </Link>
      </div>
      {loading ? (
        <div style={{ textAlign: 'center', color: '#4a4e69', fontWeight: 500 }}>Loading...</div>
      ) : (
        <div style={{ display: 'flex', flexWrap: 'wrap', justifyContent: 'center' }}>
          {balances.map(balance => (
            <div key={balance.id} style={balanceCardStyle}>
              <div style={{ fontWeight: 700, fontSize: 18, color: '#22223b', marginBottom: 8 }}>Balance #{balance.id}</div>
              <div style={{ marginBottom: 6 }}><span style={{ color: '#4a4e69', fontWeight: 600 }}>User ID:</span> {balance.user_id}</div>
              <div style={{ marginBottom: 6 }}><span style={{ color: '#4a4e69', fontWeight: 600 }}>Card ID:</span> {balance.card_id}</div>
              <div style={{ marginBottom: 6 }}><span style={{ color: '#4a4e69', fontWeight: 600 }}>Balance:</span> {balance.balance}</div>
              <div style={{ marginBottom: 6 }}><span style={{ color: '#4a4e69', fontWeight: 600 }}>Ride Cost:</span> {balance.ride_cost}</div>
              <div style={{ marginBottom: 6 }}><span style={{ color: '#4a4e69', fontWeight: 600 }}>Updated At:</span> {balance.updated_at ? new Date(balance.updated_at).toLocaleString() : 'N/A'}</div>
            </div>
          ))}
        </div>
      )}
    </div>
  );
};

export default BalancesListPage; 