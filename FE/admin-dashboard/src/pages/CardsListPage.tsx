import React, { useEffect, useState } from 'react';
import { fetchCardsList, getCardById, updateCard, deleteCard } from '../api';
import ReactModal from 'react-modal';
import { Link } from 'react-router-dom';
import GetByIdCard from '../components/GetByIdCard';

const cardStyle: React.CSSProperties = {
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
    await updateCard(editCard.id, editForm);
    setActionMsg('Card updated!');
    setEditModalOpen(false);
    setTimeout(() => setActionMsg(''), 2000);
    setLoading(true);
    fetchCardsList({ limit: 1000, offset: 0 }).then(data => {
      const cardsArr = Array.isArray(data) ? data : data.cards || [];
      setCards(cardsArr);
      setLoading(false);
    });
  };

  const handleDelete = async (cardId: number) => {
    await deleteCard(cardId);
    setActionMsg('Card deleted!');
    setTimeout(() => setActionMsg(''), 2000);
    setLoading(true);
    fetchCardsList({ limit: 1000, offset: 0 }).then(data => {
      const cardsArr = Array.isArray(data) ? data : data.cards || [];
      setCards(cardsArr);
      setLoading(false);
    });
  };

  return (
    <div className="fade-in" style={{ minHeight: '100vh', background: 'var(--color-bg)', padding: '40px 0', display: 'flex', flexDirection: 'column', alignItems: 'center', justifyContent: 'center' }}>
      <div style={{ maxWidth: 1200, margin: '0 auto', display: 'flex', justifyContent: 'space-between', alignItems: 'center', marginBottom: 32, width: '100%' }}>
        <h1 style={{ color: 'var(--color-primary)', letterSpacing: 1, fontWeight: 800, fontSize: 32 }}>All Cards</h1>
        <Link
          to="/cards"
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
          Back to Card Assignment
        </Link>
      </div>
      <div style={{ maxWidth: 400, width: '100%', marginLeft: 'auto', marginRight: 'auto', marginBottom: 32 }}>
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
        <div style={{ textAlign: 'center', color: 'var(--color-primary)', fontWeight: 500 }}>Loading...</div>
      ) : (
        <div style={{ display: 'flex', flexWrap: 'wrap', justifyContent: 'center', width: '100%', maxWidth: 1200 }}>
          {cards.map(card => (
            <div key={card.id} className="scale-hover fade-in" style={cardStyle}>
              <div style={{ fontWeight: 700, fontSize: 18, color: 'var(--color-bg-dark)', marginBottom: 8 }}>Card #{card.id}</div>
              <div style={{ marginBottom: 6 }}><span style={{ color: 'var(--color-primary)', fontWeight: 600 }}>User ID:</span> {card.user_id}</div>
              <div style={{ marginBottom: 6 }}><span style={{ color: 'var(--color-primary)', fontWeight: 600 }}>Card Number:</span> {card.card_id}</div>
              <div style={{ marginBottom: 6 }}><span style={{ color: 'var(--color-primary)', fontWeight: 600 }}>Type:</span> {card.type || 'N/A'}</div>
              <div style={{ marginBottom: 6 }}><span style={{ color: 'var(--color-primary)', fontWeight: 600 }}>Device ID:</span> {card.device_id || 'N/A'}</div>
              <div style={{ display: 'flex', gap: 8, marginTop: 12 }}>
                <button onClick={() => handleEdit(card.id)} className="scale-hover" style={{ background: 'var(--color-primary)', color: '#fff', border: 'none', borderRadius: 6, padding: '8px 16px', fontWeight: 600, cursor: 'pointer', transition: 'background var(--transition), transform var(--transition)' }}>Edit</button>
                <button onClick={() => handleDelete(card.id)} className="scale-hover" style={{ background: '#c72c41', color: '#fff', border: 'none', borderRadius: 6, padding: '8px 16px', fontWeight: 600, cursor: 'pointer', transition: 'background var(--transition), transform var(--transition)' }}>Delete</button>
              </div>
            </div>
          ))}
        </div>
      )}
      <ReactModal isOpen={editModalOpen} onRequestClose={() => setEditModalOpen(false)} ariaHideApp={false} style={{ content: { maxWidth: 400, margin: 'auto', borderRadius: 16, padding: 0, boxShadow: '0 4px 24px var(--color-shadow)', background: 'linear-gradient(135deg, var(--color-card) 60%, var(--color-accent) 100%)', border: 'none' } }}>
        <div style={{ padding: 24, borderRadius: 16 }}>
          <h2 style={{ fontWeight: 700, fontSize: 20, color: 'var(--color-bg-dark)', marginBottom: 18 }}>Edit Card</h2>
          <form onSubmit={e => { e.preventDefault(); handleEditSave(); }}>
            <label style={{ fontWeight: 600, color: 'var(--color-primary)' }}>Card ID:</label>
            <input value={editForm.card_id} onChange={e => setEditForm(f => ({ ...f, card_id: e.target.value }))} required style={{ width: '100%', marginBottom: 12, padding: 8, borderRadius: 8, border: '1px solid #bcbcbc', fontSize: 16 }} />
            <label style={{ fontWeight: 600, color: 'var(--color-primary)' }}>User ID:</label>
            <input value={editForm.user_id} onChange={e => setEditForm(f => ({ ...f, user_id: e.target.value }))} required style={{ width: '100%', marginBottom: 12, padding: 8, borderRadius: 8, border: '1px solid #bcbcbc', fontSize: 16 }} />
            <label style={{ fontWeight: 600, color: 'var(--color-primary)' }}>Device ID:</label>
            <input value={editForm.device_id} onChange={e => setEditForm(f => ({ ...f, device_id: e.target.value }))} required style={{ width: '100%', marginBottom: 12, padding: 8, borderRadius: 8, border: '1px solid #bcbcbc', fontSize: 16 }} />
            <label style={{ fontWeight: 600, color: 'var(--color-primary)' }}>Type:</label>
            <input value={editForm.type} onChange={e => setEditForm(f => ({ ...f, type: e.target.value }))} required style={{ width: '100%', marginBottom: 18, padding: 8, borderRadius: 8, border: '1px solid #bcbcbc', fontSize: 16 }} />
            <div style={{ display: 'flex', justifyContent: 'flex-end', gap: 12 }}>
              <button type="button" onClick={() => setEditModalOpen(false)} className="scale-hover" style={{ background: '#bcbcbc', color: '#fff', border: 'none', borderRadius: 6, padding: '8px 16px', fontWeight: 600, cursor: 'pointer', transition: 'background var(--transition), transform var(--transition)' }}>Cancel</button>
              <button type="submit" className="scale-hover" style={{ background: 'var(--color-primary)', color: '#fff', border: 'none', borderRadius: 6, padding: '8px 16px', fontWeight: 600, cursor: 'pointer', transition: 'background var(--transition), transform var(--transition)' }}>Save</button>
            </div>
          </form>
        </div>
      </ReactModal>
    </div>
  );
};

export default CardsListPage; 