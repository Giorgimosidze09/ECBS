import React, { useState } from 'react';

interface GetByIdCardProps<T> {
  label: string;
  fetchById: (id: number) => Promise<T>;
  renderResult: (result: T) => React.ReactNode;
}

const cardStyles: Record<string, React.CSSProperties> = {
  User: {
    background: 'linear-gradient(135deg, #fff 60%, #9a8c98 100%)',
    borderRadius: 16,
    boxShadow: '0 4px 24px #9a8c9822',
    padding: 24,
    margin: 16,
    minWidth: 260,
    maxWidth: 320,
    display: 'flex',
    flexDirection: 'column',
    alignItems: 'flex-start',
    position: 'relative',
    transition: 'box-shadow 0.2s',
  },
  Card: {
    background: 'linear-gradient(135deg, #fff 60%, #c9ada7 100%)',
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
  },
  Device: {
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
  },
};

function GetByIdCard<T>({ label, fetchById, renderResult }: GetByIdCardProps<T>) {
  const [id, setId] = useState('');
  const [result, setResult] = useState<T | null>(null);
  const [error, setError] = useState<string | null>(null);
  const [loading, setLoading] = useState(false);

  const handleFetch = async (e: React.FormEvent) => {
    e.preventDefault();
    setError(null);
    setResult(null);
    setLoading(true);

    const numId = Number(id);
    if (!id || isNaN(numId) || numId <= 0 || !Number.isInteger(numId)) {
      setError('Please enter a valid numeric ID.');
      setLoading(false);
      return;
    }

    try {
      const data = await fetchById(numId);
      setResult(data);
    } catch (err) {
      setError('Not found or error fetching.');
    } finally {
      setLoading(false);
    }
  };

  const style = cardStyles[label] || cardStyles.User;

  return (
    <div style={style}>
      <div style={{ width: '100%' }}>
        <div style={{ fontWeight: 700, fontSize: 18, color: '#22223b', marginBottom: 12 }}>{label} Lookup</div>
        <form onSubmit={handleFetch} style={{ display: 'flex', gap: 8, marginBottom: 12, width: '100%' }}>
          <input
            type="number"
            value={id}
            onChange={e => setId(e.target.value)}
            required
            placeholder={`Enter ${label.toLowerCase()} ID`}
            style={{ flex: 1, padding: 8, borderRadius: 6, border: '1px solid #bcbcbc', fontSize: 15 }}
          />
          <button type="submit" disabled={loading || !id} style={{ padding: '8px 16px', borderRadius: 6, background: '#4a4e69', color: '#fff', fontWeight: 600, cursor: loading ? 'not-allowed' : 'pointer', fontSize: 15 }}>
            {loading ? 'Fetching...' : 'Fetch'}
          </button>
        </form>
        {error && <div style={{ color: '#c72c41', marginBottom: 8 }}>{error}</div>}
        {result && (
          <div style={{ marginTop: 8, width: '100%' }}>{renderResult(result)}</div>
        )}
      </div>
    </div>
  );
}

export default GetByIdCard; 