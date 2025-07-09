import React, { useEffect, useState, useRef } from 'react';
import { topUpBalance, fetchUsersList, fetchCardsList, fetchBalanceList } from '../api';

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

const popupStyle: React.CSSProperties = {
    position: 'absolute',
    zIndex: 10,
    background: '#fff',
    border: '1px solid #bcbcbc',
    borderRadius: 6,
    boxShadow: '0 4px 16px #0001',
    width: '100%',
    maxHeight: 180,
    overflowY: 'auto',
    marginTop: 2
};

const popupItemStyle: React.CSSProperties = {
    padding: '10px 12px',
    cursor: 'pointer',
    borderBottom: '1px solid #eee',
    background: '#fff'
};

const popupItemActiveStyle: React.CSSProperties = {
    ...popupItemStyle,
    background: '#f2e9e4'
};

const BalanceTopUpForm: React.FC = () => {
    const [users, setUsers] = useState<{ id: number; name: string }[]>([]);
    const [cards, setCards] = useState<{ id: number; card_id: string; user_id: number }[]>([]);
    const [userId, setUserId] = useState('');
    const [cardId, setCardId] = useState('');
    const [amount, setAmount] = useState('');
    const [defaultRideCost, setDefaultRideCost] = useState('0');
    const [error, setError] = useState('');
    const [success, setSuccess] = useState('');
    const [userSearch, setUserSearch] = useState('');
    const [showPopup, setShowPopup] = useState(false);
    const [activeIndex, setActiveIndex] = useState(-1);
    const inputRef = useRef<HTMLInputElement>(null);

    useEffect(() => {
        fetchUsersList({ limit: 1000, offset: 0 }).then(data => {
            const usersArr = Array.isArray(data) ? data : data.users || [];
            setUsers(usersArr.map((u: any) => ({
                id: u.id ?? u.ID,
                name: u.name ?? u.Name
            })));
        });
        fetchCardsList({ limit: 1000, offset: 0 }).then(data => {
            const cardsArr = Array.isArray(data) ? data : data.cards || [];
            setCards(cardsArr.map((c: any) => ({
                id: c.id ?? c.ID,
                card_id: c.card_id ?? c.CardID,
                user_id: c.user_id ?? c.UserID
            })));
        });
    }, []);

    useEffect(() => {
        async function loadDefaultRideCost() {
            try {
                const data = await fetchBalanceList({ limit: 1, offset: 0 });
                const rows = Array.isArray(data) ? data : data.rows || data.balances || [];
                if (rows.length > 0) {
                    const cost = rows[0].ride_cost ?? rows[0].rideCost ?? rows[0].RideCost ?? 0;
                    setDefaultRideCost(String(cost));
                }
            } catch {
                setDefaultRideCost('0');
            }
        }
        loadDefaultRideCost();
    }, []);

    const filteredUsers = users.filter(user =>
        user.name.toLowerCase().includes(userSearch.trim().toLowerCase()) ||
        String(user.id).includes(userSearch.trim())
    );

    useEffect(() => {
        const found = users.find(
            u =>
                `${u.name} (ID: ${u.id})`.toLowerCase() === userSearch.trim().toLowerCase() ||
                String(u.id) === userSearch.trim()
        );
        if (found) {
            setUserId(String(found.id));
        } else {
            setUserId('');
        }
    }, [userSearch, users]);

    const userCards = cards.filter(card => String(card.user_id) === userId);

    useEffect(() => {
        if (userId && userCards.length > 0) {
            setCardId(String(userCards[0].id));
        } else {
            setCardId('');
        }
    }, [userId, userCards]);

    const handleKeyDown = (e: React.KeyboardEvent<HTMLInputElement>) => {
        if (!showPopup) return;
        if (e.key === 'ArrowDown') {
            setActiveIndex(i => Math.min(i + 1, filteredUsers.length - 1));
        } else if (e.key === 'ArrowUp') {
            setActiveIndex(i => Math.max(i - 1, 0));
        } else if (e.key === 'Enter' && activeIndex >= 0 && activeIndex < filteredUsers.length) {
            const user = filteredUsers[activeIndex];
            setUserSearch(`${user.name} (ID: ${user.id})`);
            setShowPopup(false);
            setActiveIndex(-1);
        } else if (e.key === 'Escape') {
            setShowPopup(false);
        }
    };

    useEffect(() => {
        const handleClick = (e: MouseEvent) => {
            if (inputRef.current && !inputRef.current.contains(e.target as Node)) {
                setShowPopup(false);
            }
        };
        if (showPopup) {
            document.addEventListener('mousedown', handleClick);
        }
        return () => document.removeEventListener('mousedown', handleClick);
    }, [showPopup]);

    const handleUserSelect = (user: { id: number; name: string }) => {
        setUserSearch(`${user.name} (ID: ${user.id})`);
        setShowPopup(false);
        setActiveIndex(-1);
    };

    const handleInputFocus = () => {
        setShowPopup(true);
        setActiveIndex(-1);
    };

    const handleInputChange = (e: React.ChangeEvent<HTMLInputElement>) => {
        setUserSearch(e.target.value);
        setShowPopup(true);
        setActiveIndex(-1);
    };

    const handleInputBlur = () => {
        setTimeout(() => setShowPopup(false), 150);
    };

    const handleInputClick = () => {
        setShowPopup(true);
    };

    const handleMouseEnter = (idx: number) => setActiveIndex(idx);

    const handleMouseLeave = () => setActiveIndex(-1);

    const handlePopupItemClick = (user: { id: number; name: string }) => {
        handleUserSelect(user);
        if (inputRef.current) inputRef.current.blur();
    };

    const handleInputKeyDown = (e: React.KeyboardEvent<HTMLInputElement>) => {
        handleKeyDown(e);
    };

    const isFormReady = userId !== '' && cardId !== '' && amount.trim() !== '';

   const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setError('');
    setSuccess('');

    try {
        const response = await topUpBalance({
            user_id: Number(userId),
            card_id: Number(cardId),
            balance: Number(amount),
            ride_cost: Number(defaultRideCost),
        });

        if (response && response.user_id && response.balance >= 0) {
      setSuccess('Balance topped up successfully!');
      setUserId('');
      setCardId('');
      setAmount('');
      setUserSearch('');
    } else {
      setError('Top up failed. Please try again.');
    }
  } catch (err) {
    setError('Failed to top up balance. Please try again.');
  }
};


    return (
        <form
            onSubmit={handleSubmit}
            style={{
                width: '100%',
                display: 'flex',
                flexDirection: 'column',
                alignItems: 'center',
                gap: 0
            }}
        >
            <h2 style={{ textAlign: 'center', color: '#4a4e69', marginBottom: 20 }}>Top Up Balance</h2>
            {error && <p style={{ color: 'red', textAlign: 'center', marginBottom: 12 }}>{error}</p>}
            {success && <p style={{ color: 'green', textAlign: 'center', marginBottom: 12 }}>{success}</p>}

            {/* User input */}
            <div style={{ width: '100%', position: 'relative' }}>
                <label className="styled-label" style={{ display: 'block', marginBottom: 6, color: '#4a4e69' }}>User:</label>
                <input
                    ref={inputRef}
                    className="styled-input"
                    type="text"
                    value={userSearch}
                    onChange={handleInputChange}
                    onFocus={handleInputFocus}
                    onBlur={handleInputBlur}
                    onClick={handleInputClick}
                    onKeyDown={handleInputKeyDown}
                    required
                    style={inputStyle}
                    placeholder="Type or select user"
                    autoComplete="off"
                />
                {showPopup && filteredUsers.length > 0 && (
                    <div style={popupStyle}>
                        {filteredUsers.map((user, idx) => (
                            <div
                                key={user.id}
                                style={activeIndex === idx ? popupItemActiveStyle : popupItemStyle}
                                onMouseDown={() => handlePopupItemClick(user)}
                                onMouseEnter={() => handleMouseEnter(idx)}
                                onMouseLeave={handleMouseLeave}
                            >
                                {user.name} <span style={{ color: '#888' }}>(ID: {user.id})</span>
                            </div>
                        ))}
                    </div>
                )}
            </div>

            {/* Show Card select ONLY if user selected */}
            {userId && (
                <div style={{ width: '100%' }}>
                    <label className="styled-label" style={{ display: 'block', marginBottom: 6, color: '#4a4e69' }}>Card:</label>
                    <select
                        className="styled-input"
                        value={cardId}
                        onChange={e => setCardId(e.target.value)}
                        required
                        style={inputStyle}
                    >
                        <option value="">Select card</option>
                        {userCards.map(card => (
                            <option key={card.id} value={card.id}>
                                {card.card_id} (ID: {card.id})
                            </option>
                        ))}
                    </select>
                </div>
            )}

            {/* Show Amount input ONLY if card selected */}
            {cardId && (
                <div style={{ width: '100%' }}>
                    <label className="styled-label" style={{ display: 'block', marginBottom: 6, color: '#4a4e69' }}>Amount:</label>
                    <input
                        className="styled-input"
                        type="number"
                        value={amount}
                        onChange={e => setAmount(e.target.value)}
                        required
                        style={inputStyle}
                    />
                </div>
            )}

            <input type="hidden" name="ride_cost" value={defaultRideCost} />

            {/* Submit button enabled only if all fields filled */}
            <button
                className="styled-button"
                type="submit"
                disabled={!isFormReady}
                style={{
                    width: '100%',
                    padding: '12px 0',
                    borderRadius: 6,
                    background: isFormReady ? '#22223b' : '#999',
                    color: '#fff',
                    fontWeight: 600,
                    fontSize: '1rem',
                    marginTop: 8,
                    cursor: isFormReady ? 'pointer' : 'not-allowed',
                    transition: 'background 0.3s ease-in-out'
                }}
            >
                Top Up
            </button>
        </form>
    );
};

export default BalanceTopUpForm;
