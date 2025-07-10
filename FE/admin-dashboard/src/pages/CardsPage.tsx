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

    return (
        <div className="fade-in" style={{ minHeight: '100vh', background: 'var(--color-bg)', padding: '40px 0', display: 'flex', flexDirection: 'column', alignItems: 'center', justifyContent: 'center' }}>
            <div style={{ maxWidth: 1200, margin: '0 auto', display: 'flex', justifyContent: 'space-between', alignItems: 'center', marginBottom: 32, width: '100%' }}>
                <h1 style={{ color: 'var(--color-primary)', letterSpacing: 1, fontWeight: 800, fontSize: 32 }}>Card Assignment</h1>
                <Link
                    to="/cards/list"
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
                    View All Cards
                </Link>
            </div>
            <div className="fade-in scale-hover" style={{ maxWidth: 420, width: '100%', margin: '0 auto', background: 'linear-gradient(135deg, var(--color-primary) 30%, var(--color-accent) 90%)', borderRadius: 'var(--radius)', boxShadow: '0 4px 24px var(--color-shadow)', padding: 32, marginBottom: 32, display: 'flex', flexDirection: 'column', alignItems: 'center', transition: 'box-shadow var(--transition)' }}>
                <CardAssignmentForm/>
            </div>
        </div>
    );
};

export default CardsPage;