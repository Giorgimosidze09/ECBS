import React, { useState } from 'react';
import { activateCard } from '../api';
import { CardActivationRequest } from '../types';

const inputStyle: React.CSSProperties = {
  width: '100%',
  padding: '10px 12px',
  borderRadius: 6,
  border: '1px solid #bcbcbc',
  fontSize: '1rem',
  background: '#fff',
  boxSizing: 'border-box',
  marginBottom: 12
};

const CardActivationForm: React.FC = () => {
  const [cardId, setCardId] = useState('');
  const [activationStart, setActivationStart] = useState('');
  const [activationEnd, setActivationEnd] = useState('');
  const [loading, setLoading] = useState(false);
  const [success, setSuccess] = useState<string | null>(null);
  const [error, setError] = useState<string | null>(null);

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setLoading(true);
    setSuccess(null);
    setError(null);
    try {
      const req: CardActivationRequest = {
        cardId: Number(cardId),
        activationStart,
        activationEnd
      };
      await activateCard(req);
      setSuccess('Card activated successfully!');
      setCardId('');
      setActivationStart('');
      setActivationEnd('');
    } catch {
      setError('Failed to activate card.');
    } finally {
      setLoading(false);
    }
  };

  return (
    <form onSubmit={handleSubmit} style={{ width: '100%', display: 'flex', flexDirection: 'column', alignItems: 'center' }}>
      <label style={{ ...inputStyle, marginBottom: 0, background: 'none', border: 'none', color: '#4a4e69' }}>Card ID</label>
      <input
        type="number"
        value={cardId}
        onChange={e => setCardId(e.target.value)}
        required
        style={inputStyle}
        placeholder="Enter Card ID"
      />
      <label style={{ ...inputStyle, marginBottom: 0, background: 'none', border: 'none', color: '#4a4e69' }}>Activation Start</label>
      <input
        type="date"
        value={activationStart}
        onChange={e => setActivationStart(e.target.value)}
        required
        style={inputStyle}
      />
      <label style={{ ...inputStyle, marginBottom: 0, background: 'none', border: 'none', color: '#4a4e69' }}>Activation End</label>
      <input
        type="date"
        value={activationEnd}
        onChange={e => setActivationEnd(e.target.value)}
        required
        style={inputStyle}
      />
      <button
        type="submit"
        className="styled-button"
        style={{
          width: '100%',
          padding: '12px 0',
          borderRadius: 6,
          background: '#22223b',
          color: '#fff',
          fontWeight: 600,
          fontSize: '1rem',
          marginTop: 8,
          cursor: loading ? 'not-allowed' : 'pointer',
          opacity: loading ? 0.6 : 1
        }}
        disabled={loading}
      >
        {loading ? 'Activating...' : 'Activate Card'}
      </button>
      {success && <div style={{ color: 'lightgreen', marginTop: 12 }}>{success}</div>}
      {error && <div style={{ color: 'salmon', marginTop: 12 }}>{error}</div>}
    </form>
  );
};

export default CardActivationForm; 