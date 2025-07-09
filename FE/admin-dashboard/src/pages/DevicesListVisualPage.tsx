import React, { useEffect, useState } from 'react';
import { fetchDevicesList, getDeviceById, updateDevice, deleteDevice } from '../api';
import { Device } from '../types';
import ReactModal from 'react-modal';
import { Link } from 'react-router-dom';
import GetByIdCard from '../components/GetByIdCard';

const deviceCardStyle: React.CSSProperties = {
  background: 'linear-gradient(135deg, #fff 60%, #b5ead7 100%)',
  borderRadius: 16,
  boxShadow: '0 4px 24px #b5ead722',
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

const DevicesListVisualPage: React.FC = () => {
  const [devices, setDevices] = useState<Device[]>([]);
  const [loading, setLoading] = useState(true);
  const [editDevice, setEditDevice] = useState<Device | null>(null);
  const [editModalOpen, setEditModalOpen] = useState(false);
  const [editForm, setEditForm] = useState({ device_id: '', location: '', active: true });
  const [actionMsg, setActionMsg] = useState('');

  const reloadDevices = () => {
    setLoading(true);
    fetchDevicesList({ limit: 1000, offset: 0 }).then(data => {
      const devicesArr = Array.isArray(data) ? data : data.devices || data.rows || [];
      setDevices(devicesArr);
      setLoading(false);
    });
  };

  useEffect(() => {
    reloadDevices();
  }, []);

  const handleEdit = async (deviceId: number) => {
    const device = await getDeviceById(deviceId);
    setEditDevice(device);
    setEditForm({
      device_id: device.device_id || '',
      location: device.location || '',
      active: device.active !== undefined ? device.active : true
    });
    setEditModalOpen(true);
  };

  const handleEditSave = async () => {
    if (!editDevice) return;
    try {
      await updateDevice(editDevice.id, { ...editForm, id: editDevice.id });
      setActionMsg('Device updated!');
      setEditModalOpen(false);
      setEditDevice(null);
      setEditForm({ device_id: '', location: '', active: true });
      setTimeout(() => setActionMsg(''), 2000);
      reloadDevices();
    } catch {
      setActionMsg('Failed to update device.');
    }
  };

  const handleDelete = async (deviceId: number) => {
    if (!window.confirm('Are you sure you want to delete this device?')) return;
    try {
      await deleteDevice(deviceId);
      setActionMsg('Device deleted!');
      setTimeout(() => setActionMsg(''), 2000);
      reloadDevices();
    } catch {
      setActionMsg('Failed to delete device.');
    }
  };

  return (
    <div style={{ minHeight: '100vh', background: '#f2e9e4', padding: '40px 0' }}>
      <div style={{ maxWidth: 1200, margin: '0 auto', display: 'flex', justifyContent: 'space-between', alignItems: 'center', marginBottom: 32 }}>
        <h1 style={{ color: '#4a4e69', letterSpacing: 1 }}>All Devices</h1>
        <Link
          to="/devices"
          style={{
            display: 'inline-block',
            background: 'linear-gradient(90deg, #4a4e69 60%, #b5ead7 100%)',
            color: '#fff',
            border: 'none',
            borderRadius: 8,
            padding: '10px 28px',
            fontWeight: 700,
            fontSize: 16,
            boxShadow: '0 2px 8px #b5ead722',
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
          <span style={{ position: 'relative', zIndex: 2 }}>Back to Device Management</span>
        </Link>
      </div>
      <div style={{ maxWidth: 600, margin: '0 auto', marginBottom: 32 }}>
        <GetByIdCard
          label="Device"
          fetchById={getDeviceById}
          renderResult={device => (
            <div>
              <div><b>ID:</b> {device.id}</div>
              <div><b>Device ID:</b> {device.device_id}</div>
              <div><b>Location:</b> {device.location}</div>
              <div><b>Installed At:</b> {device.installed_at ? new Date(device.installed_at).toLocaleString() : 'N/A'}</div>
              <div><b>Active:</b> {device.active ? 'Yes' : 'No'}</div>
            </div>
          )}
        />
      </div>
      {actionMsg && <div style={{ color: actionMsg.includes('Failed') ? 'red' : 'green', textAlign: 'center', marginBottom: 12 }}>{actionMsg}</div>}
      {loading ? (
        <div style={{ textAlign: 'center', color: '#4a4e69', fontWeight: 500 }}>Loading...</div>
      ) : (
        <div style={{ display: 'flex', flexWrap: 'wrap', justifyContent: 'center' }}>
          {devices.map(device => (
            <div key={device.id} style={deviceCardStyle}>
              <div style={{ fontWeight: 700, fontSize: 18, color: '#22223b', marginBottom: 8 }}>Device #{device.id}</div>
              <div style={{ marginBottom: 6 }}><span style={{ color: '#4a4e69', fontWeight: 600 }}>Device ID:</span> {device.device_id}</div>
              <div style={{ marginBottom: 6 }}><span style={{ color: '#4a4e69', fontWeight: 600 }}>Location:</span> {device.location}</div>
              <div style={{ marginBottom: 6 }}><span style={{ color: '#4a4e69', fontWeight: 600 }}>Installed At:</span> {device.installed_at ? new Date(device.installed_at).toLocaleString() : 'N/A'}</div>
              <div style={{ marginBottom: 6 }}><span style={{ color: '#4a4e69', fontWeight: 600 }}>Active:</span> {device.active ? 'Yes' : 'No'}</div>
              <div style={{ display: 'flex', gap: 8, marginTop: 12 }}>
                <button onClick={() => handleEdit(device.id)} style={{ background: '#4a4e69', color: '#fff', border: 'none', borderRadius: 6, padding: '8px 16px', fontWeight: 600, cursor: 'pointer' }}>Edit</button>
                <button onClick={() => handleDelete(device.id)} style={{ background: '#c72c41', color: '#fff', border: 'none', borderRadius: 6, padding: '8px 16px', fontWeight: 600, cursor: 'pointer' }}>Delete</button>
              </div>
            </div>
          ))}
        </div>
      )}
      <ReactModal isOpen={editModalOpen} onRequestClose={() => setEditModalOpen(false)} ariaHideApp={false} style={{ content: { maxWidth: 400, margin: 'auto', borderRadius: 16, padding: 0, boxShadow: '0 4px 24px #b5ead722', background: 'linear-gradient(135deg, #fff 60%, #b5ead7 100%)', border: 'none' } }}>
        <div style={{ padding: 24, borderRadius: 16 }}>
          <h2 style={{ fontWeight: 700, fontSize: 20, color: '#22223b', marginBottom: 18 }}>Edit Device</h2>
          <form onSubmit={e => { e.preventDefault(); handleEditSave(); }}>
            <label style={{ fontWeight: 600, color: '#4a4e69' }}>Device ID:</label>
            <input value={editForm.device_id} onChange={e => setEditForm(f => ({ ...f, device_id: e.target.value }))} required style={{ width: '100%', marginBottom: 12, padding: 8, borderRadius: 8, border: '1px solid #bcbcbc', fontSize: 16 }} />
            <label style={{ fontWeight: 600, color: '#4a4e69' }}>Location:</label>
            <input value={editForm.location} onChange={e => setEditForm(f => ({ ...f, location: e.target.value }))} required style={{ width: '100%', marginBottom: 12, padding: 8, borderRadius: 8, border: '1px solid #bcbcbc', fontSize: 16 }} />
            <label style={{ fontWeight: 600, color: '#4a4e69' }}>Active:</label>
            <select value={editForm.active ? 'true' : 'false'} onChange={e => setEditForm(f => ({ ...f, active: e.target.value === 'true' }))} style={{ width: '100%', marginBottom: 18, padding: 8, borderRadius: 8, border: '1px solid #bcbcbc', fontSize: 16 }}>
              <option value="true">Yes</option>
              <option value="false">No</option>
            </select>
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

export default DevicesListVisualPage; 