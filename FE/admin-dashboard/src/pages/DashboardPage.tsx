import React, { useEffect, useState } from 'react';
import { PieChart, Pie, Cell, Tooltip, ResponsiveContainer } from 'recharts';
import axios from 'axios';
import { fetchUsersList, fetchCardsList } from '../api';

const API_URL = 'http://localhost:8080';
const COLORS = ['var(--color-primary)', 'var(--color-secondary)', 'var(--color-accent)', 'var(--color-bg-dark)'];

const cardContainerStyle: React.CSSProperties = {
  flex: 1,
  maxWidth: 340,
  minWidth: 220,
  margin: 8,
  borderRadius: 'var(--radius)',
  boxShadow: '0 4px 24px var(--color-shadow)',
  background: 'var(--color-card)',
  transition: 'box-shadow var(--transition)',
  cursor: 'pointer',
  overflow: 'hidden',
  position: 'relative',
};

const expandableContentStyle: React.CSSProperties = {
  background: 'var(--color-bg)',
  borderTop: '1px solid var(--color-accent)',
  padding: '12px 20px',
  animation: 'fadeIn 0.4s',
};

const DashboardPage: React.FC = () => {
  const [stats, setStats] = useState({ users: 0, cards: 0, totalBalance: 0 });
  const [usersPreview, setUsersPreview] = useState<any[]>([]);
  const [cardsPreview, setCardsPreview] = useState<any[]>([]);
  const [expanded, setExpanded] = useState<'users' | 'cards' | null>(null);

  useEffect(() => {
    Promise.all([
      axios.get(`${API_URL}/stats/users`),
      axios.get(`${API_URL}/stats/cards`),
      axios.get(`${API_URL}/stats/total-balance`),
    ]).then(([usersRes, cardsRes, balanceRes]) => {
      setStats({
        users: usersRes.data.count,
        cards: cardsRes.data.count,
        totalBalance: balanceRes.data.total,
      });
    });
    fetchUsersList({ limit: 5, offset: 0 }).then(data => {
      setUsersPreview(Array.isArray(data) ? data : data.users || []);
    });
    fetchCardsList({ limit: 5, offset: 0 }).then(data => {
      setCardsPreview(Array.isArray(data) ? data : data.cards || []);
    });
  }, []);

  const pieData = [
    { name: 'Users', value: stats.users },
    { name: 'Cards', value: stats.cards },
  ];

  return (
    <div className="fade-in" style={{ padding: '32px 0', minHeight: '100vh', display: 'flex', flexDirection: 'column', alignItems: 'center', justifyContent: 'center' }}>
      <h1 style={{ color: 'var(--color-primary)', marginBottom: 32, textAlign: 'center', letterSpacing: 1, fontWeight: 800, fontSize: 36 }}>Dashboard</h1>
      <div style={{ display: 'flex', gap: 32, marginBottom: 32, flexWrap: 'wrap', justifyContent: 'center', width: '100%', maxWidth: 1200 }}>
        {/* Users Expandable Card */}
        <div
          className={`scale-hover fade-in`}
          style={{ ...cardContainerStyle, background: 'linear-gradient(135deg, var(--color-primary) 30%, var(--color-secondary) 90%)', color: '#fff' }}
          onClick={() => setExpanded(expanded === 'users' ? null : 'users')}
        >
          <h3 style={{ margin: '24px 0 0 0', fontWeight: 600, textAlign: 'center' }}>Total Users</h3>
          <p style={{ fontSize: 40, margin: 0, fontWeight: 700, textAlign: 'center' }}>{stats.users}</p>
          <div style={{ textAlign: 'center', fontSize: 14, opacity: 0.8, marginBottom: 8 }}>Click to preview</div>
          {expanded === 'users' && (
            <div style={expandableContentStyle}>
              <div style={{ fontWeight: 700, marginBottom: 8 }}>Top Users</div>
              <ul style={{ listStyle: 'none', padding: 0, margin: 0 }}>
                {usersPreview.map((user, idx) => (
                  <li key={user.id} style={{ padding: '4px 0', borderBottom: idx < usersPreview.length - 1 ? '1px solid #eee' : 'none' }}>
                    <span style={{ fontWeight: 600 }}>{user.name}</span> <span style={{ color: '#eee', fontSize: 13 }}>({user.email})</span>
                  </li>
                ))}
              </ul>
              <a href="/users/list" style={{ display: 'inline-block', marginTop: 12, color: 'var(--color-primary)', fontWeight: 700, textDecoration: 'none', background: '#fff', borderRadius: 8, padding: '6px 18px', boxShadow: '0 2px 8px var(--color-shadow)', transition: 'background 0.2s' }}>View All Users</a>
            </div>
          )}
        </div>
        {/* Cards Expandable Card */}
        <div
          className={`scale-hover fade-in`}
          style={{ ...cardContainerStyle, background: 'linear-gradient(135deg, var(--color-secondary) 30%, var(--color-accent) 90%)', color: '#22223b' }}
          onClick={() => setExpanded(expanded === 'cards' ? null : 'cards')}
        >
          <h3 style={{ margin: '24px 0 0 0', fontWeight: 600, textAlign: 'center' }}>Total Cards</h3>
          <p style={{ fontSize: 40, margin: 0, fontWeight: 700, textAlign: 'center' }}>{stats.cards}</p>
          <div style={{ textAlign: 'center', fontSize: 14, opacity: 0.8, marginBottom: 8 }}>Click to preview</div>
          {expanded === 'cards' && (
            <div style={expandableContentStyle}>
              <div style={{ fontWeight: 700, marginBottom: 8 }}>Top Cards</div>
              <ul style={{ listStyle: 'none', padding: 0, margin: 0 }}>
                {cardsPreview.map((card, idx) => (
                  <li key={card.id} style={{ padding: '4px 0', borderBottom: idx < cardsPreview.length - 1 ? '1px solid #eee' : 'none' }}>
                    <span style={{ fontWeight: 600 }}>Card #{card.card_id}</span> <span style={{ color: '#888', fontSize: 13 }}>User: {card.user_id}</span>
                  </li>
                ))}
              </ul>
              <a href="/cards/list" style={{ display: 'inline-block', marginTop: 12, color: 'var(--color-primary)', fontWeight: 700, textDecoration: 'none', background: '#fff', borderRadius: 8, padding: '6px 18px', boxShadow: '0 2px 8px var(--color-shadow)', transition: 'background 0.2s' }}>View All Cards</a>
            </div>
          )}
        </div>
        {/* Total Balance Card (not expandable) */}
        <div className="scale-hover fade-in" style={{ ...cardContainerStyle, background: 'linear-gradient(135deg, var(--color-accent) 30%, var(--color-bg) 90%)', color: '#22223b', cursor: 'default' }}>
          <h3 style={{ margin: '24px 0 0 0', fontWeight: 600, textAlign: 'center' }}>Total Balance</h3>
          <p style={{ fontSize: 40, margin: 0, fontWeight: 700, textAlign: 'center' }}>{stats.totalBalance}</p>
        </div>
      </div>
      <div className="fade-in" style={{ width: '100%', maxWidth: 480, background: '#fff', borderRadius: 'var(--radius)', boxShadow: '0 4px 24px var(--color-shadow)', padding: 24, marginTop: 16 }}>
        <h3 style={{ textAlign: 'center', color: 'var(--color-primary)', fontWeight: 700, marginBottom: 16 }}>Users vs Cards</h3>
        <ResponsiveContainer width="100%" height={220}>
          <PieChart>
            <Pie data={pieData} dataKey="value" nameKey="name" cx="50%" cy="50%" outerRadius={80} label>
              {pieData.map((entry, idx) => (
                <Cell key={`cell-${idx}`} fill={COLORS[idx % COLORS.length]} />
              ))}
            </Pie>
            <Tooltip />
          </PieChart>
        </ResponsiveContainer>
      </div>
    </div>
  );
};

export default DashboardPage;