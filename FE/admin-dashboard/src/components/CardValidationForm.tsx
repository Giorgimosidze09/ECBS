import React, { useState } from 'react';
import { validateCard } from '../api';

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

const CardValidationForm: React.FC = () => {
    const [cardId, setCardId] = useState('');
    const [result, setResult] = useState<any>(null);
    const [error, setError] = useState('');

    const handleSubmit = async (e: React.FormEvent) => {
        e.preventDefault();
        setError('');
        setResult(null);

        try {
            const response = await validateCard({ card_id: Number(cardId) });
            setResult(response);
        } catch (err) {
            setError('Failed to validate card. Please try again.');
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
            <label className="styled-label" style={{ display: 'block', marginBottom: 6, color: '#4a4e69', width: '100%' }}>
                Card ID:
            </label>
            <input
                className="styled-input"
                type="number"
                value={cardId}
                onChange={e => setCardId(e.target.value)}
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
                Validate Card
            </button>
            {error && <div style={{ color: 'red', marginTop: 12 }}>{error}</div>}
            {result && (
                <div style={{
                    marginTop: 16,
                    background: '#fff',
                    padding: 16,
                    borderRadius: 8,
                    color: result.isValid ? 'green' : 'red',
                    textAlign: 'center'
                }}>
                    {result.message || (result.isValid ? 'Card is valid!' : 'Card is invalid!')}
                </div>
            )}
        </form>
    );
};

export default CardValidationForm;