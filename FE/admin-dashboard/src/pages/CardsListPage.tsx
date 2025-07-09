import React, { useEffect, useState } from 'react';
import { fetchCardsList, getCardById, updateCard, deleteCard } from '../api';
import ReactModal from 'react-modal';
import { Link } from 'react-router-dom';
import GetByIdCard from '../components/GetByIdCard';

const cardStyle: React.CSSProperties = {
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
};

const CardsListPage: React.FC = () => {
  const [cards, setCards] = useState<any[]>([]);
  const [loading, setLoading] = useState(true);
  const [editCard, setEditCard] = useState<any | null>(null);
  const [editModalOpen, setEditModalOpen] = useState(false);
  const [editForm, setEditForm] = useState({ card_id: '', user_id: '', device_id: '', type: '' });
  const [actionMsg, setActionMsg] = useState('');

  useEffect(() => {
    setLoading(true);
    fetchCardsList({ limit: 1000, offset: 0 }).then(data => {
      const cardsArr = Array.isArray(data) ? data : data.cards || [];
      setCards(cardsArr);
      setLoading(false);
    });
  }, []);

  const handleEdit = async (cardId: number) => {
    const card = await getCardById(cardId);
    setEditCard(card);
    setEditForm({
      card_id: card.card_id || '',
      user_id: card.user_id ? String(card.user_id) : '',
      device_id: card.device_id ? String(card.device_id) : '',
      type: card.type || ''
    });
    setEditModalOpen(true);
  };

  const handleEditSave = async () => {
    if (!editCard) return;
    try {
      await updateCard(editCard.id, {
        ...editForm,
        id: editCard.id,
        user_id: Number(editForm.user_id),
        device_id: Number(editForm.device_id),
      });
      setActionMsg('Card updated!');
      setEditModalOpen(false);
      setEditCard(null);
      setEditForm({ card_id: '', user_id: '', device_id: '', type: '' });
      setTimeout(() => setActionMsg(''), 2000);
      // reload cards
      setLoading(true);
      fetchCardsList({ limit: 1000, offset: 0 }).then(data => {
        const cardsArr = Array.isArray(data) ? data : data.cards || [];
        setCards(cardsArr);
        setLoading(false);
      });
    } catch {
      setActionMsg('Failed to update card.');
    }
  };

  const handleDelete = async (cardId: number) => {
    if (!window.confirm('Are you sure you want to delete this card?')) return;
    try {
      await deleteCard(cardId);
      setActionMsg('Card deleted!');
      setTimeout(() => setActionMsg(''), 2000);
      // reload cards
      setLoading(true);
      fetchCardsList({ limit: 1000, offset: 0 }).then(data => {
        const cardsArr = Array.isArray(data) ? data : data.cards || [];
        setCards(cardsArr);
        setLoading(false);
      });
    } catch {
      setActionMsg('Failed to delete card.');
    }
  };

  return (
    <div style={{ minHeight: '100vh', background: '#f2e9e4', padding: '40px 0' }}>
      <div style={{ maxWidth: 1200, margin: '0 auto', display: 'flex', justifyContent: 'space-between', alignItems: 'center', marginBottom: 32 }}>
        <h1 style={{ color: '#4a4e69', letterSpacing: 1 }}>All Cards</h1>
        <Link
          to="/cards"
          style={{
            display: 'inline-block',
            background: 'linear-gradient(90deg, #4a4e69 60%, #c9ada7 100%)',
            color: '#fff',
            border: 'none',
            borderRadius: 8,
            padding: '10px 28px',
            fontWeight: 700,
            fontSize: 16,
            boxShadow: '0 2px 8px #c9ada722',
            textDecoration: 'none',
            transition: 'transform 0.15s cubic-bezier(.4,2,.6,1), box-shadow 0.15s',
            cursor: 'pointer',
            outline: 'none',
            position: 'relative',
            overflow: 'hidden',
          }}
          onMouseEnter={e => (e.currentTarget.style.transform = 'scale(1.06)')}
          onMouseLeave={e => (e.currentTarget.style.transform = 'scale(1)')}
        >
          <span style={{ position: 'relative', zIndex: 2 }}>Back to Card Assignment</span>
        </Link>
      </div>
      <div style={{ display: 'flex', justifyContent: 'center', marginBottom: 32 }}>
        <GetByIdCard
          label="Card"
          fetchById={getCardById}
          renderResult={card => (
            <div>
              <div><b>ID:</b> {card.id}</div>
              <div><b>User ID:</b> {card.user_id}</div>
              <div><b>Card Number:</b> {card.card_id}</div>
              {card.type && <div><b>Type:</b> {card.type}</div>}
              {card.device_id && <div><b>Device ID:</b> {card.device_id}</div>}
            </div>
          )}
        />
      </div>
      {actionMsg && <div style={{ color: actionMsg.includes('Failed') ? 'red' : 'green', textAlign: 'center', marginBottom: 12 }}>{actionMsg}</div>}
      {loading ? (
        <div style={{ textAlign: 'center', color: '#4a4e69', fontWeight: 500 }}>Loading...</div>
      ) : (
        <div style={{ display: 'flex', flexWrap: 'wrap', justifyContent: 'center' }}>
          {cards.map(card => (
            <div key={card.id} style={cardStyle}>
              <div style={{ fontWeight: 700, fontSize: 18, color: '#22223b', marginBottom: 8 }}>Card #{card.id}</div>
              <div style={{ marginBottom: 6 }}><span style={{ color: '#4a4e69', fontWeight: 600 }}>User ID:</span> {card.user_id}</div>
              <div style={{ marginBottom: 6 }}><span style={{ color: '#4a4e69', fontWeight: 600 }}>Card Number:</span> {card.card_id}</div>
              <div style={{ marginBottom: 6 }}><span style={{ color: '#4a4e69', fontWeight: 600 }}>Type:</span> {card.type || 'N/A'}</div>
              <div style={{ marginBottom: 6 }}><span style={{ color: '#4a4e69', fontWeight: 600 }}>Device ID:</span> {card.device_id || 'N/A'}</div>
              <div style={{ display: 'flex', gap: 8, marginTop: 12 }}>
                <button onClick={() => handleEdit(card.id)} style={{ background: '#4a4e69', color: '#fff', border: 'none', borderRadius: 6, padding: '8px 16px', fontWeight: 600, cursor: 'pointer' }}>Edit</button>
                <button onClick={() => handleDelete(card.id)} style={{ background: '#c72c41', color: '#fff', border: 'none', borderRadius: 6, padding: '8px 16px', fontWeight: 600, cursor: 'pointer' }}>Delete</button>
              </div>
            </div>
          ))}
        </div>
      )}
      <ReactModal isOpen={editModalOpen} onRequestClose={() => setEditModalOpen(false)} ariaHideApp={false} style={{ content: { maxWidth: 400, margin: 'auto', borderRadius: 16, padding: 0, boxShadow: '0 4px 24px #c9ada722', background: 'linear-gradient(135deg, #fff 60%, #c9ada7 100%)', border: 'none' } }}>
        <div style={{ padding: 24, borderRadius: 16 }}>
          <h2 style={{ fontWeight: 700, fontSize: 20, color: '#22223b', marginBottom: 18 }}>Edit Card</h2>
          <form onSubmit={e => { e.preventDefault(); handleEditSave(); }}>
            <label style={{ fontWeight: 600, color: '#4a4e69' }}>Card ID:</label>
            <input value={editForm.card_id} onChange={e => setEditForm(f => ({ ...f, card_id: e.target.value }))} required style={{ width: '100%', marginBottom: 12, padding: 8, borderRadius: 8, border: '1px solid #bcbcbc', fontSize: 16 }} />
            <label style={{ fontWeight: 600, color: '#4a4e69' }}>User ID:</label>
            <input value={editForm.user_id} onChange={e => setEditForm(f => ({ ...f, user_id: e.target.value }))} required style={{ width: '100%', marginBottom: 12, padding: 8, borderRadius: 8, border: '1px solid #bcbcbc', fontSize: 16 }} />
            <label style={{ fontWeight: 600, color: '#4a4e69' }}>Device ID:</label>
            <input value={editForm.device_id} onChange={e => setEditForm(f => ({ ...f, device_id: e.target.value }))} required style={{ width: '100%', marginBottom: 12, padding: 8, borderRadius: 8, border: '1px solid #bcbcbc', fontSize: 16 }} />
            <label style={{ fontWeight: 600, color: '#4a4e69' }}>Type:</label>
            <input value={editForm.type} onChange={e => setEditForm(f => ({ ...f, type: e.target.value }))} required style={{ width: '100%', marginBottom: 18, padding: 8, borderRadius: 8, border: '1px solid #bcbcbc', fontSize: 16 }} />
            <div style={{ display: 'flex', justifyContent: 'flex-end', gap: 10 }}>
              <button type="button" onClick={() => setEditModalOpen(false)} style={{ background: '#bcbcbc', color: '#22223b', border: 'none', borderRadius: 8, padding: '8px 18px', fontWeight: 600, fontSize: 15, cursor: 'pointer' }}>Cancel</button>
              <button type="submit" style={{ background: '#4a4e69', color: '#fff', border: 'none', borderRadius: 8, padding: '8px 18px', fontWeight: 600, fontSize: 15, cursor: 'pointer' }}>Save</button>
            </div>
          </form>
        </div>
      </ReactModal>
    </div>
  );
};

export default CardsListPage; 