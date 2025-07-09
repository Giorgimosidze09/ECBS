import React, { useEffect, useState } from 'react';
import { PieChart, Pie, Cell, Tooltip, ResponsiveContainer } from 'recharts';
import axios from 'axios';

const API_URL = 'http://localhost:8080';
//const API_URL = 'http://localhost:5204';

const COLORS = ['#4a4e69', '#9a8c98', '#c9ada7', '#22223b'];

const DashboardPage: React.FC = () => {
  const [stats, setStats] = useState({
    users: 0,
    cards: 0,
    totalBalance: 0,
  });

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
  }, []);

  const pieData = [
    { name: 'Users', value: stats.users },
    { name: 'Cards', value: stats.cards },
  ];

  return (
    <div style={{ padding: 32, background: '#f2e9e4', minHeight: '100vh' }}>
      <h1 style={{ color: '#4a4e69', marginBottom: 32, textAlign: 'center', letterSpacing: 1 }}>Dashboard</h1>
      <div style={{
        display: 'flex',
        gap: 32,
        marginBottom: 32,
        flexWrap: 'wrap',
        justifyContent: 'center'
      }}>
        <div style={{
          background: 'linear-gradient(135deg, #4a4e69 30%, #9a8c98 90%)',
          color: '#fff',
          borderRadius: 16,
          padding: 32,
          minWidth: 200,
          boxShadow: '0 4px 24px #4a4e6922',
          textAlign: 'center'
        }}>
          <h3 style={{ margin: 0, fontWeight: 600 }}>Total Users</h3>
          <p style={{ fontSize: 40, margin: 0, fontWeight: 700 }}>{stats.users}</p>
        </div>
        <div style={{
          background: 'linear-gradient(135deg, #9a8c98 30%, #c9ada7 90%)',
          color: '#22223b',
          borderRadius: 16,
          padding: 32,
          minWidth: 200,
          boxShadow: '0 4px 24px #9a8c9822',
          textAlign: 'center'
        }}>
          <h3 style={{ margin: 0, fontWeight: 600 }}>Total Cards</h3>
          <p style={{ fontSize: 40, margin: 0, fontWeight: 700 }}>{stats.cards}</p>
        </div>
        <div style={{
          background: 'linear-gradient(135deg, #c9ada7 30%, #f2e9e4 90%)',
          color: '#22223b',
          borderRadius: 16,
          padding: 32,
          minWidth: 200,
          boxShadow: '0 4px 24px #c9ada722',
          textAlign: 'center'
        }}>
          <h3 style={{ margin: 0, fontWeight: 600 }}>Total Balance</h3>
          <p style={{ fontSize: 40, margin: 0, fontWeight: 700 }}>{stats.totalBalance}</p>
        </div>
      </div>
      <div style={{
        background: '#fff',
        borderRadius: 16,
        padding: 32,
        maxWidth: 440,
        margin: '0 auto',
        boxShadow: '0 4px 24px #22223b22'
      }}>
        <h3 style={{ color: '#4a4e69', marginBottom: 16, textAlign: 'center' }}>Users vs Cards</h3>
        <ResponsiveContainer width="100%" height={250}>
          <PieChart>
            <Pie
              data={pieData}
              dataKey="value"
              nameKey="name"
              cx="50%"
              cy="50%"
              outerRadius={80}
              fill="#8884d8"
              label
            >
              {pieData.map((entry, index) => (
                <Cell key={`cell-${index}`} fill={COLORS[index % COLORS.length]} />
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