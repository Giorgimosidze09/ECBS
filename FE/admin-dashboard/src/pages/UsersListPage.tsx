import React, { useEffect, useState } from 'react';
import { fetchUsersList, getUserById, updateUser, deleteUser } from '../api';
import ReactModal from 'react-modal';
import { Link } from 'react-router-dom';
import GetByIdCard from '../components/GetByIdCard';

const userCardStyle: React.CSSProperties = {
  background: 'linear-gradient(135deg, var(--color-card) 60%, var(--color-secondary) 100%)',
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
    await updateUser(editUser.id, editForm);
    setActionMsg('User updated!');
    setEditModalOpen(false);
    setTimeout(() => setActionMsg(''), 2000);
    setLoading(true);
    fetchUsersList({ limit: 1000, offset: 0 }).then(data => {
      const usersArr = Array.isArray(data) ? data : data.users || [];
      setUsers(usersArr);
      setLoading(false);
    });
  };

  const handleDelete = async (userId: number) => {
    await deleteUser(userId);
    setActionMsg('User deleted!');
    setTimeout(() => setActionMsg(''), 2000);
    setLoading(true);
    fetchUsersList({ limit: 1000, offset: 0 }).then(data => {
      const usersArr = Array.isArray(data) ? data : data.users || [];
      setUsers(usersArr);
      setLoading(false);
    });
  };

  return (
    <div className="fade-in" style={{ minHeight: '100vh', background: 'var(--color-bg)', padding: '40px 0', display: 'flex', flexDirection: 'column', alignItems: 'center', justifyContent: 'center' }}>
      <div style={{ maxWidth: 1200, margin: '0 auto', display: 'flex', justifyContent: 'space-between', alignItems: 'center', marginBottom: 32, width: '100%' }}>
        <h1 style={{ color: 'var(--color-primary)', letterSpacing: 1, fontWeight: 800, fontSize: 32 }}>All Users</h1>
        <Link
          to="/users"
          className="scale-hover"
          style={{
            display: 'inline-block',
            background: 'linear-gradient(90deg, var(--color-primary) 60%, var(--color-secondary) 100%)',
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
          Back to User Creation
        </Link>
      </div>
      <div style={{ maxWidth: 400, width: '100%', marginLeft: 'auto', marginRight: 'auto', marginBottom: 32 }}>
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
        <div style={{ textAlign: 'center', color: 'var(--color-primary)', fontWeight: 500 }}>Loading...</div>
      ) : (
        <div style={{ display: 'flex', flexWrap: 'wrap', justifyContent: 'center', width: '100%', maxWidth: 1200 }}>
          {users.map(user => (
            <div key={user.id} className="scale-hover fade-in" style={userCardStyle}>
              <div style={{ fontWeight: 700, fontSize: 18, color: 'var(--color-bg-dark)', marginBottom: 8 }}>{user.name}</div>
              <div style={{ marginBottom: 6 }}><span style={{ color: 'var(--color-primary)', fontWeight: 600 }}>Email:</span> {user.email}</div>
              <div style={{ marginBottom: 6 }}><span style={{ color: 'var(--color-primary)', fontWeight: 600 }}>Phone:</span> {user.phone}</div>
              <div style={{ marginBottom: 6 }}><span style={{ color: 'var(--color-primary)', fontWeight: 600 }}>Cards:</span> {user.card_count}</div>
              <div style={{ marginBottom: 6 }}><span style={{ color: 'var(--color-primary)', fontWeight: 600 }}>Total Balance:</span> {user.total_balance}</div>
              <div style={{ display: 'flex', gap: 8, marginTop: 12 }}>
                <button onClick={() => handleEdit(user.id)} className="scale-hover" style={{ background: 'var(--color-primary)', color: '#fff', border: 'none', borderRadius: 6, padding: '8px 16px', fontWeight: 600, cursor: 'pointer', transition: 'background var(--transition), transform var(--transition)' }}>Edit</button>
                <button onClick={() => handleDelete(user.id)} className="scale-hover" style={{ background: '#c72c41', color: '#fff', border: 'none', borderRadius: 6, padding: '8px 16px', fontWeight: 600, cursor: 'pointer', transition: 'background var(--transition), transform var(--transition)' }}>Delete</button>
              </div>
            </div>
          ))}
        </div>
      )}
      <ReactModal isOpen={editModalOpen} onRequestClose={() => setEditModalOpen(false)} ariaHideApp={false} style={{ content: { maxWidth: 400, margin: 'auto', borderRadius: 16, padding: 0, boxShadow: '0 4px 24px var(--color-shadow)', background: 'linear-gradient(135deg, var(--color-card) 60%, var(--color-secondary) 100%)', border: 'none' } }}>
        <div style={{ padding: 24, borderRadius: 16 }}>
          <h2 style={{ fontWeight: 700, fontSize: 20, color: 'var(--color-bg-dark)', marginBottom: 18 }}>Edit User</h2>
          <form onSubmit={e => { e.preventDefault(); handleEditSave(); }}>
            <label style={{ fontWeight: 600, color: 'var(--color-primary)' }}>Name:</label>
            <input value={editForm.name} onChange={e => setEditForm(f => ({ ...f, name: e.target.value }))} required style={{ width: '100%', marginBottom: 12, padding: 8, borderRadius: 8, border: '1px solid #bcbcbc', fontSize: 16 }} />
            <label style={{ fontWeight: 600, color: 'var(--color-primary)' }}>Email:</label>
            <input value={editForm.email} onChange={e => setEditForm(f => ({ ...f, email: e.target.value }))} required style={{ width: '100%', marginBottom: 12, padding: 8, borderRadius: 8, border: '1px solid #bcbcbc', fontSize: 16 }} />
            <label style={{ fontWeight: 600, color: 'var(--color-primary)' }}>Phone:</label>
            <input value={editForm.phone} onChange={e => setEditForm(f => ({ ...f, phone: e.target.value }))} required style={{ width: '100%', marginBottom: 18, padding: 8, borderRadius: 8, border: '1px solid #bcbcbc', fontSize: 16 }} />
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

export default UsersListPage; 