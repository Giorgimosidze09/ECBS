import React, { useEffect, useState } from 'react';
import { fetchChargesList } from '../api';
import { Link } from 'react-router-dom';

const chargeCardStyle: React.CSSProperties = {
  background: 'linear-gradient(135deg, var(--color-card) 60%, var(--color-bg) 100%)',
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
    <div className="fade-in" style={{ minHeight: '100vh', background: 'var(--color-bg)', padding: '40px 0', display: 'flex', flexDirection: 'column', alignItems: 'center', justifyContent: 'center' }}>
      <div style={{ maxWidth: 1200, margin: '0 auto', display: 'flex', justifyContent: 'space-between', alignItems: 'center', marginBottom: 32, width: '100%' }}>
        <h1 style={{ color: 'var(--color-primary)', letterSpacing: 1, fontWeight: 800, fontSize: 32 }}>All Charges</h1>
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
          {charges.map(charge => (
            <div key={charge.id} className="scale-hover fade-in" style={chargeCardStyle}>
              <div style={{ fontWeight: 700, fontSize: 18, color: 'var(--color-bg-dark)', marginBottom: 8 }}>Charge #{charge.id}</div>
              <div style={{ marginBottom: 6 }}><span style={{ color: 'var(--color-primary)', fontWeight: 600 }}>User ID:</span> {charge.user_id}</div>
              <div style={{ marginBottom: 6 }}><span style={{ color: 'var(--color-primary)', fontWeight: 600 }}>Amount:</span> {charge.amount}</div>
              <div style={{ marginBottom: 6 }}><span style={{ color: 'var(--color-primary)', fontWeight: 600 }}>Type:</span> {charge.type}</div>
              <div style={{ marginBottom: 6 }}><span style={{ color: 'var(--color-primary)', fontWeight: 600 }}>Description:</span> {charge.description}</div>
              <div style={{ marginBottom: 6 }}><span style={{ color: 'var(--color-primary)', fontWeight: 600 }}>Created At:</span> {charge.created_at ? new Date(charge.created_at).toLocaleString() : 'N/A'}</div>
            </div>
          ))}
        </div>
      )}
    </div>
  );
};

export default ChargesListPage; 