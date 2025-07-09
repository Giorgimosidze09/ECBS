import React, { useEffect, useState } from 'react';
import CardAssignmentForm from '../components/CardAssignmentForm';
import { fetchCardsList, CardsListOutput, fetchDevicesList, updateCard, deleteCard, getCardById } from '../api';
import ReactModal from 'react-modal';
import { Link } from 'react-router-dom';

const CardsPage: React.FC = () => {
    const [cards, setCards] = useState<CardsListOutput[]>([]);
    const [total, setTotal] = useState(0);
    const [loading, setLoading] = useState(true);
    const [page, setPage] = useState(0);
    const [limit, setLimit] = useState(25);
    const [search, setSearch] = useState('');
    const [editCard, setEditCard] = useState<any | null>(null);
    const [editModalOpen, setEditModalOpen] = useState(false);
    const [editForm, setEditForm] = useState({ card_id: '', user_id: '', device_id: '', type: '' });
    const [actionMsg, setActionMsg] = useState('');

    useEffect(() => {
        setLoading(true);
        const offset = page * limit;
        fetchCardsList({ limit, offset }).then(data => {
            const cardsArr = Array.isArray(data) ? data : data.cards || [];
            const totalVal =
                cardsArr.length > 0
                    ? cardsArr[0].total || cardsArr[0].Total || cardsArr.length
                    : 0;
            setCards(cardsArr);
            setTotal(totalVal);
            setLoading(false);
        });
    }, [page, limit]);

    const totalPages = Math.max(1, Math.ceil(total / limit));

   


    // Frontend filter: search by user_id (exact or partial match)
    const filteredCards = cards.filter(card =>
        !search.trim() ||
        (card.user_id && String(card.user_id).toLowerCase().includes(search.trim().toLowerCase()))
    );

    // For the datalist in the input, get unique user_ids
    const userIdOptions = Array.from(new Set(cards.map(card => card.user_id).filter(Boolean)));

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
            await updateCard(editCard.id, { ...editForm, id: editCard.id });
            setActionMsg('Card updated!');
            setEditModalOpen(false);
            setEditCard(null);
            setEditForm({ card_id: '', user_id: '', device_id: '', type: '' });
            setTimeout(() => setActionMsg(''), 2000);
            // reload cards
            const offset = page * limit;
            fetchCardsList({ limit, offset }).then(data => {
                const cardsArr = Array.isArray(data) ? data : data.cards || [];
                setCards(cardsArr);
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
            const offset = page * limit;
            fetchCardsList({ limit, offset }).then(data => {
                const cardsArr = Array.isArray(data) ? data : data.cards || [];
                setCards(cardsArr);
            });
        } catch {
            setActionMsg('Failed to delete card.');
        }
    };

    return (
        <div
            style={{
                minHeight: '100vh',
                background: '#f2e9e4',
                padding: '40px 0',
            }}
        >
            <div style={{ maxWidth: 1200, margin: '0 auto', display: 'flex', justifyContent: 'space-between', alignItems: 'center', marginBottom: 32 }}>
                <h1 style={{ color: '#4a4e69', letterSpacing: 1 }}>Card Assignment</h1>
                <Link
                    to="/cards/list"
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
                    }}
                    onMouseOver={e => { (e.currentTarget as HTMLAnchorElement).style.transform = 'scale(1.06)'; (e.currentTarget as HTMLAnchorElement).style.boxShadow = '0 4px 16px #c9ada744'; }}
                    onMouseOut={e => { (e.currentTarget as HTMLAnchorElement).style.transform = 'scale(1)'; (e.currentTarget as HTMLAnchorElement).style.boxShadow = '0 2px 8px #c9ada722'; }}
                >
                    View All Cards 
                </Link>
            </div>
            <div
                style={{
                    maxWidth: 420,
                    margin: '0 auto',
                    background: 'linear-gradient(135deg, #4a4e69 30%, #9a8c98 90%)',
                    borderRadius: 16,
                    boxShadow: '0 4px 24px #4a4e6922',
                    padding: 32,
                    marginBottom: 32,
                    display: 'flex',
                    flexDirection: 'column',
                    alignItems: 'center'
                }}
            >
                <CardAssignmentForm/>
            </div>
            <ReactModal isOpen={editModalOpen} onRequestClose={() => setEditModalOpen(false)} ariaHideApp={false} style={{ content: { maxWidth: 400, margin: 'auto', borderRadius: 12 } }}>
                <h2>Edit Card</h2>
                <form onSubmit={e => { e.preventDefault(); handleEditSave(); }}>
                    <label>Card ID:</label>
                    <input value={editForm.card_id} onChange={e => setEditForm(f => ({ ...f, card_id: e.target.value }))} required style={{ width: '100%', marginBottom: 8 }} />
                    <label>User ID:</label>
                    <input value={editForm.user_id} onChange={e => setEditForm(f => ({ ...f, user_id: e.target.value }))} required style={{ width: '100%', marginBottom: 8 }} />
                    <label>Device ID:</label>
                    <input value={editForm.device_id} onChange={e => setEditForm(f => ({ ...f, device_id: e.target.value }))} required style={{ width: '100%', marginBottom: 8 }} />
                    <label>Type:</label>
                    <input value={editForm.type} onChange={e => setEditForm(f => ({ ...f, type: e.target.value }))} required style={{ width: '100%', marginBottom: 8 }} />
                    <div style={{ display: 'flex', justifyContent: 'flex-end', gap: 8 }}>
                        <button type="button" onClick={() => setEditModalOpen(false)}>Cancel</button>
                        <button type="submit">Save</button>
                    </div>
                </form>
            </ReactModal>
            {actionMsg && <div style={{ color: actionMsg.includes('Failed') ? 'red' : 'green', textAlign: 'center', marginTop: 12 }}>{actionMsg}</div>}
        </div>
    );
};

export default CardsPage;