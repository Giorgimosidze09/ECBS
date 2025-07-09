import React, { useEffect, useState } from 'react';
import { fetchChargesList } from '../api';
import { Link } from 'react-router-dom';

const chargeCardStyle: React.CSSProperties = {
  background: 'linear-gradient(135deg, #fff 60%, #f2e9e4 100%)',
  borderRadius: 16,
  boxShadow: '0 4px 24px #c9ada722',
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

const ChargesListPage: React.FC = () => {
  const [charges, setCharges] = useState<any[]>([]);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    setLoading(true);
    fetchChargesList({ limit: 1000, offset: 0 }).then(data => {
      const rows = Array.isArray(data) ? data : data.rows || data.charges || [];
      setCharges(rows);
      setLoading(false);
    });
  }, []);

  return (
    <div style={{ minHeight: '100vh', background: '#f2e9e4', padding: '40px 0' }}>
      <div style={{ maxWidth: 1200, margin: '0 auto', display: 'flex', justifyContent: 'space-between', alignItems: 'center', marginBottom: 32 }}>
        <h1 style={{ color: '#4a4e69', letterSpacing: 1 }}>All Charges</h1>
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
          {charges.map(charge => (
            <div key={charge.id} style={chargeCardStyle}>
              <div style={{ fontWeight: 700, fontSize: 18, color: '#22223b', marginBottom: 8 }}>Charge #{charge.id}</div>
              <div style={{ marginBottom: 6 }}><span style={{ color: '#4a4e69', fontWeight: 600 }}>User ID:</span> {charge.user_id}</div>
              <div style={{ marginBottom: 6 }}><span style={{ color: '#4a4e69', fontWeight: 600 }}>Amount:</span> {charge.amount}</div>
              <div style={{ marginBottom: 6 }}><span style={{ color: '#4a4e69', fontWeight: 600 }}>Type:</span> {charge.type}</div>
              <div style={{ marginBottom: 6 }}><span style={{ color: '#4a4e69', fontWeight: 600 }}>Description:</span> {charge.description}</div>
              <div style={{ marginBottom: 6 }}><span style={{ color: '#4a4e69', fontWeight: 600 }}>Created At:</span> {charge.created_at ? new Date(charge.created_at).toLocaleString() : 'N/A'}</div>
            </div>
          ))}
        </div>
      )}
    </div>
  );
};

export default ChargesListPage; 