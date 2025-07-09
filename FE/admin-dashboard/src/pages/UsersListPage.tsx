import React, { useEffect, useState } from 'react';
import { fetchUsersList, getUserById, updateUser, deleteUser } from '../api';
import ReactModal from 'react-modal';
import { Link } from 'react-router-dom';
import GetByIdCard from '../components/GetByIdCard';

const userCardStyle: React.CSSProperties = {
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
};

const UsersListPage: React.FC = () => {
  const [users, setUsers] = useState<any[]>([]);
  const [loading, setLoading] = useState(true);
  const [editUser, setEditUser] = useState<any | null>(null);
  const [editModalOpen, setEditModalOpen] = useState(false);
  const [editForm, setEditForm] = useState({ name: '', email: '', phone: '' });
  const [actionMsg, setActionMsg] = useState('');

  useEffect(() => {
    setLoading(true);
    fetchUsersList({ limit: 1000, offset: 0 }).then(data => {
      const usersArr = Array.isArray(data) ? data : data.users || [];
      setUsers(usersArr);
      setLoading(false);
    });
  }, []);

  const handleEdit = async (userId: number) => {
    const user = await getUserById(userId);
    setEditUser(user);
    setEditForm({
      name: user.name || '',
      email: user.email || '',
      phone: user.phone || ''
    });
    setEditModalOpen(true);
  };

  const handleEditSave = async () => {
    if (!editUser) return;
    try {
      await updateUser(editUser.id, { ...editForm, id: editUser.id });
      setActionMsg('User updated!');
      setEditModalOpen(false);
      setEditUser(null);
      setEditForm({ name: '', email: '', phone: '' });
      setTimeout(() => setActionMsg(''), 2000);
      // reload users
      setLoading(true);
      fetchUsersList({ limit: 1000, offset: 0 }).then(data => {
        const usersArr = Array.isArray(data) ? data : data.users || [];
        setUsers(usersArr);
        setLoading(false);
      });
    } catch {
      setActionMsg('Failed to update user.');
    }
  };

  const handleDelete = async (userId: number) => {
    if (!window.confirm('Are you sure you want to delete this user?')) return;
    try {
      await deleteUser(userId);
      setActionMsg('User deleted!');
      setTimeout(() => setActionMsg(''), 2000);
      // reload users
      setLoading(true);
      fetchUsersList({ limit: 1000, offset: 0 }).then(data => {
        const usersArr = Array.isArray(data) ? data : data.users || [];
        setUsers(usersArr);
        setLoading(false);
      });
    } catch {
      setActionMsg('Failed to delete user.');
    }
  };

  return (
    <div style={{ minHeight: '100vh', background: '#f2e9e4', padding: '40px 0' }}>
      <div style={{ maxWidth: 1200, margin: '0 auto', display: 'flex', justifyContent: 'space-between', alignItems: 'center', marginBottom: 32 }}>
        <h1 style={{ color: '#4a4e69', letterSpacing: 1 }}>All Users</h1>
        <Link
          to="/users"
          style={{
            display: 'inline-block',
            background: 'linear-gradient(90deg, #4a4e69 60%, #9a8c98 100%)',
            color: '#fff',
            border: 'none',
            borderRadius: 8,
            padding: '10px 28px',
            fontWeight: 700,
            fontSize: 16,
            boxShadow: '0 2px 8px #9a8c9822',
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
          <span style={{ position: 'relative', zIndex: 2 }}>Back to User Creation</span>
        </Link>
      </div>
      <div style={{ display: 'flex', justifyContent: 'center', marginBottom: 32 }}>
        <GetByIdCard
          label="User"
          fetchById={getUserById}
          renderResult={user => (
            <div>
              <div><b>ID:</b> {user.id}</div>
              <div><b>Name:</b> {user.name}</div>
              <div><b>Email:</b> {user.email}</div>
              {user.phone && <div><b>Phone:</b> {user.phone}</div>}
            </div>
          )}
        />
      </div>
      {actionMsg && <div style={{ color: actionMsg.includes('Failed') ? 'red' : 'green', textAlign: 'center', marginBottom: 12 }}>{actionMsg}</div>}
      {loading ? (
        <div style={{ textAlign: 'center', color: '#4a4e69', fontWeight: 500 }}>Loading...</div>
      ) : (
        <div style={{ display: 'flex', flexWrap: 'wrap', justifyContent: 'center' }}>
          {users.map(user => (
            <div key={user.id} style={userCardStyle}>
              <div style={{ fontWeight: 700, fontSize: 18, color: '#22223b', marginBottom: 8 }}>{user.name}</div>
              <div style={{ marginBottom: 6 }}><span style={{ color: '#4a4e69', fontWeight: 600 }}>Email:</span> {user.email}</div>
              <div style={{ marginBottom: 6 }}><span style={{ color: '#4a4e69', fontWeight: 600 }}>Phone:</span> {user.phone}</div>
              <div style={{ marginBottom: 6 }}><span style={{ color: '#4a4e69', fontWeight: 600 }}>Cards:</span> {user.card_count}</div>
              <div style={{ marginBottom: 6 }}><span style={{ color: '#4a4e69', fontWeight: 600 }}>Total Balance:</span> {user.total_balance}</div>
              <div style={{ display: 'flex', gap: 8, marginTop: 12 }}>
                <button onClick={() => handleEdit(user.id)} style={{ background: '#4a4e69', color: '#fff', border: 'none', borderRadius: 6, padding: '8px 16px', fontWeight: 600, cursor: 'pointer' }}>Edit</button>
                <button onClick={() => handleDelete(user.id)} style={{ background: '#c72c41', color: '#fff', border: 'none', borderRadius: 6, padding: '8px 16px', fontWeight: 600, cursor: 'pointer' }}>Delete</button>
              </div>
            </div>
          ))}
        </div>
      )}
      <ReactModal isOpen={editModalOpen} onRequestClose={() => setEditModalOpen(false)} ariaHideApp={false} style={{ content: { maxWidth: 400, margin: 'auto', borderRadius: 16, padding: 0, boxShadow: '0 4px 24px #9a8c9822', background: 'linear-gradient(135deg, #fff 60%, #9a8c98 100%)', border: 'none' } }}>
        <div style={{ padding: 24, borderRadius: 16 }}>
          <h2 style={{ fontWeight: 700, fontSize: 20, color: '#22223b', marginBottom: 18 }}>Edit User</h2>
          <form onSubmit={e => { e.preventDefault(); handleEditSave(); }}>
            <label style={{ fontWeight: 600, color: '#4a4e69' }}>Name:</label>
            <input value={editForm.name} onChange={e => setEditForm(f => ({ ...f, name: e.target.value }))} required style={{ width: '100%', marginBottom: 12, padding: 8, borderRadius: 8, border: '1px solid #bcbcbc', fontSize: 16 }} />
            <label style={{ fontWeight: 600, color: '#4a4e69' }}>Email:</label>
            <input value={editForm.email} onChange={e => setEditForm(f => ({ ...f, email: e.target.value }))} required style={{ width: '100%', marginBottom: 12, padding: 8, borderRadius: 8, border: '1px solid #bcbcbc', fontSize: 16 }} />
            <label style={{ fontWeight: 600, color: '#4a4e69' }}>Phone:</label>
            <input value={editForm.phone} onChange={e => setEditForm(f => ({ ...f, phone: e.target.value }))} required style={{ width: '100%', marginBottom: 18, padding: 8, borderRadius: 8, border: '1px solid #bcbcbc', fontSize: 16 }} />
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

export default UsersListPage; 