import React from 'react';
import DeviceCreationForm from '../components/DeviceCreationForm';
import { Link } from 'react-router-dom';

const DevicesListPage: React.FC = () => {
  return (
    <div style={{ minHeight: '100vh', background: '#f2e9e4', padding: '40px 0' }}>
      <div style={{ maxWidth: 1200, margin: '0 auto', display: 'flex', justifyContent: 'space-between', alignItems: 'center', marginBottom: 32 }}>
        <h1 style={{ color: '#22223b', letterSpacing: 1 }}>Device Management</h1>
        <Link
          to="/devices/list"
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
          }}
          onMouseOver={e => { (e.currentTarget as HTMLAnchorElement).style.transform = 'scale(1.06)'; (e.currentTarget as HTMLAnchorElement).style.boxShadow = '0 4px 16px #b5ead744'; }}
          onMouseOut={e => { (e.currentTarget as HTMLAnchorElement).style.transform = 'scale(1)'; (e.currentTarget as HTMLAnchorElement).style.boxShadow = '0 2px 8px #b5ead722'; }}
        >
          View All Devices 
        </Link>
      </div>
      <div style={{ maxWidth: 420, margin: '0 auto', background: 'linear-gradient(135deg, #4a4e69 30%, #b5ead7 90%)', borderRadius: 16, boxShadow: '0 4px 24px #4a4e6922', padding: 32, marginBottom: 32, display: 'flex', flexDirection: 'column', alignItems: 'center' }}>
        <DeviceCreationForm />
      </div>
    </div>
  );
};

export default DevicesListPage; 