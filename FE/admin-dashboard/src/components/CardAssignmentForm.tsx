import React, { useEffect, useState, useRef } from 'react';
import { fetchUsersList, fetchDevicesList, assignCard } from '../api';


const inputStyle: React.CSSProperties = {
    width: '100%',
    padding: '10px 12px',
    borderRadius: 6,
    border: '1px solid #bcbcbc',
    fontSize: '1rem',
    background: '#fff',
    boxSizing: 'border-box',
    transition: 'border 0.2s ease-in-out'
};

const popupStyle: React.CSSProperties = {
    position: 'absolute',
    zIndex: 10,
    background: '#fff',
    border: '1px solid #bcbcbc',
    borderRadius: 6,
    boxShadow: '0 4px 16px #0001',
    width: '100%',
    maxHeight: 180,
    overflowY: 'auto',
    marginTop: 2
};

const popupItemStyle: React.CSSProperties = {
    padding: '10px 12px',
    cursor: 'pointer',
    borderBottom: '1px solid #eee',
    background: '#fff'
};

const popupItemActiveStyle: React.CSSProperties = {
    ...popupItemStyle,
    background: '#f2e9e4'
};

const CardAssignmentForm: React.FC = () => {
    const [users, setUsers] = useState<{ id: number; name: string }[]>([]);
    const [devices, setDevices] = useState<{ id: number; location: string }[]>([]);
    const [selectedUser, setSelectedUser] = useState<string>('');
    const [selectedDevice, setSelectedUDevice] = useState<string>('');
    const [cardId, setCardId] = useState('');
    const [loading, setLoading] = useState(false);
    const [success, setSuccess] = useState<string | null>(null);
    const [error, setError] = useState<string | null>(null);
    const [showPopup, setShowPopup] = useState(false);
    const [activeIndex, setActiveIndex] = useState(-1);
    const inputRef = useRef<HTMLInputElement>(null);
    const [cardType, setCardType] = useState('balance');
    const [activationStart, setActivationStart] = useState('');
    const [activationEnd, setActivationEnd] = useState('');

    useEffect(() => {
        fetchUsersList({ limit: 1000, offset: 0 }).then(data => {
            const usersArr = Array.isArray(data) ? data : data.users || [];
            setUsers(usersArr.map((u: any) => ({
                id: u.id ?? u.ID,
                name: u.name ?? u.Name
            })));
        });

         fetchDevicesList({ limit: 1000, offset: 0 }).then(data => {
        const devicesArr = Array.isArray(data) ? data : data.devices || [];
        setDevices(devicesArr.map((d: any) => ({
            id: d.id ?? d.ID,
            location: d.location ?? d.Location
        })));
    });
    }, []);

    const filteredUsers = users.filter(user =>
        user.name.toLowerCase().includes(selectedUser.trim().toLowerCase()) ||
        String(user.id).includes(selectedUser.trim())
    );

    const handleKeyDown = (e: React.KeyboardEvent<HTMLInputElement>) => {
        if (!showPopup) return;
        if (e.key === 'ArrowDown') {
            setActiveIndex(i => Math.min(i + 1, filteredUsers.length - 1));
        } else if (e.key === 'ArrowUp') {
            setActiveIndex(i => Math.max(i - 1, 0));
        } else if (e.key === 'Enter' && activeIndex >= 0) {
            const user = filteredUsers[activeIndex];
            setSelectedUser(`${user.name} (ID: ${user.id})`);
            setShowPopup(false);
            setActiveIndex(-1);
        } else if (e.key === 'Escape') {
            setShowPopup(false);
        }
    };

    useEffect(() => {
        const handleClick = (e: MouseEvent) => {
            if (inputRef.current && !inputRef.current.contains(e.target as Node)) {
                setShowPopup(false);
            }
        };
        if (showPopup) {
            document.addEventListener('mousedown', handleClick);
        }
        return () => document.removeEventListener('mousedown', handleClick);
    }, [showPopup]);

    const handleUserSelect = (user: { id: number; name: string }) => {
        setSelectedUser(`${user.name} (ID: ${user.id})`);
        setShowPopup(false);
        setActiveIndex(-1);
    };

    const handleSubmit = async (e: React.FormEvent) => {
        e.preventDefault();
        setLoading(true);
        setSuccess(null);
        setError(null);

        const userObj = users.find(u => `${u.name} (ID: ${u.id})` === selectedUser || String(u.id) === selectedUser);
        if (!userObj) {
            setError('Please select a valid user.');
            setLoading(false);
            return;
        }

          const deviceOjb = devices.find(u => `${u.location} (ID: ${u.id})` === selectedDevice || String(u.id) === selectedDevice);
        if (!deviceOjb) {
            setError('Please select a valid user.');
            setLoading(false);
            return;
        }

        try {
            const payload: any = {
                user_id: userObj.id,
                card_id: cardId,
                device_id: deviceOjb?.id,
                type: cardType
            };
            if (cardType === 'activation') {
                payload.activation_start = activationStart;
                payload.activation_end = activationEnd;
            }
            await assignCard(payload);
            setSuccess('Card assigned successfully!');
            setSelectedUser('');
            setCardId('');
        } catch {
            setError('Failed to assign card.');
        } finally {
            setLoading(false);
        }
    };

    const isFormReady =
    selectedUser.trim() !== '' &&
    cardId.trim() !== '' &&
    selectedDevice.trim() !== '';


    return (
        <form
            onSubmit={handleSubmit}
            style={{
                width: '100%',
                display: 'flex',
                flexDirection: 'column',
                alignItems: 'center',
                gap: 18,
                transition: 'opacity 0.3s ease-in-out'
            }}
        >
            {/* User Field with Popup */}
            <div style={{ width: '100%', position: 'relative' }}>
                <label className="styled-label" style={{ display: 'block', marginBottom: 6, color: '#fff' }}>
                    Select User
                </label>
                <input
                    ref={inputRef}
                    type="text"
                    value={selectedUser}
                    onChange={e => {
                        setSelectedUser(e.target.value);
                        setShowPopup(true);
                        setActiveIndex(-1);
                    }}
                    onFocus={() => setShowPopup(true)}
                    onBlur={() => setTimeout(() => setShowPopup(false), 150)}
                    onClick={() => setShowPopup(true)}
                    onKeyDown={handleKeyDown}
                    required
                    style={inputStyle}
                    placeholder="Type or select user"
                    autoComplete="off"
                />
                {showPopup && filteredUsers.length > 0 && (
                    <div style={popupStyle}>
                        {filteredUsers.map((user, idx) => (
                            <div
                                key={user.id}
                                style={activeIndex === idx ? popupItemActiveStyle : popupItemStyle}
                                onMouseDown={() => handleUserSelect(user)}
                                onMouseEnter={() => setActiveIndex(idx)}
                                onMouseLeave={() => setActiveIndex(-1)}
                            >
                                {user.name} <span style={{ color: '#888' }}>(ID: {user.id})</span>
                            </div>
                        ))}
                    </div>
                )}
            </div>


{/* User Field with Popup (always visible) */}
<div style={{ width: '100%', position: 'relative' }}>
  {/* ...user input code remains unchanged */}
</div>

{/* Show Card ID input only if user selected */}
{selectedUser.trim() !== '' && (
  <div style={{ width: '100%' }}>
    <label className="styled-label" style={{ display: 'block', marginBottom: 6, color: '#fff' }}>
      Enter Card ID
    </label>
    <input
      type="text"
      value={cardId}
      onChange={e => setCardId(e.target.value)}
      required
      style={inputStyle}
      placeholder="Enter Card ID"
      autoComplete="off"
    />
  </div>
)}

{/* Show Device dropdown only if cardId is filled */}
{cardId.trim() !== '' && (
  <div style={{ width: '100%' }}>
    <label className="styled-label" style={{ display: 'block', marginBottom: 6, color: '#fff' }}>
      Select Device
    </label>
    <select
      value={selectedDevice}
      onChange={e => setSelectedUDevice(e.target.value)}
      required
      style={{
        ...inputStyle,
        background: '#fff',
        fontSize: '1rem'
      }}
    >
      <option value="">-- Select a device --</option>
      {devices.map(device => (
        <option key={device.id} value={`${device.location} (ID: ${device.id})`}>
          {device.location} (ID: {device.id})
        </option>
      ))}
    </select>
  </div>
)}

{/* Card Type Selection */}
<div style={{ width: '100%' }}>
  <label className="styled-label" style={{ display: 'block', marginBottom: 6, color: '#fff' }}>
    Card Type
  </label>
  <select
    value={cardType}
    onChange={e => setCardType(e.target.value)}
    style={inputStyle}
    required
  >
    <option value="balance">Balance</option>
    <option value="activation">Activation</option>
  </select>
</div>

{/* Activation Dates if type is activation */}
{cardType === 'activation' && (
  <>
    <div style={{ width: '100%' }}>
      <label className="styled-label" style={{ display: 'block', marginBottom: 6, color: '#fff' }}>
        Activation Start
      </label>
      <input
        type="date"
        value={activationStart}
        onChange={e => setActivationStart(e.target.value)}
        required={cardType === 'activation'}
        style={inputStyle}
      />
    </div>
    <div style={{ width: '100%' }}>
      <label className="styled-label" style={{ display: 'block', marginBottom: 6, color: '#fff' }}>
        Activation End
      </label>
      <input
        type="date"
        value={activationEnd}
        onChange={e => setActivationEnd(e.target.value)}
        required={cardType === 'activation'}
        style={inputStyle}
      />
    </div>
  </>
)}

            {/* Submit Button */}
            <button
                type="submit"
                className="styled-button"
                style={{
                    width: '100%',
                    padding: '12px 0',
                    borderRadius: 6,
                    background: isFormReady ? '#22223b' : '#999',
                    color: '#fff',
                    fontWeight: 600,
                    fontSize: '1rem',
                    marginTop: 8,
                    cursor: loading || !isFormReady ? 'not-allowed' : 'pointer',
                    opacity: loading ? 0.6 : 1,
                    transition: 'background 0.3s ease-in-out'
                }}
                disabled={loading || !isFormReady}
            >
                {loading ? 'Assigning...' : 'Assign Card'}
            </button>

            {/* Feedback */}
            {success && <div style={{ color: 'lightgreen', marginTop: 12 }}>{success}</div>}
            {error && <div style={{ color: 'salmon', marginTop: 12 }}>{error}</div>}
        </form>
    );
};

export default CardAssignmentForm;
