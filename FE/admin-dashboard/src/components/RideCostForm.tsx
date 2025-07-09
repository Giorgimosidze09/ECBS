import React, { useState } from 'react';
import { rideCost } from '../api';

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

const RideCostForm: React.FC = () => {
  const [rideCostChange, SetRideCost] = useState('');
  const [result, setResult] = useState<boolean | null>(null); // explicitly boolean or null
  const [error, setError] = useState('');

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setError('');
    setResult(null);

    try {
      const response = await rideCost({ ride_cost: Number(rideCostChange) });
      // assuming backend returns `true` or `false`
      setResult(response);
    } catch (err) {
      setError('Failed to Change Ride Cost. Please try again.');
    }
  };

  return (
    <form
      onSubmit={handleSubmit}
      style={{
        width: '100%',
        display: 'flex',
        flexDirection: 'column',
        alignItems: 'center'
      }}
    >
      <label
        className="styled-label"
        style={{ display: 'block', marginBottom: 6, color: '#4a4e69', width: '100%' }}
      >
        Ride Cost Amount:
      </label>
      <input
        className="styled-input"
        type="number"
        value={rideCostChange}
        onChange={e => SetRideCost(e.target.value)}
        required
        style={inputStyle}
      />
      <button
        className="styled-button"
        type="submit"
        style={{
          width: '100%',
          padding: '12px 0',
          borderRadius: 6,
          background: '#22223b',
          color: '#fff',
          fontWeight: 600,
          fontSize: '1rem',
          marginTop: 8
        }}
      >
        Update Ride Cost
      </button>

      {error && <div style={{ color: 'red', marginTop: 12 }}>{error}</div>}

      {result !== null && (
        <div
          style={{
            marginTop: 16,
            background: '#fff',
            padding: 16,
            borderRadius: 8,
            color: result ? 'green' : 'red',
            textAlign: 'center'
          }}
        >
          {result ? 'Ride cost updated successfully!' : 'Failed to update ride cost.'}
        </div>
      )}
    </form>
  );
};

export default RideCostForm;
